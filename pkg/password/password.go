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

//
//func Verify(hashedPassword string, password string, algorithm string) (string, error) {
//	if algorithm == "bcrypt" {
//		hash, err := bcryptHash([]byte(password), 10)
//		if err != nil {
//			return "", err
//		}
//
//		return hash, nil
//	}
//
//	return "", nil
//}
