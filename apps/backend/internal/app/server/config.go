package server

type Config struct {
	Port           int
	DSN            string
	AllowedOrigins []string
}

func InitConfig() Config {
	return Config{
		Port:           8080,
		DSN:            "postgres://postgres:postgres@postgres:5432/postgres?sslmode=disable&client_encoding=UTF8",
		AllowedOrigins: []string{"http://localhost", "http://tramondi.ru"},
	}
}
