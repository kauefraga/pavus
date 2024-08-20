<table align="center">
<tr>
<td><img src="internal/server/static/icon.png" alt="Pavus' logo" height="128" /></td>
<td>

# Pavus

Pavus is the **next-generation markdown tool**.

Need to preview your markdown? Need templates for your README? It got your back!

[![GitHub Release](https://img.shields.io/github/v/release/kauefraga/pavus)](https://github.com/kauefraga/pavus/releases/latest)
[![GitHub Downloads](https://img.shields.io/github/downloads/kauefraga/pavus/total)](https://github.com/kauefraga/pavus/releases)
[![GitHub License](https://img.shields.io/github/license/kauefraga/pavus)](https://github.com/kauefraga/pavus/blob/main/LICENSE)
[![GitHub Last Commit](https://img.shields.io/github/last-commit/kauefraga/pavus/main)](https://github.com/kauefraga/pavus/commits/main/)
[![Go Reference](https://pkg.go.dev/badge/github.com/kauefraga/pavus.svg)](https://pkg.go.dev/github.com/kauefraga/pavus)

[Getting started](#usage) •
[Installation](#installation) •
[Configuration](#configuration) •
[Pix me a coffee](#pix-me-a-coffee)
<br/><br/>

</td>
</tr>
</table>

## Key Features

- Pleasant DX: intuitive, descriptive messages and colorful outputs
- Preview your markdown with hot reloading out-of-the-box
- Just a single lightweight binary with all the batteries included
<!--
- Configurable, but awesome defaults
- Templates, so you don't need to come with the new idea
-->

> [!IMPORTANT]
> Work in progress... Just wait!

## Usage

### Installation

### Configuration

## Mastering the CLI

Like in [Ruke](https://github.com/kauefraga/ruke#mastering-the-cli), here you will find all about the command-line interface (available commands, descriptions, flags and aliases).

## Roadmap

A list of the features I want to implement, this should serve as a guide during development.

- Markdown preview
  - [x] `pavus` or `pavus path/to.md`
  - [x] Rendering (markdown to html, [markdown-it](https://github.com/markdown-it/markdown-it))
  - [x] Hot reloading ([gorilla/websocket](https://github.com/gorilla/websocket))
  - [x] Syntax highlighting ([shiki](https://github.com/shikijs/shiki)) 
  - [ ] LaTeX ([KaTeX](https://katex.org/)?) and [Mermaidjs](https://mermaid.js.org/)
  - [ ] Images
  - [ ] Default theme
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

Feel free to contribute, opening an issue to report a bug, suggest a CLI change, an improvement or a feature.

### How to contribute

1. Fork this repository
2. Clone your fork on your machine
3. Make your changes, commit and push them
4. Open a pull request (write a descriptive message about what you changed)

## License

This project is licensed under the MIT License - See the [LICENSE](https://github.com/kauefraga/pavus/blob/main/LICENSE) for more information.

---

## Pix me a coffee

If this project helped you, consider to [pix me a coffee](https://www.pixme.bio/kauefraga)! 💜 🇧🇷
