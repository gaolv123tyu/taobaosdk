package taobaosdk

import "fmt"

type Param interface {
	// 用于提供访问的 method
	APIName() string

	// 返回参数列表
	Params() map[string]string
}

type Code int32
func (c Code) IsSuccess() bool {
	return c == CodeSuccess
}

const (
	CodeSuccess          Code = 0 // 接口调用成功
)

type ErrorRsp struct {
	Code          Code   `json:"code"`
	Msg           string `json:"msg"`
	SubCode       string `json:"sub_code"`
	SubMsg        string `json:"sub_msg"`
	RequestId     string `json:"request_id"`
	ErrorResponse string `json:"error_response"`
}

func (this *ErrorRsp) Error() string {
	return fmt.Sprintf("%s - %s", this.Code, this.SubMsg)
}

const (
	Format = "json"

	Charset = "UTF-8"

	Version = "2.0"

	PartnerId = "apidoc"

	TimeFormat = "2006-01-02 15:04:05"

	SignMethod = "hmac"

	kContentType = "application/x-www-form-urlencoded;charset=utf-8"
)
