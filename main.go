package main

import (
    "encoding/json"
    "log"
    "github.com/tiny-ci/core/parser"
    "github.com/tiny-ci/core/pipe"
    "github.com/tiny-ci/core/schema"
)

func main() {
    configContent, err := pipe.Fetch(&pipe.GitRef{
        Name: "master",
        URL:  "https://github.com/tiny-ci/example",
        Hash: "961240b68021e7d8ecb53c0b5b4e1e8e097fb319",
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

    matchedJobs := pipe.Filter(config, "master")
}
