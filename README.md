# Task Tracker CLI

A simple command-line task tracker written in Go. Track what you need to do, what you've done, and what you're working on — all from your terminal. Tasks are stored in a local `tasks.json` file.

## Features

- Add, update, and delete tasks
- Mark tasks as `in-progress` or `done`
- List all tasks, or filter by status (`todo`, `in-progress`, `done`)
- Tasks persist in a `tasks.json` file in the current directory
- No external dependencies — Go standard library only