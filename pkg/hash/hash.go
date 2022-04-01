package hash

import (
	"crypto"
	_ "crypto/md5"
	_ "crypto/sha1"
	_ "crypto/sha256"
	_ "crypto/sha512"
	"encoding/hex"
	"errors"
	"fmt"
	_ "golang.org/x/crypto/blake2b"
	_ "golang.org/x/crypto/blake2s"
	_ "golang.org/x/crypto/md4"
	_ "golang.org/x/crypto/ripemd160"
	_ "golang.org/x/crypto/sha3"
	"sort"
)

func GetSupportedHashingAlgorithms() []string {
	ha := hashAlgorithms()
	a := make([]string, 0, len(ha))

	for key, _ := range ha {
		a = append(a, key)
	}

	sort.Strings(a)

	return a
}

func hashAlgorithms() map[string]crypto.Hash {
	m := make(map[string]crypto.Hash)

	m["MD5"] = crypto.MD5
	m["MD4"] = crypto.MD4
	m["SHA1"] = crypto.SHA1
	m["SHA224"] = crypto.SHA224
	m["SHA256"] = crypto.SHA256
	m["SHA384"] = crypto.SHA384
	m["SHA512"] = crypto.SHA512
	m["RIPEMD160"] = crypto.RIPEMD160
	m["SHA3_224"] = crypto.SHA3_224
	m["SHA3_256"] = crypto.SHA3_256
	m["SHA3_384"] = crypto.SHA3_384
	m["SHA3_512"] = crypto.SHA3_512
	m["SHA512_224"] = crypto.SHA3_512
	m["SHA512_256"] = crypto.SHA3_512
	m["BLAKE2S_256"] = crypto.BLAKE2s_256
	m["BLAKE2B_256"] = crypto.BLAKE2b_256
	m["BLAKE2B_384"] = crypto.BLAKE2b_384
	m["BLAKE2B_512"] = crypto.BLAKE2b_512

	return m
}

func Execute(value []byte, algorithm string) (string, error) {
	var (
		ha = hashAlgorithms()
	)

	if val, ok := ha[algorithm]; ok {
		return cryptoHash(value, val), nil
	}

	return "", errors.New(fmt.Sprintf("%v is not supported", algorithm))
}

func cryptoHash(data []byte, hash crypto.Hash) string {
	h := hash.New()
	h.Write(data)

	return hex.EncodeToString(h.Sum(nil))
}
