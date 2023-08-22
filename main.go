package main

import (
	"fmt"

	"github.com/OneOf-Inc/hcp-vault-tezos-util/utils"
)

func main() {
	tezKey, err := utils.GetTezosKeyFromBase64("3RqoUFWP2DKoXGJvsJ1sNXx2S2Vq5MpU1g57LCltEpw=")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Tezos Public Key:", tezKey)
	fmt.Println("Tezos Public Key Address:", tezKey.Address())
}
