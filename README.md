<p align="center"><img src="logo.svg" alt="" height="128"></p>
<h1 align="center">rzjd</h1>
<p align="center">Razza's <a href="https://johnnydecimal.com/">Johnny.Decimal</a> Management System</p>

<p align="center">
    <a href="#download"><img src="https://img.shields.io/github/downloads/itisrazza/rzjd/total" alt="Downloads"></a>
    <a href="LICENSE"><img src="https://img.shields.io/github/license/itisrazza/rzjd" alt="Licence"></a>
    <a href="https://github.com/itisrazza/rzjd/actions/workflows/nightlies.yml"><img src="https://img.shields.io/github/actions/workflow/status/itisrazza/rzjd/nightlies.yml?label=nightly" alt="Nightly Builds"></a>
    <a href="https://github.com/itisrazza/rzjd/actions/workflows/build-test.yml"><img src="https://img.shields.io/github/actions/workflow/status/itisrazza/rzjd/build-test.yml?label=tests" alt="Tests Passing"></a>
</p>

`rzjd` is a command-line utility for managing your notes and files organised with the [Johnny.Decimal] system.

(TODO: add asciinema demo of how it works)

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

## Built On

- [Johnny.Decimal]
- [Kong]
- [Charm]'s [Huh] & [Bubble Tea]

[Johnny.Decimal]: https://johnnydecimal.com/
[Kong]: github.com/alecthomas/kong
[Charm]: https://charm.sh/
[Huh]: https://github.com/charmbracelet/huh
[Bubble Tea]: https://github.com/charmbracelet/bubbletea
