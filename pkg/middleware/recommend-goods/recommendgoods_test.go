package recommendgoods

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"
	deviceinfo "github.com/NpoolPlatform/cloud-hashing-goods/pkg/crud/device-info"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/crud/fee"
	feetype "github.com/NpoolPlatform/cloud-hashing-goods/pkg/crud/fee-type"
	goodinfo "github.com/NpoolPlatform/cloud-hashing-goods/pkg/crud/good-info"
	pricecurrency "github.com/NpoolPlatform/cloud-hashing-goods/pkg/crud/price-currency"
	recommendgood "github.com/NpoolPlatform/cloud-hashing-goods/pkg/crud/recommend-good"
	vendorlocation "github.com/NpoolPlatform/cloud-hashing-goods/pkg/crud/vendor-location"
	testinit "github.com/NpoolPlatform/cloud-hashing-goods/pkg/test-init" //nolint
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestRecommendGoodsCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
	ctx := context.Background()

	nano := time.Now().UnixNano()

	deviceInfo := npool.DeviceInfo{
		Type:            fmt.Sprintf("S19-%v", nano),
		Manufacturer:    fmt.Sprintf("Ant-%v", nano),
		PowerComsuption: 120,
		ShipmentAt:      int32(time.Now().Unix()),
	}

	deviceResp, err := deviceinfo.Create(context.Background(), &npool.CreateDeviceInfoRequest{
		Info: &deviceInfo,
	})
	assert.Nil(t, err)

	vendorLocation := npool.VendorLocationInfo{
		Country:  fmt.Sprintf("China-%v", nano),
		Province: fmt.Sprintf("Shanghai-%v", nano),
		City:     fmt.Sprintf("Shanghai-%v", nano),
		Address:  fmt.Sprintf("Shanghai-%v", nano),
	}
	vendorLocationResp, err := vendorlocation.Create(context.Background(), &npool.CreateVendorLocationRequest{
		Info: &vendorLocation,
	})
	assert.Nil(t, err)

	currency := npool.PriceCurrency{
		Name:   fmt.Sprintf("USDT-%v", nano),
		Unit:   "USDT",
		Symbol: "$",
	}

	currencyResp, err := pricecurrency.Create(context.Background(), &npool.CreatePriceCurrencyRequest{
		Info: &currency,
	})
	assert.Nil(t, err)

	feeType1 := npool.FeeType{
		FeeType: fmt.Sprintf("Gas Fee-%v", nano),
		PayType: "amount",
	}
	feeTypeResp1, err := feetype.Create(context.Background(), &npool.CreateFeeTypeRequest{
		Info: &feeType1,
	})
	assert.Nil(t, err)

	fee1 := npool.Fee{
		AppID:     uuid.New().String(),
		FeeTypeID: feeTypeResp1.Info.ID,
		Value:     20,
	}
	feeResp1, err := fee.Create(context.Background(), &npool.CreateFeeRequest{
		Info: &fee1,
	})
	assert.Nil(t, err)

	feeType2 := npool.FeeType{
		FeeType: fmt.Sprintf("Gas Fee 1-%v", nano),
		PayType: "amount",
	}
	feeTypeResp2, err := feetype.Create(context.Background(), &npool.CreateFeeTypeRequest{
		Info: &feeType2,
	})
	assert.Nil(t, err)

	fee2 := npool.Fee{
		AppID:     fee1.AppID,
		FeeTypeID: feeTypeResp2.Info.ID,
		Value:     20,
	}
	feeResp2, err := fee.Create(context.Background(), &npool.CreateFeeRequest{
		Info: &fee2,
	})
	assert.Nil(t, err)

	goodInfo := npool.GoodInfo{
		Title:              "Ant Miner S19 Pro",
		Unit:               "TH/s",
		DeviceInfoID:       deviceResp.Info.ID,
		SeparateFee:        true,
		UnitPower:          10,
		DurationDays:       180,
		CoinInfoID:         uuid.New().String(),
		Actuals:            true,
		DeliveryAt:         uint32(time.Now().Unix()),
		VendorLocationID:   vendorLocationResp.Info.ID,
		InheritFromGoodID:  uuid.UUID{}.String(),
		Price:              13.0,
		PriceCurrency:      currencyResp.Info.ID,
		BenefitType:        "platform",
		Classic:            true,
		SupportCoinTypeIDs: []string{uuid.New().String(), uuid.New().String()},
		Total:              100,
		FeeIDs:             []string{feeResp1.Info.ID, feeResp2.Info.ID},
	}
	goodInfoResp, err := goodinfo.Create(context.Background(), &npool.CreateGoodRequest{
		Info: &goodInfo,
	})
	assert.Nil(t, err)

	testGood := &npool.RecommendGood{
		RecommenderID: uuid.NewString(),
		GoodID:        goodInfoResp.Info.ID,
		Message:       "created in unit test",
	}
	resp, err := recommendgood.Create(ctx, &npool.CreateRecommendGoodRequest{
		Info: testGood,
	})
	if assert.Nil(t, err) {
		assert.NotNil(t, resp.Info)
		assert.Equal(t, testGood.GoodID, resp.Info.GoodID)
		assert.Equal(t, testGood.RecommenderID, resp.Info.RecommenderID)
	}

	resp1, err := Get(ctx, &npool.GetRecommendGoodsRequest{
		PageInfo: &npool.PageInfo{
			PageIndex: 0,
			PageSize:  10,
		},
	})
	if assert.Nil(t, err) {
		assert.NotNil(t, resp1)
		assert.Positive(t, len(resp1.Infos))
		assert.NotNil(t, resp1.Infos[0])
	}

	resp2, err := GetByRecommender(ctx, &npool.GetRecommendGoodsByRecommenderRequest{
		RecommenderID: testGood.RecommenderID,
		PageInfo:      &npool.PageInfo{PageIndex: 0, PageSize: 10},
	})
	if assert.Nil(t, err) {
		assert.NotNil(t, resp2)
		assert.Positive(t, len(resp2.Infos))
		assert.NotNil(t, resp1.Infos[0])
	}
}
