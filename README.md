# test && commit || revert

[![Go Report Card](https://goreportcard.com/badge/github.com/heymatthew/tcr)](https://goreportcard.com/report/github.com/go-git/go-git)
[![Maintainability](https://api.codeclimate.com/v1/badges/2161eb11259c773f1e61/maintainability)](https://codeclimate.com/github/heymatthew/tcr/maintainability)
[![Build Status](https://img.shields.io/endpoint.svg?url=https%3A%2F%2Factions-badge.atrox.dev%2Fheymatthew%2Ftcr%2Fbadge%3Fref%3Dmaster&style=flat)](https://actions-badge.atrox.dev/heymatthew/tcr/goto?ref=master)

Test, commit on success, reset on failure. This workflow was pioneered by Kent Beck in 2019.

https://www.youtube.com/watch?v=FFzHOyFeovE

# Goals and non-goals

The 'test' step is tricky. I'd really like a command which
- monitors the file system for changes
- when changes are found, runs command supplied by the user
- if command succeeds, calls `git add` on all files
- if command fails, calls `git checkout` and resets all files

This command should not commit your changes. Changes are added to the staging area and can be
committed when the user is ready using `git commit`

# Usage [draft]

```sh
tcr -c 'go test'
tcr -c 'bundle exec rspec'
```
