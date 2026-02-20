# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

A single-file Go CLI that checks whether Max Verstappen won the most recent F1 race. Uses the `f1apireader` library to fetch race results from an F1 API.

## Git Workflow

Never commit directly to main. Always create a feature branch for changes.

## Build and Run

```bash
go build -o did_verstappen_win did_verstappen_win.go
./did_verstappen_win
```

## Makefile Targets

- `make` / `make all` — runs fmt + test
- `make fmt` — checks formatting with gofmt
- `make lint` — runs golangci-lint
- `make install_deps` — downloads dependencies
- `make clean` — removes `./bin`

No test suite currently exists (test targets are commented out).

## Architecture

- **Single file:** `did_verstappen_win.go` — all logic lives here
- **Module name:** `github.com/rpunt/did_verstappen_win`
- **Key dependency:** `github.com/rpunt/f1apireader` — provides `RaceResults()`, race status checking, and winner extraction
- **Flow:** Fetch race results → if race completed, check if winner TLA == "VER" → print YES/NO; otherwise print current race status
