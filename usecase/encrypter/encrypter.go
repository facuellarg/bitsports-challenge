package encrypter

import "golang.org/x/crypto/bcrypt"

type EncrypterI interface {
	EncryptPassword(password []byte) ([]byte, error)
	ComparePassword(password, hash []byte) error
}

type bcryptEncrypter struct{}

func NewBcryptEncripter() EncrypterI {
	return &bcryptEncrypter{}
}

func (e *bcryptEncrypter) EncryptPassword(password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
}

func (e *bcryptEncrypter) ComparePassword(password, hash []byte) error {
	return bcrypt.CompareHashAndPassword(hash, password)
}
