# hcp-vault-tezos-util
Tool to convert HCP Vault transit key to Tezos format

1. Set base64 encoded HCP Vault transit key
```
tezKey, err := utils.GetTezosKeyFromBase64("3RqoUFWP2DKoXGJvsJ1sNXx2S2Vq5MpU1g57LCltEpw=")
```

2. Run ``` go run main.go ```
   
3. Get a public key in tezos format with an associated address
```
Tezos Public Key: edpkvKbkAxSa7JJJ8j95xJxpg3LoCjuNgj1JzKtU9NttjdgabyTJzZ
Tezos Public Key Address: tz1eXM1uGi5THR7Aj8VnkteA5nrBmPyKAufM
```
