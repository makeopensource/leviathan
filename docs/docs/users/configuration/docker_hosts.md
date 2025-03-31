---
sidebar_position: 3
title: Docker Hosts
---

# Connecting to Docker Hosts

Leviathan uses SSH to connect to remote docker hosts and supports multiple ways to connect to Docker hosts.

This can be configured in the `clients` section in `config.yaml`
Here is an example ssh section

```yaml
clients:
  enable_local_docker: true
  ssh:
    example: # key is a nickname for this machine
      enable: false
      host: 192.168.1.69
      port: 22
      user: test
      password: ""
      remotepublickey: ""
      usepublickeyauth: false
```

## SSH user considerations

You must ensure that the ssh user that the leviathan uses can access the docker demon properly,

You can create a user with limited permissions that can only control the Docker daemon remotely. This approach follows
the principle of the least privilege, providing only the specific access needed.

## Connection hierarchy

When leviathan connects to a docker host it follows the following order

1. Checks if  `use_public_key_auth` is enabled then uses [public key auth](#ssh-public-key-authentication-recommended)
2. If that is disabled, it checks the `passoword` section and uses [password auth](#ssh-password-authentication)
3. If password is empty, then it will fallback to [host public key auth](#ssh-public-key-authentication-host-dependent)

Regardless of the auth method, if the host public key is automatically stored on first
connection and verified on subsequent connections.

If the `remotepublickey` is set before first connect then it will be used to verify.

**Leviathan requires at least 1 docker host connected at anytime,
if it is unable to connect to any docker host then it will exit with a fatal error.**

## Authentication Methods

### SSH Public Key Authentication (Recommended)

```yaml
use_public_key_auth: true
```

- Uses auto-generated key pair stored in Leviathan's config directory, found in `appdata/config/ssh_config`
- Requires adding Leviathan's public key to remote host's `~/.ssh/authorized_keys`
- No password required

### SSH Password Authentication

```yaml
password: "some password"
```

- Basic authentication method
- Recommended for initial setup/testing only

### SSH Public Key Authentication (Host dependent)

- Uses key pair stored on the machine running leviathan typically at `~/.ssh`
- Requires mounting the `.ssh` dir into the docker container or
- transferring the keys manually and setting up `.ssh` inside the container file system

### Local Docker Connection

- Enabled by default on first run.
- Requires Docker daemon running on local machine
- Requires mounting `/var/run/docker.sock:/var/run/docker.sock` into the container
- Disable with `enable_local_docker: false` in config.yaml

# Configuration Parameters

| Parameter             | Required | Default | Description                      | Notes                                              |
|-----------------------|:--------:|:-------:|----------------------------------|----------------------------------------------------|
| `enable`              |    No    | `true`  | Host availability                | Set to `false` to disable host                     |
| `host`                |   Yes    |    -    | Host machine IP                  | `http://192.168.1.1` or `https:remote.docker.com`  |
| `port`                |    No    |  `22`   | SSH port number                  | Standard SSH port used if not specified            |
| `user`                |   Yes    |    -    | SSH login username               |                                                    |                                           
| `password`            |    No    |    -    | SSH password                     | Used for password-based authentication             |
| `remote_public_key`   |    No    |    -    | Remote host's public key         | Automatically set by Leviathan on first connection |
| `use_public_key_auth` |    No    | `false` | Enable public key authentication | Set to `true` to use leviathan's SSH key pairs     |


