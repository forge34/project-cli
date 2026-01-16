# project-cli (pjc)

![Go](https://img.shields.io/badge/Go-%2300ADD8.svg?style=flat&logo=go&logoColor=white)
[![License](https://img.shields.io/github/license/forge34/project-cli?style=flat-square)](./LICENSE)
A compact CLI to scaffold projects from embedded templates with interactive variable substitution.

Quick links

- List templates: `pjc list`
- Create project: `pjc create <template> <destination>`
- Run locally: `go run ./cmd/app create <template> <destination>`

Install

From source

1. Clone:

```bash
git clone <https://github.com/forge34/project-cli.git>```
```

2. Build:

```bash
go build -o bin/pjc ./cmd/app
```

1. (Optional) Install system-wide:
`sudo cp bin/pjc /usr/local/bin/pjc && sudo chmod +x /usr/local/bin/pjc`

Makefile conveniences

- `make build` — build to `bin/pjc`
- `make test` — run tests
- `make install` — build and install to `/usr/local/bin/pjc` (uses sudo)

Usage

Show help:

```bash
pjc help
```

List templates:

```bash
pjc list
```
Create a project:
```bash
pjc create <template> <destination>
```

Behavior notes

- Templates are embedded via Go's embed.FS (internals/templates).
- If a template contains `template.json` the CLI prompts for variables.
- Files ending with `.tmpl` are processed and written without the `.tmpl` suffix.
- Existing files in the destination are skipped.

Template format (template.json)

```json
{
  "prompts": [
    { "name": "projectName", "prompt": "Project name", "default": "my-app" },
    { "name": "author", "prompt": "Author name" }
  ]
}
````

Example

```bash
pjc create express-ts server
```
