package passwordvalidator

import (
	"bitsports/usecase/encrypter"
	"errors"
	"fmt"
	"regexp"
)

type PasswordValidatorI interface {
	encrypter.EncrypterI
	ValidatePassword(s []byte) error
}

type passwordValidator struct {
	encrypter.EncrypterI
	rules     map[string]error
	minLen    int
	errMinLen error
}

type Options func(pv *passwordValidator)

func WithMinLen(minLen int) Options {
	return func(pv *passwordValidator) {
		pv.minLen = minLen
		pv.errMinLen = fmt.Errorf("must contain at least %d characters", minLen)
	}
}

func WithRules(rules map[string]error) Options {
	return func(pv *passwordValidator) {
		pv.rules = rules
	}
}

func WithEncrypter(e encrypter.EncrypterI) Options {
	return func(pv *passwordValidator) {
		pv.EncrypterI = e
	}
}

func NewPasswordValidator(options ...Options) *passwordValidator {
	defaultPasswordvalidator := &passwordValidator{
		rules:      regexRules,
		minLen:     8,
		errMinLen:  ErrMinLen,
		EncrypterI: encrypter.NewBcryptEncripter(),
	}

	for _, option := range options {
		option(defaultPasswordvalidator)
	}
	return defaultPasswordvalidator
}

var (
	ErrUpperCase         = errors.New("must contain uppercase letter")
	ErrNumber            = errors.New("must contain a number")
	ErrSpecialCharacters = errors.New("must contain special character")
	ErrMinLen            = errors.New("must contain at least 8 characters")
	regexRules           = map[string]error{
		`[A-Z]`:              ErrUpperCase,         //uppercase
		`[0-9]`:              ErrNumber,            //numbers
		`[!@#$%^&()_+?><>?]`: ErrSpecialCharacters, //specialcharacters
	}
)

func (pv *passwordValidator) ValidatePassword(s []byte) error {
	if len(s) < pv.minLen {
		return pv.errMinLen
	}
	for rule, err := range pv.rules {
		r := regexp.MustCompile(rule)
		if !r.Match(s) {
			return err
		}
	}
	return nil
}
