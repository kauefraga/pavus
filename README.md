# `pavus` [![GitHub Release](https://img.shields.io/github/v/release/kauefraga/pavus?label=latest+version)](https://github.com/kauefraga/pavus/releases/latest) [![GitHub Last Commit](https://img.shields.io/github/last-commit/kauefraga/pavus/main)](https://github.com/kauefraga/pavus/commits/main/) [![Go Reference](https://pkg.go.dev/badge/github.com/kauefraga/pavus)](https://pkg.go.dev/github.com/kauefraga/pavus)

**Next-generation markdown tool**

Preview your markdown file, generate one from templates and more. Add pavus in your toolchain!

[Getting started](#usage) •
[Installation](#installation) •
[Templates](#templates) •
[Contributing](#contributing)

## Key Features

- Simplicity and power - just a single binary with all the batteries included
- Tailored experience - descriptive messages, interactive mode and colorful outputs
- Templates - just focus on your content
- Configurable - but awesome defaults (not available yet)

Want to know what's coming? Take a look at the [roadmap](https://github.com/kauefraga/pavus/issues/3).

## Usage

### Installation

- Via [pkg.go.dev](https://pkg.go.dev/github.com/kauefraga/pavus) (not available yet)
- Prebuilt binary

Check the [latest release](https://github.com/kauefraga/pavus/releases/latest) page to install the prebuilt binary for your platform.

If there is no binary for your platform you can build it yourself, see [how to build](#how-to-build).

### Templates

### Mastering the CLI

Reference about the command-line interface: available commands, examples, flags and aliases.

Try `pavus --help` and `pavus help [command]`.

##### Available commands

- `pavus [markdown]` - preview a markdown file in the browser with hot reload
- `pavus init` - create a markdown file based in a template
- `pavus templates` - list available markdown templates

##### Aliases

- `pavus init`, `pavus i`, `pavus create`, `pavus c`

##### Examples

###### Root command (preview)

Preview a markdown file, in this case, "README.md"

```sh
pavus README.md
```

If no file is passed to pavus, it will automatically look for one

```sh
pavus
```

The flag `-o --open-browser` can be used to open the preview in your default browser

```sh
pavus README.md --open-browser
```

The flag `-a --asset-directory <directory>` can be used to specify the assets directory (images and other resources being used in the markdown)

```sh
# - README.md
# - public/
#   - demo.gif
#   - header.png
#   - ...

pavus README.md --asset-directory public/
```

###### Init command

Create a markdown file based in a template interactively

It'll ask you for a template name and an output file name

```sh
pavus init
```

Looking for flags?

```sh
# non-interactive
pavus init --template "tooling" --output "my-docs.md"

# shorthand flags
pavus init -t tooling -o my-docs.md

# fill the missing extension (.md) and only prompts for the template name
pavus init -o my-docs
```

###### Templates command

List pavus available markdown templates

```sh
pavus templates
```

### Mastering the configs

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

## How to build

With [Go](https://go.dev/) installed, building pavus should be as easy as running the following command

```sh
go build cmd/main.go -o pavus
```

However, running the command below should generate a more lightweight binary

```sh
CGO_ENABLED=0 go build -ldflags='-w -s' cmd/main.go -o pavus
```

In the [`build.sh`](build.sh) you can see how the release binaries are being built.

## License

This project is licensed under the MIT License - See the [LICENSE](https://github.com/kauefraga/pavus/blob/main/LICENSE) for more information.

---

Se você gostou do projeto e ele te ajudou, considere [me apoiar um café](https://www.pixme.bio/kauefraga) ☕ 🇧🇷
