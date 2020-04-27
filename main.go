package main

import (
    "io"
    "os"
    "log"
    "github.com/go-git/go-billy/v5/memfs"
    "github.com/go-git/go-git/v5"
    "github.com/go-git/go-git/v5/storage/memory"
)

func main() {
    fs := memfs.New()
    storer := memory.NewStorage()

    _, err := git.Clone(storer, fs, &git.CloneOptions{
        URL: "https://github.com/tiny-ci/example",
    })

    if err != nil {
        log.Fatal(err)
    }

    config, err := fs.Open(".tyci.yml")
    if err != nil {
        log.Fatal(err)
    }

    io.Copy(os.Stdout, config)
}
