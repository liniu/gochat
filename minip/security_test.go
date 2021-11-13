package minip

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shenghui0779/yiigo"
	"github.com/stretchr/testify/assert"

	"github.com/shenghui0779/gochat/wx"
)

func TestImageSecCheck(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Upload(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/wxa/img_sec_check?access_token=ACCESS_TOKEN", gomock.AssignableToTypeOf(yiigo.NewUploadForm())).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", ImageSecCheck("../test/test.jpg"))

	assert.Nil(t, err)
}

func TestMediaCheckAsync(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/wxa/media_check_async?access_token=ACCESS_TOKEN", []byte(`{"media_type":2,"media_url":"https://developers.weixin.qq.com/miniprogram/assets/images/head_global_z_@all.png"}`)).Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"trace_id": "967e945cd8a3e458f3c74dcb886068e9"
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	params := &ParamsMediaCheckAsync{
		MediaType: SecMediaImage,
		MediaURL:  "https://developers.weixin.qq.com/miniprogram/assets/images/head_global_z_@all.png",
	}
	result := new(ResultMediaCheckAsync)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", MediaCheckAsync(params, result))

	assert.Nil(t, err)
	assert.Equal(t, "967e945cd8a3e458f3c74dcb886068e9", result.TraceID)
}

func TestMsgSecCheck(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/wxa/msg_sec_check?access_token=ACCESS_TOKEN", []byte(`{"content":"hello world!"}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", MsgSecCheck("hello world!"))

	assert.Nil(t, err)
}
