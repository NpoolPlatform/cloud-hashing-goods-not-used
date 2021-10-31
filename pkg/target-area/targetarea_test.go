package targetarea

import (
	"context"
	"fmt"
	"os"
	"path"
	"runtime"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"
	db "github.com/NpoolPlatform/cloud-hashing-goods/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/service-name"

	"github.com/NpoolPlatform/go-service-framework/pkg/app"

	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	_, myPath, _, ok := runtime.Caller(0)
	if !ok {
		return
	}

	configPath := fmt.Sprintf("%s/../../cmd/cloud-hashing-goods", path.Dir(myPath))

	app.Init(servicename.ServiceName, "", "", "", configPath, nil, nil)
	db.Init()
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	resp, err := Create(context.Background(), &npool.CreateTargetAreaRequest{
		Info: &npool.TargetAreaInfo{
			Continent: "Asia",
			Country:   "China",
		},
	})
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		if assert.NotNil(t, resp.Info) {
			assert.Equal(t, resp.Info.Continent, "Asia")
			assert.Equal(t, resp.Info.Country, "China")
		}
	}

	resp1, err := Update(context.Background(), &npool.UpdateTargetAreaRequest{
		Info: resp.Info,
	})
	assert.Nil(t, err)
	if assert.NotNil(t, resp1) {
		if assert.NotNil(t, resp1.Info) {
			assert.Equal(t, resp1.Info.ID, resp.Info.ID)
			assert.Equal(t, resp1.Info.Continent, "Asia")
			assert.Equal(t, resp1.Info.Country, "China")
		}
	}
}

func TestGetAll(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	resp, err := GetAll(context.Background(), &npool.GetTargetAreasRequest{})
	assert.NotNil(t, resp)
	assert.Nil(t, err)
}
