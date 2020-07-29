package main

import (
	"fmt"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/config"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

func LocalClone() {
	url := ConfigFile.VCS
	path := ConfigFile.Environment["local"].WpPath

	_, err := git.PlainClone(path, false, &git.CloneOptions{URL: url})
	if err != nil {
		panic(err)
	}
}

func LocalFetch() {
	path := ConfigFile.Environment["local"].WpPath

	r, err := git.PlainOpen(path)
	if err != nil {
		panic(err)
	}

	w, err := r.Worktree()
	if err != nil {
		panic(err)
	}

	/* TODO */
	/* Should properly handle 'already up-to-date' error later */
	err = w.Pull(&git.PullOptions{RemoteName: "origin"})
	if err != nil {
		panic(err)
	}
}

func SwitchBranch() {
	path := ConfigFile.Environment["local"].WpPath

	r, err := git.PlainOpen(path)
	if err != nil {
		panic(err)
	}

	w, err := r.Worktree()
	if err != nil {
		panic(err)
	}

	err = r.Fetch(&git.FetchOptions{
		RefSpecs: []config.RefSpec{"refs/*:refs/*", "HEAD:refs/heads/HEAD"},
	})
	if err != nil {
		fmt.Println(err)
	}

	/* TODO */
	/* Fix issue 'reference not found' */
	err = w.Checkout(&git.CheckoutOptions{
		Branch: plumbing.NewBranchReferenceName(fmt.Sprintf("refs/heads/%s", CurrentEnv)),
		Force:  true,
	})
	if err != nil {
		panic(err)
	}
}

func RemoteClone() {
	/* TODO */
}

func RemoteFetch() {
	/* TODO */
}
