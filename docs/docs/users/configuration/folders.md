---
sidebar_position: 5
title: Folders
---

# Folder config

Leviathan uses the following folders for its operations.

| Environment Variable    | Description                  | Purpose/Usage                                                                  | Default Value                 |
|-------------------------|------------------------------|--------------------------------------------------------------------------------|-------------------------------|
| `TMP_SUBMISSION_DIR`    | Temporary submission storage | Stores job submissions temporarily; contents are deleted after job completion  | `./appdata/submissions`       |
| `SUBMISSION_OUTPUT_DIR` | Job output storage           | Preserves stdout and logs from containers where jobs executed                  | `./appdata/output`            |
| `SSH_CONFIG_DIR`        | SSH configuration storage    | Contains Leviathan-managed SSH keys for remote Docker host connections         | `./appdata/config/ssh_config` |
| `LABS_DIR`              | Laboratory files storage     | Stores lab files                                                               | `./appdata/labs`              |
| `TMP_UPLOAD_DIR`        | Temporary upload storage     | Facilitates file transfers to Leviathan; contents are deleted after processing | `./appdata/tmp_uploads`       |

# Folder Configuration

Leviathan uses a structured directory system for its operations.

## Hardlinks

Leviathan uses hardlinks for efficient job processing. Changing folder locations requires special consideration:

- 🚫 **Avoid network mounts** - NFS/SMB volumes often lack proper hardlink support
- 📂 **Maintain same filesystem** - All paths must reside on the same physical storage device

### Docker Deployment Considerations

When deploying via Docker, volume mounting requires special attention. For example:

`/some_host_mount:/appdata/folder1`

`/some_host_mount:/appdata/folder2`

Docker treats these paths as separate filesystems, preventing hardlinks between them.

❗ Changing these paths incorrectly may lead to:

- Failed job submissions
- Corrupted submission artifacts
- Incomplete output generation
