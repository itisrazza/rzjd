<p align="center"><img src="logo.svg" alt="" height="128"></p>
<h1 align="center">rzjd</h1>
<p align="center">Razza's <a href="https://johnnydecimal.com/">Johnny.Decimal</a> Management System</p>

`rzjd` is a command-line utility for managing your notes and files organised with the [Johnny.Decimal] system.

(TODO: add asciinema demo of how it works)

(TODO: add shields with the language, build status, licence etc)

[Johnny.Decimal]: https://johnnydecimal.com/

## Download

(TODO: add how to download prebuilt binaries)

## Contributing to rzjd

Clone the repository with `git clone` and you can use the standard go tooling
to make changes.

```bash
go run ./... <args>
```

For development, you can set the path to the data store using the `RZJD_STORE`
environment variable, or by using the `-C`/`--store` flags.
