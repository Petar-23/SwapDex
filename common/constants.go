package common

import (
	"math/big"
)

const (
	RewardMasterPercent        = 51
	RewardVoterPercent         = 19
	RewardFoundationPercent    = 30
	HexSignMethod              = "e341eaa4"
	HexSetSecret               = "34d38600"
	HexSetOpening              = "e11f5ba2"
	EpocBlockSecret            = 800
	EpocBlockOpening           = 850
	EpocBlockRandomize         = 900
	MaxMasternodes             = 100000000000000
	LimitPenaltyEpoch          = 4
	BlocksPerYear              = uint64(15768000)
	LimitThresholdNonceInQueue = 10
	DefaultMinGasPrice         = 250000000
	MergeSignRange             = 15
	RangeReturnSigner          = 150
	MinimunMinerBlockPerEpoch  = 1
	//IgnoreSignerCheckBlock     = uint64(14458500) // Not Needed
	OneYear                    = uint64(365 * 86400)
	LiquidateLendingTradeBlock = uint64(100)
)

var Rewound = uint64(0)

// hardforks
var TIP2019Block = big.NewInt(1050000)
var TIPSigning = big.NewInt(1648000)
var TIPRandomize = big.NewInt(1648000)
var BlackListHFNumber = uint64(1648000)
var TIPSdxX = big.NewInt(1648000)
var TIPSdxXLending = big.NewInt(1648000)
var TIPSdxXCancellationFee = big.NewInt(1648000)
var TIPSdxXTestnet = big.NewInt(0)
var IsTestnet bool = false
var StoreRewardFolder string
var RollbackHash Hash
var BasePrice = big.NewInt(1000000000000000000)                       // 1
var RelayerLockedFund = big.NewInt(20000)                             // 20000 SDX
var RelayerFee = big.NewInt(1000000000000000)                         // 0.001
var SdxXBaseFee = big.NewInt(10000)                                   // 1 / SdxXBaseFee
var RelayerCancelFee = big.NewInt(100000000000000)                    // 0.0001
var SdxXBaseCancelFee = new(big.Int).Mul(SdxXBaseFee, big.NewInt(10)) // 1/ (SdxXBaseFee *10)
var RelayerLendingFee = big.NewInt(10000000000000000)                 // 0.01
var RelayerLendingCancelFee = big.NewInt(1000000000000000)            // 0.001
var BaseLendingInterest = big.NewInt(100000000)                       // 1e8

var MinGasPrice = big.NewInt(DefaultMinGasPrice)
var RelayerRegistrationSMC = "0x0" // These need to be added at a later date once all contracts are deployed
var RelayerRegistrationSMCTestnet = "0x0"
var LendingRegistrationSMC = "0x0" // These need to be added at a later date once all contracts are deployed
var LendingRegistrationSMCTestnet = "0x0"
var SRC21IssuerSMCTestNet = HexToAddress("0x0")
var SRC21IssuerSMC = HexToAddress("0x0") // These need to be added at a later date once all contracts are deployed
var SdxXListingSMC = HexToAddress("0x0") // These need to be added at a later date once all contracts are deployed
var SdxXListingSMCTestNet = HexToAddress("0x0")
var SRC21GasPriceBefore = big.NewInt(2500)
var SRC21GasPrice = big.NewInt(250000000)
var RateTopUp = big.NewInt(90) // 90%
var BaseTopUp = big.NewInt(100)
var BaseRecall = big.NewInt(100)
var Blacklist = map[Address]bool{
	HexToAddress(""): true,
}
var TIPSRC21Fee = big.NewInt(1648000)
var LimitTimeFinality = uint64(30) // limit in 30 block
