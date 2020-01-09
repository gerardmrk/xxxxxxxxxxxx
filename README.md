# Go + PG Boilerplate

## One time setup

### Homebrew (mac)

```bash
$ brew install go-task/tap/go-task
```

### Other

```bash
$ curl -sL https://taskfile.dev/install.sh | sh
```

or visit https://taskfile.dev/#/installation for more instructions.

## Quickstart

Your task file is `Taskfile.yml` if on a Mac or any Debian-based Linux distros, or `Taskfile_windows.yml` if on Windows.

1. Set `vars.PROJECT_NAME` in your task file.
2. Run `task start`.

## Teardown

1. Run `make purge`.
