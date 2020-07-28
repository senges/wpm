# How to contribute
## Commit convention

WPM repo uses a commit naming convention inspired by [Angular Commit Message Format](https://github.com/angular/angular/blob/22b96b9/CONTRIBUTING.md#-commit-message-guidelines).

### Commit message format

```
<type>(<scope>): <subject>
```

**Type**

* common: Misc changes (ex: config file update)
* feat: Feature update
* doc: Documentation related changes
* fix: A bug fix
* refactor: Cosmetic code or file update
* test: Test related changes
* perf: Performance improvement
* build: Changes that affect the build system
* ci: Changes to our CI configuration files and scripts

**Scope**

Targeted file, subject, code or function related to commit `<type>`.

**subject**

The subject contains a succinct description of the change:

* use the imperative, present tense: "change" not "changed" nor "changes"
* don't capitalize the first letter
* no dot (.) at the end

**Exemple**

```bash
doc(readme.md): add "Config file syntax" part
```

```bash
refactor(config.toml): make overall syntax more understandable
```