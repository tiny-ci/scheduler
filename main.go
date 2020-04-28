package main

import (
    "encoding/json"
    "log"
    "github.com/tiny-ci/core/parser"
    "github.com/tiny-ci/core/pipeconf"
    "github.com/tiny-ci/core/schema"
)

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

    config, err := parser.ParsePipeConfig(configContent.Bytes())
    if err != nil {
        log.Println("parser error")
        log.Fatal(err)
    }

    var marshalledConfig []byte
    marshalledConfig, err = json.Marshal(config)
    if err != nil {
        log.Println("marshall error")
        log.Fatal(err)
    }

    err = schema.ValidateConfig(marshalledConfig)
    if err != nil {
        log.Println("validation error")
        log.Fatal(err)
    }

    log.Println(config.Jobs[0].When.Tag)
}
