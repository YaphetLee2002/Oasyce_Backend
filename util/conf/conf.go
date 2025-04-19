package conf

import (
	"errors"
	"sync"

	"github.com/jinzhu/configor"
)

var globalConfig *Config

func GlobalConfig() *Config {
	var err error
	if globalConfig == nil {
		globalConfig, err = loadConfig()
		if err != nil {
			panic(err)
		}
	}
	return globalConfig
}

// Config contains all the configuration, including APPConfig.
type Config struct {
	API APIConfig `yaml:"api"`
	RPC RPCConfig `yaml:"rpc"`
	DB  DBConfig  `yaml:"db"`
	JWT JWTConfig `yaml:"jwt"`
}

type APIConfig struct {
	ENV   string `yaml:"env"`
	Hertz struct {
		Addr    string `yaml:"addr"`
		APIPath string `yaml:"api_path"`
		SSOPath string `yaml:"sso_path"`
	} `yaml:"hertz"`
	Kitex struct {
		Addr string `yaml:"addr"`
	} `yaml:"kitex"`
	ProxyHeaders []string `yaml:"proxy_headers"` // The names of HTTP headers that will be extracted from Hertz request and delivered to KiteX.
}

type RPCConfig struct {
	ENV string `yaml:"env"` // byted/local/volc
}

type DBConfig struct {
	Dialect  string `yaml:"dialect"`
	DSN      string `yaml:"dsn"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
	// Used for bytedance only.
	GlobalDSNRead  string `yaml:"global_dsn_read"`
	GlobalDSNWrite string `yaml:"global_dsn_write"`
	RepoDSNRead    string `yaml:"repo_dsn_read"`
	RepoDSNWrite   string `yaml:"repo_dsn_write"`
}
type JWTConfig struct {
	PublicKey  string `yaml:"public_key"`
	PrivateKey string `yaml:"private_key" kms:"private_key"`
}

var (
	configOnce sync.Once
	config     *Config
	configErr  error
)

func loadConfig() (*Config, error) {
	configOnce.Do(func() {
		cfg := new(Config)
		err := configor.New(&configor.Config{
			ENVPrefix: "CODE",
		}).Load(cfg, "conf/conf.yml")
		config = cfg
		configErr = err
		if configErr != nil {
			return
		}

		// Check DB config.
		if config.DB.Dialect != "mysql" {
			configErr = errors.New("invalid database dialect")
			return
		}
	})
	return config, configErr
}
