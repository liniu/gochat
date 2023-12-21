package corp

import (
	"encoding/json"

	"github.com/liniu/gochat/urls"
	"github.com/liniu/gochat/wx"
)

// AuthScope 应用授权作用域
type AuthScope string

// 企业微信支持的应用授权作用域
const (
	ScopeSnsapiBase        AuthScope = "snsapi_base"        // 企业自建应用固定填写：snsapi_base
	ScopeSnsapiPrivateinfo AuthScope = "snsapi_privateinfo" // 获取访问用户敏感信息填写：snsapi_privateinfo
)

// AccessToken 企业微信AccessToken
type AccessToken struct {
	Token     string `json:"access_token"`
	ExpiresIn int64  `json:"expires_in"`
}

type ResultIP struct {
	IPList []string `json:"ip_list"`
}

func GetAPIDomainIP(result *ResultIP) wx.Action {
	return wx.NewGetAction(urls.CorpCgiBinAPIDomainIP,
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

func GetCallbackIP(result *ResultIP) wx.Action {
	return wx.NewGetAction(urls.CorpCgiBinCallbackIP,
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

type ResultOAuthUser struct {
	UserID         string `json:"UserId"`
	OpenID         string `json:"OpenId"`
	DeviceID       string `json:"DeviceId"`
	ExternalUserID string `json:"external_userid"`
	UserTicket     string `json:"user_ticket"`
}

// GetOAuthUser 获取访问用户身份
func GetOAuthUser(code string, result *ResultOAuthUser) wx.Action {
	return wx.NewGetAction(urls.CorpCgiBinUserInfo,
		wx.WithQuery("code", code),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

// UserAuthSucc 二次验证
func UserAuthSucc(userID string) wx.Action {
	return wx.NewGetAction(urls.CorpCgiBinUserAuthSucc,
		wx.WithQuery("userid", userID),
	)
}

// OAuthUserDetail 访问用户敏感信息
type ResultOAuthUserDetail struct {
	UserId  string `json:"userid"`
	Mobile  string `json:"mobile"`
	Gender  string `json:"gender"`
	Email   string `json:"email"`
	Avatar  string `json:"avatar"`
	QrCode  string `json:"qr_code"`
	BizMail string `json:"biz_mail"`
	Address string `json:"address"`
}

// ParamsOAuthUserDetail 获取访问用户敏感信息参数
type ParamsOAuthUserDetail struct {
	UserTicket string `json:"user_ticket"`
}

// GetOAuthUserDetail 获取访问用户敏感信息，需要在后台通讯录配置开启
func GetOAuthUserDetail(userTicket string, result *ResultOAuthUserDetail) wx.Action {
	params := &ParamsOAuthUserDetail{
		UserTicket: userTicket,
	}
	return wx.NewPostAction(urls.CorpCgiBinUserDetail,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}
