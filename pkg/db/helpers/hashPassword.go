package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(possword string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(possword), 14)
	return string(bytes), err

}

func CompareHashPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
