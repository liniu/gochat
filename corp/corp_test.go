package corp

import (
	"context"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/liniu/gochat/mock"
	"github.com/stretchr/testify/assert"
)

func TestOAuth2URL(t *testing.T) {
	cp := New("CORPID")

	authURL := cp.OAuth2URL(ScopeSnsapiBase, "REDIRECT_URI", "STATE", "AGENTID")

	assert.Equal(t, "https://open.weixin.qq.com/connect/oauth2/authorize?appid=CORPID&redirect_uri=REDIRECT_URI&response_type=code&scope=snsapi_base&state=STATE&agentid=AGENTID#wechat_redirect", authURL)
}

func TestQRCodeAuthURL(t *testing.T) {
	cp := New("CORPID")

	authURL := cp.QRCodeAuthURL("AGENTID", "REDIRECT_URI", "STATE")

	assert.Equal(t, "https://open.work.weixin.qq.com/wwopen/sso/qrConnect?appid=CORPID&agentid=AGENTID&redirect_uri=REDIRECT_URI&state=STATE", authURL)
}

func TestAccessToken(t *testing.T) {
	resp := []byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"access_token": "accesstoken000001",
	"expires_in": 7200
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=CORPID&corpsecret=SECRET", nil).Return(resp, nil)

	cp := New("CORPID", WithMockClient(client))

	accessToken, err := cp.AccessToken(context.TODO(), "SECRET")

	assert.Nil(t, err)
	assert.Equal(t, &AccessToken{
		Token:     "accesstoken000001",
		ExpiresIn: 7200,
	}, accessToken)
}
