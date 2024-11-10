#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
  DO \$\$
  BEGIN
      IF NOT EXISTS (SELECT FROM pg_catalog.pg_user WHERE usename = 'docker') THEN
          CREATE USER docker;
      END IF;
  END
  \$\$;

  GRANT ALL PRIVILEGES ON DATABASE gator TO docker;
EOSQL