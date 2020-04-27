package main

import (
    "io"
    "os"
    "log"
    "github.com/go-git/go-billy/v5/memfs"
    "github.com/go-git/go-git/v5"
    "github.com/go-git/go-git/v5/storage/memory"
    "github.com/go-git/go-git/v5/plumbing"
)

func main() {
    fs := memfs.New()
    storer := memory.NewStorage()

    reference := plumbing.NewBranchReferenceName("master")

    log.Println("clonning repository in memory")
    repo, err := git.Clone(storer, fs, &git.CloneOptions{
        URL: "https://github.com/tiny-ci/example",
        ReferenceName: reference,
        SingleBranch: true,
        Depth: 50,
        Tags: git.NoTags,
    })

    if err != nil {
        log.Fatal(err)
    }

    worktree, err := repo.Worktree()
    if err != nil {
        log.Fatal(err)
    }

    err = worktree.Checkout(&git.CheckoutOptions{
        Hash: plumbing.NewHash("dba8a3250ff364a8a1ccfe0ca0b1bdeb43adadcb"),
    })

    if err != nil {
        log.Println("cannot checkout")
        log.Fatal(err)
    }

    config, err := fs.Open(".tyci.yml")
    if err != nil {
        log.Fatal(err)
    }

    io.Copy(os.Stdout, config)
}
