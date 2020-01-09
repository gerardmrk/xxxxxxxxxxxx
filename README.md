# Go + PG Boilerplate

## Requirements

You need to install [**Taskfile**](https://taskfile.dev/).

If you are on a Mac with Homebrew:

```bash
$ brew install go-task/tap/go-task
```

For other OSes, follow installation instructions [**here**](https://taskfile.dev/#/installation) or run (for Unix):

```bash
$ curl -sL https://taskfile.dev/install.sh | sh
```

## Setup

Change the value of `vars.PROJECT_NAME` in your **Taskfile**.

Your task file is:

- `Taskfile.yml` for MacOS and Debian-based distros.
- `Taskfile_windows.yml` for Windows (**_TODO_**).

Then run the following to setup git, among other stuff:

```bash
task init
```

## Quickstart

```bash
$ task start # or task start-bg; to run all containers.
$ task stop # or task stop-bg; to stop all containers.
```
