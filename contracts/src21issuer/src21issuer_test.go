package src21issuer

import (
	"math/big"
	"testing"

	"github.com/69th-byte/SmartDex-Chain/accounts/abi/bind"
	"github.com/69th-byte/SmartDex-Chain/accounts/abi/bind/backends"
	"github.com/69th-byte/SmartDex-Chain/common"
	"github.com/69th-byte/SmartDex-Chain/core"
	"github.com/69th-byte/SmartDex-Chain/crypto"
)

var (
	mainKey, _ = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	mainAddr   = crypto.PubkeyToAddress(mainKey.PublicKey)

	airdropKey, _ = crypto.HexToECDSA("49a7b37aa6f6645917e7b807e9d1c00d4fa71f18343b0d4122a4d2df64dd6fee")
	airdropAddr   = crypto.PubkeyToAddress(airdropKey.PublicKey)

	subKey, _ = crypto.HexToECDSA("5bb98c5f937d176aa399ea6e6541f4db8f8db5a4ee1a8b56fb8beb41f2d755e3")
	subAddr   = crypto.PubkeyToAddress(subKey.PublicKey) //0x21292d56E2a8De3cC4672dB039AAA27f9190B1f6

	token = common.HexToAddress("0000000000000000000000000000000000000089")

	delay    = big.NewInt(30 * 48)
	minApply = big.NewInt(0).Mul(big.NewInt(1000), big.NewInt(100000000000000000)) // 100 SDX
)

func TestFeeTxWithSRC21Token(t *testing.T) {

	// init genesis
	contractBackend := backends.NewSimulatedBackend(core.GenesisAlloc{
		mainAddr: {Balance: big.NewInt(0).Mul(big.NewInt(10000000000000), big.NewInt(10000000000000))},
	})
	transactOpts := bind.NewKeyedTransactor(mainKey)
	// deploy payer swap SMC
	src21IssuerAddr, src21Issuer, err := DeploySRC21Issuer(transactOpts, contractBackend, minApply)

	//set contract address to config
	common.SRC21IssuerSMC = src21IssuerAddr
	if err != nil {
		t.Fatal("can't deploy smart contract: ", err)
	}
	contractBackend.Commit()
	cap := big.NewInt(0).Mul(big.NewInt(10000000), big.NewInt(10000000000000))
	SRC21fee := big.NewInt(100)
	//  deploy a SRC21 SMC
	src21TokenAddr, src21, err := DeploySRC21(transactOpts, contractBackend, "TEST", "SDX", 18, cap, SRC21fee)
	if err != nil {
		t.Fatal("can't deploy smart contract: ", err)
	}
	contractBackend.Commit()
	// add src21 address to list token src21Issuer
	src21Issuer.TransactOpts.Value = minApply
	_, err = src21Issuer.Apply(src21TokenAddr)
	if err != nil {
		t.Fatal("can't add a token in  smart contract pay swap: ", err)
	}
	contractBackend.Commit()

	//check src21 SMC balance
	balance, err := contractBackend.BalanceAt(nil, src21IssuerAddr, nil)
	if err != nil || balance.Cmp(minApply) != 0 {
		t.Fatal("can't get balance  in src21Issuer SMC: ", err, "got", balance, "wanted", minApply)
	}

	//check balance fee
	balanceIssuerFee, err := src21Issuer.GetTokenCapacity(src21TokenAddr)
	if err != nil || balanceIssuerFee.Cmp(minApply) != 0 {
		t.Fatal("can't get balance token fee in  smart contract: ", err, "got", balanceIssuerFee, "wanted", minApply)
	}
	src21Issuer.TransactOpts.Value = big.NewInt(0)
	airDropAmount := big.NewInt(1000000000)
	// airdrop token src21 to a address no sdx
	tx, err := src21.Transfer(airdropAddr, airDropAmount)
	if err != nil {
		t.Fatal("can't execute transfer in tr20: ", err)
	}
	contractBackend.Commit()
	receipt, err := contractBackend.TransactionReceipt(nil, tx.Hash())
	if err != nil {
		t.Fatal("can't transaction's receipt ", err, "hash", tx.Hash())
	}
	fee := big.NewInt(0).SetUint64(receipt.GasUsed)
	if receipt.Logs[0].BlockNumber > common.TIPSRC21Fee.Uint64() {
		fee = fee.Mul(fee, common.SRC21GasPrice)
	}
	remainFee := big.NewInt(0).Sub(minApply, fee)

	// check balance src21 again
	balance, err = src21.BalanceOf(airdropAddr)
	if err != nil || balance.Cmp(airDropAmount) != 0 {
		t.Fatal("check balance after fail transfer in tr20: ", err, "get", balance, "transfer", airDropAmount)
	}

	//check balance fee
	balanceIssuerFee, err = src21Issuer.GetTokenCapacity(src21TokenAddr)
	if err != nil || balanceIssuerFee.Cmp(remainFee) != 0 {
		t.Fatal("can't get balance token fee in  smart contract: ", err, "got", balanceIssuerFee, "wanted", remainFee)
	}
	//check src21 SMC balance
	balance, err = contractBackend.BalanceAt(nil, src21IssuerAddr, nil)
	if err != nil || balance.Cmp(remainFee) != 0 {
		t.Fatal("can't get balance token fee in  smart contract: ", err, "got", balanceIssuerFee, "wanted", remainFee)
	}

	// access to address which received token src21 but dont have sdx
	key1TransactOpts := bind.NewKeyedTransactor(airdropKey)
	key1Trc20, _ := NewSRC21(key1TransactOpts, src21TokenAddr, contractBackend)

	transferAmount := big.NewInt(100000)
	// execute transfer trc to other address
	tx, err = key1Trc20.Transfer(subAddr, transferAmount)
	if err != nil {
		t.Fatal("can't execute transfer in tr20:", err)
	}
	contractBackend.Commit()

	balance, err = src21.BalanceOf(subAddr)
	if err != nil || balance.Cmp(transferAmount) != 0 {
		t.Fatal("check balance after fail transfer in tr20: ", err, "get", balance, "transfer", transferAmount)
	}

	remainAirDrop := big.NewInt(0).Sub(airDropAmount, transferAmount)
	remainAirDrop = remainAirDrop.Sub(remainAirDrop, SRC21fee)
	// check balance src21 again
	balance, err = src21.BalanceOf(airdropAddr)
	if err != nil || balance.Cmp(remainAirDrop) != 0 {
		t.Fatal("check balance after fail transfer in tr20: ", err, "get", balance, "wanted", remainAirDrop)
	}

	receipt, err = contractBackend.TransactionReceipt(nil, tx.Hash())
	if err != nil {
		t.Fatal("can't transaction's receipt ", err, "hash", tx.Hash())
	}
	fee = big.NewInt(0).SetUint64(receipt.GasUsed)
	if receipt.Logs[0].BlockNumber > common.TIPSRC21Fee.Uint64() {
		fee = fee.Mul(fee, common.SRC21GasPrice)
	}
	remainFee = big.NewInt(0).Sub(remainFee, fee)
	//check balance fee
	balanceIssuerFee, err = src21Issuer.GetTokenCapacity(src21TokenAddr)
	if err != nil || balanceIssuerFee.Cmp(remainFee) != 0 {
		t.Fatal("can't get balance token fee in  smart contract: ", err, "got", balanceIssuerFee, "wanted", remainFee)
	}
	//check src21 SMC balance
	balance, err = contractBackend.BalanceAt(nil, src21IssuerAddr, nil)
	if err != nil || balance.Cmp(remainFee) != 0 {
		t.Fatal("can't get balance token fee in  smart contract: ", err, "got", balanceIssuerFee, "wanted", remainFee)
	}
}
