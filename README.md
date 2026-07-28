# freckles

[![PkgGoDev](https://pkg.go.dev/badge/github.com/carapace-sh/freckles)](https://pkg.go.dev/github.com/carapace-sh/freckles)
[![GoReportCard](https://goreportcard.com/badge/github.com/carapace-sh/freckles)](https://goreportcard.com/report/github.com/carapace-sh/freckles)
[![Coverage Status](https://coveralls.io/repos/github/carapace-sh/freckles/badge.svg?branch=master)](https://coveralls.io/github/carapace-sh/freckles?branch=master)
[![documentation](https://img.shields.io/badge/&zwnj;-documentation-blue?logo=gitbook)](https://freckles.carapace.sh)
[![Packaging status](https://repology.org/badge/tiny-repos/freckles.svg)](https://repology.org/project/freckles/versions)

A simple dotfile manager based on [carapace](https://github.com/rsteube/carapace) using the symlink approach.

[![asciicast](https://asciinema.org/a/499658.svg)](https://asciinema.org/a/499658)

## Completion

```sh
# bash
source <(freckles _carapace)

# elvish
eval (freckles _carapace|slurp)

# fish
freckles _carapace | source

# oil
source <(freckles _carapace)

# nushell
freckles _carapace | save freckles.nu ; nu -c 'source freckles.nu'

# powershell
freckles _carapace | Out-String | Invoke-Expression

# tcsh
eval `freckles _carapace`

# xonsh
exec($(freckles _carapace))

# zsh
source <(freckles _carapace)
```
