package taobaosdk

import "encoding/json"

// https://open.taobao.com/api.htm?spm=a219a.7386797.0.0.24f6669aDO9AKk&source=search&docId=48589&docType=2

type TaobaoUsergrowthDhhDeliveryAskRequest struct {
	Profile            string `json:"profile"`
	OaidMd5            string `json:"oaid_md5"`
	IdfaMd5            string `json:"idfa_md5"`
	ImeiMd5            string `json:"imei_md5"`
	Oaid               string `json:"oaid"`
	Idfa               string `json:"idfa"`
	Imei               string `json:"imei"`
	Os                 string `json:"os"`
	Channel            string `json:"channel"`
	AdvertisingSpaceID string `json:"advertising_space_id"`
}

func (this *TaobaoUsergrowthDhhDeliveryAskRequest) GetBody() (str string) {
	b, _ := json.Marshal(this)
	return string(b)
}

func (this TaobaoUsergrowthDhhDeliveryAskRequest) APIName() string {
	return "taobao.usergrowth.dhh.delivery.ask"
}

func (this TaobaoUsergrowthDhhDeliveryAskRequest) Params() map[string]string {
	m := make(map[string]string)
	m["profile"] = this.Profile
	m["oaid_md5"] = this.OaidMd5
	m["idfa_md5"] = this.IdfaMd5
	m["imei_md5"] = this.ImeiMd5
	m["oaid"] = this.Oaid
	m["idfa"] = this.Idfa
	m["imei"] = this.ImeiMd5
	m["os"] = this.Os
	m["channel"] = this.Channel
	m["advertising_space_id"] = this.AdvertisingSpaceID
	return m
}

type TaobaoUsergrowthDhhDeliveryAskResponse struct {
	ErrorResponse struct {
		SubMsg  string `json:"sub_msg"`
		Code    Code   `json:"code"`
		SubCode string `json:"sub_code"`
		Msg     string `json:"msg"`
	} `json:"error_response"`
	UsergrowthDhhDeliveryAskResponse struct {
		Result  bool   `json:"result"`
		TaskID  string `json:"task_id"`
		Errcode int    `json:"errcode"`
	} `json:"usergrowth_dhh_delivery_ask_response"`
}

func (this *TaobaoUsergrowthDhhDeliveryAskResponse) IsSuccess() bool {
	if this.ErrorResponse.Code == CodeSuccess {
		return true
	}
	return false
}

func (this *TaobaoUsergrowthDhhDeliveryAskResponse) GetBody() (str string) {
	b, _ := json.Marshal(this)
	return string(b)
}

func (this *TaobaoUsergrowthDhhDeliveryAskResponse) ToJson() (str string) {
	b, _ := json.Marshal(this.UsergrowthDhhDeliveryAskResponse)
	return string(b)
}

func (this *Client) TaobaoUsergrowthDhhDeliveryAskQuery(param TaobaoUsergrowthDhhDeliveryAskRequest) (result *TaobaoUsergrowthDhhDeliveryAskResponse, err error) {
	err = this.DoRequest("POST", param, &result)
	return result, err
}