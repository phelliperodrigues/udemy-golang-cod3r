package arquitetura

import (
	"runtime"
	"testing"
)

func TestDependente(t *testing.T) {

	if runtime.GOARCH == "amd64" {
		t.Skip("Não funciona em arquiteura amd64")
	}

	if runtime.GOARCH == "arm" {
		t.Skip("Não funciona em arquiteura arm")
	}

	if runtime.GOARCH == "arm64" {
		t.Skip("Não funciona em arquiteura arm64")
	}

	if runtime.GOARCH == "386" {
		t.Skip("Não funciona em arquiteura 386")
	}

	t.Fail()
}
