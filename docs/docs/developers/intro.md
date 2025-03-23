---
sidebar_position: 1
title: Introduction
---

# Getting Started

We're excited that you're interested in contributing to our project. This guide will help you set up your development
environment and understand the project's licensing.

## License

TBD

## Development Prerequisites

Before you begin, please ensure you have the following tools installed:

### Required Technologies

* Go >= v1.24
    - Powers the main application
    - [Installation Guide](https://go.dev/learn/)

* Docker >= 28.0.1,

* Node.js and NPM >= 22
    - Used to develop the testing frontend ([kraken](kraken))
    - [Installation Guide](https://docs.npmjs.com/downloading-and-installing-node-js-and-npm)

* Just 1.38.0 and above (Optional)
    - Just is a command runner to save and run project-specific commands.
    - Think of it as an alternative to `npm run` commands or `Makefiles`
    - [Installation Guide](https://just.systems/man/en/)
    - We use it to run a few helpful commands, but you can run them manually if you want.

Please install and configure all these dependencies before proceeding with Gouda development. For version-specific
installation instructions, refer to each tool's official documentation.

## Cloning

1. [Fork](https://github.com/makeopensource/leviathan/fork) the repository
2. Clone your fork onto your machine.

### Folders

1. `/src`: main Go app 
2. `/kraken`: testing frontend 
3. `/spec`: grpc specification
4. `/example`: example graders for testing
5. `/docs`: docs

### Building

leviathan can be built using
```bash
go run src/main.go
```



## Development Workflow

### Branch Management

- Always create branches from `dev`
- Keep your local `dev` branch synced with upstream:
- Create feature/fix branches
  ```bash
  git checkout -b your-feature-name
  ```

### Making Changes

1. Make commits with clear messages:
   ```bash
   git commit -m "feat: add new functionality"
   ```
2. Push your changes:
   ```bash
   git push origin your-feature-name
   ```

## Pull Requests

### Opening a PR

1. Create PR targeting the `dev` branch
2. Use the PR template if provided
3. Include:
    - Clear title describing the change
    - Detailed description of changes
    - Link to related issues using GitHub keywords (Fixes #123, Closes #456)
    - Add appropriate labels/tags

### PR Guidelines

- Keep changes focused and atomic
- Update documentation if needed
- Add tests for new features
- Ensure CI checks pass
- Respond to review feedback promptly

## Labels

Common labels used in the project:

- `bug`: Bug fixes
- `feature`: New features
- `documentation`: Documentation updates
- `breaking`: Breaking changes
- `dependencies`: Dependency updates
- `UI`: If working on the frontend

## Questions?

Feel free to ask any questions on the devU channel in [discord](https://discord.gg/ChrT2DfcDT).
