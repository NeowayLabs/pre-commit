# pre-commit

The `pre-commit` will validate your commit messages before the commit being accepted.
That will prevent you from having to rebase your commit history to adapt your commit messages to [semantic-release](https://github.com/NeowayLabs/semantic-release) standards.

# Requirements
- [Golang installation](https://go.dev/doc/install)

# How to install `pre-commit`?

Clone this project and run the following `make` command on its root path.

```
make install
```

# How to use it?

After adding new changes with the `git add` command, you can run `commit .` on any git project root path and follow CLI steps.

```
commit .
```
