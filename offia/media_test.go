package offia

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shenghui0779/yiigo"
	"github.com/stretchr/testify/assert"

	"github.com/shenghui0779/gochat/wx"
)

func TestUploadMedia(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Upload(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/media/upload?access_token=ACCESS_TOKEN&type=image", gomock.AssignableToTypeOf(yiigo.NewUploadForm())).Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"type": "image",
		"media_id": "MEDIA_ID",
		"created_at": 1606717010
	  }`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	params := &ParamsMediaUpload{
		MediaType: MediaImage,
		Path:      "../test/test.jpg",
	}
	result := new(ResultMediaUpload)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", UploadMedia(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMediaUpload{
		Type:      "image",
		MediaID:   "MEDIA_ID",
		CreatedAt: 1606717010,
	}, result)
}

func TestUploadMediaByURL(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Upload(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/media/upload?access_token=ACCESS_TOKEN&type=image", gomock.AssignableToTypeOf(yiigo.NewUploadForm())).Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"type": "image",
		"media_id": "MEDIA_ID",
		"created_at": 1606717010
	  }`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	params := &ParamsMediaUploadByURL{
		MediaType: MediaImage,
		Filename:  "test.png",
		URL:       "https://golang.google.cn/doc/gopher/pkg.png",
	}
	result := new(ResultMediaUpload)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", UploadMediaByURL(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMediaUpload{
		Type:      "image",
		MediaID:   "MEDIA_ID",
		CreatedAt: 1606717010,
	}, result)
}

func TestAddMaterial(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Upload(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/material/add_material?access_token=ACCESS_TOKEN&type=image", gomock.AssignableToTypeOf(yiigo.NewUploadForm())).Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"media_id": "MEDIA_ID",
		"url": "URL"
	  }`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	params := &ParamsMaterialAdd{
		MediaType: MediaImage,
		Path:      "../test/test.jpg",
	}
	result := new(ResultMaterialAdd)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", AddMaterial(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMaterialAdd{
		MediaID: "MEDIA_ID",
		URL:     "URL",
	}, result)
}

func TestAddMaterialByURL(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Upload(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/material/add_material?access_token=ACCESS_TOKEN&type=image", gomock.AssignableToTypeOf(yiigo.NewUploadForm())).Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"media_id": "MEDIA_ID",
		"url": "URL"
	  }`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	params := &ParamsMaterialAddByURL{
		MediaType: MediaImage,
		Filename:  "test.png",
		URL:       "https://golang.google.cn/doc/gopher/pkg.png",
	}
	result := new(ResultMaterialAdd)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", AddMaterialByURL(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMaterialAdd{
		MediaID: "MEDIA_ID",
		URL:     "URL",
	}, result)
}

func TestDeleteMaterial(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/material/del_material?access_token=ACCESS_TOKEN", []byte(`{"media_id":"MEDIA_ID"}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", DeleteMaterial("MEDIA_ID"))

	assert.Nil(t, err)
}

func TestUploadImage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Upload(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/media/uploadimg?access_token=ACCESS_TOKEN", gomock.AssignableToTypeOf(yiigo.NewUploadForm())).Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"url": "URL"
	  }`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	result := new(ResultMaterialAdd)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", UploadImage("../test/test.jpg", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMaterialAdd{
		URL: "URL",
	}, result)
}

func TestUploadImageByURL(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Upload(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/media/uploadimg?access_token=ACCESS_TOKEN", gomock.AssignableToTypeOf(yiigo.NewUploadForm())).Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"url": "URL"
	  }`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	params := &ParamsImageUploadByURL{
		Filename: "test.png",
		URL:      "https://golang.google.cn/doc/gopher/pkg.png",
	}
	result := new(ResultMaterialAdd)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", UploadImageByURL(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMaterialAdd{
		URL: "URL",
	}, result)
}

func TestUploadVideo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Upload(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/material/add_material?access_token=ACCESS_TOKEN&type=video", gomock.AssignableToTypeOf(yiigo.NewUploadForm())).Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"media_id": "MEDIA_ID",
		"url": "URL"
	  }`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	params := &ParamsVideoUpload{
		Path:        "../test/test.mp4",
		Title:       "TITLE",
		Description: "INTRODUCTION",
	}
	result := new(ResultMaterialAdd)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", UploadVideo(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMaterialAdd{
		MediaID: "MEDIA_ID",
		URL:     "URL",
	}, result)
}

func TestUploadVideoByURL(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Upload(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/material/add_material?access_token=ACCESS_TOKEN&type=video", gomock.AssignableToTypeOf(yiigo.NewUploadForm())).Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"media_id": "MEDIA_ID",
		"url": "URL"
	  }`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	params := &ParamsVideoUploadByURL{
		Filename:    "test.mp4",
		Title:       "TITLE",
		Description: "INTRODUCTION",
		URL:         "https://video.ivwen.com/users/4576112/46e9506e35534ddb961772727f32399d.mp4",
	}
	result := new(ResultMaterialAdd)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", UploadVideoByURL(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMaterialAdd{
		MediaID: "MEDIA_ID",
		URL:     "URL",
	}, result)
}

func TestAddNews(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/material/add_news?access_token=ACCESS_TOKEN", []byte(`{"articles":[{"title":"TITLE","thumb_media_id":"THUMB_MEDIA_ID","author":"AUTHOR","digest":"DIGEST","show_cover_pic":1,"content":"CONTENT","content_source_url":"CONTENT_SOURCE_URL","need_open_comment":1,"only_fans_can_comment":1}]}`)).Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"media_id": "MEDIA_ID"
	  }`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	params := &ParamsNewsAdd{
		Articles: []*NewsArticle{
			{
				Title:              "TITLE",
				ThumbMediaID:       "THUMB_MEDIA_ID",
				Author:             "AUTHOR",
				Digest:             "DIGEST",
				ShowCoverPic:       1,
				Content:            "CONTENT",
				ContentSourceURL:   "CONTENT_SOURCE_URL",
				NeedOpenComment:    1,
				OnlyFansCanComment: 1,
			},
		},
	}
	result := new(ResultMaterialAdd)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", AddNews(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMaterialAdd{
		MediaID: "MEDIA_ID",
	}, result)
}
