package pricecurrency

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

func validatePriceCurrency(t *testing.T, actual, expected *npool.PriceCurrency) {
	assert.Equal(t, actual.Name, expected.Name)
	assert.Equal(t, actual.Unit, expected.Unit)
	assert.Equal(t, actual.Symbol, expected.Symbol)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	nano := time.Now().UnixNano()
	currency := npool.PriceCurrency{
		Name:   fmt.Sprintf("USDT-%v", nano),
		Unit:   "USDT",
		Symbol: "$",
	}

	resp, err := Create(context.Background(), &npool.CreatePriceCurrencyRequest{
		Info: &currency,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		validatePriceCurrency(t, resp.Info, &currency)
	}

	currency.Unit = "USDT1"
	currency.ID = resp.Info.ID

	resp1, err := Update(context.Background(), &npool.UpdatePriceCurrencyRequest{
		Info: &currency,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp1.Info.ID, uuid.UUID{}.String())
		validatePriceCurrency(t, resp1.Info, &currency)
	}

	resp2, err := Get(context.Background(), &npool.GetPriceCurrencyRequest{
		ID: resp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp2.Info.ID, uuid.UUID{}.String())
		validatePriceCurrency(t, resp2.Info, &currency)
	}
}

func TestGetAll(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	resp, err := GetAll(context.Background(), &npool.GetPriceCurrencysRequest{})
	assert.NotNil(t, resp)
	assert.Nil(t, err)
}
