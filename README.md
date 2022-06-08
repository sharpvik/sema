# `sema` - semantic commit tool

This is a simple command line tool inspired by [this gist][gist] about semantic
commit messages. In short, it proposed to use labelled commit messages that
derive their format from [Angular's commit rules][angular].

[gist]: https://gist.github.com/joshbuchea/6f47e86d2510bce28f8e7f42ae84c716
[angular]: https://github.com/angular/angular.js/blob/master/DEVELOPERS.md#commits

The `sema` command will help you follow these guidelines with no effort on your
part to memorise labels or double-check things.

## Contents

1. [Format](#format)
2. [Installation](#install)
3. [Usage](#usage)
4. [Screenshots](#demo)

## <a name="format"></a> Format

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

> You can see the list of these labels with explanations using `sema --more`.

## <a name="install"></a> Installation

### From AUR (for Arch-based Linux)

```bash
yay -S sema
```

### Using the `go` tool

```bash
go install github.com/sharpvik/sema  # => $GOPATH/bin/sema
```

> Both installation methods put `sema` binary into your `$GOPATH/bin` so make
> sure that your `$GOPATH/bin` is in `$PATH`!

## <a name="usage"></a> Usage

### Overview

```bash
Usage: sema [flags]

 -a, --add      Begin with running 'git add'
 -p, --push     Run 'git push' on successful commit
 -f, --force    Add force push flag '-f' during 'git push'

 -h, --help     Display help message
 -m, --more     Explain commit types
 -n, --version  Display installed version of sema
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

### Commit Hooks

Sometimes we'd like to run a script before every commit. For example, I often
forget to run `go fmt ./...` before publishing changes. To combat this issue,
introducing **commit hooks**.

Every time you run `sema`, it will look for a file called `hooks.sema` in the
current working directory and attempt to execute it. Make sure to give executive
permissions to the hooks file (consider `chmod +x hooks.sema`).

Of course, using `hooks.sema` is optional and its absence won't break anything.
For a basic example of such a file, take a look at [`hooks.sema`](./hooks.sema).

## <a name="demo"></a> Screenshots

![label](img/label.png)
![scope](img/scope.png)
![message](img/message.png)
![result](img/result.png)
