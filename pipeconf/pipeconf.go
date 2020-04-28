package pipeconf

import (
    "bytes"
    "io"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage/memory"
)

type GitRef struct {
    Name  string
    URL   string
    Hash  string
    IsTag bool
}

func Fetch(ref *GitRef) (*bytes.Buffer, error) {
    fs := memfs.New()
    storer := memory.NewStorage()

    reference := plumbing.NewBranchReferenceName(ref.Name)
    repo, err := git.Clone(storer, fs, &git.CloneOptions{
        URL: ref.URL,
        ReferenceName: reference,
        SingleBranch: true,
        Depth: 50,
        Tags: git.NoTags,
    })

    if err != nil {
        return nil, err
    }

    worktree, err := repo.Worktree()
    if err != nil {
        return nil, err
    }

    err = worktree.Checkout(&git.CheckoutOptions{
        Hash: plumbing.NewHash(ref.Hash),
    })

    if err != nil {
        return nil, err
    }

    configFile, err := fs.Open(".tyci.yml")
    if err != nil {
        return nil, err
    }

    var content bytes.Buffer
    io.Copy(&content, configFile)

    return &content, nil
}
