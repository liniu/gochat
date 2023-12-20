package jsapi

import (
	"fmt"
	"time"

	"github.com/liniu/gochat/wx"
)

type ResultSign struct {
	NonceStr  string `json:"noncestr"`
	Timestamp int64  `json:"timestamp"`
	Signature string `json:"signature"`
}

func Sign(ticket, url string) *ResultSign {
	ret := &ResultSign{
		NonceStr:  wx.Nonce(16),
		Timestamp: time.Now().Unix(),
	}

	signStr := fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%d&url=%s", ticket, ret.NonceStr, ret.Timestamp, url)
	ret.Signature = wx.SHA1(signStr)

	return ret
}
