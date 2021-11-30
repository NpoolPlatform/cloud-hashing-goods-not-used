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

func assertFeeDurationEqual(t *testing.T, actual, expected *npool.FeeDuration) {
	assert.Equal(t, actual.ID, expected.ID)
	assert.Equal(t, actual.Duration, expected.Duration)
	assert.Equal(t, actual.FeeTypeID, expected.FeeTypeID)
}

func TestFeeDurationCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	cli := resty.New()

	newFeeDurationRequest := &npool.CreateFeeDurationRequest{
		Info: &npool.FeeDuration{
			ID:        uuid.New().String(),
			FeeTypeID: uuid.New().String(),
			Duration:  30,
		},
	}

	// create
	respFeeDurationResponse := npool.CreateFeeDurationResponse{}
	restyFeeDurationTest(cli, t, "http://localhost:50020/v1/create/fee/duration", newFeeDurationRequest, &respFeeDurationResponse)
	newFeeDurationRequest.Info.ID = respFeeDurationResponse.Info.ID
	assert.Equal(t, newFeeDurationRequest.Info.Duration, respFeeDurationResponse.Info.Duration)
	assert.Equal(t, newFeeDurationRequest.Info.FeeTypeID, respFeeDurationResponse.Info.FeeTypeID)

	// update
	respFeeDurationResponse.Info.Duration = 0
	resp2 := npool.UpdateFeeDurationResponse{}
	restyFeeDurationTest(cli, t, "http://localhost:50020/v1/update/fee/duration", &npool.UpdateFeeDurationRequest{
		Info: respFeeDurationResponse.Info,
	}, &resp2)
	assertFeeDurationEqual(t, respFeeDurationResponse.Info, resp2.Info)

	// get
	resp3 := npool.GetFeeDurationResponse{}
	restyFeeDurationTest(cli, t, "http://localhost:50020/v1/get/fee/duration", &npool.GetFeeDurationRequest{
		ID: newFeeDurationRequest.Info.ID,
	}, &resp3)
	assertFeeDurationEqual(t, newFeeDurationRequest.Info, resp3.Info)

	// get by fee type
	resp4 := npool.GetFeeDurationsByFeeTypeResponse{
		Infos: []*npool.FeeDuration{},
	}
	restyFeeDurationTest(cli, t, "http://localhost:50020/v1/get/fee/duration", &npool.GetFeeDurationsByFeeTypeRequest{
		FeeTypeID: newFeeDurationRequest.Info.FeeTypeID,
	}, &resp4)
	assert.NotNil(t, resp4.Infos)
	assert.Positive(t, len(resp4.Infos))

	// delete
	resp5 := npool.DeleteFeeDurationResponse{}
	restyFeeDurationTest(cli, t, "http://localhost:50020/v1/get/fee/duration", &npool.DeleteFeeDurationRequest{
		ID: newFeeDurationRequest.Info.ID,
	}, &resp5)
	assert.NotNil(t, resp5.Info)
}

func restyFeeDurationTest(cli *resty.Client, t *testing.T, url string, body interface{ String() string }, respStructPointer interface{}) {
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
