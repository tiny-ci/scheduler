package main

import (
    "encoding/json"
    "log"
    "github.com/tiny-ci/core/db"
    "github.com/tiny-ci/core/parser"
    "github.com/tiny-ci/core/pipe"
    "github.com/tiny-ci/core/schema"
    "github.com/tiny-ci/core/types"
)

func main() {
    ntf := types.ApiNotification{
        PipelineId: "507f1f77bcf86cd799439011",
        Info: types.GitInfo{
            URL: "https://github.com/caiertl/tmp",
            RefName: "master",
            IsTag: false,
            CommitHash: "64954853e0169c930024d24563883f9ea59a0c9a",
        },
    }

    configContent, err := pipe.Fetch(&pipe.GitRef{
        Name:  ntf.Info.RefName,
        URL:   ntf.Info.URL,
        Hash:  ntf.Info.CommitHash,
        IsTag: ntf.Info.IsTag,
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

    matchedJobs, err := pipe.Filter(config, ntf.Info.RefName)
    if err != nil {
        log.Println("filter error")
        log.Fatal(err)
    }

    if len(matchedJobs) == 0 {
        log.Fatal("no jobs to be processed")
    }

    rdb, err := db.New("localhost:6379")
    if err != nil {
        log.Println("redis conn error")
        log.Fatal(err)
    }

    err = rdb.Populate(&ntf, &matchedJobs)
    if err != nil {
        log.Println("redis population error")
        log.Fatal(err)
    }
}
