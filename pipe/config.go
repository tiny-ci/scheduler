package pipe

import (
	"bytes"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage/memory"
	"io"
)

type GitRef struct {
	Name  string
	URL   string
	Hash  string
	IsTag bool
}

func getRefName(name string, isTag bool) plumbing.ReferenceName {
	if isTag {
		return plumbing.NewTagReferenceName(name)
	}

	return plumbing.NewBranchReferenceName(name)
}

func Fetch(ref *GitRef) (*bytes.Buffer, error) {
	fs := memfs.New()
	storer := memory.NewStorage()

	reference := getRefName(ref.Name, ref.IsTag)
	repo, err := git.Clone(storer, fs, &git.CloneOptions{
		URL:           ref.URL,
		ReferenceName: reference,
		SingleBranch:  true,
		Depth:         50,
		Tags:          git.NoTags,
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
		return nil, nil
	}

	var content bytes.Buffer
	io.Copy(&content, configFile)

	return &content, nil
}
