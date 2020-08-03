package main

import (
	"errors"
	"fmt"
	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"os"
)

var (
	ErrEnvDoesNotExists = errors.New("target environment does not exist")
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

/* Push local changes to current branch */
func PushToCurrent() {
	INFO("Retrieving local repository information")
	r, err := git.PlainOpen( getWD() )
	CheckIfError(err)

	w, err := r.Worktree()
	CheckIfError(err)

	INFO("Building commit object")
	commit, err := w.Commit("managed by wpm", &git.CommitOptions{
		All: true,
	})
	CheckIfError(err)

	obj, err := r.CommitObject(commit)
	CheckIfError(err)

	OK("---- COMMIT ----\n%s\n------------", obj.String())

	/* Not proper usage of CurrentEnv */
	INFO("Pushing local changes to remote repository (%s)", getCurrentRefName(r))
	err = r.Push(&git.PushOptions{
		RemoteName: "origin",
		Progress: os.Stdout,
	})
	CheckIfError(err)

	showHead(r)
}

/* Switch to branch envName */
func SwitchToBranch(envName string) error {
	ref := fmt.Sprintf("refs/heads/%s", envName)
	spec := fmt.Sprintf("%s:%s", ref, ref)

	INFO("Retrieving local repository information")
	r, err := git.PlainOpen( getWD() )
	CheckIfError(err)

	w, err := r.Worktree()
	CheckIfError(err)

	if !branchExists(r, envName) {
		return ErrEnvDoesNotExists
	}

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

	return nil
}

/* Check if branch ref exists */
/* Quite ugly but quite working as well for the moment */
/* Might use r.Branch() later */
func branchExists(r *git.Repository, branchName string) bool {
	fullRef := fmt.Sprintf("refs/remotes/origin/%s", branchName)
	exists := false

	refs, _ := r.References()
	refs.ForEach(func(ref *plumbing.Reference) error {
		if ref.Type() == plumbing.HashReference && string(ref.Name()) == fullRef {
			exists = true
		}

		return nil
	})

	return exists
}

/* Show commit object pointed by HEAD */
func showHead(r *git.Repository) {

	ref, err := r.Head()
	CheckIfError(err)

	commit, err := r.CommitObject(ref.Hash())
	CheckIfError(err)

	OK("HEAD -> %sRef: %s", commit, ref.Name())
}

/* Get ref of current local branch */
func getCurrentRefName(r *git.Repository) string {
	ref, err := r.Head()
	CheckIfError(err)

	return string(ref.Name())
}