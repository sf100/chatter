package app

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
