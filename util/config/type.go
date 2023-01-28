package config

type DbConfig struct {
	Host      string `yaml:"Host"`
	Port      string `yaml:"Port"`
	User      string `yaml:"User"`
	Password  string `yaml:"Password"`
	DB_Name   string `yaml:"DB_Name"`
	Migration bool   `yaml:"MIGRATION"`
}

type SERVER struct {
	Port int `yaml:"Port"`
}

type Struct struct {
	Port      int
	Root_Path string
	DB        DbConfig `yaml:"DB"`
	SERVER    SERVER   `yaml:"SERVER"`
}
