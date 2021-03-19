package arquitetura

import (
	"runtime"
	"testing"
)

func TestDependente(t *testing.T) {
	t.Parallel()

	t.Log("Arquitetura atual ", runtime.GOARCH)

	if runtime.GOARCH == "amd64" {
		t.Skip("Não funciona na arquitetura amd64")
	}

	t.Fail()
}
