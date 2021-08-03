# `sema` - semantic commit tool

This is a simple command line tool inspired by [this gist][gist] about semantic
commit messages. In short, it proposed to use labelled commit messages that
derive their format from [Angular's commit rules][angular].

The `sema` command will help you follow these guidelines with no effort on your
part to memorise labels or double-check things.

[gist]: https://gist.github.com/joshbuchea/6f47e86d2510bce28f8e7f42ae84c716
[angular]: https://github.com/angular/angular.js/blob/master/DEVELOPERS.md#commits

## Format

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

## Installation

```bash
go install github.com/sharpvik/sema  # => $GOPATH/bin/sema
```

**NOTICE:** Make sure that your `$GOPATH/bin` is in `$PATH`!

**HACK:** After the default installation, the `sema` command will be available.
However, if you rename that binary file to `git-sema`, you will be able to use
it as follows (as if it's part of the default `git` tools):

```bash
git sema
```

## Usage

```bash
sema --help  # if you need a usage hint
sema --more  # see all label descriptions
sema --add   # run 'git add .' before all else
sema --push  # run 'git push' in the end
sema         # commit changes in current repo
```

**NOTE:** the `--push` and `--add` flags can be combined, which will be
equivalent to running the following:

```bash
git add .
git commit -m "feat(*): commit description"
git push
```

![label](img/label.png)
![scope](img/scope.png)
![message](img/message.png)
![result](img/result.png)
