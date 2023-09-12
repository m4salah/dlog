# Download latest binary

Github has a release for each dlog version tag. it has binaries built for (Windows, Linux, MacOS) for several architectures. you can download the latest version from this page: https://github.com/m4salah/dlog/releases/latest

# Using Go

```bash
go install github.com/m4salah/dlog/cmd/dlog@latest
```

# From source

```bash
git clone git@github.com:emad-elsaid/dlog.git
cd dlog
go run ./cmd/dlog # to run it
go install ./cmd/dlog # to install it to Go bin.
```

# Arch Linux (AUR)

* dlog is published to AUR: https://aur.archlinux.org/packages/dlog-git
* Using `yay` for example:

```bash
yay -S dlog-git
```

# From source with docker-compose

```bash
git clone git@github.com:emad-elsaid/dlog.git
cd dlog
docker-composer build
docker-composer run
```

```info
dlog container attach `~/.dlog` as a volume and will write pages to it.
```

# Docker

Releases are packaged as docker images and pushed to GitHub 

```bash
docker pull ghcr.io/emad-elsaid/dlog:latest
docker run -p 3000:3000 -v ~/.dlog:/files ghcr.io/emad-elsaid/dlog:latest
```