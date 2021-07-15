package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/69th-byte/SmartDex-Chain/accounts/abi/bind"
	"github.com/69th-byte/SmartDex-Chain/common"
	"github.com/69th-byte/SmartDex-Chain/contracts/src21issuer"
	"github.com/69th-byte/SmartDex-Chain/contracts/src21issuer/simulation"
	"github.com/69th-byte/SmartDex-Chain/ethclient"
)

func main() {
	fmt.Println("========================")
	fmt.Println("mainAddr", simulation.MainAddr.Hex())
	fmt.Println("airdropAddr", simulation.AirdropAddr.Hex())
	fmt.Println("receiverAddr", simulation.ReceiverAddr.Hex())
	fmt.Println("========================")
	client, err := ethclient.Dial(simulation.RpcEndpoint)
	if err != nil {
		fmt.Println(err, client)
	}
	nonce, _ := client.NonceAt(context.Background(), simulation.MainAddr, nil)

	// init sdx21 issuer
	auth := bind.NewKeyedTransactor(simulation.MainKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(4000000) // in units
	auth.GasPrice = big.NewInt(210000000000000)
	src21IssuerAddr, src21Issuer, err := src21issuer.DeploySRC21Issuer(auth, client, simulation.MinApply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("main address", simulation.MainAddr.Hex(), "nonce", nonce)
	fmt.Println("===> sdx21 issuer address", src21IssuerAddr.Hex())

	auth.Nonce = big.NewInt(int64(nonce + 1))

	// init trc20
	src21TokenAddr, src21Token, err := src21issuer.DeploySRC21(auth, client, "TEST", "SDX", 18, simulation.Cap, simulation.Fee)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("===>  src21 token address", src21TokenAddr.Hex(), "cap", simulation.Cap)

	fmt.Println("wait 10s to execute init smart contract")
	time.Sleep(10 * time.Second)

	src21Issuer.TransactOpts.Nonce = big.NewInt(int64(nonce + 2))
	src21Issuer.TransactOpts.Value = simulation.MinApply
	src21Issuer.TransactOpts.GasPrice = big.NewInt(common.DefaultMinGasPrice)
	src21Token.TransactOpts.GasPrice = big.NewInt(common.DefaultMinGasPrice)
	src21Token.TransactOpts.GasLimit = uint64(4000000)
	auth.GasPrice = big.NewInt(common.DefaultMinGasPrice)
	// get balance init src21 smart contract
	balance, err := src21Token.BalanceOf(simulation.MainAddr)
	if err != nil || balance.Cmp(simulation.Cap) != 0 {
		log.Fatal(err, "\tget\t", balance, "\twant\t", simulation.Cap)
	}
	fmt.Println("balance", balance, "mainAddr", simulation.MainAddr.Hex())

	// add trc20 list token sdx21 issuer
	_, err = src21Issuer.Apply(src21TokenAddr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("wait 10s to add token to list issuer")
	time.Sleep(10 * time.Second)

	//check src21 SMC balance
	balance, err = client.BalanceAt(context.Background(), src21IssuerAddr, nil)
	if err != nil || balance.Cmp(simulation.MinApply) != 0 {
		log.Fatal("can't get balance  in src21Issuer SMC: ", err, "got", balance, "wanted", simulation.MinApply)
	}

	//check balance fee
	balanceIssuerFee, err := src21Issuer.GetTokenCapacity(src21TokenAddr)
	if err != nil || balanceIssuerFee.Cmp(simulation.MinApply) != 0 {
		log.Fatal("can't get balance token fee in  smart contract: ", err, "got", balanceIssuerFee, "wanted", simulation.MinApply)
	}
}
