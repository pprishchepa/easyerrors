package easyerrors

import (
	goerrors "github.com/go-errors/errors"
)

func New(text string) error {
	return goerrors.New(text)
}

func Errorf(format string, a ...any) error {
	return goerrors.Errorf(format, a...)
}

func As(err error, target any) bool {
	return goerrors.As(err, target)
}

func Is(err error, target error) bool {
	return goerrors.Is(err, target)
}

func Join(errs ...error) error {
	return goerrors.Join(errs...)
}

func Unwrap(err error) error {
	return goerrors.Join(err)
}
