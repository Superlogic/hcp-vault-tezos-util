package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"

	"blockwatch.cc/tzgo/base58"
	"blockwatch.cc/tzgo/tezos"
)

func GetTezosKeyFromBase64(encodedPubKey string) (*tezos.Key, error) {
	ed25519PublicKeyBytes, err := base64.StdEncoding.DecodeString(encodedPubKey)
	if err != nil {
		return nil, err
	}

	return convertPubKeyToTezosFormat(ed25519PublicKeyBytes)
}

func GetTezosKeyFromHex(hashedPubKey string) (*tezos.Key, error) {
	ed25519PublicKeyBytes, err := hex.DecodeString(hashedPubKey)
	if err != nil {
		return nil, err
	}

	return convertPubKeyToTezosFormat(ed25519PublicKeyBytes)
}

func convertPubKeyToTezosFormat(ed25519PublicKeyBytes []byte) (*tezos.Key, error) {
	tezosPublicKeyPrefix := []byte{13, 15, 37, 217}
	tezosPublicKeyBytes := append(tezosPublicKeyPrefix, ed25519PublicKeyBytes...)
	checksum := calculateChecksum(tezosPublicKeyBytes)
	tezosPublicKeyWithChecksum := append(tezosPublicKeyBytes, checksum[:]...)

	base58EncodedPublicKey := base58.Encode(tezosPublicKeyWithChecksum)

	tezosKey, err := tezos.ParseKey(base58EncodedPublicKey)
	if err != nil {
		return nil, err
	}
	return &tezosKey, nil
}

// calculateChecksum: first four bytes of sha256^2
func calculateChecksum(input []byte) (cksum [4]byte) {
	h := sha256.Sum256(input)
	h2 := sha256.Sum256(h[:])
	copy(cksum[:], h2[:4])
	return
}
