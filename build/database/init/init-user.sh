#!/bin/bash

set -e

# Perform all actions as $POSTGRES_USER
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
EOSQL
