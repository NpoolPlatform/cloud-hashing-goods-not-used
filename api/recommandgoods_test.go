package api

import (
	"encoding/json"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"
)

func TestRecommendGoodsCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	cli := resty.New()

	Recommendgood := npool.RecommendGood{
		RecommenderID: uuid.New().String(),
		GoodID:        getGoodID(t, cli),
		Message:       "test recommend good",
	}
	firstCreateInfo := npool.CreateRecommendGoodResponse{}

	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.CreateRecommendGoodRequest{
			Info: &Recommendgood,
		}).
		Post("http://localhost:50020/v1/create/recommend/good")
	if err != nil {
		assert.Equal(t, 200, resp.StatusCode())
		err := json.Unmarshal(resp.Body(), &firstCreateInfo)
		if assert.Nil(t, err) {
			assert.NotEqual(t, firstCreateInfo.Info.ID, uuid.UUID{})
			assert.Equal(t, firstCreateInfo.Info.RecommenderID, Recommendgood.RecommenderID)
			assert.Equal(t, firstCreateInfo.Info.GoodID, Recommendgood.GoodID)
			assert.Equal(t, firstCreateInfo.Info.Message, Recommendgood.Message)
		}
	}

	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetRecommendGoodsRequest{
			PageInfo: &npool.PageInfo{},
		}).
		Post("http://localhost:50020/v1/get/recommend/goods")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		resp3 := &npool.GetRecommendGoodsResponse{}
		err = json.Unmarshal(resp.Body(), resp3)
		if assert.Nil(t, err) {
			assert.NotEmpty(t, resp3.Infos)
			assert.NotEmpty(t, resp3.Recommends)
		}
	}

	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetRecommendGoodsByRecommenderRequest{
			RecommenderID: Recommendgood.RecommenderID,
			PageInfo:      &npool.PageInfo{},
		}).
		Post("http://localhost:50020/v1/get/recommend/goods/by/recommender")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		resp4 := &npool.GetRecommendGoodsByRecommenderResponse{}
		err = json.Unmarshal(resp.Body(), resp4)
		if assert.Nil(t, err) {
			assert.NotEmpty(t, resp4.Recommends)
			assert.NotNil(t, resp4.Infos[0])
		}
	}
}

func getGoodID(t *testing.T, cli *resty.Client) string {
	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetGoodsRequest{
			PageInfo: &npool.PageInfo{
				PageIndex: 0,
				PageSize:  10,
			},
		}).
		Post("http://localhost:50020/v1/get/goods")
	if err == nil {
		assert.Equal(t, 200, resp.StatusCode())
		resp2 := &npool.GetGoodsResponse{}
		err = json.Unmarshal(resp.Body(), resp2)
		if assert.Nil(t, err) {
			assert.NotEmpty(t, resp2.Infos)
			assert.NotNil(t, resp2.Infos[0])
		}
		return resp2.Infos[0].ID
	}
	return createGood(t, cli)
}

func createGood(t *testing.T, cli *resty.Client) string {
	deviceInfoID := uuid.New().String()
	separateFee := true
	unitPower := int32(100)
	duration := int32(180)
	coinInfoID := uuid.New().String()
	actuals := true
	deliveryAt := uint32(time.Now().Unix())
	inheritFromGoodID := uuid.New().String()
	vendorLocationID := uuid.New().String()
	price := 100.8
	benefitType := "pool"
	classic := true
	supportCoinTypeIDs := []string{uuid.New().String(), uuid.New().String()}
	total := int32(1700)

	goodInfo := npool.GoodInfo{
		DeviceInfoID:       deviceInfoID,
		SeparateFee:        separateFee,
		UnitPower:          unitPower,
		DurationDays:       duration,
		CoinInfoID:         coinInfoID,
		DeliveryAt:         deliveryAt,
		Actuals:            actuals,
		InheritFromGoodID:  inheritFromGoodID,
		VendorLocationID:   vendorLocationID,
		Price:              price,
		PriceCurrency:      uuid.New().String(),
		BenefitType:        benefitType,
		Classic:            classic,
		SupportCoinTypeIDs: supportCoinTypeIDs,
		Total:              total,
		Title:              "Ant Miner S19 Pro",
		Unit:               "TH/s",
	}
	firstCreateInfo := npool.CreateGoodResponse{}

	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.CreateGoodRequest{
			Info: &goodInfo,
		}).
		Post("http://localhost:50020/v1/create/good")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		err := json.Unmarshal(resp.Body(), &firstCreateInfo)
		if assert.Nil(t, err) {
			assert.NotEqual(t, firstCreateInfo.Info.ID, uuid.New())
		}
	}
	return firstCreateInfo.Info.ID
}
