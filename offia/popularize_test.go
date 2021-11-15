package offia

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/shenghui0779/gochat/wx"
)

func TestCreateQRCode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=ACCESS_TOKEN", []byte(`{"action_info":{"scene":{"scene_id":123}},"action_name":"QR_SCENE","expire_seconds":60}`)).Return([]byte(`{
		"ticket": "gQH47joAAAAAAAAAASxodHRwOi8vd2VpeGluLnFxLmNvbS9xL2taZ2Z3TVRtNzJXV1Brb3ZhYmJJAAIEZ23sUwMEmm3sUw==",
		"expire_seconds": 60,
		"url": "http://weixin.qq.com/q/kZgfwMTm72WWPkovabbI"
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	params := &ParamsQRCodeCreate{
		ActionName: QRScene,
		ActionInfo: &QRCodeActionInfo{
			Scene: &QRCodeScene{
				SceneID: "123",
			},
		},
		ExpireSeconds: 60,
	}
	result := new(ResultQRCodeCreate)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", CreateQRCode(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultQRCodeCreate{
		Ticket:        "gQH47joAAAAAAAAAASxodHRwOi8vd2VpeGluLnFxLmNvbS9xL2taZ2Z3TVRtNzJXV1Brb3ZhYmJJAAIEZ23sUwMEmm3sUw==",
		ExpireSeconds: 60,
		URL:           "http://weixin.qq.com/q/kZgfwMTm72WWPkovabbI",
	}, result)
}

func TestShortURL(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/shorturl?access_token=ACCESS_TOKEN", []byte(`{"action":"long2short","long_url":"http://wap.koudaitong.com/v2/showcase/goods?alias=128wi9shh&spm=h56083&redirect_count=1"}`)).Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"short_url": "http:\/\/w.url.cn\/s\/AvCo6Ih"
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	result := new(ResultShortURL)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", ShortURL("http://wap.koudaitong.com/v2/showcase/goods?alias=128wi9shh&spm=h56083&redirect_count=1", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultShortURL{
		ShortURL: "http://w.url.cn/s/AvCo6Ih",
	}, result)
}
