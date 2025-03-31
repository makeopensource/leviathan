---
sidebar_position: 2
title: config.yaml
---

# Configuration file

Leviathan uses a `config.yaml` file for configration, this file is automatically created on first run with
some default values, everything about leviathan can be configured here ([except logging](#logging))

You can also modify some of the options using environment variables.

Below is the default configuration file with explanations for each option.

```yaml
clients:
  # Enable local Docker execution for jobs
  enable_local_docker: true
  # SSH client configurations for remote execution
  ssh:
    example:
      enable: false              # Enable/disable this SSH client
      host: 192.168.1.69        # Remote host address
      port: 22                  # SSH port
      user: test                # SSH username
      password: ""              # SSH password (if not using key auth)
      remotepublickey: ""       # Remote public key path
      usepublickeyauth: false   # Whether to use public key authentication
db:
  postgres:
    enable_postgres: false        # Enable PostgreSQL instead of SQLite
    postgres_db: leviathan       # Database name
    postgres_host: localhost     # PostgreSQL host
    postgres_pass: postgres      # Database password
    postgres_port: "5432"        # Database port
    postgres_ssl: disable        # SSL mode (disable, require, verify-ca, verify-full)
    postgres_user: postgres      # Database user
  sqlite:
    db_path: ./appdata/config/leviathan.db  # SQLite database file path
folder:
  # All folder paths should remain within ./appdata to ensure proper hardlinking
  job_output_dir: ./appdata/output           # Directory for job outputs
  labs: ./appdata/labs                       # Directory for lab files
  log_dir: ./appdata/config/logs/leviathan.log  # Log file location
  ssh_config: ./appdata/config/ssh_config    # SSH configuration directory
  tmp_submission_dir: ./appdata/submissions  # Temporary submission directory
  tmp_uploads: ./appdata/tmp_uploads        # Temporary upload directory
jobs:
  concurrent_jobs: 50              # Maximum number of concurrent jobs
server:
  log_level: info                 # Logging level (debug, info, warn, error)
  port: "9221"                    # Server port
  apikey: ""                      # API key for authentication
```

# Environment Variables

Some configuration options can be overridden using environment variables. Here are the available environment variables:

## Database Configuration

Read more [here](database.md)

## Directory Configuration

Read more [here](folders.md)

## Logging

You can set the following logger options these are available using only env vars

| Environment Variable             | Description                                        | Default Value |
|----------------------------------|----------------------------------------------------|---------------|
| `LEVIATHAN_LOG_LEVEL`            | Set log level (debug, info, warn, error, disabled) | `info`        |
| `LEVIATHAN_LOG_SHOW_CALLER_FILE` | Shows the file where the log was generated from    | `false`       |

The logging system has the following behaviour:

- Log rotation: Maximum file size of 10MB
- Backup retention: Keeps up to 5 backup files
- Age limit: Logs are kept for up to 30 days
- Compression: Old log files are automatically compressed
