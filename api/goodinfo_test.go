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

func TestGoodCRUD(t *testing.T) { //nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	deviceInfoID := uuid.New().String()
	gasPrice := 0.13
	separateGasFee := true
	unitPower := int32(100)
	duration := int32(180)
	coinInfoID := uuid.New().String()
	actuals := true
	deliveryAt := int32(time.Now().UnixNano() / 1000000)
	inheritFromGoodID := uuid.New().String()
	vendorLocationID := uuid.New().String()
	price := 100.8
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
	firstCreateInfo := npool.CreateGoodResponse{}

	cli := resty.New()

	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.CreateGoodRequest{
			Info: &goodInfo,
		}).
		Post("http://localhost:33759/v1/create/good")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		err := json.Unmarshal(resp.Body(), &firstCreateInfo)
		if assert.Nil(t, err) {
			assert.NotEqual(t, firstCreateInfo.Info.ID, uuid.New())
			assert.Equal(t, firstCreateInfo.Info.DeviceInfoID, goodInfo.DeviceInfoID)
			assert.Equal(t, firstCreateInfo.Info.GasPrice, goodInfo.GasPrice)
			assert.Equal(t, firstCreateInfo.Info.SeparateGasFee, goodInfo.SeparateGasFee)
			assert.Equal(t, firstCreateInfo.Info.UnitPower, goodInfo.UnitPower)
			assert.Equal(t, firstCreateInfo.Info.DurationDays, goodInfo.DurationDays)
			assert.Equal(t, firstCreateInfo.Info.CoinInfoID, goodInfo.CoinInfoID)
			assert.Equal(t, firstCreateInfo.Info.DeliveryAt, goodInfo.DeliveryAt)
			assert.Equal(t, firstCreateInfo.Info.Actuals, goodInfo.Actuals)
			assert.Equal(t, firstCreateInfo.Info.InheritFromGoodID, goodInfo.InheritFromGoodID)
			assert.Equal(t, firstCreateInfo.Info.VendorLocationID, goodInfo.VendorLocationID)
			assert.Equal(t, firstCreateInfo.Info.Price, goodInfo.Price)
			assert.Equal(t, firstCreateInfo.Info.BenefitType, goodInfo.BenefitType)
			assert.Equal(t, firstCreateInfo.Info.Classic, goodInfo.Classic)
			assert.Equal(t, firstCreateInfo.Info.SupportCoinTypeIDs, goodInfo.SupportCoinTypeIDs)
			assert.Equal(t, firstCreateInfo.Info.Total, goodInfo.Total)
		}
	}

	goodInfo.BenefitType = "platform"
	goodInfo.ID = firstCreateInfo.Info.ID

	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.UpdateGoodRequest{
			Info: &goodInfo,
		}).
		Post("http://localhost:33759/v1/update/good")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		info := npool.UpdateGoodResponse{}
		err := json.Unmarshal(resp.Body(), &info)
		if assert.Nil(t, err) {
			assert.Equal(t, firstCreateInfo.Info.ID, info.Info.ID)
			assert.Equal(t, info.Info.DeviceInfoID, goodInfo.DeviceInfoID)
			assert.Equal(t, info.Info.GasPrice, goodInfo.GasPrice)
			assert.Equal(t, info.Info.SeparateGasFee, goodInfo.SeparateGasFee)
			assert.Equal(t, info.Info.UnitPower, goodInfo.UnitPower)
			assert.Equal(t, info.Info.DurationDays, goodInfo.DurationDays)
			assert.Equal(t, info.Info.CoinInfoID, goodInfo.CoinInfoID)
			assert.Equal(t, info.Info.DeliveryAt, goodInfo.DeliveryAt)
			assert.Equal(t, info.Info.Actuals, goodInfo.Actuals)
			assert.Equal(t, info.Info.InheritFromGoodID, goodInfo.InheritFromGoodID)
			assert.Equal(t, info.Info.VendorLocationID, goodInfo.VendorLocationID)
			assert.Equal(t, info.Info.Price, goodInfo.Price)
			assert.Equal(t, info.Info.BenefitType, goodInfo.BenefitType)
			assert.Equal(t, info.Info.Classic, goodInfo.Classic)
			assert.Equal(t, info.Info.SupportCoinTypeIDs, goodInfo.SupportCoinTypeIDs)
			assert.Equal(t, info.Info.Total, goodInfo.Total)
		}
	}

	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetGoodRequest{
			ID: goodInfo.ID,
		}).
		Post("http://localhost:33759/v1/get/good")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		info := npool.GetGoodResponse{}
		err := json.Unmarshal(resp.Body(), &info)
		if assert.Nil(t, err) {
			assert.Equal(t, firstCreateInfo.Info.ID, info.Info.ID)
			assert.Equal(t, info.Info.DeviceInfoID, goodInfo.DeviceInfoID)
			assert.Equal(t, info.Info.GasPrice, goodInfo.GasPrice)
			assert.Equal(t, info.Info.SeparateGasFee, goodInfo.SeparateGasFee)
			assert.Equal(t, info.Info.UnitPower, goodInfo.UnitPower)
			assert.Equal(t, info.Info.DurationDays, goodInfo.DurationDays)
			assert.Equal(t, info.Info.CoinInfoID, goodInfo.CoinInfoID)
			assert.Equal(t, info.Info.DeliveryAt, goodInfo.DeliveryAt)
			assert.Equal(t, info.Info.Actuals, goodInfo.Actuals)
			assert.Equal(t, info.Info.InheritFromGoodID, goodInfo.InheritFromGoodID)
			assert.Equal(t, info.Info.VendorLocationID, goodInfo.VendorLocationID)
			assert.Equal(t, info.Info.Price, goodInfo.Price)
			assert.Equal(t, info.Info.BenefitType, goodInfo.BenefitType)
			assert.Equal(t, info.Info.Classic, goodInfo.Classic)
			assert.Equal(t, info.Info.SupportCoinTypeIDs, goodInfo.SupportCoinTypeIDs)
			assert.Equal(t, info.Info.Total, goodInfo.Total)
		}
	}

	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetGoodRequest{
			ID: goodInfo.ID,
		}).
		Post("http://localhost:33759/v1/delete/good")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		info := npool.DeleteGoodResponse{}
		err := json.Unmarshal(resp.Body(), &info)
		if assert.Nil(t, err) {
			assert.Equal(t, firstCreateInfo.Info.ID, info.Info.ID)
			assert.Equal(t, info.Info.DeviceInfoID, goodInfo.DeviceInfoID)
			assert.Equal(t, info.Info.GasPrice, goodInfo.GasPrice)
			assert.Equal(t, info.Info.SeparateGasFee, goodInfo.SeparateGasFee)
			assert.Equal(t, info.Info.UnitPower, goodInfo.UnitPower)
			assert.Equal(t, info.Info.DurationDays, goodInfo.DurationDays)
			assert.Equal(t, info.Info.CoinInfoID, goodInfo.CoinInfoID)
			assert.Equal(t, info.Info.DeliveryAt, goodInfo.DeliveryAt)
			assert.Equal(t, info.Info.Actuals, goodInfo.Actuals)
			assert.Equal(t, info.Info.InheritFromGoodID, goodInfo.InheritFromGoodID)
			assert.Equal(t, info.Info.VendorLocationID, goodInfo.VendorLocationID)
			assert.Equal(t, info.Info.Price, goodInfo.Price)
			assert.Equal(t, info.Info.BenefitType, goodInfo.BenefitType)
			assert.Equal(t, info.Info.Classic, goodInfo.Classic)
			assert.Equal(t, info.Info.SupportCoinTypeIDs, goodInfo.SupportCoinTypeIDs)
			assert.Equal(t, info.Info.Total, goodInfo.Total)
		}
	}

	_, err = cli.R().
		SetHeader("Content-Type", "application/json").
		Get("http://localhost:33759/v1/get/goods")
	assert.Nil(t, err)
}
