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
	"golang.org/x/xerrors"
)

func TestFeeDurationCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	testFeeDurationInfo := &npool.FeeDuration{
		ID:        uuid.New().String(),
		FeeTypeID: uuid.New().String(),
		Duration:  30,
	}

	cli := resty.New()

	newFeeDurationRequest := &npool.CreateFeeDurationRequest{
		Info: testFeeDurationInfo,
	}

	// create
	resp1 := &npool.CreateFeeDurationResponse{
		Info: &npool.FeeDuration{},
	}
	err := restyFeeDurationTest(cli, "http://localhost:50020/v1/create/fee/duration", newFeeDurationRequest, resp1)
	assert.Nil(t, err)
	testFeeDurationInfo.ID = resp1.Info.ID
	assert.Equal(t, testFeeDurationInfo.Duration, resp1.Info.Duration)
	assert.Equal(t, testFeeDurationInfo.FeeTypeID, resp1.Info.FeeTypeID)

	// update
	testFeeDurationInfo.Duration = 0
	resp2 := &npool.UpdateFeeDurationResponse{
		Info: &npool.FeeDuration{},
	}
	err = restyFeeDurationTest(cli, "http://localhost:50020/v1/update/fee/duration", &npool.UpdateFeeDurationRequest{
		Info: testFeeDurationInfo,
	}, resp2)
	assert.Nil(t, err)
	assert.Equal(t, testFeeDurationInfo.ID, resp2.Info.ID)
	assert.Equal(t, testFeeDurationInfo.Duration, resp2.Info.Duration)
	assert.Equal(t, testFeeDurationInfo.FeeTypeID, resp2.Info.FeeTypeID)

	// get
	resp3 := &npool.GetFeeDurationResponse{}
	err = restyFeeDurationTest(cli, "http://localhost:50020/v1/get/fee/duration", &npool.GetFeeDurationRequest{
		ID: testFeeDurationInfo.ID,
	}, resp3)
	assert.Nil(t, err)
	assert.Equal(t, testFeeDurationInfo.ID, resp3.Info.ID)
	assert.Equal(t, testFeeDurationInfo.Duration, resp3.Info.Duration)
	assert.Equal(t, testFeeDurationInfo.FeeTypeID, resp3.Info.FeeTypeID)

	// get by fee type
	resp4 := &npool.GetFeeDurationsByFeeTypeResponse{
		Infos: []*npool.FeeDuration{},
	}
	err = restyFeeDurationTest(cli, "http://localhost:50020/v1/get/fee/duration", &npool.GetFeeDurationsByFeeTypeRequest{
		FeeTypeID: testFeeDurationInfo.FeeTypeID,
	}, resp4)
	assert.Nil(t, err)
	assert.NotNil(t, resp4.Infos)
	assert.Positive(t, len(resp4.Infos))

	// delete
	resp5 := &npool.DeleteFeeDurationResponse{}
	err = restyFeeDurationTest(cli, "http://localhost:50020/v1/get/fee/duration", &npool.DeleteFeeDurationRequest{
		ID: testFeeDurationInfo.ID,
	}, resp5)
	assert.Nil(t, err)
	assert.NotNil(t, resp5.Info)
}

func restyFeeDurationTest(cli *resty.Client, url string, body interface{ String() string }, respStructPointer interface{}) (err error) {
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
