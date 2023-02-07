package test

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	taskBeforeAll()
	cryptoBeforeAll()
	reconBeforeAll()
	kittenBeforeAll()
	apiBeforeAll()
	configBeforeAll()
	exitCode := m.Run()
	os.Exit(exitCode)
}
