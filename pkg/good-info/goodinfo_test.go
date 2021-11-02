package goodinfo

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

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

func TestGoodInfoCRUD(t *testing.T) { //nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	deviceInfoID := uuid.New().String()
	gasPrice := int64(0.13 * 1000000000)
	separateGasFee := true
	unitPower := int32(100)
	duration := int32(180)
	coinInfoID := uuid.New().String()
	actuals := true
	deliveryAt := int32(time.Now().UnixNano() / 1000000)
	inheritFromGoodID := uuid.New().String()
	vendorLocationID := uuid.New().String()
	price := int64(100.8 * 1000000000)
	benefitType := "pool"
	classic := true
	supportCoinTypeIDs := []string{uuid.New().String(), uuid.New().String()}
	total := int32(1700)

	goodInfo := npool.GoodInfo{
		DeviceInfoID:       deviceInfoID,
		GasPrice:           gasPrice,
		SeparateGasFee:     separateGasFee,
		UnitPower:          unitPower,
		DurationDays:       duration,
		CoinInfoID:         coinInfoID,
		DeliveryAt:         deliveryAt,
		Actuals:            actuals,
		InheritFromGoodID:  inheritFromGoodID,
		VendorLocationID:   vendorLocationID,
		Price:              price,
		BenefitType:        benefitType,
		Classic:            classic,
		SupportCoinTypeIDs: supportCoinTypeIDs,
		Total:              total,
	}

	resp, err := Create(context.Background(), &npool.CreateGoodRequest{
		Info: &goodInfo,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{})
		assert.Equal(t, resp.Info.DeviceInfoID, deviceInfoID)
		assert.Equal(t, resp.Info.GasPrice, gasPrice)
		assert.Equal(t, resp.Info.SeparateGasFee, separateGasFee)
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
	}

	goodInfo.BenefitType = "platform"
	goodInfo.ID = resp.Info.ID

	resp1, err := Update(context.Background(), &npool.UpdateGoodRequest{
		Info: &goodInfo,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assert.Equal(t, resp1.Info.DeviceInfoID, deviceInfoID)
		assert.Equal(t, resp1.Info.GasPrice, gasPrice)
		assert.Equal(t, resp1.Info.SeparateGasFee, separateGasFee)
		assert.Equal(t, resp1.Info.UnitPower, unitPower)
		assert.Equal(t, resp1.Info.DurationDays, duration)
		assert.Equal(t, resp1.Info.CoinInfoID, coinInfoID)
		assert.Equal(t, resp1.Info.Actuals, actuals)
		assert.Equal(t, resp1.Info.InheritFromGoodID, inheritFromGoodID)
		assert.Equal(t, resp1.Info.VendorLocationID, vendorLocationID)
		assert.Equal(t, resp1.Info.Price, price)
		assert.Equal(t, resp1.Info.BenefitType, goodInfo.BenefitType)
		assert.Equal(t, resp1.Info.Classic, classic)
		assert.Equal(t, resp1.Info.SupportCoinTypeIDs, supportCoinTypeIDs)
		assert.Equal(t, resp1.Info.Total, total)
	}
}

func TestGetAll(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	_, err := GetAll(context.Background(), &npool.GetGoodsRequest{})
	assert.Nil(t, err)
}
