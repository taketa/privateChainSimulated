package main

import "fmt"
import (
	"github.com/ethereum/go-ethereum/crypto"
	//"log"
	//"github.com/ethereum/go-ethereum/crypto/sha3"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
str := "hello"
bs := []byte(str)
sha3:= [32]byte(crypto.Keccak256Hash(bs[:]))
fmt.Println("sha3===",sha3)
fmt.Println("HEX===",common.Bytes2Hex(sha3[:]))


	//key,err:= crypto.ToECDSA(sha3Key[:])
	//if err!=nil{
	//	log.Fatalln(err)
	//}
	//fmt.Println("ToECDSA",key)
}