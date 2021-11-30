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
	"golang.org/x/xerrors"
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
	err := restyFeeTest(cli, "http://localhost:50020/v1/create/fee", newFeeRequest, &respFeeResponse)
	newFeeRequest.Info.ID = respFeeResponse.Info.ID
	if err != nil {
		logger.Sugar().Warn(err.Error())
	}

	// get
	err = restyFeeTest(cli, "http://localhost:50020/v1/get/fee", &npool.GetFeeRequest{
		ID: newFeeRequest.Info.ID,
	}, &respFeeResponse)
	if err != nil {
		logger.Sugar().Warn(err.Error())
	}
}

func restyFeeTest(cli *resty.Client, url string, body interface{ String() string }, respStructPointer interface{}) (err error) {
	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(url)
	if err != nil || resp.StatusCode() != 200 {
		err = xerrors.New("code not 200 - " + url)
		return
	}
	err = json.Unmarshal(resp.Body(), respStructPointer)
	return
}
