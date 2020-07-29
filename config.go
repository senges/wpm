package main

import (
	"errors"
	"fmt"
	toml "github.com/pelletier/go-toml"
	"io/ioutil"
	"os/user"
)

const localEnvName string = "local"

/* Global config file */
var ConfigFile Config
var CurrentEnv string

type Config struct {
	VCS string
	Environment map[string]*Environment
}

type Environment struct {
	AuthMethod string
	Username   string
	Host       string
	WpPath     string
	KeyPath    string
}

/* Load toml config file from disk and map it */
func LoadConfigFileFromDisk() {

	/* Read config file */
	buffer, err := ioutil.ReadFile("config.toml")
	if err != nil {
		fmt.Println("Could not open config file")
		panic(err)
	}

	/* Map config file */
	if err := toml.Unmarshal(buffer, &ConfigFile); err != nil {
		fmt.Println("Error Unmarshal file")
		panic(err)
	}

	/* Default Current environment is set to local */
	if err := SwitchToEnv(localEnvName); err != nil {
		fmt.Println("No local environment in config file")
		panic(err)
	}

	/* Configure local env with proper user / host */
	localUser, err := user.Current()

	ConfigFile.Environment[localEnvName].Username = localUser.Username
	ConfigFile.Environment[localEnvName].Host = "localhost"

}

/* Switch from a custom env to another */
func SwitchToEnv(env string) error {
	if _, exists := ConfigFile.Environment[env]; exists {
		CurrentEnv = env
		return nil
	}

	return errors.New(env + ": no such environment")
}
