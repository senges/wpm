package main

import (
	"fmt"
	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"os"
)

/* Clone repository to defined path */
func CloneTo(envName string) {
	env := ConfigFile.Environment[envName]
	url := ConfigFile.VCS

	INFO("Environment (%s) appears to be empty", envName)
	INFO("Deploying last wordpress version to `%v' on host [%v@%v]", env.WpPath, env.Username, env.Host)
	CMD("git -C %s clone %s", env.WpPath, url)

	r, err := git.PlainClone(env.WpPath, false, &git.CloneOptions{
		URL: url,
		Progress: os.Stdout,
	})
	CheckIfError(err)

	showHead(r)
}

/* Switch to branch envName */
func SwitchToBranch(envName string) {
	spec := fmt.Sprintf("refs/heads/%s:refs/heads/%s", envName, envName)

	INFO("Retrieving local repository information")
	r, err := git.PlainOpen(ConfigFile.Environment[localEnvName].WpPath)
	CheckIfError(err)

	w, err := r.Worktree()
	CheckIfError(err)

	INFO("Fetching remote branch %v", envName)
	err = r.Fetch(&git.FetchOptions{
		RemoteName: "origin",
		RefSpecs: []config.RefSpec{ config.RefSpec(spec) },
	})
	// Already-Up-To-Date should not be treated as error and panic program
	if err != nil && err != git.NoErrAlreadyUpToDate {
		CheckIfError(err)
	}

	INFO("Switching to branch %s", envName)
	err = w.Checkout(&git.CheckoutOptions{Branch:plumbing.NewBranchReferenceName(envName)})
	CheckIfError(err)

	showHead(r)
}

/* Check if branch ref exists */
/* Quite ugly but quite working as well for the moment */
func branchExists(r *git.Repository, refName string) bool {
	exists := false

	refs, _ := r.References()
	refs.ForEach(func(ref *plumbing.Reference) error {
		if ref.Type() == plumbing.HashReference && string(ref.Name()) == refName {
			exists = true
		}

		return nil
	})

	return exists
}

func showHead(r *git.Repository) {

	ref, err := r.Head()
	CheckIfError(err)

	commit, err := r.CommitObject(ref.Hash())
	CheckIfError(err)

	OK("HEAD -> %sRef: %s", commit, ref.Name())
}