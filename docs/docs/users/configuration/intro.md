---
sidebar_position: 1
title: Intro
---

# Leviathan Configuration

Leviathan creates an `appdata` directory in its working directory upon startup. This directory stores all configuration and runtime data.

For Docker deployments, the `appdata` directory is located at `/app/appdata` inside the container. To persist data between container restarts, you should mount this directory as a volume.

## Docker Compose Example

Below is a recommended `docker-compose.yml` configuration:

```yaml
services:
  leviathan:
    container_name: leviathan
    image: ghcr.io/makeopensource/leviathan:master
    volumes:
      - ./appdata:/app/appdata
    ports:
      - "9221:9221"
    restart: unless-stopped
```
