package kv

import (
	"errors"
	"testing"

	govalidator "github.com/go-playground/validator/v10"
)

func TestNewRejectsMissingRequiredOptions(t *testing.T) {
	_, err := New(Options{})
	if err == nil {
		t.Fatal("expected empty Options to be rejected")
	}

	var validationErrs govalidator.ValidationErrors
	if !errors.As(err, &validationErrs) {
		t.Fatalf("expected validator.ValidationErrors, got %T: %v", err, err)
	}
}
