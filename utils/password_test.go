package utils

import (
	"fmt"
	"testing"
)

func TestHashPassword(t *testing.T) {
	pass := "123hihi"
	hashedPass, err := HashPassword(pass)
	if err != nil {
		t.Error(err)
	}
	fmt.Print(hashedPass)
	fmt.Println()
}
