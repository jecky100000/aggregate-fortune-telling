package ay

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Yaml struct {
	Domain string `yaml:"domain"`
	Mysql  YamlMysql
	Sms    YamlSms
}

type YamlMysql struct {
	Localhost string `yaml:"localhost"`
	Port      string `yaml:"port"`
	User      string `yaml:"user"`
	Password  string `yaml:"password"`
	Database  string `yaml:"database"`
}

type YamlSms struct {
	Ak           string `yaml:"ak"`
	Sk           string `yaml:"sk"`
	Sign         string `yaml:"sign"`
	TemplateCode string `yaml:"template_code"`
	Database     string `yaml:"database"`
}

func (c *Yaml) GetConf() *Yaml {
	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}
