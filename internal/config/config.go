package config

type (
	Config struct {
		Postgres Pgsql `yaml:"pgsql"`
	}

	Pgsql struct {
		Url      string `yaml:"url"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Dbname   string `yaml:"dbname"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Sslmode  string `yaml:"sslmode"`
	}
)
