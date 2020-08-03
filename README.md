# WPM, a comprehensive tool for advanced wordpress management

wpm is a tool built to make wordpress versionning, deployment and migration between any environment easier.

**:construction:  Caution :** poject is under development, and definitely not production ready. Feel free to contribute !

## How does it works ?

The global idea is to isolate wordpress core from your project. In this configuration, it's far easier to keep your wordpress installation updated and secure, while being consistent in your project. No fancy files everywhere, no overwrite, no confusion, your project is 100% independent from worpress core version.

```
wordpress/
├── core
└── data
```

As wordpress core is released publicly [on github](https://github.com/wordpress/wordpress), `core` folder can be fully managed using wpm. The tool is built to make sure your installation is up-to-date and that it has no integrity issue. 

If you already have a wordpress project you would like to integrate wpm in, feel free to read the [step by step guide](MELT_ME.md) to accomplish such cool things.

## What's gonna be in ?

* Website source and database versionning (migration, backups, rollback)
* Multiple environment (test / prod / staging...)
* Works over SSH
* WP core management

## Requirements

* SSH access to any environment
* git repository as a VCS

## Config file syntax

wpm config file uses TOML syntax

```toml
VCS = "github.com/user/repo.git"

[Environment]

    [Environment.dev]
    WpPath = "/var/www/html/wordpress"
    Username = "web-manager"
    Host = "192.168.0.1:22"
    AuthMethod = "password"

    [Environment.prod]
    WpPath = "/var/www/html/wordpress"
    Username = "web-manager"
    Host = "192.168.0.1:22"
    AuthMethod = "publickey"
    KeyPath = "/home/mike/.ssh/key.pem"
```

## Contribute

To contribute, take a closer look at the [CONTRIBUTING.md](CONTRIBUTING.md) file.

## Dependencies

* TOML handler : https://github.com/pelletier/go-toml
* Git handler : https://gopkg.in/src-d/go-git.v4
* Basic CLI : https://github.com/urfave/cli

