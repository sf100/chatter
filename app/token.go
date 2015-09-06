package app

import (
	log "code.google.com/p/log4go"
	"github.com/sf100/go-uuid/uuid"
	"time"
)

// 令牌生成.
func genToken(user *User) (string, error) {

	conn := rs.getConn("token")
	if conn == nil {
		return "", RedisNoConnErr
	}

	defer conn.Close()

	confExpire := int64(Conf.TokenExpire)
	expire := confExpire + time.Now().Unix()
	token := user.Id + "_" + uuid.New()

	// 使用 Redis Hash 结构保存用户令牌值
	if err := conn.Send("HSET", token, "expire", expire); err != nil {
		log.Error(err)
		return "", err
	}

	// 设置令牌过期时间
	if err := conn.Send("EXPIRE", token, confExpire); err != nil {
		log.Error(err)
		return "", err
	}

	if err := conn.Flush(); err != nil {
		log.Error(err)
		return "", err
	}

	_, err := conn.Receive()
	if err != nil {
		log.Error(err)
		return "", err
	}

	return token, nil
}
