package simulation

import (
	"math/big"
	"os"

	"github.com/69th-byte/SmartDex-Chain/common"
	"github.com/69th-byte/SmartDex-Chain/crypto"
)

var (
	BaseSDX     = big.NewInt(0).Mul(big.NewInt(10), big.NewInt(100000000000000000)) // 1 SDX
	RpcEndpoint = "http://127.0.0.1:8501/"
	MainKey, _  = crypto.HexToECDSA(os.Getenv("MAIN_ADDRESS_KEY"))
	MainAddr    = crypto.PubkeyToAddress(MainKey.PublicKey) //0x17F2beD710ba50Ed27aEa52fc4bD7Bda5ED4a037

	// SRC21 Token
	MinSRC21Apply  = big.NewInt(0).Mul(big.NewInt(10), BaseSDX) // 10 SDX
	SRC21TokenCap  = big.NewInt(0).Mul(big.NewInt(1000000000000), BaseSDX)
	SRC21TokenFee  = big.NewInt(0)
	SdxXListingFee = big.NewInt(0).Mul(big.NewInt(1000), BaseSDX) // 1000 SDX

	// SDXX
	MaxRelayers               = big.NewInt(200)
	MaxTokenList              = big.NewInt(200)
	MinDeposit                = big.NewInt(0).Mul(big.NewInt(25000), BaseSDX) // 25000 SDX
	CollateralDepositRate     = big.NewInt(150)
	CollateralLiquidationRate = big.NewInt(110)
	CollateralRecallRate      = big.NewInt(200)
	TradeFee                  = uint16(10)  // trade fee decimals 10^4
	LendingTradeFee           = uint16(100) // lending trade fee decimals 10^4
	// 1m , 1d,7d,30d
	Terms                 = []*big.Int{big.NewInt(60), big.NewInt(86400), big.NewInt(7 * 86400), big.NewInt(30 * 86400)}
	RelayerCoinbaseKey, _ = crypto.HexToECDSA(os.Getenv("RELAYER_COINBASE_KEY")) //
	RelayerCoinbaseAddr   = crypto.PubkeyToAddress(RelayerCoinbaseKey.PublicKey) // 0x0D3ab14BBaD3D99F4203bd7a11aCB94882050E7e

	OwnerRelayerKey, _ = crypto.HexToECDSA(os.Getenv("RELAYER_OWNER_KEY"))
	OwnerRelayerAddr   = crypto.PubkeyToAddress(OwnerRelayerKey.PublicKey) //0x703c4b2bD70c169f5717101CaeE543299Fc946C7

	SDXNative = common.HexToAddress("0x0000000000000000000000000000000000000001")

	TokenNameList = []string{"BTC", "ETH", "XRP", "LTC", "BNB", "ADA", "ETC", "BCH", "EOS", "USDT"}
	TeamAddresses = []common.Address{
		common.HexToAddress("0x0"), // MM team
		common.HexToAddress("0x0"), // MM team
		common.HexToAddress("0x0"), // CTO
		common.HexToAddress("0x0"), // DEX team
		common.HexToAddress("0x0"), // HaiDV
		common.HexToAddress("0x0"), // Can
		common.HexToAddress("0x0"), // Nien
		common.HexToAddress("0x0"), // Vu Pham
		common.HexToAddress("0x0"), // BTCSDX
		common.HexToAddress("0x0"), // ETHSDX
		common.HexToAddress("0x0"), // XRPSDX
		common.HexToAddress("0x0"), // LTCSDX
		common.HexToAddress("0x0"), // BNBSDX
		common.HexToAddress("0x0"), // ADASDX
		common.HexToAddress("0x0"), // ETCSDX
		common.HexToAddress("0x0"), // BCHSDX
		common.HexToAddress("0x0"), // EOSSDX
		common.HexToAddress("0x0"), // ETHBTC
		common.HexToAddress("0x0"), // XRPBTC
		common.HexToAddress("0x0"), // USDSDX
		common.HexToAddress("0x0"), // BTCUSD
		common.HexToAddress("0x0"), // ETHUSD
		common.HexToAddress("0x0"), // rosetta-cli testing account
	}

	Required = big.NewInt(2)
	Owners   = []common.Address{
		common.HexToAddress("0x0"),
		common.HexToAddress("0x0"),
		common.HexToAddress("0x0"),
	}
)
