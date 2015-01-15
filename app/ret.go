package app

import (
	log "code.google.com/p/log4go"
	"encoding/json"
	"net/http"
	"time"
)

const (
	OK               = 0
	NotFoundServer   = 1001
	NotFound         = 65531
	TooLong          = 65532
	AuthErr          = 65533
	ParamErr         = 65534
	InternalErr      = 65535
	OverQuotaApicall = 65536
	OverQuotaPush    = 65537
)

// 响应基础结构.
type baseResponse struct {
	Ret    int    `json:"ret"`
	ErrMsg string `json:"errMsg"`
}

// 带 Content-Type=application/json 头 写 JSON 数据，为了保持兼容性，所以新加了这个函数.
func RetPWriteJSON(w http.ResponseWriter, r *http.Request, res map[string]interface{}, body *string, start time.Time) {
	w.Header().Set("Content-Type", "application/json")

	RetPWrite(w, r, res, body, start)
}

func RetPWrite(w http.ResponseWriter, r *http.Request, res map[string]interface{}, body *string, start time.Time) {
	data, err := json.Marshal(res)
	if err != nil {
		log.Error("json.Marshal(\"%v\") error(%v)", res, err)
		return
	}
	dataStr := string(data)
	if n, err := w.Write([]byte(dataStr)); err != nil {
		log.Error("w.Write(\"%s\") error(%v)", dataStr, err)
	} else {
		log.Trace("w.Write(\"%s\") write %d bytes", dataStr, n)
	}

	log.Trace("req: \"%s\", post: \"%s\", res:\"%s\", ip:\"%s\", time:\"%fs\"", r.URL.String(), *body, dataStr, r.RemoteAddr, time.Now().Sub(start).Seconds())
}
