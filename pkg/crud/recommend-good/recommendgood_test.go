package recommendgood

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"
	goodinfo "github.com/NpoolPlatform/cloud-hashing-goods/pkg/crud/good-info"
	testinit "github.com/NpoolPlatform/cloud-hashing-goods/pkg/test-init" //nolint

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

func assertRecommendGood(t *testing.T, actual, expected *npool.RecommendGood) {
	assert.Equal(t, actual.RecommenderID, expected.RecommenderID)
	assert.Equal(t, actual.GoodID, expected.GoodID)
	assert.Equal(t, actual.Message, expected.Message)
}

func TestRecommendGoodCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	ctx := context.Background()

	goodInfos, err := goodinfo.GetAll(ctx, &npool.GetGoodsRequest{
		PageInfo: &npool.PageInfo{},
	})
	if assert.Nil(t, err) {
		assert.NotNil(t, goodInfos.Infos)
	}
	var goodID string
	if len(goodInfos.Infos) > 0 {
		goodID = goodInfos.Infos[0].ID
	} else {
		goodID = createGood(ctx, t)
	}

	Recommendgood := npool.RecommendGood{
		RecommenderID: uuid.New().String(),
		GoodID:        goodID,
		Message:       "test recommend good",
	}

	resp, err := Create(context.Background(), &npool.CreateRecommendGoodRequest{
		Info: &Recommendgood,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{})
		assertRecommendGood(t, resp.Info, &Recommendgood)
	}

	Recommendgood.ID = resp.Info.ID
	Recommendgood.Message += "."

	resp1, err := Update(context.Background(), &npool.UpdateRecommendGoodRequest{
		Info: &Recommendgood,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertRecommendGood(t, resp1.Info, &Recommendgood)
	}

	_, err = GetByRecommender(context.Background(), &npool.GetRecommendGoodsByRecommenderRequest{
		RecommenderID: Recommendgood.RecommenderID,
		PageInfo:      &npool.PageInfo{},
	})
	assert.Nil(t, err)

	_, err = Get(context.Background(), &npool.GetRecommendGoodsRequest{
		PageInfo: &npool.PageInfo{},
	})
	assert.Nil(t, err)
}

func createGood(ctx context.Context, t *testing.T) string {
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

	resp, err := goodinfo.Create(ctx, &npool.CreateGoodRequest{
		Info: &goodInfo,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{})
		assert.Equal(t, resp.Info.DeviceInfoID, deviceInfoID)
		assert.Equal(t, resp.Info.SeparateFee, separateFee)
		assert.Equal(t, resp.Info.UnitPower, unitPower)
		assert.Equal(t, resp.Info.DurationDays, duration)
		assert.Equal(t, resp.Info.CoinInfoID, coinInfoID)
		assert.Equal(t, resp.Info.Actuals, actuals)
		assert.Equal(t, resp.Info.InheritFromGoodID, inheritFromGoodID)
		assert.Equal(t, resp.Info.VendorLocationID, vendorLocationID)
		assert.Equal(t, resp.Info.Price, price)
		assert.Equal(t, resp.Info.BenefitType, benefitType)
		assert.Equal(t, resp.Info.Classic, classic)
		assert.Equal(t, resp.Info.SupportCoinTypeIDs, supportCoinTypeIDs)
		assert.Equal(t, resp.Info.Total, total)
		assert.Equal(t, resp.Info.PriceCurrency, goodInfo.PriceCurrency)
		assert.Equal(t, resp.Info.Title, goodInfo.Title)
		assert.Equal(t, resp.Info.Unit, goodInfo.Unit)
	}
	return resp.Info.ID
}
