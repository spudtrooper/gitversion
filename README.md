# gitversion

Tool to add a `--version` flag to your binary the prints the git version.

This will add some files to the `gitversion` package that can be changed with the `--pkg` flag. You should add the following to your main:

```go
if gitversion.CheckVersionFlag() {
    return
}
```

## Usage

Install:

```bash
go install github.com/spudtrooper/gitversion
```

Create the files:

```bash
~/go/bin/gitversion
```

Increment the tag's most minor version (e.g. `v0.0.1` to `v0.0.2`)

```bash
~/go/bin/gitversion --inc
```