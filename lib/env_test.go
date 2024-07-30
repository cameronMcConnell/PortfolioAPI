package lib

import (
	"testing"
	"github.com/joho/godotenv"
)

func TestReadEnv(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		t.Error(err)
	}

	got := ReadEnv("SERVER_ADDRESS")

	want := ":8080"

	if got != want {
		t.Errorf("Got %s, want %s", got, want)
	}
}