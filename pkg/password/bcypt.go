package password

import "golang.org/x/crypto/bcrypt"

type Password interface {
	Hash(password []byte, cost int) (string, error)
	Verify(hashedPassword []byte, password []byte) error
}

type Bcrypt struct {
	Password
}

func (b Bcrypt) Hash(password []byte, cost int) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(password, cost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (b Bcrypt) Verify(hashedPassword []byte, password []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	if err != nil {
		return err
	}

	return nil
}
