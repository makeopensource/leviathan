---
sidebar_position: 4
title: Database
---

# Database options

Leviathan uses sqlite by default this makes development and testing easier, but it also supports postgres.

## PostgresSQL

To enable either modify the config.yaml or pass in the following environment variables

| Environment Variable | Description              | Default Value |
|----------------------|--------------------------|---------------|
| `POSTGRES_ENABLE`    | Enable PostgreSQL        | `false`       |
| `POSTGRES_HOST`      | PostgreSQL host          | `"localhost"` |
| `POSTGRES_PORT`      | PostgreSQL port          | `"5432"`      |
| `POSTGRES_USER`      | PostgreSQL user          | `"postgres"`  |
| `POSTGRES_PASSWORD`  | PostgreSQL password      | `"postgres"`  |
| `POSTGRES_DB`        | PostgreSQL database name | `"postgres"`  |
| `POSTGRES_SSL`       | PostgreSQL SSL mode      | `"disable"`   |

## Sqlite

No additional setup is required for sqlite.

If you wish to keep using sqlite, set `postgres_enable: false` and leave the postgres option as is.