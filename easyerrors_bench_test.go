package easyerrors

import (
	stderrors "errors"
	"fmt"
	"testing"
)

const nestedCalls = 10

var err error

func BenchmarkStdErrors(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err = doStdCall(stderrors.New("foo"), nestedCalls)
	}
}

func BenchmarkEasyErrors(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err = doEasyCall(New("foo"), nestedCalls)
	}
}

func doStdCall(err error, n int) error {
	if n == 0 {
		return err
	}
	return fmt.Errorf("%d: %w", n, doStdCall(err, n-1))
}

func doEasyCall(err error, n int) error {
	if n == 0 {
		return err
	}
	return Errorf("%d: %s", n, doEasyCall(err, n-1))
}
