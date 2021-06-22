package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/tomochain/tomochain/accounts/abi/bind"
	"github.com/tomochain/tomochain/common"
	"github.com/tomochain/tomochain/common/hexutil"
	"github.com/tomochain/tomochain/contracts/src21issuer"
	"github.com/tomochain/tomochain/contracts/src21issuer/simulation"
	"github.com/tomochain/tomochain/ethclient"
)

var (
	src21TokenAddr = common.HexToAddress("0x80430A33EaB86890a346bCf64F86CFeAC73287f3")
)

func airDropTokenToAccountNoTomo() {
	client, err := ethclient.Dial(simulation.RpcEndpoint)
	if err != nil {
		fmt.Println(err, client)
	}
	nonce, _ := client.NonceAt(context.Background(), simulation.MainAddr, nil)
	mainAccount := bind.NewKeyedTransactor(simulation.MainKey)
	mainAccount.Nonce = big.NewInt(int64(nonce))
	mainAccount.Value = big.NewInt(0)      // in wei
	mainAccount.GasLimit = uint64(4000000) // in units
	mainAccount.GasPrice = big.NewInt(0).Mul(common.SRC21GasPrice, big.NewInt(2))
	src21Instance, _ := src21issuer.NewSRC21(mainAccount, src21TokenAddr, client)
	src21IssuerInstance, _ := src21issuer.NewSRC21Issuer(mainAccount, common.SRC21IssuerSMC, client)
	// air drop token
	remainFee, _ := src21IssuerInstance.GetTokenCapacity(src21TokenAddr)
	tx, err := src21Instance.Transfer(simulation.AirdropAddr, simulation.AirDropAmount)
	if err != nil {
		log.Fatal("can't air drop to ", err)
	}
	// check balance after transferAmount
	fmt.Println("wait 10s to airdrop success ", tx.Hash().Hex())
	time.Sleep(10 * time.Second)

	_, receiptRpc, err := client.GetTransactionReceiptResult(context.Background(), tx.Hash())
	receipt := map[string]interface{}{}
	err = json.Unmarshal(receiptRpc, &receipt)
	if err != nil {
		log.Fatal("can't transaction's receipt ", err, "hash", tx.Hash().Hex())
	}
	fee := big.NewInt(0).SetUint64(hexutil.MustDecodeUint64(receipt["gasUsed"].(string)))
	if hexutil.MustDecodeUint64(receipt["blockNumber"].(string)) > common.TIPSRC21Fee.Uint64() {
		fee = fee.Mul(fee, common.SRC21GasPrice)
	}
	fmt.Println("fee", fee.Uint64(), "number", hexutil.MustDecodeUint64(receipt["blockNumber"].(string)))
	remainFee = big.NewInt(0).Sub(remainFee, fee)
	//check balance fee
	balanceIssuerFee, err := src21IssuerInstance.GetTokenCapacity(src21TokenAddr)
	if err != nil || balanceIssuerFee.Cmp(remainFee) != 0 {
		log.Fatal("can't get balance token fee in  smart contract: ", err, "got", balanceIssuerFee, "wanted", remainFee)
	}
	if err != nil {
		log.Fatal("can't execute transferAmount in tr21:", err)
	}
}
func testTransferSRC21TokenWithAccountNoTomo() {
	client, err := ethclient.Dial(simulation.RpcEndpoint)
	if err != nil {
		fmt.Println(err, client)
	}

	// access to address which received token trc20 but dont have sdx
	nonce, _ := client.NonceAt(context.Background(), simulation.AirdropAddr, nil)
	airDropAccount := bind.NewKeyedTransactor(simulation.AirdropKey)
	airDropAccount.Nonce = big.NewInt(int64(nonce))
	airDropAccount.Value = big.NewInt(0)      // in wei
	airDropAccount.GasLimit = uint64(4000000) // in units
	airDropAccount.GasPrice = big.NewInt(0).Mul(common.SRC21GasPrice, big.NewInt(2))
	src21Instance, _ := src21issuer.NewSRC21(airDropAccount, src21TokenAddr, client)
	src21IssuerInstance, _ := src21issuer.NewSRC21Issuer(airDropAccount, common.SRC21IssuerSMC, client)

	remainFee, _ := src21IssuerInstance.GetTokenCapacity(src21TokenAddr)
	airDropBalanceBefore, err := src21Instance.BalanceOf(simulation.AirdropAddr)
	receiverBalanceBefore, err := src21Instance.BalanceOf(simulation.ReceiverAddr)
	// execute transferAmount trc to other address
	tx, err := src21Instance.Transfer(simulation.ReceiverAddr, simulation.TransferAmount)
	if err != nil {
		log.Fatal("can't execute transferAmount in tr21:", err)
	}

	// check balance after transferAmount
	fmt.Println("wait 10s to transferAmount success ")
	time.Sleep(10 * time.Second)

	balance, err := src21Instance.BalanceOf(simulation.ReceiverAddr)
	wantedBalance := big.NewInt(0).Add(receiverBalanceBefore, simulation.TransferAmount)
	if err != nil || balance.Cmp(wantedBalance) != 0 {
		log.Fatal("check balance after fail receiverAmount in tr21: ", err, "get", balance, "wanted", wantedBalance)
	}

	remainAirDrop := big.NewInt(0).Sub(airDropBalanceBefore, simulation.TransferAmount)
	remainAirDrop = remainAirDrop.Sub(remainAirDrop, simulation.Fee)
	// check balance src21 again
	balance, err = src21Instance.BalanceOf(simulation.AirdropAddr)
	if err != nil || balance.Cmp(remainAirDrop) != 0 {
		log.Fatal("check balance after fail transferAmount in tr21: ", err, "get", balance, "wanted", remainAirDrop)
	}
	_, receiptRpc, err := client.GetTransactionReceiptResult(context.Background(), tx.Hash())
	receipt := map[string]interface{}{}
	err = json.Unmarshal(receiptRpc, &receipt)
	if err != nil {
		log.Fatal("can't transaction's receipt ", err, "hash", tx.Hash().Hex())
	}
	fee := big.NewInt(0).SetUint64(hexutil.MustDecodeUint64(receipt["gasUsed"].(string)))
	if hexutil.MustDecodeUint64(receipt["blockNumber"].(string)) > common.TIPSRC21Fee.Uint64() {
		fee = fee.Mul(fee, common.SRC21GasPrice)
	}
	fmt.Println("fee", fee.Uint64(), "number", hexutil.MustDecodeUint64(receipt["blockNumber"].(string)))
	remainFee = big.NewInt(0).Sub(remainFee, fee)
	//check balance fee
	balanceIssuerFee, err := src21IssuerInstance.GetTokenCapacity(src21TokenAddr)
	if err != nil || balanceIssuerFee.Cmp(remainFee) != 0 {
		log.Fatal("can't get balance token fee in  smart contract: ", err, "got", balanceIssuerFee, "wanted", remainFee)
	}
	//check src21 SMC balance
	balance, err = client.BalanceAt(context.Background(), common.SRC21IssuerSMC, nil)
	if err != nil || balance.Cmp(remainFee) != 0 {
		log.Fatal("can't get balance token fee in  smart contract: ", err, "got", balanceIssuerFee, "wanted", remainFee)
	}
}
func testTransferTrc21Fail() {
	client, err := ethclient.Dial(simulation.RpcEndpoint)
	if err != nil {
		fmt.Println(err, client)
	}
	nonce, _ := client.NonceAt(context.Background(), simulation.AirdropAddr, nil)
	airDropAccount := bind.NewKeyedTransactor(simulation.AirdropKey)
	airDropAccount.Nonce = big.NewInt(int64(nonce))
	airDropAccount.Value = big.NewInt(0)      // in wei
	airDropAccount.GasLimit = uint64(4000000) // in units
	airDropAccount.GasPrice = big.NewInt(0).Mul(common.SRC21GasPrice, big.NewInt(2))
	src21Instance, _ := src21issuer.NewSRC21(airDropAccount, src21TokenAddr, client)
	src21IssuerInstance, _ := src21issuer.NewSRC21Issuer(airDropAccount, common.SRC21IssuerSMC, client)
	balanceIssuerFee, err := src21IssuerInstance.GetTokenCapacity(src21TokenAddr)

	minFee, err := src21Instance.MinFee()
	if err != nil {
		log.Fatal("can't get minFee of src21 smart contract:", err)
	}
	ownerBalance, err := src21Instance.BalanceOf(simulation.MainAddr)
	remainFee, err := src21IssuerInstance.GetTokenCapacity(src21TokenAddr)
	airDropBalanceBefore, err := src21Instance.BalanceOf(simulation.AirdropAddr)

	tx, err := src21Instance.Transfer(common.Address{}, big.NewInt(1))
	if err != nil {
		log.Fatal("can't execute test transfer to zero address in tr21:", err)
	}
	fmt.Println("wait 10s to transfer to zero address")
	time.Sleep(10 * time.Second)

	fmt.Println("airDropBalanceBefore", airDropBalanceBefore)
	// check balance src21 again
	airDropBalanceBefore = big.NewInt(0).Sub(airDropBalanceBefore, minFee)
	balance, err := src21Instance.BalanceOf(simulation.AirdropAddr)
	if err != nil || balance.Cmp(airDropBalanceBefore) != 0 {
		log.Fatal("check balance after fail transferAmount in tr21: ", err, "get", balance, "wanted", airDropBalanceBefore)
	}

	ownerBalance = big.NewInt(0).Add(ownerBalance, minFee)
	//check balance fee
	balance, err = src21Instance.BalanceOf(simulation.MainAddr)
	if err != nil || balance.Cmp(ownerBalance) != 0 {
		log.Fatal("can't get balance token fee in  smart contract: ", err, "got", balanceIssuerFee, "wanted", remainFee)
	}
	_, receiptRpc, err := client.GetTransactionReceiptResult(context.Background(), tx.Hash())
	receipt := map[string]interface{}{}
	err = json.Unmarshal(receiptRpc, &receipt)
	if err != nil {
		log.Fatal("can't transaction's receipt ", err, "hash", tx.Hash().Hex())
	}
	fee := big.NewInt(0).SetUint64(hexutil.MustDecodeUint64(receipt["gasUsed"].(string)))
	if hexutil.MustDecodeUint64(receipt["blockNumber"].(string)) > common.TIPSRC21Fee.Uint64() {
		fee = fee.Mul(fee, common.SRC21GasPrice)
	}
	fmt.Println("fee", fee.Uint64(), "number", hexutil.MustDecodeUint64(receipt["blockNumber"].(string)))
	remainFee = big.NewInt(0).Sub(remainFee, fee)
	//check balance fee
	balanceIssuerFee, err = src21IssuerInstance.GetTokenCapacity(src21TokenAddr)
	if err != nil || balanceIssuerFee.Cmp(remainFee) != 0 {
		log.Fatal("can't get balance token fee in  smart contract: ", err, "got", balanceIssuerFee, "wanted", remainFee)
	}
	//check src21 SMC balance
	balance, err = client.BalanceAt(context.Background(), common.SRC21IssuerSMC, nil)
	if err != nil || balance.Cmp(remainFee) != 0 {
		log.Fatal("can't get balance token fee in  smart contract: ", err, "got", balanceIssuerFee, "wanted", remainFee)
	}

}
func main() {
	fmt.Println("========================")
	fmt.Println("airdropAddr", simulation.AirdropAddr.Hex())
	fmt.Println("receiverAddr", simulation.ReceiverAddr.Hex())
	fmt.Println("========================")

	start := time.Now()
	for i := 0; i < 10000000; i++ {
		airDropTokenToAccountNoTomo()
		fmt.Println("Finish airdrop token to a account")
		testTransferSRC21TokenWithAccountNoTomo()
		fmt.Println("Finish transfer src21 token with a account no sdx")
		testTransferTrc21Fail()
		fmt.Println("Finish testing ! Success transferAmount token trc20 with a account no sdx")
	}
	fmt.Println(common.PrettyDuration(time.Since(start)))
}
