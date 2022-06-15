package config

type (
	Config struct {
		Database database
	}

	database struct {
		Url      string `yaml:"url"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Dbname   string `yaml:"dbname"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}
)
