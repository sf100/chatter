package controllers

const (
	USER_SUFFIX      = "@user"
	QUN_SUFFIX       = "@qun"
	APP_SUFFIX       = "@app"
	SOURCE_TYPE_USER = 1
	SOURCE_TYPE_QUN  = 2
	SOURCE_TYPE_APP  = 3
)

type Result struct {
	Success bool
	Msg     string
	Data    interface{}
}
