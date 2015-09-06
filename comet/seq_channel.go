// Copyright © 2014 Terry Mao, LiuDing All rights reserved.
// This file is part of gopush-cluster.

// gopush-cluster is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// gopush-cluster is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with gopush-cluster.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	log "code.google.com/p/log4go"
	"errors"
	"github.com/sf100/chatter/hlist"
	"github.com/sf100/chatter/id"
	myrpc "github.com/sf100/chatter/rpc"
	"sync"
)

var (
	ErrMessageSave   = errors.New("Message set failed")
	ErrMessageGet    = errors.New("Message get failed")
	ErrMessageRPC    = errors.New("Message RPC not init")
	ErrAssectionConn = errors.New("Assection type Connection failed")
)

// Sequence Channel struct.
type SeqChannel struct {
	// Mutex
	mutex *sync.Mutex
	// client conn double linked-list
	conn *hlist.Hlist
	// Remove time id or lazy New
	// timeID *id.TimeID
	// token
	token *Token
}

// New a user seq stored message channel.
func NewSeqChannel() *SeqChannel {
	ch := &SeqChannel{
		mutex: &sync.Mutex{},
		conn:  hlist.New(),
		//timeID: id.NewTimeID(),
		token: nil,
	}
	// save memory
	if Conf.Auth {
		ch.token = NewToken()
	}
	return ch
}

// AddToken implements the Channel AddToken method.
func (c *SeqChannel) AddToken(key, token string) error {
	if !Conf.Auth {
		return nil
	}
	c.mutex.Lock()
	if err := c.token.Add(token); err != nil {
		c.mutex.Unlock()
		log.Error("user_key:\"%s\" c.token.Add(\"%s\") error(%v)", key, token, err)
		return err
	}
	c.mutex.Unlock()
	return nil
}

// AuthToken implements the Channel AuthToken method.
func (c *SeqChannel) AuthToken(key, token string) bool {
	if !Conf.Auth {
		return true
	}
	c.mutex.Lock()
	if err := c.token.Auth(token); err != nil {
		c.mutex.Unlock()
		log.Error("user_key:\"%s\" c.token.Auth(\"%s\") error(%v)", key, token, err)
		return false
	}
	c.mutex.Unlock()
	return true
}

// WriteMsg implements the Channel WriteMsg method.
func (c *SeqChannel) WriteMsg(key string, m *myrpc.Message) (err error) {
	c.mutex.Lock()
	err = c.writeMsg(key, m)
	c.mutex.Unlock()
	return
}

// writeMsg write msg to conn.
func (c *SeqChannel) writeMsg(key string, m *myrpc.Message) (err error) {
	var (
		oldMsg, msg, sendMsg []byte
	)
	// push message
	for e := c.conn.Front(); e != nil; e = e.Next() {
		conn, _ := e.Value.(*Connection)
		// if version empty then use old protocol
		if conn.Version == "" {
			if oldMsg == nil {
				if oldMsg, err = m.OldBytes(); err != nil {
					return
				}
			}
			sendMsg = oldMsg
		} else {
			if msg == nil {
				if msg, err = m.Bytes(); err != nil {
					return
				}
			}
			sendMsg = msg
		}
		// TODO use goroutine
		conn.Write(key, sendMsg)
	}
	return
}

// PushMsg implements the Channel PushMsg method.
func (c *SeqChannel) PushMsg(key, fkey string, m *myrpc.Message, expire uint) (err error) {
	client := myrpc.MessageRPC.Get()
	if client == nil {
		return ErrMessageRPC
	}
	c.mutex.Lock()
	// private message need persistence
	// if message expired no need persistence, only send online message
	// rewrite message id
	//m.MsgId = c.timeID.ID()
	m.MsgId = id.Get()
	if m.GroupId != myrpc.PublicGroupId && expire > 0 {
		args := &myrpc.MessageSavePrivateArgs{Key: key, Msg: m.Msg, MsgId: m.MsgId, Expire: expire}
		ret := 0
		if err = client.Call(myrpc.MessageServiceSavePrivate, args, &ret); err != nil {
			c.mutex.Unlock()
			log.Error("%s(\"%s\", \"%v\", &ret) error(%v)", myrpc.MessageServiceSavePrivate, key, args, err)
			return
		}
	}
	// push message
	if err = c.writeMsg(key, m); err != nil {
		c.mutex.Unlock()
		log.Error("c.WriteMsg(\"%s\", m) error(%v)", key, err)
		return
	}
	c.mutex.Unlock()
	return
}

// AddConn implements the Channel AddConn method.
func (c *SeqChannel) AddConn(key string, conn *Connection) (*hlist.Element, error) {
	c.mutex.Lock()
	if c.conn.Len()+1 > Conf.MaxSubscriberPerChannel {
		c.mutex.Unlock()
		log.Error("user_key:\"%s\" exceed conn", key)
		return nil, ErrMaxConn
	}
	// send first heartbeat to tell client service is ready for accept heartbeat
	if _, err := conn.Conn.Write(HeartbeatReply); err != nil {
		c.mutex.Unlock()
		log.Error("user_key:\"%s\" write first heartbeat to client error(%v)", key, err)
		return nil, err
	}
	// add conn
	conn.Buf = make(chan []byte, Conf.MsgBufNum)
	conn.HandleWrite(key)
	e := c.conn.PushFront(conn)
	c.mutex.Unlock()
	ConnStat.IncrAdd()
	log.Info("user_key:\"%s\" add conn = %d", key, c.conn.Len())
	return e, nil
}

// RemoveConn implements the Channel RemoveConn method.
func (c *SeqChannel) RemoveConn(key string, e *hlist.Element) error {
	c.mutex.Lock()
	tmp := c.conn.Remove(e)
	c.mutex.Unlock()
	conn, ok := tmp.(*Connection)
	if !ok {
		return ErrAssectionConn
	}
	close(conn.Buf)
	ConnStat.IncrRemove()
	log.Info("user_key:\"%s\" remove conn = %d", key, c.conn.Len())
	return nil
}

// Close implements the Channel Close method.
func (c *SeqChannel) Close() error {
	c.mutex.Lock()
	for e := c.conn.Front(); e != nil; e = e.Next() {
		if conn, ok := e.Value.(*Connection); !ok {
			c.mutex.Unlock()
			return ErrAssectionConn
		} else {
			if err := conn.Conn.Close(); err != nil {
				// ignore close error
				log.Warn("conn.Close() error(%v)", err)
			}
		}
	}
	c.mutex.Unlock()
	return nil
}
