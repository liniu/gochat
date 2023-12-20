package minip

import (
	"context"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/liniu/gochat/mock"
	"github.com/liniu/gochat/wx"
)

func TestUploadMedia(t *testing.T) {
	resp := []byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"type": "image",
	"media_id": "MEDIA_ID",
	"created_at": 1606717010
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Upload(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/media/upload?access_token=ACCESS_TOKEN&type=image", gomock.AssignableToTypeOf(wx.NewUploadForm())).Return(resp, nil)

	mp := New("APPID", "APPSECRET", WithMockClient(client))

	result := new(ResultMediaUpload)

	err := mp.Do(context.TODO(), "ACCESS_TOKEN", UploadTempMedia(MediaImage, "../mock/test.jpg", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMediaUpload{
		Type:      "image",
		MediaID:   "MEDIA_ID",
		CreatedAt: 1606717010,
	}, result)
}

func TestUploadMediaByURL(t *testing.T) {
	resp := []byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"type": "image",
	"media_id": "MEDIA_ID",
	"created_at": 1606717010
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Upload(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/media/upload?access_token=ACCESS_TOKEN&type=image", gomock.AssignableToTypeOf(wx.NewUploadForm())).Return(resp, nil)

	mp := New("APPID", "APPSECRET", WithMockClient(client))

	result := new(ResultMediaUpload)

	err := mp.Do(context.TODO(), "ACCESS_TOKEN", UploadTempMediaByURL(MediaImage, "test.png", "https://golang.google.cn/doc/gopher/pkg.png", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMediaUpload{
		Type:      "image",
		MediaID:   "MEDIA_ID",
		CreatedAt: 1606717010,
	}, result)
}

func TestGetMedia(t *testing.T) {
	resp := []byte("BUFFER")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://api.weixin.qq.com/cgi-bin/media/get?access_token=ACCESS_TOKEN&media_id=MEDIA_ID", nil).Return(resp, nil)

	mp := New("APPID", "APPSECRET", WithMockClient(client))

	result := new(Media)

	err := mp.Do(context.TODO(), "ACCESS_TOKEN", GetTempMedia("MEDIA_ID", result))

	assert.Nil(t, err)
	assert.Equal(t, "BUFFER", string(result.Buffer))
}
