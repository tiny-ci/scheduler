package main

import (
	"log"
	"gopkg.in/yaml.v2"
	"github.com/tiny-ci/core/pipeconf"
)

type Config struct {
    Jobs []struct {
        Name  string      `yaml:"name"`
        Image string      `yaml:"image"`
        Steps interface{} `yaml:"steps"`
        When  struct {
            Branch interface{} `yaml:"branch"`
            Tag    string      `yaml:"tag"`
        }
    }
}

func main() {
    configContent, err := pipeconf.Fetch(&pipeconf.GitRef{
        Name: "master",
        URL:  "https://github.com/tiny-ci/example",
        Hash: "dba8a3250ff364a8a1ccfe0ca0b1bdeb43adadcb",
        IsTag: false,
    })

    if err != nil {
        log.Fatal(err)
    }

    config := Config{}
    err = yaml.Unmarshal(configContent.Bytes(), &config)
    if err != nil {
        log.Fatal(err)
    }

    log.Println(config.Jobs[0].When.Tag)
}
