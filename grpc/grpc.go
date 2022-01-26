package grpc

import (
	"github.com/gocurr/good/consts"
	"github.com/gocurr/good/pre"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// ServerPort returns server.port in grpc field and
// reports the state of process which must be checked.
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

// AddrTimeout wraps Addr and Timeout of Client configuration.
type AddrTimeout struct {
	Addr    string
	Timeout int
}

// ClientAddrTimeout returns client.addr and client.timeout in grpc field and
// reports the state of process which must be checked.
func ClientAddrTimeout(i interface{}) (bool, map[string]AddrTimeout) {
	if i == nil {
		return false, nil
	}

	var c reflect.Value
	if reflect.TypeOf(i).Kind() == reflect.Ptr {
		c = reflect.ValueOf(i).Elem()
	} else {
		c = reflect.ValueOf(i)
	}

	grpcField := c.FieldByName(pre.GRPC)
	if !grpcField.IsValid() {
		return false, nil
	}

	enableField := grpcField.FieldByName(consts.Enable)
	if !enableField.IsValid() {
		return false, nil
	}

	enable := enableField.Bool()
	if !enable {
		return false, nil
	}

	clientField := grpcField.FieldByName(consts.Client)
	if !clientField.IsValid() {
		return false, nil
	}

	var result = make(map[string]AddrTimeout)
	iter := clientField.MapRange()
	for iter.Next() {
		name := iter.Key().String()
		c := iter.Value().String()

		split := strings.Split(c, ",")
		if len(split) == 2 {
			timeout, err := strconv.Atoi(split[1])
			if err != nil {
				continue
			}
			result[name] = AddrTimeout{
				Addr:    addrFromEnv(name, split[0]),
				Timeout: timeout,
			}
		}
	}

	return true, result
}

// addrFromEnv returns a new addr from environment variables.
func addrFromEnv(name, old string) string {
	const separator = "_"
	key := strings.ToUpper(strings.Join([]string{pre.GRPC, name, consts.Addr}, separator))
	newVal := os.Getenv(key)
	if newVal == "" {
		return old
	}
	return newVal
}
