package grpcconf

import (
	"github.com/gocurr/good/consts"
	"github.com/gocurr/good/pre"
	"reflect"
)

func ServerPort(i interface{}) (bool, int) {
	if i == nil {
		return false, 0
	}

	var c reflect.Value
	if reflect.TypeOf(i).Kind() == reflect.Ptr {
		c = reflect.ValueOf(i).Elem()
	} else {
		c = reflect.ValueOf(i)
	}

	grpcField := c.FieldByName(pre.GRPC)
	if !grpcField.IsValid() {
		return false, 0
	}

	enableField := grpcField.FieldByName(consts.Enable)
	if !enableField.IsValid() {
		return false, 0
	}

	enable := enableField.Bool()
	if !enable {
		return false, 0
	}

	serverField := grpcField.FieldByName(consts.Server)
	if !serverField.IsValid() {
		return false, 0
	}

	portField := serverField.FieldByName(consts.Port)
	if !portField.IsValid() {
		return false, 0
	}

	port := portField.Int()
	if port == 0 {
		return false, 0
	}

	return true, int(port)
}

func ClientAddrTimeout(i interface{}) (bool, string, int) {
	if i == nil {
		return false, "", 0
	}

	var c reflect.Value
	if reflect.TypeOf(i).Kind() == reflect.Ptr {
		c = reflect.ValueOf(i).Elem()
	} else {
		c = reflect.ValueOf(i)
	}

	grpcField := c.FieldByName(pre.GRPC)
	if !grpcField.IsValid() {
		return false, "", 0
	}

	enableField := grpcField.FieldByName(consts.Enable)
	if !enableField.IsValid() {
		return false, "", 0
	}

	enable := enableField.Bool()
	if !enable {
		return false, "", 0
	}

	clientField := grpcField.FieldByName(consts.Client)
	if !clientField.IsValid() {
		return false, "", 0
	}

	addrField := clientField.FieldByName(consts.Addr)
	if !addrField.IsValid() {
		return false, "", 0
	}

	addr := addrField.String()
	if addr == "" {
		return false, "", 0
	}

	timeoutField := clientField.FieldByName(consts.Timeout)
	if !timeoutField.IsValid() {
		return false, "", 0
	}

	timeout := timeoutField.Int()

	return true, addr, int(timeout)
}
