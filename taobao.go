package taobaosdk

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"sync"
	"time"
)

type Client struct {
	mu                 sync.Mutex
	url                string
	appKey             string
	secret             string
	notifyVerifyDomain string
	Client             *http.Client
	location           *time.Location
}

type OptionFunc func(c *Client)

func WithTimeLocation(location *time.Location) OptionFunc {
	return func(c *Client) {
		c.location = location
	}
}

func WithHTTPClient(client *http.Client) OptionFunc {
	return func(c *Client) {
		c.Client = client
	}
}

func NewTaobaoClient(url string, appKey string, secret string, opts ...OptionFunc) (client *Client, err error) {

	location, err := time.LoadLocation("Asia/Chongqing")
	if err != nil {
		return nil, err
	}

	client = &Client{}
	client.appKey = appKey
	client.secret = secret
	client.url = url
	client.Client = http.DefaultClient
	//client.Client = createHTTPClient()
	client.location = location

	for _, opt := range opts {
		opt(client)
	}

	return client, nil
}

func (this *Client) URLValues(param Param) (value url.Values, err error) {
	var p = url.Values{}

	p.Add("app_key", this.appKey)
	p.Add("method", param.APIName())
	p.Add("format", Format)
	p.Add("partner_id", PartnerId)
	p.Add("timestamp", time.Now().In(this.location).Format(TimeFormat))
	p.Add("sign_method", SignMethod)
	p.Add("v", Version)

	// 添加业务参数
	var ps = param.Params()
	if ps != nil {
		for key, value := range ps {
			if value != "" {
				p.Add(key, value)
			}
		}
	}

	sign, err := this.sign(p, this.secret)
	if err != nil {
		return nil, err
	}
	p.Add("sign", sign)
	return p, nil
}

//计算签名
func (this *Client) hmacMd5(text, key string) string {
	hashed := hmac.New(md5.New, []byte(key))
	hashed.Write([]byte(text))
	return hex.EncodeToString(hashed.Sum(nil))
}

func (this *Client) sign(args url.Values, secKey string) (string, error) {

	if args == nil {
		args = make(url.Values, 0)
	}

	var keys []string
	for key := range args {
		keys = append(keys, key)
	}

	// 排序
	sort.Strings(keys)

	// 拼接请求串
	var signString string
	for _, key := range keys {
		var value = strings.TrimSpace(args.Get(key))
		//var value = args.Get(key)
		if len(key) != 0 && len(value) > 0 {
			signString += fmt.Sprintf("%s%s", key, args.Get(key))
		}
	}

	// 加密串
	encText := this.hmacMd5(signString, secKey)

	return strings.ToUpper(encText), nil
}

func (this *Client) doRequest(method string, param Param, result interface{}) (err error) {
	var buf io.Reader
	if param != nil {
		p, err := this.URLValues(param)
		if err != nil {
			return err
		}

		idx := 0
		var newArgs string
		for key := range p {
			if idx != 0 {
				newArgs += "&"
			}
			newArgs += fmt.Sprintf("%s=%s", key, url.QueryEscape(p.Get(key)))
			idx++ // 不能作为表达式
		}
		fmt.Println(newArgs)
		buf = bytes.NewReader([]byte(newArgs))
	}

	req, err := http.NewRequest(method, this.url, buf)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", kContentType)

	resp, err := this.Client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, result)
	if err != nil {
		return err
	}

	return err
}

func (this *Client) DoRequest(method string, param Param, result interface{}) (err error) {
	return this.doRequest(method, param, result)
}
