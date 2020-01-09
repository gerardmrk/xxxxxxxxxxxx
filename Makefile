# docker-compose command pointing to a custom config path
DKRC_CMD = docker-compose -f ./docker-compose.yml

# start runs the local infra.
start:
	@export $$(grep -v '^#' .env | xargs) && \
	$(DKRC_CMD) up --build

# start-d is similar to start, but runs all processes as
# daemons (in background mode). Logs can be accessed by `make show-logs`.
start-d:
	@export $$(grep -v '^#' .env | xargs) && \
	$(DKRC_CMD) up -d --build

# stop shutsdown all local infra.
stop:
	@$(DKRC_CMD) stop && \
	unset $$(grep -v '^#' .env | cut -d '=' -f 1 | xargs)

# logs feeds all infra's outputs to stdout and stderr.
# this is useful for if the infras are running in background mode.
logs:
	@$(DKRC_CMD) logs --follow

# purge removes all traces of project's infra artifacts
# (currently: Docker images, containers, and volumes).
purge:
	@$(DKRC_CMD) rm --force --stop -v && \
	rm -rf /usr/local/var/localdb && \
	unset $$(grep -v '^#' .env | cut -d '=' -f 1 | xargs)

.PHONY: install start start-d stop logs purge
