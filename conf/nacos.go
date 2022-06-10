package conf

import (
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

// NewNacos returns a configuration by the given NacosOptions.
func NewNacos(opt *NacosOptions) (*Configuration, error) {
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
	cache := []byte(content)

	var cfg Configuration
	if err := yaml.Unmarshal(cache, &cfg); err != nil {
		return nil, err
	}

	cfg.cache = cache
	return &cfg, nil
}
