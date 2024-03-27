package config

type Hash struct {
	Salt string `env:"salt" envDefault:"salt"`
}
