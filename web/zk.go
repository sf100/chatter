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
	"github.com/samuel/go-zookeeper/zk"
	myrpc "github.com/sf100/chatter/rpc"
	myzk "github.com/sf100/chatter/zk"
)

func InitZK() (*zk.Conn, error) {
	conn, err := myzk.Connect(Conf.ZookeeperAddr, Conf.ZookeeperTimeout)
	if err != nil {
		log.Error("zk.Connect() error(%v)", err)
		return nil, err
	}
	myrpc.InitComet(conn, Conf.ZookeeperMigratePath, Conf.ZookeeperCometPath, Conf.RPCRetry, Conf.RPCPing)
	myrpc.InitMessage(conn, Conf.ZookeeperMessagePath, Conf.RPCRetry, Conf.RPCPing)
	return conn, nil
}
