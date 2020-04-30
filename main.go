package main

import (
    "encoding/json"
    "log"
    "github.com/tiny-ci/core/db"
    "github.com/tiny-ci/core/parser"
    "github.com/tiny-ci/core/pipe"
    "github.com/tiny-ci/core/schema"
)

func main() {
    ref := "master"
    configContent, err := pipe.Fetch(&pipe.GitRef{
        Name: ref,
        URL:  "https://github.com/caiertl/tmp",
        Hash: "64954853e0169c930024d24563883f9ea59a0c9a",
        IsTag: false,
    })

    if err != nil {
        log.Println("fetch error")
        log.Fatal(err)
    }

    if err == nil && configContent == nil {
        log.Println("pipe config not found")
        return
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

    matchedJobs, err := pipe.Filter(config, ref)
    if err != nil {
        log.Println("filter error")
        log.Fatal(err)
    }

    log.Println(matchedJobs)

    rdb := db.RedisDatabase{}
    err = rdb.Connect()
    if err != nil {
        log.Println("redis error")
        log.Fatal(err)
    }
}
