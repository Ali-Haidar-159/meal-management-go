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

The file paths are relative to the directory where you run the application. Keep `.env` private; it is excluded from Git.

## Run

From the project root:

```sh
go run .
```

Choose **1** to sign up, **2** to log in, or **0** to exit. User records are saved to the file configured by `USERS_FILE`. After a successful login, the user's ID is written to the file configured by `LOGIN_FILE`.

The dashboard currently displays a menu; its actions are not implemented yet.
