package config

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

func (s *Struct) Init() {
	log.Println(s.RootPath)
	filePath := fmt.Sprintf("%s/env.yaml", s.RootPath)
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic("failed to connect database")
	}
	err = yaml.Unmarshal(yamlFile, s)
	if err != nil {
		log.Printf("Unmarshal: %v", err)
	}
}
