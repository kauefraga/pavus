# `pavus` [![GitHub Release](https://img.shields.io/github/v/release/kauefraga/pavus?label=latest+version)](https://github.com/kauefraga/pavus/releases/latest) [![GitHub Last Commit](https://img.shields.io/github/last-commit/kauefraga/pavus/main)](https://github.com/kauefraga/pavus/commits/main/) [![Go Reference](https://pkg.go.dev/badge/github.com/kauefraga/pavus)](https://pkg.go.dev/github.com/kauefraga/pavus)

**Next-generation markdown tool**

Need to preview your markdown? Templates for your README? It got your back!

[Getting started](#usage) •
[Installation](#installation) •
[Configuration](#configuration) •
[Contributing](#contributing)

## Key Features

- Preview your markdown with hot reloading out-of-the-box
- Just a single binary with all the batteries included
- Tailored DX: descriptive messages, interactive mode, flags and colorful outputs
- Configurable, but awesome defaults (not available yet)
- Templates, so you can focus on content (not available yet)

> [!IMPORTANT]
> Work in progress... Just wait!

## Usage

### Installation

- Via [pkg.go.dev](https://pkg.go.dev/github.com/kauefraga/pavus) (not available yet)
- Prebuilt binary

Check the [latest release](https://github.com/kauefraga/pavus/releases/latest) page to install the prebuilt binary for your platform.

If there is no binary for your platform you can build it yourself, see [how to build](#how-to-build).

### Configuration

### Mastering the CLI

Reference about the command-line interface: available commands, descriptions, flags and aliases.

## Contributing

Feel free to contribute [opening an issue](https://github.com/kauefraga/pavus/issues/new) to report a bug or suggesting a CLI change, an improvement or a feature.

### How to contribute

1. Fork this project
2. Clone your fork on your machine
3. Setup the [dev environment](#how-to-setup-dev-environment)
4. Make your changes and commit them following [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/)
5. Run `git push` to sync the changes
6. Open a pull request specifying what you did

### How to setup dev environment

- Have [Go](https://go.dev/) installed (Preferably [1.22.5](go.mod))

Install the dependencies

```sh
go mod download
```

And/or just run the project

```sh
go run cmd/main.go
```

### How to build

With [Go](https://go.dev/) installed, building pavus should be as easy as running the following command

```sh
go build cmd/main.go -o pavus
```

However, running the command below should generate a more lightweight binary

```sh
CGO_ENABLED=0 go build -ldflags='-w -s' cmd/main.go -o pavus
```

## License

This project is licensed under the MIT License - See the [LICENSE](https://github.com/kauefraga/pavus/blob/main/LICENSE) for more information.

---

Se você gostou do projeto e ele te ajudou, considere [me apoiar um café](https://www.pixme.bio/kauefraga) ☕ 🇧🇷
