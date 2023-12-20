package offia

import (
	"encoding/json"

	"github.com/liniu/gochat/urls"
	"github.com/liniu/gochat/wx"
)

// OAuthToken 公众号网页授权Token
type OAuthToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
	OpenID       string `json:"openid"`
	Scope        string `json:"scope"`
}

// AccessToken 公众号普通AccessToken
type AccessToken struct {
	Token     string `json:"access_token"`
	ExpiresIn int64  `json:"expires_in"`
}

// CheckOAuthToken 检验授权凭证（access_token）是否有效
func CheckOAuthToken(openid string) wx.Action {
	return wx.NewGetAction(urls.OffiaSnsCheckAccessToken,
		wx.WithQuery("openid", openid),
	)
}

// ResultOAuthUser 授权用户信息
type ResultOAuthUser struct {
	OpenID     string   `json:"openid"`
	UnionID    string   `json:"unionid"`
	Nickname   string   `json:"nickname"`
	Sex        int      `json:"sex"`
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	HeadImgURL string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
}

// GetOAuthUser 获取授权用户信息（注意：使用网页授权的access_token）
func GetOAuthUser(openid string, result *ResultOAuthUser) wx.Action {
	return wx.NewGetAction(urls.OffiaSnsUserInfo,
		wx.WithQuery("openid", openid),
		wx.WithQuery("lang", "zh_CN"),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

// TicketType JSApi ticket 类型
type TicketType string

// 微信支持的 JSApi ticket
const (
	WXCardTicket TicketType = "wx_card"
	JSAPITicket  TicketType = "jsapi"
)

// JSApiSign JSApi签名
type JSApiSign struct {
	Signature string `json:"signature"`
	NonceStr  string `json:"noncestr"`
	Timestamp int64  `json:"timestamp"`
}

// ResultApiTicket 公众号 api ticket
type ResultApiTicket struct {
	Ticket    string `json:"ticket"`
	ExpiresIn int64  `json:"expires_in"`
}

// GetApiTicket 获取 JSApi ticket (注意：使用普通access_token)
func GetApiTicket(ticketType TicketType, result *ResultApiTicket) wx.Action {
	return wx.NewGetAction(urls.OffiaCgiBinTicket,
		wx.WithQuery("type", string(ticketType)),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}
