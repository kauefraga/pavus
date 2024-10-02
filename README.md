<div align="center">

<img src="internal/server/static/icon.png" alt="Pavus' logo" height="128" />

# Pavus

Pavus is the **next-generation markdown tool**.

Need to preview your markdown? Templates for your README? It got your back!

[![GitHub Release](https://img.shields.io/github/v/release/kauefraga/pavus)](https://github.com/kauefraga/pavus/releases/latest)
[![GitHub Downloads](https://img.shields.io/github/downloads/kauefraga/pavus/total)](https://github.com/kauefraga/pavus/releases)
[![GitHub License](https://img.shields.io/github/license/kauefraga/pavus)](https://github.com/kauefraga/pavus/blob/main/LICENSE)
[![GitHub Last Commit](https://img.shields.io/github/last-commit/kauefraga/pavus/main)](https://github.com/kauefraga/pavus/commits/main/)
[![Go Reference](https://pkg.go.dev/badge/github.com/kauefraga/pavus.svg)](https://pkg.go.dev/github.com/kauefraga/pavus)

[Getting started](#usage) •
[Installation](#installation) •
[Configuration](#configuration) •
[Contributing](#contributing)

</div>

## Key Features

Here's what I'll be working on, what to expect

- Preview your markdown with hot reloading out-of-the-box
- Just a single binary with all the batteries included
- Tailored DX: descriptive messages, interactive mode, flags and colorful outputs
- Configurable, but awesome defaults (not available yet)
- Templates, so you can focus on content (not available yet)

> [!IMPORTANT]
> Work in progress... Just wait!

## Usage

### Installation

Installing via [pkg.go.dev](https://pkg.go.dev/github.com/kauefraga/pavus) is not available yet, but you can install a prebuilt binary for your platform.

Check the [latest release page](https://github.com/kauefraga/pavus/releases/latest).

### Configuration

## Mastering the CLI

Here you will find all about the command-line interface (available commands, descriptions, flags and aliases).

## Roadmap

A list of the features I want to implement, this should serve as a guide during development.

- Markdown preview
  - [x] `pavus` or `pavus path/to.md`
  - [x] Rendering (markdown to html, [markdown-it](https://github.com/markdown-it/markdown-it))
  - [x] Hot reloading ([fsnotify](https://github.com/fsnotify/fsnotify) and server-sent events)
  - [x] Syntax highlighting ([shiki](https://github.com/shikijs/shiki))
  - [ ] LaTeX ([KaTeX](https://katex.org/)?) and [Mermaidjs](https://mermaid.js.org/)
  - [x] Images (need to adjust the image path in the markdown to point to the server)
  - [x] Default theme
- Markdown templates
  - [ ] `init` or `init --template "x"` - create markdown with a template and generate prompts to fill it (interactive)
  - [ ] `add` or `add section` - append section (add at the end, interactive, available sections should be documented)
- Configuration
  - [ ] [Toml file](https://toml.io/en/v1.0.0) ([go-toml](https://github.com/pelletier/go-toml))
  - [ ] Themes (markdown preview)
  - [ ] Port (markdown preview)
  - [ ] AI token
- Markdown + AI
  - [ ] Fill template gaps

## Contributing

Feel free to contribute [opening an issue](https://github.com/kauefraga/pavus/issues/new) to report a bug or suggesting a CLI change, an improvement or a feature.

### How to contribute

1. Fork this project
2. Clone your fork on your machine
3. Setup the development environment
4. Make your changes and commit them following [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/)
5. Run `git push` to sync the changes
6. Open a pull request specifying what you did

## License

This project is licensed under the MIT License - See the [LICENSE](https://github.com/kauefraga/pavus/blob/main/LICENSE) for more information.

---

Se você gostou do projeto e ele te ajudou, considere [me apoiar um café](https://www.pixme.bio/kauefraga) ☕ 🇧🇷
