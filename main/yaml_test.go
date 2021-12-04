package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"testing"
)

type T struct {
	F int `yaml:"a,omitempty"`
	B int `yaml:"b,omitempty"`
}

func Test_Yml(t *testing.T) {
	var v T
	_ = yaml.Unmarshal([]byte("a: 1\nb: x"), &v)
	fmt.Println(v)
}
