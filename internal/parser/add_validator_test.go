package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddValidator(t *testing.T) {
	successCase := []string{"arg1", "arg2"}
	errorCase := []string{"arg1"}

	validator := addValidator{}
	err := validator.Validate(successCase)

	assert.NoError(t, err)

	anotherErr := validator.Validate(errorCase)

	assert.EqualError(t, anotherErr, "CLIENT_ERROR invalid number of arguments\r\n")
}
