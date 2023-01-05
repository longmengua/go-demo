package config

type DbConfig struct {
	Host     string `yaml:"Host"`
	Port     string `yaml:"Port"`
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
}

type Struct struct {
	DB DbConfig `yaml:"DB"`
}
