package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	// 查询区块
	// 编写 Go 代码，使用 ethclient 连接到 Sepolia 测试网络。
	// 实现查询指定区块号的区块信息，包括区块的哈希、时间戳、交易数量等。
	// 输出查询结果到控制台。
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/5b20ca5b9ddf4b4c851351696ad9d564")
	if err != nil {
		log.Fatal(err)
	}

	blockNumber := big.NewInt(5671744)

	// 查询区块头
	header, err := client.HeaderByNumber(context.Background(), blockNumber)
	fmt.Println("区块号", header.Number.Uint64())      // 5671744
	fmt.Println("区块时间戳", header.Time)               // 1712798400
	fmt.Println("区块难度", header.Difficulty.Uint64()) // 0
	fmt.Println("区块哈希", header.Hash().Hex())        // 0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5

	if err != nil {
		log.Fatal(err)
	}

	// 查询完整区块
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("区块号", block.Number().Uint64())      // 5671744
	fmt.Println("区块时间戳", block.Time())               // 1712798400
	fmt.Println("区块难度", block.Difficulty().Uint64()) // 0
	fmt.Println("区块哈希", block.Hash().Hex())          // 0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5
	fmt.Println("区块交易数量", len(block.Transactions())) // 70

	// 查询区块的交易数目
	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("区块交易数量", count) // 70

	// 发送交易
	// 准备一个 Sepolia 测试网络的以太坊账户，并获取其私钥。
	// 编写 Go 代码，使用 ethclient 连接到 Sepolia 测试网络。
	// 构造一笔简单的以太币转账交易，指定发送方、接收方和转账金额。
	// 对交易进行签名，并将签名后的交易发送到网络。
	// 输出交易的哈希值。
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("e31b9ce0393ea6812b08726092069390e01e53d1cd62e598a8e3915c10f6af42")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(1000000000000000) // in wei (0.001 eth)
	gasLimit := uint64(21000)             // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0x22d5D0526fE1bdF74ee567682A34AC87E9E31889")
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("交易的哈希值: %s", signedTx.Hash().Hex())
}
