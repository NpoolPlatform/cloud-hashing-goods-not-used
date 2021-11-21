package goodextrainfo

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/test-init" //nolint

	"github.com/stretchr/testify/assert"

	"github.com/google/uuid"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

func assertGoodExtraInfo(t *testing.T, actual, expected *npool.GoodExtraInfo) {
	assert.Equal(t, actual.GoodID, expected.GoodID)
	assert.Equal(t, actual.Posters, expected.Posters)
	assert.Equal(t, actual.Labels, expected.Labels)
}

func TestGoodExtraInfoCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	goodExtraInfo := npool.GoodExtraInfo{
		GoodID:  uuid.New().String(),
		Posters: []string{"http://xxxaad.areafdsaf.info/a.img", "http://xxafda.areadfavvvv.info/b.img"},
		Labels:  []string{"aaa", "bbbb"},
	}

	resp, err := Create(context.Background(), &npool.CreateGoodExtraInfoRequest{
		Info: &goodExtraInfo,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{})
		assertGoodExtraInfo(t, resp.Info, &goodExtraInfo)
	}

	goodExtraInfo.Posters = append(goodExtraInfo.Posters, "http://adfafdsa.aresadf.info/c.img")
	goodExtraInfo.Labels = append(goodExtraInfo.Labels, "xxxx")
	goodExtraInfo.ID = resp.Info.ID

	resp1, err := Update(context.Background(), &npool.UpdateGoodExtraInfoRequest{
		Info: &goodExtraInfo,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertGoodExtraInfo(t, resp1.Info, &goodExtraInfo)
	}

	resp2, err := Get(context.Background(), &npool.GetGoodExtraInfoRequest{
		GoodID: goodExtraInfo.GoodID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp2.Info.ID, resp.Info.ID)
		assertGoodExtraInfo(t, resp2.Info, &goodExtraInfo)
	}
}
