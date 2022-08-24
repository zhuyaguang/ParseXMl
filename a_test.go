package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestStringsJoin(t *testing.T) {
	s := []string{"a", "b", "c"}
	fmt.Println(strings.Join(s, ","))
}
