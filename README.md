# is-older

A tool used in combination with rsync and other standard Unix tools to make backup easy.

## Building

```sh
go build -o ,is-older main.go
```

## Usage

Add this to your crontab jobs (once a day should work):

```sh
,is-older -m 1 machine-1-snapshots/ | rm -r
```

## TODO

- [ ] add tests

## Licence

This code is licensed under the terms of a 2-clause BSD licence (see `LICENCE` for details).

