package config

type Server struct {
	Name string `json:"name" mapstructure:"name"`
	Port int    `json:"port" mapstructure:"port"`
}

type DB struct {
	Host     string `json:"host,omitempty"                     mapstructure:"host"`
	Username string `json:"username,omitempty"                 mapstructure:"username"`
	Password string `json:"-"                                  mapstructure:"password"`
	Database string `json:"database"                           mapstructure:"database"`
}

type Config struct {
	Server *Server `json:"server"  mapstructure:"server"`
	DB     *DB     `json:"db"  mapstructure:"db"`
}

var (
	defaultServer = &Server{
		Name: "world",
		Port: 8080,
	}

	defaultDB = &DB{
		Host:     "127.0.0.1",
		Username: "root",
		Password: "root",
		Database: "world",
	}
)

func New() *Config {
	return &Config{
		Server: defaultServer,
		DB:     defaultDB,
	}
}

func (c *Config) Validate() {}
