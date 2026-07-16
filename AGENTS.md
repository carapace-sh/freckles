# AGENTS.md

Guidance for AI agents working on the freckles codebase.

## Project Overview

freckles is a simple dotfile manager built with [Cobra] and [Carapace].
It uses the symlink approach: dotfiles are moved into a managed directory and
symlinked back to their original location in `$HOME`.

[Cobra]: https://github.com/spf13/cobra
[Carapace]: https://carapace.sh

## Repository Layout

```
cmd/freckles/
  main.go              # entrypoint, passes version to cmd.Execute
  cmd/
    root.go            # root cobra command + carapace/spec registration
    add.go             # add dotfiles to managed directory
    edit.go            # open a managed dotfile in $EDITOR
    git.go             # run git against the managed directory
    init.go            # initialize or clone the managed directory
    link.go            # create symlinks for all managed dotfiles
    list.go            # list all managed dotfiles (with style)
    verify.go          # verify symlink status of all managed dotfiles
    action/
      freckle.go       # carapace Action for completing freckle paths
pkg/freckles/
  freckles.go          # core domain logic: Freckle type, Walk, ignore
docs/                  # mdbook documentation
```

## Key Concepts

- **Managed directory**: `$HOME/.local/share/freckles/` (see `pkg/freckles.Dir()`).
  Use `filepath.Join` for all path construction, never string concatenation.
- **Freckle**: a single managed dotfile, identified by its relative path
  within the managed directory. `FrecklePath()` is the real file,
  `HomePath()` is the symlink target.
- **`.frecklesignore`**: gitignore-style file in the managed directory root,
  parsed via `go-git`'s `gitignore` matcher.
- **Carapace completion**: completion is explicit, not implicit.
  `carapace.Gen(cmd)` must be called on the root command (done in `root.go`
  `init()`).
- **Spec macros**: `spec.AddMacro` and `spec.Register` in `root.go` wire up
  the `freckles` macro for cross-application completion.

## Build & Test Commands

```sh
go build ./...
go test ./...
go vet ./...
gofmt -d -s .
```

There are currently no test files. CI also runs `staticcheck ./...`.

## Conventions

- Commands use `RunE` (not `Run`) and return errors to Cobra for display.
- Errors are wrapped with `fmt.Errorf("%s: %w", ...)` to include context.
- No `println` / `fmt.Println` for error reporting — return the error.
- Import grouping: stdlib, then external, then internal (separated by blank
  lines).
- Doc anchors (`// ANCHOR: name` / `// ANCHOR_END: name`) are used by the
  mdbook `include` shortcode — do not remove or rename without checking
  `docs/src/`.

## Common Tasks

### Adding a new subcommand

1. Create `cmd/freckles/cmd/<name>.go` with a `*cobra.Command`.
2. Register it in `init()`: `rootCmd.AddCommand(<name>Cmd)`.
3. Call `carapace.Gen(<name>Cmd)` if adding completions.
4. Add a docs page at `docs/src/freckles/cmd/<name>.md` and update
   `docs/src/SUMMARY.md`.

### Changing the managed directory path

Update `pkg/freckles.Dir()` — it is the single source of truth. All code
should call `Dir()`, not hardcode the path.

## CI

GitHub Actions workflow (`.github/workflows/go.yml`) runs on every push and
PR: build, test, format check, coverage, staticcheck, and (on tag) goreleaser.
The `doc` job builds mdbook docs and pushes `gh-pages` on master.
