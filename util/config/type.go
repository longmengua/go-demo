package config

type DbConfig struct {
	Host     string `yaml:"Host"`
	Port     string `yaml:"Port"`
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
	DB_Name  string `yaml:"DB_Name"`
}

type Struct struct {
	RootPath   string
	DB         DbConfig `yaml:"DB"`
	MIGRATE_UP bool     `yaml:"MIGRATE_UP"`
}
