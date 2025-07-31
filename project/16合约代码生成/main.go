package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/test/init_project/store"
)

func main() {

	// 任务 2：合约代码生成 任务目标
	// 使用 abigen 工具自动生成 Go 绑定代码，用于与 Sepolia 测试网络上的智能合约进行交互。
	//  具体任务
	// 编写智能合约
	// 使用 Solidity 编写一个简单的智能合约，例如一个计数器合约。
	// 编译智能合约，生成 ABI 和字节码文件。
	// 使用 abigen 生成 Go 绑定代码
	// 安装 abigen 工具。
	// 使用 abigen 工具根据 ABI 和字节码文件生成 Go 绑定代码。
	// 使用生成的 Go 绑定代码与合约交互
	// 编写 Go 代码，使用生成的 Go 绑定代码连接到 Sepolia 测试网络上的智能合约。
	// 调用合约的方法，例如增加计数器的值。
	// 输出调用结果。
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/5b20ca5b9ddf4b4c851351696ad9d564")
	if err != nil {
		log.Fatal(err)
	}

	// privateKey, err := crypto.GenerateKey()
	// privateKeyBytes := crypto.FromECDSA(privateKey)
	// privateKeyHex := hex.EncodeToString(privateKeyBytes)
	// fmt.Println("Private Key:", privateKeyHex)
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

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	input := "1.0"
	address, tx, instance, err := store.DeployStore(auth, client, input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex())   // 0x9930918653beE8E6817396E2CBB65E8bF9101E62
	fmt.Println(tx.Hash().Hex()) // 0x98b089b3d039d1bf2ed3c1630c1b3239213a27457656c4a4a47647dd4f67500b

	_ = instance

	storeContract, err := store.NewStore(common.HexToAddress("0x9930918653beE8E6817396E2CBB65E8bF9101E62"), client)
	if err != nil {
		log.Fatal(err)
	}

	// privateKey, err := crypto.HexToECDSA("<your private key>")
	if err != nil {
		log.Fatal(err)
	}

	var key [32]byte
	var value [32]byte

	copy(key[:], []byte("demo_save_key"))
	copy(value[:], []byte("demo_save_value11111"))

	opt, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111))
	if err != nil {
		log.Fatal(err)
	}
	tx, err = storeContract.SetItem(opt, key, value)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tx hash:", tx.Hash().Hex())

	callOpt := &bind.CallOpts{Context: context.Background()}
	valueInContract, err := storeContract.Items(callOpt, key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("value:", value)
	fmt.Println("valueInContract:", valueInContract)
	fmt.Println("is value saving in contract equals to origin value:", valueInContract == value)
}
