![Go icon](https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/512px-Go_Logo_Blue.svg.png)

---

# The discord bot in Go

## Description

This is a discord bot written in Go. It is a work in progress.

## Getting Started

The discord bot uses the following environment variables:

| Variable Name | Description                      |
|---------------|----------------------------------|
| `token` | The discord authentication token |

## Todo
- [ ] soon

## Installation

### Installations with Docker

The docker image is available on [GitHub Container Registry](https://github.com/d151l/discordbot/pkgs/container/discordbot)

##### Docker CLI

```bash
docker run -d \
    --name discordbot \
    -e token=your_token \
    ghcr.io/d151l/discordbot:master
```

#### Docker Compose

```yaml
version: "3.8"
services:
  discordbot:
        container_name: "discordbot"
        image: ghcr.io/d151l/discordbot:master
        restart: always
        pull_policy: always
        environment:
            - token=your_token
```