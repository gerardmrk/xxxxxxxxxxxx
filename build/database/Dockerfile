FROM postgres:12.0

RUN localedef -i en_GB -c -f UTF-8 -A /usr/share/locale/locale.alias en_GB.UTF-8

ENV LANG en_GB.UTF-8
ENV LANGUAGE=en_GB.UTF-8
ENV LC_ALL=en_GB.UTF-8
ENV PGDATA /var/lib/postgresql/data
ENV POSTGRES_INITDB_ARGS -T template0 -E UTF8 --locale=en_GB.UTF-8 --lc-collate=en_GB.UTF-8 --lc-ctype=en_GB.UTF-8

ENV POSTGRES_DB postgres
ENV POSTGRES_USER postgres
ENV POSTGRES_PASSWORD postgres

RUN apt-get update && \
    apt-get install -y --no-install-recommends postgresql-$PG_MAJOR-pgaudit && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

EXPOSE 5432/tcp

VOLUME ["/var/lib/postgresql/data", "/docker-entrypoint-initdb.d"]
