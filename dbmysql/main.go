package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Conf 配置文件的结构
type Conf struct {
	Host   string `yaml:"host"`
	User   string `yaml:"user"`
	Pwd    string `yaml:"pwd"`
	Dbname string `yaml:"dbname"`
}

// InitConf 初始化配置文件
func InitConf() *Conf {
	c := &Conf{}
	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}

func main() {
	conf := InitConf()

	fmt.Println(*conf)
}
