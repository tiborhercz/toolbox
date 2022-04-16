package password

func SupportedHashingAlgorithms() string {
	return "bcrypt"
}

func Hash(password string, algorithm string, rounds int) (string, error) {
	var bcrypt Password
	bcrypt = Bcrypt{}

	if algorithm == "bcrypt" {
		hash, err := bcrypt.Hash([]byte(password), rounds)
		if err != nil {
			return "", err
		}

		return hash, nil
	}

	return "", nil
}

func Verify(hashedPassword string, password string) error {
	var bcrypt Password
	bcrypt = Bcrypt{}

	err := bcrypt.Verify([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}

	return nil
}
