package conf

import (
	"errors"
	"os"
	"reflect"
	"strconv"

	"github.com/gocurr/good/consts"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"gopkg.in/yaml.v3"
)

// NacosOptions the options to connect to nacos.
type NacosOptions struct {
	Ip               string
	Port             uint64
	NamespaceId      string
	DataId           string
	Group            string
	Scheme           string
	ContextPath      string
	TimeoutMs        uint64 // millisecond
	LoadCacheAtStart bool
}

var errNacosOptions = errors.New("conf: bad Nacos options")

// NewNacos returns a configuration by the given NacosOptions.
func NewNacos(opt *NacosOptions) (*Configuration, error) {
	if opt == nil || reflect.ValueOf(*opt).IsZero() {
		var err error
		if opt, err = readEnv(); err != nil {
			return nil, err
		}
	}
	if opt.Scheme == "" {
		opt.Scheme = constant.DEFAULT_SERVER_SCHEME
	}
	if opt.ContextPath == "" {
		opt.ContextPath = constant.DEFAULT_CONTEXT_PATH
	}
	if opt.TimeoutMs == 0 {
		opt.TimeoutMs = 5000
	}

	sc := []constant.ServerConfig{
		*constant.NewServerConfig(
			opt.Ip,
			opt.Port,
			constant.WithScheme(opt.Scheme),
			constant.WithContextPath(opt.ContextPath)),
	}

	cc := constant.NewClientConfig(
		constant.WithNamespaceId(opt.NamespaceId),
		constant.WithTimeoutMs(opt.TimeoutMs),
		constant.WithNotLoadCacheAtStart(!opt.LoadCacheAtStart),
	)

	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		return nil, err
	}

	content, err := client.GetConfig(vo.ConfigParam{
		DataId: opt.DataId,
		Group:  opt.Group,
	})
	if err != nil {
		return nil, err
	}

	if content == "" {
		return nil, errNacosOptions
	}
	cache := []byte(content)

	var cfg Configuration
	if err := yaml.Unmarshal(cache, &cfg); err != nil {
		return nil, err
	}

	if reflect.ValueOf(cfg).IsZero() {
		return nil, errNacosOptions
	}

	cfg.cache = cache
	return &cfg, nil
}

func readEnv() (*NacosOptions, error) {
	ip := os.Getenv(consts.NacosIp)
	port := os.Getenv(consts.NacosPort)
	dataId := os.Getenv(consts.NacosDataId)
	group := os.Getenv(consts.NacosGroup)
	namespaceId := os.Getenv(consts.NacosNamespaceId)

	if ip == "" || port == "" ||
		dataId == "" || group == "" || namespaceId == "" {
		return nil, errNacosOptions
	}

	portUint, err := strconv.ParseUint(port, 10, 64)
	if err != nil {
		return nil, err
	}

	var timeoutMsUint uint64
	if timeoutMs := os.Getenv(consts.NacosTimeoutMs); timeoutMs != "" {
		timeoutMsUint, err = strconv.ParseUint(timeoutMs, 10, 64)
		if err != nil {
			return nil, err
		}
	}
	scheme := os.Getenv(consts.NacosScheme)
	contextPath := os.Getenv(consts.NacosContextPath)

	var loadCacheAtStartBool bool
	if loadCacheAtStart := os.Getenv(consts.NacosLoadCacheAtStart); loadCacheAtStart != "" {
		loadCacheAtStartBool, err = strconv.ParseBool(loadCacheAtStart)
		if err != nil {
			return nil, err
		}
	}

	return &NacosOptions{
		Ip:               ip,
		Port:             portUint,
		DataId:           dataId,
		Group:            group,
		NamespaceId:      namespaceId,
		TimeoutMs:        timeoutMsUint,
		Scheme:           scheme,
		ContextPath:      contextPath,
		LoadCacheAtStart: loadCacheAtStartBool,
	}, nil
}
