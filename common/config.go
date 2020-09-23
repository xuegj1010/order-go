package common

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

type Cfg struct {
	MySQL string `yaml:"MySQL"`
}

func LoadCfg() *Cfg {
	file, err := os.Open("./configs/application.yml")
	if err != nil {
		panic(err.Error())
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	cfg := Cfg{}
	err = yaml.Unmarshal(bytes, &cfg)
	if err != nil {
		panic(err)
	}
	return &cfg
}
