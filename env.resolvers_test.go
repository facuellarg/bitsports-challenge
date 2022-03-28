package bitsports

import (
	"bitsports/ent"
	"bitsports/ent/enttest"
	"bitsports/mocks"
	"context"
	"errors"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DefaultValues struct {
	suite.Suite
	client   *ent.Client
	assert   *assert.Assertions
	resolver *Resolver
	pv       *mocks.PasswordValidatorI
}

func (suite *DefaultValues) SetupTest() {

	suite.client = enttest.Open(suite.T(), "sqlite3", "file::memory:?cache=shared&_fk=1")
	suite.pv = new(mocks.PasswordValidatorI)
	suite.resolver = &Resolver{
		client:            suite.client,
		passwordValidator: suite.pv,
	}
	suite.assert = assert.New(suite.T())
}

func (suite *DefaultValues) TestCreateUser() {
	user := suite.client.User.Create().
		SetEmail("email@test.com").SetName("test").SetPassword("test").SaveX(context.TODO())
	suite.assert.Equal(user.Email, "email@test.com")
}

func (suite *DefaultValues) TestResolverCreateuser() {
	userInput := &UserInputSingUp{
		Email:    "email@testing",
		Name:     "test",
		Password: "test1@SDASD",
	}
	suite.pv.On("ValidatePassword", []byte(userInput.Password)).Return(nil)
	suite.pv.On("EncryptPassword", []byte(userInput.Password)).Return([]byte("test"), nil)

	user, err := suite.resolver.Mutation().CreateUser(context.TODO(), *userInput)
	suite.assert.Nil(err)
	suite.assert.Equal(user.Email, "email@testing")
	suite.assert.Equal(user.Name, "test")
}

func (suite *DefaultValues) TestResolverCreateUserFail() {
	userInput := &UserInputSingUp{
		Email:    "email@testing",
		Name:     "test",
		Password: "test1@SDASD",
	}

	errExpected := errors.New("test error")
	suite.pv.On("ValidatePassword", []byte(userInput.Password)).Return(errExpected)

	user, err := suite.resolver.Mutation().CreateUser(context.TODO(), *userInput)
	suite.assert.NotNil(err)
	suite.assert.Nil(user)
}

func TestResolver(t *testing.T) {
	suite.Run(t, new(DefaultValues))
}
