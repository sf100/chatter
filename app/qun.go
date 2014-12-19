package app

import (
	log "code.google.com/p/log4go"
	"github.com/sf100/chatter/db"
	"time"
)

const (
	SELECT_QUN_BY_ID     = "select id , name,type_id,creator_id,liveness,description,max_member,avatar,created,updated from qun where id =?"
	SELECT_QUN_MEMBERIDS = "select   user_id  from  qun_user where qun_id = ?"
)

type Qun struct {
	Id          string
	Name        string
	TypeId      string
	CreatorId   string
	Liveness    string
	Description string
	MaxMember   string
	Avatar      string
	Created     time.Time
	Updated     time.Time
}

//获取群信息
func GetQunById(id string) *Qun {
	row := db.MySQL.QueryRow(SELECT_QUN_BY_ID, id)
	qun := &Qun{}
	if err := row.Scan(qun.Id, qun.Name, qun.TypeId, qun.CreatorId, qun.Liveness, qun.Description, qun.MaxMember, qun.Avatar, qun.Created, qun.Updated); err != nil {
		log.Error(err)
		return nil
	}
	return qun
}

//获取群成员
func GetQunMemberIDs(id string) []string {

	rows, err := db.MySQL.Query(SELECT_QUN_MEMBERIDS, id)
	if err != nil {
		log.Error(err)
		return nil
	}
	if rows != nil {
		defer rows.Close()
	}

	userIds := []string{}
	for rows.Next() {
		var userId string
		if err := rows.Scan(&userId); err != nil {
			log.Error(err)
			return nil
		}
		userIds = append(userIds, userId)
	}
	if err := rows.Err(); err != nil {
		log.Error(err)
		return nil
	}
	return userIds
}
