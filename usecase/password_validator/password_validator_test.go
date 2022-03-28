package passwordvalidator

// Basic imports
import (
	"bitsports/mocks"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type DefaultPasswordValidatorTest struct {
	suite.Suite
	PassWordValidator PasswordValidatorI
	password          string
	assert            *assert.Assertions
	me                *mocks.EncrypterI
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (suite *DefaultPasswordValidatorTest) SetupTest() {
	// suite.me = new(MockEncrypter)
	suite.me = new(mocks.EncrypterI)
	suite.PassWordValidator = NewPasswordValidator(WithEncrypter(suite.me))
	suite.password = "MyPassword1@"
	suite.assert = assert.New(suite.T())

}

// All methods that begin with "Test" are run as tests within a
// suite.
func (suite *DefaultPasswordValidatorTest) TestEncryptPassword() {

	errExpected := fmt.Errorf("mocked error")
	suite.me.On("EncryptPassword", []byte(suite.password)).Return([]byte(""), errExpected)
	passwordEncrypted, err := suite.PassWordValidator.EncryptPassword([]byte(suite.password))
	suite.assert.ErrorIs(err, errExpected)

	suite.assert.Equal([]byte(""), passwordEncrypted)

}

func (suite *DefaultPasswordValidatorTest) TestEncryptPasswordFail() {

	hashed := []byte("hash")
	suite.me.On("EncryptPassword", []byte(suite.password)).Return(hashed, nil)

	passwordEncrypted, err := suite.PassWordValidator.EncryptPassword([]byte(suite.password))
	suite.assert.NoError(err)

	suite.assert.Equal(hashed, passwordEncrypted)
}

func (suite *DefaultPasswordValidatorTest) TestValidatePassword() {
	suite.assert.NoError(
		suite.PassWordValidator.ValidatePassword([]byte(suite.password)),
	)
}

func (suite *DefaultPasswordValidatorTest) TestWithoutUpperCasePassoword() {
	suite.password = "noupper1@"
	suite.assert.ErrorIs(suite.PassWordValidator.ValidatePassword(
		[]byte(suite.password),
	), ErrUpperCase)
}

func (suite *DefaultPasswordValidatorTest) TestWithoutNumberPassoword() {
	suite.password = "noupperA@"
	suite.assert.ErrorIs(suite.PassWordValidator.ValidatePassword(
		[]byte(suite.password),
	), ErrNumber)
}

func (suite *DefaultPasswordValidatorTest) TestWithoutSpecialCharactersPassoword() {
	suite.password = "noupperA1"
	suite.assert.ErrorIs(suite.PassWordValidator.ValidatePassword(
		[]byte(suite.password),
	), ErrSpecialCharacters)
}

func (suite *DefaultPasswordValidatorTest) TestWithoutMinLenPassoword() {
	suite.password = "noup"
	suite.assert.ErrorIs(suite.PassWordValidator.ValidatePassword(
		[]byte(suite.password),
	), ErrMinLen)
}

//function that sum two values

func (suite *DefaultPasswordValidatorTest) TestComparePassword() {
	hash := "hash"
	suite.me.On("ComparePassword", []byte(suite.password), []byte(hash)).Return(nil)

	suite.assert.NoError(suite.PassWordValidator.ComparePassword([]byte(suite.password), []byte(hash)))

}

func (suite *DefaultPasswordValidatorTest) TestComparePasswordFail() {
	hash := "hash"
	errExpected := fmt.Errorf("mocked error")
	suite.me.On("ComparePassword", []byte(suite.password), []byte(hash)).Return(errExpected)

	suite.assert.ErrorIs(suite.PassWordValidator.ComparePassword([]byte(suite.password), []byte(hash)), errExpected)

}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestEncryptSuite(t *testing.T) {
	suite.Run(t, new(DefaultPasswordValidatorTest))
}
