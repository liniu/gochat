package oa

import (
	"context"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/shenghui0779/gochat/wx"
	"github.com/stretchr/testify/assert"
)

func TestCreateTempQRCode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=ACCESS_TOKEN", gomock.AssignableToTypeOf(postBody)).Return([]byte(`{
		"ticket": "gQH47joAAAAAAAAAASxodHRwOi8vd2VpeGluLnFxLmNvbS9xL2taZ2Z3TVRtNzJXV1Brb3ZhYmJJAAIEZ23sUwMEmm3sUw==",
		"expire_seconds": 60,
		"url": "http://weixin.qq.com/q/kZgfwMTm72WWPkovabbI"
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	dest := new(QRCode)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", CreateTempQRCode(dest, 123, 60))

	assert.Nil(t, err)
	assert.Equal(t, &QRCode{
		Ticket:        "gQH47joAAAAAAAAAASxodHRwOi8vd2VpeGluLnFxLmNvbS9xL2taZ2Z3TVRtNzJXV1Brb3ZhYmJJAAIEZ23sUwMEmm3sUw==",
		ExpireSeconds: 60,
		URL:           "http://weixin.qq.com/q/kZgfwMTm72WWPkovabbI",
	}, dest)
}

func TestCreatePermQRCode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=ACCESS_TOKEN", gomock.AssignableToTypeOf(postBody)).Return([]byte(`{
		"ticket": "gQH47joAAAAAAAAAASxodHRwOi8vd2VpeGluLnFxLmNvbS9xL2taZ2Z3TVRtNzJXV1Brb3ZhYmJJAAIEZ23sUwMEmm3sUw==",
		"expire_seconds": 60,
		"url": "http://weixin.qq.com/q/kZgfwMTm72WWPkovabbI"
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	dest := new(QRCode)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", CreatePermQRCode(dest, 123, 60))

	assert.Nil(t, err)
	assert.Equal(t, &QRCode{
		Ticket:        "gQH47joAAAAAAAAAASxodHRwOi8vd2VpeGluLnFxLmNvbS9xL2taZ2Z3TVRtNzJXV1Brb3ZhYmJJAAIEZ23sUwMEmm3sUw==",
		ExpireSeconds: 60,
		URL:           "http://weixin.qq.com/q/kZgfwMTm72WWPkovabbI",
	}, dest)
}
