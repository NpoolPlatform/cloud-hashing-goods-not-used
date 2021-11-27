package gooddetail

import (
	"context"
	"math/rand"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/crud/device-info"     //nolint
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/crud/fee"             //nolint
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/crud/good-extra-info" //nolint
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/crud/good-fee"        //nolint
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/crud/good-info"       //nolint
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/crud/price-currency"  //nolint
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/crud/vendor-location" //nolint

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func Get(ctx context.Context, in *npool.GetGoodDetailRequest) (*npool.GetGoodDetailResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid good id: %v", err)
	}

	goodInfo, err := goodinfo.Get(ctx, &npool.GetGoodRequest{
		ID: id.String(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get good: %v", err)
	}

	var inheritGoodInfo *npool.GoodInfo

	inheritFromGoodInfo, err := goodinfo.Get(ctx, &npool.GetGoodRequest{
		ID: goodInfo.Info.InheritFromGoodID,
	})
	if err == nil {
		inheritGoodInfo = inheritFromGoodInfo.Info
	}

	vendorLocation, err := vendorlocation.Get(ctx, &npool.GetVendorLocationRequest{
		ID: goodInfo.Info.VendorLocationID,
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get vendor location: %v", err)
	}

	deviceInfo, err := deviceinfo.Get(ctx, &npool.GetDeviceInfoRequest{
		ID: goodInfo.Info.DeviceInfoID,
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get device info: %v", err)
	}

	extraInfo := &npool.GoodExtraInfo{
		Rating:    4.6,
		VoteCount: uint32(rand.Intn(140)),
	}

	extra, err := goodextrainfo.Get(ctx, &npool.GetGoodExtraInfoRequest{
		GoodID: id.String(),
	})
	if err == nil {
		extraInfo = extra.Info
	}

	priceCurrency, err := pricecurrency.Get(ctx, &npool.GetPriceCurrencyRequest{
		ID: goodInfo.Info.PriceCurrency,
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get price currency %v: %v", goodInfo.Info.PriceCurrency, err)
	}

	fees := []*npool.FeeDetail{}
	for _, id := range goodInfo.Info.FeeIDs {
		myFee, err := fee.Get(ctx, &npool.GetFeeRequest{
			ID: id,
		})
		if err != nil {
			return nil, xerrors.Errorf("fail get fee: %v", err)
		}

		myGoodFee, err := goodfee.Get(ctx, &npool.GetGoodFeeRequest{
			ID: myFee.Info.FeeTypeID,
		})
		if err != nil {
			return nil, xerrors.Errorf("fail get good fee: %v", err)
		}
		fees = append(fees, &npool.FeeDetail{
			Fee:   myGoodFee.Info,
			Value: myFee.Info.Value,
		})
	}

	return &npool.GetGoodDetailResponse{
		Detail: &npool.GoodDetail{
			ID:                 id.String(),
			DeviceInfo:         deviceInfo.Info,
			SeparateFee:        goodInfo.Info.SeparateFee,
			UnitPower:          goodInfo.Info.UnitPower,
			DurationDays:       goodInfo.Info.DurationDays,
			CoinInfoID:         goodInfo.Info.CoinInfoID,
			Actuals:            goodInfo.Info.Actuals,
			DeliveryAt:         goodInfo.Info.DeliveryAt,
			InheritFromGood:    inheritGoodInfo,
			VendorLocation:     vendorLocation.Info,
			Price:              goodInfo.Info.Price,
			PriceCurrency:      priceCurrency.Info,
			BenefitType:        goodInfo.Info.BenefitType,
			Classic:            goodInfo.Info.Classic,
			SupportCoinTypeIDs: goodInfo.Info.SupportCoinTypeIDs,
			Total:              goodInfo.Info.Total,
			Extra:              extraInfo,
			Unit:               goodInfo.Info.Unit,
			Title:              goodInfo.Info.Title,
			Fees:               fees,
		},
	}, nil
}

func GetAll(ctx context.Context, in *npool.GetGoodsDetailRequest) (*npool.GetGoodsDetailResponse, error) {
	resp, err := goodinfo.GetAll(ctx, &npool.GetGoodsRequest{
		PageInfo: in.GetPageInfo(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get goods info: %v", err)
	}

	details := []*npool.GoodDetail{}
	for _, info := range resp.Infos {
		detail, err := Get(ctx, &npool.GetGoodDetailRequest{
			ID: info.ID,
		})
		if err != nil {
			logger.Sugar().Errorf("fail get good detail: %v", err)
			continue
		}
		details = append(details, detail.Detail)
	}

	return &npool.GetGoodsDetailResponse{
		Details: details,
	}, nil
}
