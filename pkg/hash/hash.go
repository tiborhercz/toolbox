package hash

type Hashing interface {
	Hash(password []byte, cost int) (string, error)
}

func Hash(algorithm string) (string, error) {
	return "", nil
}
