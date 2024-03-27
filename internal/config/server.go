package config

// Server ..
type Server struct {
	Port int `env:"SERVER_PORT" envDefault:"8000"`
}
