package vm

import (
	"github.com/tomochain/tomochain/common"
	"github.com/tomochain/tomochain/log"
	"github.com/tomochain/tomochain/params"
	"github.com/tomochain/tomochain/sdxx/tradingstate"
)
const TomoXPriceNumberOfBytesReturn = 32
// sdxxPrice implements a pre-compile contract to get token price in sdxx

type sdxxLastPrice struct {
	tradingStateDB *tradingstate.TradingStateDB
}
type sdxxEpochPrice struct {
	tradingStateDB *tradingstate.TradingStateDB
}

func (t *sdxxLastPrice) RequiredGas(input []byte) uint64 {
	return params.TomoXPriceGas
}

func (t *sdxxLastPrice) Run(input []byte) ([]byte, error) {
	// input includes baseTokenAddress, quoteTokenAddress
	if t.tradingStateDB != nil && len(input) == 64 {
		base := common.BytesToAddress(input[12:32]) // 20 bytes from 13-32
		quote := common.BytesToAddress(input[44:]) // 20 bytes from 45-64
		price := t.tradingStateDB.GetLastPrice(tradingstate.GetTradingOrderBookHash(base, quote))
		if price != nil {
			log.Debug("Run GetLastPrice", "base", base.Hex(), "quote", quote.Hex(), "price", price)
			return common.LeftPadBytes(price.Bytes(), TomoXPriceNumberOfBytesReturn), nil
		}
	}
	return common.LeftPadBytes([]byte{}, TomoXPriceNumberOfBytesReturn), nil
}

func (t *sdxxLastPrice) SetTradingState(tradingStateDB *tradingstate.TradingStateDB) {
	if tradingStateDB != nil {
		t.tradingStateDB = tradingStateDB.Copy()
	} else {
		t.tradingStateDB = nil
	}
}

func (t *sdxxEpochPrice) RequiredGas(input []byte) uint64 {
	return params.TomoXPriceGas
}

func (t *sdxxEpochPrice) Run(input []byte) ([]byte, error) {
	// input includes baseTokenAddress, quoteTokenAddress
	if t.tradingStateDB != nil && len(input) == 64 {
		base := common.BytesToAddress(input[12:32]) // 20 bytes from 13-32
		quote := common.BytesToAddress(input[44:]) // 20 bytes from 45-64
		price := t.tradingStateDB.GetMediumPriceBeforeEpoch(tradingstate.GetTradingOrderBookHash(base, quote))
		if price != nil {
			log.Debug("Run GetEpochPrice", "base", base.Hex(), "quote", quote.Hex(), "price", price)
			return common.LeftPadBytes(price.Bytes(), TomoXPriceNumberOfBytesReturn), nil
		}
	}
	return common.LeftPadBytes([]byte{}, TomoXPriceNumberOfBytesReturn), nil
}

func (t *sdxxEpochPrice) SetTradingState(tradingStateDB *tradingstate.TradingStateDB) {
	if tradingStateDB != nil {
		t.tradingStateDB = tradingStateDB.Copy()
	} else {
		t.tradingStateDB = nil
	}
}


