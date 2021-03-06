package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type mysql struct {
	User     string `yaml:"user"`
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
	Name     string `yaml:"name"`
}

type cache struct {
	Enable bool     `yaml:"enable"`
	List   []string `yaml:"list"`
}

type person struct {
	Name  string   `yaml:"name"`
	Age   int64    `yaml:"age"`
	Names []string `yaml:"names"`
}

type Yaml struct {
	Mysql   mysql    `yaml:"mysql"`
	Cache   cache    `yaml:"cache"`
	Persons []person `yaml:"persons"`
}

func main() {
	conf := new(Yaml)
	yamlFile, err := ioutil.ReadFile("test.yaml")
	log.Println("yamlFile:", string(yamlFile))
	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	log.Println("conf", conf)
}
