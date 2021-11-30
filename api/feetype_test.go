package api

import (
	"encoding/json"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func assertFeeTypeEqual(t *testing.T, actual, expected *npool.FeeType) {
	assert.Equal(t, actual.ID, expected.ID)
	assert.Equal(t, actual.FeeType, expected.FeeType)
	assert.Equal(t, actual.FeeDescription, expected.FeeDescription)
	assert.Equal(t, actual.PayType, expected.PayType)
}

func TestFeeTypeCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	cli := resty.New()

	newFeeTypeRequest := &npool.CreateFeeTypeRequest{
		Info: &npool.FeeType{
			ID:             uuid.New().String(),
			FeeType:        "TestFeeType",
			FeeDescription: "created by unit test",
			PayType:        "percent",
		},
	}

	// create
	respFeeTypeResponse := npool.CreateFeeTypeResponse{}
	restyFeeTypeTest(cli, t, "http://localhost:50020/v1/create/fee/type", newFeeTypeRequest, &respFeeTypeResponse)
	newFeeTypeRequest.Info.ID = respFeeTypeResponse.Info.ID
	assertFeeTypeEqual(t, newFeeTypeRequest.Info, respFeeTypeResponse.Info)

	// update
	newFeeTypeRequest.Info.FeeType = "UpdatedFeeType"
	respFeeTypeResponse = npool.CreateFeeTypeResponse{}
	restyFeeTypeTest(cli, t, "http://localhost:50020/v1/update/fee/type", newFeeTypeRequest, &respFeeTypeResponse)
	assertFeeTypeEqual(t, newFeeTypeRequest.Info, respFeeTypeResponse.Info)

	// get
	restyFeeTypeTest(cli, t, "http://localhost:50020/v1/get/fee/type", &npool.GetFeeTypeRequest{
		ID: newFeeTypeRequest.Info.ID,
	}, &respFeeTypeResponse)
	assertFeeTypeEqual(t, newFeeTypeRequest.Info, respFeeTypeResponse.Info)
}

func restyFeeTypeTest(cli *resty.Client, t *testing.T, url string, body interface{ String() string }, respStructPointer interface{}) {
	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(url)
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		err = json.Unmarshal(resp.Body(), respStructPointer)
		assert.Nil(t, err)
		assert.NotNil(t, respStructPointer)
	}
}
