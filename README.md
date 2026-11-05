# Meal Management App

A command-line meal management application written in Go. Users can register and log in; the dashboard provides options for adding bazaar purchases and meal counts. User records are stored in a JSON file.

## Requirements

- Go 1.25.6 or later

## Setup

Create a `.env` file in the project root:

```env
APP_NAME=Meal Management App
USERS_FILE=users.json
LOGIN_FILE=localstorage.txt
```

Keep `.env` private; it is excluded from Git. The app searches for `.env` in the
current directory and its parents. Relative data-file paths are resolved from
the directory containing `.env`.

## Run

From the project root:

```sh
go run ./cmd/meal-management
```

You can also run it from the command directory:

```sh
cd cmd/meal-management
go run .
```

Choose **1** to sign up, **2** to log in, or **0** to exit. User records are saved to the file configured by `USERS_FILE`. After a successful login, the user's ID is written to the file configured by `LOGIN_FILE`.

The dashboard currently displays a menu; its actions are not implemented yet.

## Project layout

```text
cmd/meal-management/  Application entry point
internal/cli/         Console menus and input prompts
internal/domain/      Core data types
internal/service/     Application operations and JSON file access
```

Packages under `internal/` are implementation details and cannot be imported by
other Go modules. Runtime JSON files remain relative to the working directory,
so existing `.env` paths and monthly data files continue to work unchanged.
