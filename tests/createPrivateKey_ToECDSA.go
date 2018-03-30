package main

import (
	"github.com/ethereum/go-ethereum/crypto"
	"fmt"
)

func main() {
	str := "hello"
	bs := []byte(str)
	sha3bs:= [32]byte(crypto.Keccak256Hash(bs[:]))
	key,_:= crypto.ToECDSA(sha3bs[:])
	fmt.Println(key)


}