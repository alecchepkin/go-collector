package app

import (
	"github.com/caarlos0/env"
	"sync"
)

type config struct {
	DbDsn          string `env:"DB_DSN" envDefault:"host=localhost port=5432 user=postgres dbname=postgres password="`
	DbServiceDsn   string `env:"DB_SERVICE_DSN" envDefault:"host=localhost port=5432 user=postgres dbname=postgres password="`
	ListenAddr     string `env:"LISTEN_ADDR" envDefault:"127.0.0.1:8080"`
	MailServerAddr string `env:"MAIL_SERVER_ADDR" envDefault:"127.0.0.1:25"`
	ErrMailFrom    string `env:"ERR_MAIL_FROM" envDefault:"sender@domain"`
	ErrMailTo      string `env:"ERR_MAIL_TO" envDefault:"recipient@domain"`
	LogFormat      string `env:"LOG_FORMAT" envDefault:"text"`
	LogLevel       string `env:"LOG_LEVEL" envDefault:"debug"`
	TokenProxyHost string `env:"TOKEN_PROXY_URL" envDefault:"http://mt-proxy.mobio.ru"`
	MyTargetHost   string `env:"MY_TARGET_HOST" envDefault:"https://target.my.com"`
}

var (
	confOnce sync.Once
	conf     *config
)

func Config() *config {
	confOnce.Do(func() {
		conf = &config{}
		err := env.Parse(conf)
		if err != nil {
			panic(err)
		}
	})
	return conf
}
