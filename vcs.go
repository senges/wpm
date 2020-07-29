package main

import (
	git "github.com/go-git/go-git/v5"
	"os"
)

func CloneTo(envName string) {
	env := ConfigFile.Environment[envName]
	url := ConfigFile.VCS

	INFO("Environment (%s) appears to be empty", envName)
	INFO("Deploying last wordpress version to `%v' on host [%v@%v]", env.WpPath, env.Username, env.Host)
	CMD("git -C %s clone %s", env.WpPath, url)

	_, err := git.PlainClone(env.WpPath, false, &git.CloneOptions{
		URL: url,
		Progress: os.Stdout,
	})

	CheckIfError(err)
}