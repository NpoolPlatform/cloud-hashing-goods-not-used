package api

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"golang.org/x/xerrors"
)

func TestFeeDurationCRUD(t *testing.T) { // nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	testFeeDurationInfo := &npool.FeeDuration{
		FeeTypeID: uuid.New().String(),
		Duration:  30,
	}

	cli := resty.New()

	newFeeDurationRequest := &npool.CreateFeeDurationRequest{
		Info: testFeeDurationInfo,
	}
	var restyRet *resty.Response

	// create
	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(newFeeDurationRequest).
		Post("http://localhost:50020/v1/create/fee/duration")
	if assert.Nil(t, err) {
		assert.NotNil(t, resp)
		restyRet = resp
		resp1 := &npool.CreateFeeDurationResponse{
			Info: &npool.FeeDuration{},
		}
		err = json.Unmarshal(restyRet.Body(), resp1)
		if assert.Nil(t, err) {
			testFeeDurationInfo.ID = resp1.Info.ID
			assert.Equal(t, testFeeDurationInfo.Duration, resp1.Info.Duration)
			assert.Equal(t, testFeeDurationInfo.FeeTypeID, resp1.Info.FeeTypeID)
			fmt.Printf("created: %v", resp1.Info)
		}
	}

	// update
	testFeeDurationInfo.Duration = 0
	restyRet, err = restyFeeDurationTest(cli, "http://localhost:50020/v1/update/fee/duration", &npool.UpdateFeeDurationRequest{
		Info: testFeeDurationInfo,
	})
	if assert.Nil(t, err) {
		resp2 := &npool.UpdateFeeDurationResponse{
			Info: &npool.FeeDuration{},
		}
		err = json.Unmarshal(restyRet.Body(), resp2)
		if assert.Nil(t, err) {
			assert.Equal(t, testFeeDurationInfo.ID, resp2.Info.ID)
			assert.Equal(t, testFeeDurationInfo.Duration, resp2.Info.Duration)
			assert.Equal(t, testFeeDurationInfo.FeeTypeID, resp2.Info.FeeTypeID)
		}
	}

	// get
	restyRet, err = restyFeeDurationTest(cli, "http://localhost:50020/v1/get/fee/duration", &npool.GetFeeDurationRequest{
		ID: testFeeDurationInfo.ID,
	})
	if assert.Nil(t, err) {
		resp3 := &npool.GetFeeDurationResponse{}
		err = json.Unmarshal(restyRet.Body(), resp3)
		if assert.Nil(t, err) {
			assert.Equal(t, testFeeDurationInfo.ID, resp3.Info.ID)
			assert.Equal(t, testFeeDurationInfo.Duration, resp3.Info.Duration)
			assert.Equal(t, testFeeDurationInfo.FeeTypeID, resp3.Info.FeeTypeID)
		}
	}

	// get by fee type
	restyRet, err = restyFeeDurationTest(cli, "http://localhost:50020/v1/get/fee/durations/feetype", &npool.GetFeeDurationsByFeeTypeRequest{
		FeeTypeID: testFeeDurationInfo.FeeTypeID,
	})
	if assert.Nil(t, err) {
		resp4 := &npool.GetFeeDurationsByFeeTypeResponse{
			Infos: []*npool.FeeDuration{},
		}
		err = json.Unmarshal(restyRet.Body(), resp4)
		if assert.Nil(t, err) {
			assert.NotNil(t, resp4.Infos)
			assert.Positive(t, len(resp4.Infos))
		}
	}

	// delete
	restyRet, err = restyFeeDurationTest(cli, "http://localhost:50020/v1/get/fee/duration", &npool.DeleteFeeDurationRequest{
		ID: testFeeDurationInfo.ID,
	})
	if assert.Nil(t, err) {
		resp5 := &npool.DeleteFeeDurationResponse{}
		err = json.Unmarshal(restyRet.Body(), resp5)
		if assert.Nil(t, err) {
			assert.NotNil(t, resp5.Info)
		}
	}
}

func restyFeeDurationTest(cli *resty.Client, url string, body interface{ String() string }) (resp *resty.Response, err error) {
	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(url)
	if err != nil || resp.StatusCode() != 200 {
		err = xerrors.New("code not 200 - " + url)
	}
	return
}
