# `sema` - semantic commit tool

This is a simple command line tool inspired by [this gist][gist] about semantic
commit messages. In short, it proposed to use labelled commit messages that
derive their format from [Angular's commit rules][angular].

[gist]: https://gist.github.com/joshbuchea/6f47e86d2510bce28f8e7f42ae84c716
[angular]: https://github.com/angular/angular.js/blob/master/DEVELOPERS.md#commits

The `sema` command will help you follow these guidelines with no effort on your
part to memorise labels or double-check things.

## 🌎 Contents

1. [Format](#format)
2. [Installation](#install)
3. [Usage](#usage)
4. [Screenshots](#demo)

## <a name="format"></a> 🍭 Format

Each commit message is supposed to be formatted in the following way:

```
TYPE(SCOPE): MESSAGE
```

Where `SCOPE` tells you about the scope of changes, `MESSAGE` summarises those
in a concise way, and `TYPE` is a short label from the following:

- `feat`: new feature for the user
- `fix`: bug fix for the user
- `docs`: changes to the documentation
- `style`: formatting with no production code change
- `refactor`: refactoring production code
- `test`: adding missing tests, refactoring tests
- `perf`: performance improvements
- `chore`: updating grunt tasks
- `infra`: infrastructural changes

## <a name="install"></a> 🚀 Installation

### Homebrew

```bash
brew install sharpvik/sema/sema
```

### AUR (for Arch-based Linux)

```bash
yay -S sema
```

### Using the `go` tool

```bash
go install github.com/sharpvik/sema/v3
```

> Both `yay` and `go` put `sema` binary into your `$GOPATH/bin` during
> installation so make sure that your `$GOPATH/bin` is in `$PATH`!

## <a name="usage"></a> 🔭 Usage

### Overview

```bash
NAME:
   sema - Semantic commits made simple

USAGE:
   sema [global options] command [command options] [arguments...]

COMMANDS:
   github   Open sema GitHub repository in browser
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --add, -a       begin by running 'git add' (default: false)
   --push, -p      run 'git push' on successful commit (default: false)
   --force, -f     force push changes with 'git push -f' (default: false)
   --long, -l      open editor to elaborate commit message (default: false)
   --breaking, -b  mark commit as introducing breaking changes (default: false)
   --tags, -t      push tags along with commits (default: false)
   --help, -h      show help
   --version, -v   print the version
```

### Flag Combos

#### Add & Push

The `--push` and `--add` flags can be combined (or `-ap`), which will be
equivalent to running the following:

```bash
git add .
git commit -m "feat(*): commit description"
git push
```

#### Force Push

Adding the `--force` flag to `--push` (or `-pf`) runs forceful push:

```bash
git push -f
```

> The `--force` used without `--push` will be ignored.

#### Breaking Changes

The `--breaking` flag will append an exclamation point to the end of your commit
label like so:

```bash
fix!(server): critical API change
```

On top of that, using `--breaking` with the [`--long`](#long) flag (or `-bl`),
appends `BREAKING CHANGE` suffix to the commit template file for your
convenience as follows:

```bash
fix!(server): critical API change

BREAKING CHANGE: [elaborate on this breaking change here]
```

### <a name="long"></a> Long Commits

By default, `git commit` opens an editor in your terminal where you can write a
commit message. For shorter commits, one could use `git commit -m "*****"`,
which is the default mode of operation for `sema`.

However, sometimes it is very beneficial to be able to elaborate your commit
message instead of just posting a semantic title. For this use case, meet the
new `--long` execution flag: after helping you come up with a semantic commit
title, it will open an editor (with your title prepended at the top) and let you
write some prose or poetry (whatever helps you get promotions).

## <a name="demo"></a> 🌌 Demo

![demo](img/demo.gif)
