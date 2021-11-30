package api

import (
	"encoding/json"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestFeeCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	cli := resty.New()

	newFeeRequest := &npool.CreateFeeRequest{
		Info: &npool.Fee{
			AppID:     uuid.New().String(),
			FeeTypeID: uuid.New().String(),
			Value:     1.23,
		},
	}

	// create
	respFeeResponse := npool.CreateFeeResponse{}
	restyFeeTest(cli, t, "http://localhost:50020/v1/create/fee", newFeeRequest, &respFeeResponse)
	newFeeRequest.Info.ID = respFeeResponse.Info.ID

	// get
	restyFeeTest(cli, t, "http://localhost:50020/v1/get/fee", &npool.GetFeeRequest{
		ID: newFeeRequest.Info.ID,
	}, &respFeeResponse)
}

func restyFeeTest(cli *resty.Client, t *testing.T, url string, body interface{ String() string }, respStructPointer interface{}) {
	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(url)
	logger.Sugar().Infof("[resty] %v - %v", url, resp.String())
	if assert.Nil(t, err) {
		if resp.StatusCode() == 200 {
			err = json.Unmarshal(resp.Body(), respStructPointer)
			assert.Nil(t, err)
			assert.NotNil(t, respStructPointer)
		} else {
			logger.Sugar().Errorf("returned 500 %v %v", err, resp.String())
			assert.Equal(t, 200, resp.StatusCode())
		}
	}
}
