package report

import (
	"context"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/liniu/gochat/corp"
	"github.com/liniu/gochat/mock"
)

func TestAddGridCata(t *testing.T) {
	body := []byte(`{"category_name":"category_name","level":2,"parent_category_id":"parent_category_id"}`)
	resp := []byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"category_id": "category_id"
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/report/grid/add_cata?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	params := &ParamsGridCataAdd{
		CategoryName:     "category_name",
		Level:            2,
		ParentCategoryID: "parent_category_id",
	}
	result := new(ResultGridCataAdd)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", AddGridCata(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultGridCataAdd{
		CategoryID: "category_id",
	}, result)
}

func TestUpdateGridCata(t *testing.T) {
	body := []byte(`{"category_id":"category_id","category_name":"category_name","level":2,"parent_category_id":"parent_category_id"}`)
	resp := []byte(`{
	"errcode": 0,
	"errmsg": "ok"
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/report/grid/update_cata?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	params := &ParamsGridCataUpdate{
		CategoryID:       "category_id",
		CategoryName:     "category_name",
		Level:            2,
		ParentCategoryID: "parent_category_id",
	}

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", UpdateGridCata(params))

	assert.Nil(t, err)
}

func TestDeleteGridCata(t *testing.T) {
	body := []byte(`{"category_id":"category_id"}`)
	resp := []byte(`{
	"errcode": 0,
	"errmsg": "ok"
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/report/grid/delete_cata?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", DeleteGridCata("category_id"))

	assert.Nil(t, err)
}

func TestListGridCata(t *testing.T) {
	resp := []byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"category_list": [
		{
			"category_id": "category_id",
			"category_name": "2222",
			"level": 1
		},
		{
			"category_id": "category_id",
			"category_name": "2222",
			"level": 2,
			"parent_category_id": "parent_category_id"
		}
	]
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/report/grid/list_cata?access_token=ACCESS_TOKEN", nil).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	result := new(ResultGridCataList)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", ListGridCata(result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultGridCataList{
		CategoryList: []*GridCategory{
			{
				CategoryID:   "category_id",
				CategoryName: "2222",
				Level:        1,
			},
			{
				CategoryID:       "category_id",
				CategoryName:     "2222",
				Level:            2,
				ParentCategoryID: "parent_category_id",
			},
		},
	}, result)
}
