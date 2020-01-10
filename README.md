# Go + PG Boilerplate

## Requirements

[**Go 1.13**](https://golang.org/dl/) or later is required. Mac users can install it via Homebrew:

```bash
$ brew install go
```

[**Taskfile**](https://taskfile.dev/) is required (see instructions [here](https://taskfile.dev/#/installation)). Mac users can install it via Homebrew:

```bash
$ brew install go-task/tap/go-task
```

[**Docker Desktop**](https://www.docker.com/products/docker-desktop) is required. For Unix users, whitelist `/usr/local/var` path in the File Sharing preferences. For Windows users, **TODO**.

## Setup

1. Update values in `.env` to your own preference. These env vars will be available to all containers (don't do this in production).

2. Update the values in `Taskvars.yml` as well. These values are referenced by proxy (via the task files) in various files.

3. Run `task init` to do all other automated initializations.

## Quickstart

Ensure your Docker client is up and running.

### Start all containers

```bash
$ task start-bg # or `start` to run in foreground.
```

### Stop all containers

```bash
$ task stop-bg # or just CTRL-F if you ran `start`.
```

### Show logs from all containers

```bash
$ task log-bg # don't need this if you ran `start`.
```

### Remove all resources

```bash
$ task purge
```
