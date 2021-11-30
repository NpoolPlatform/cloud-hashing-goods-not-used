package api

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func assertFeeTypeEqual(t *testing.T, actual, expected *npool.FeeType) {
	assert.NotNil(t, actual)
	assert.NotNil(t, expected)
	if actual.ID != "" && expected.ID != "" {
		assert.Equal(t, actual.ID, expected.ID)
	}
	assert.Equal(t, actual.FeeType, expected.FeeType)
	assert.Equal(t, actual.FeeDescription, expected.FeeDescription)
	assert.Equal(t, actual.PayType, expected.PayType)
}

func TestFeeTypeCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	cli := resty.New()

	testFeeType := &npool.FeeType{
		ID:             uuid.New().String(),
		FeeType:        fmt.Sprintf("TestFeeType - %v", time.Now().UTC().Unix()),
		FeeDescription: "created by unit test",
		PayType:        "percent",
	}

	newFeeTypeRequest := &npool.CreateFeeTypeRequest{
		Info: testFeeType,
	}

	// create
	resp1 := &npool.CreateFeeTypeResponse{
		Info: &npool.FeeType{},
	}
	err := restyFeeTypeTest(cli, "http://localhost:50020/v1/create/fee/type", newFeeTypeRequest, resp1)
	assert.Nil(t, err)
	testFeeType.ID = resp1.Info.ID
	assertFeeTypeEqual(t, testFeeType, resp1.Info)

	// update
	testFeeType.FeeType = fmt.Sprintf("UpdatedFeeType - %v", time.Now().UTC().Unix())
	resp2 := &npool.CreateFeeTypeResponse{
		Info: &npool.FeeType{},
	}
	err = restyFeeTypeTest(cli, "http://localhost:50020/v1/update/fee/type", &npool.UpdateFeeTypeRequest{
		Info: testFeeType,
	}, resp2)
	assert.Nil(t, err)
	assertFeeTypeEqual(t, testFeeType, resp2.Info)

	// get
	resp3 := &npool.GetFeeTypeResponse{
		Info: &npool.FeeType{},
	}
	err = restyFeeTypeTest(cli, "http://localhost:50020/v1/get/fee/type", &npool.GetFeeTypeRequest{
		ID: testFeeType.ID,
	}, resp3)
	assert.Nil(t, err)
	assertFeeTypeEqual(t, testFeeType, resp3.Info)
}

func restyFeeTypeTest(cli *resty.Client, url string, body interface{ String() string }, respStructPointer interface{}) (err error) {
	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(url)
	if err != nil || resp.StatusCode() != 200 {
		return
	}
	err = json.Unmarshal(resp.Body(), respStructPointer)
	return
}
