# Repository Guidelines

## Project Structure

This is a Go module at `github.com/tizips/shenhuo`.

- `server/admin` contains the admin API service entrypoint, bootstrap, config, routes, business handlers, HTTP request/response DTOs, migrations, and console commands.
- `server/web` contains the web API service entrypoint, bootstrap, config, routes, business handlers, and HTTP request/response DTOs.
- `model` contains shared GORM model definitions and table-name constants.
- `docker/admin.Dockerfile` and `docker/web.Dockerfile` build the two service binaries from `server/admin` and `server/web`.

The module currently uses a local replace:

```go
replace github.com/herhe-com/framework => ../framework
```

Keep this in mind when running commands from outside the `herhe-com` workspace.

## Commands

Run commands from this directory unless noted otherwise.

```bash
go test ./...
go run ./server/admin server
go run ./server/web server
go run ./server/admin developer
docker build -t herhe/admin:1.0.0 -f docker/admin.Dockerfile .
docker build -t herhe/web:1.0.0 -f docker/web.Dockerfile .
```

Use the repository's normal command wrapper when operating in this workspace:

```bash
rtk go test ./...
rtk go run ./server/admin server
```

## Development Notes

- Service configuration is registered through `server/*/config` packages using `facades.Config()`.
- Routes are wired in `server/*/route`, with auth middleware applied at the route-group level.
- Admin migrations live under `server/admin/migration`; keep schema changes and model changes in sync.
- Shared database tables should keep their constants in `model` and implement `TableName()`.
- Preserve the existing package layout when adding endpoints: route -> request DTO -> biz handler -> response DTO.
- Run `gofmt` on changed Go files before finishing.

## Verification

Prefer targeted checks for the code touched, then run:

```bash
go test ./...
```

## Git Workflow Constraints

- When merging feature work into `main`, keep `main` history linear. Prefer `git merge --ff-only <branch>` and do not create commits named like `Merge <branch> into main`.
- If `main` cannot fast-forward, rebase the feature branch onto `main` or use a squash commit with a conventional commit message instead of a merge commit.
- If a merge commit is accidentally pushed to `main`, remove it by resetting `main` to the merge commit's first parent, re-merge with `git merge --ff-only <branch>`, then update the remote with `git push --force-with-lease origin main`.

At the time this file was created, `go test ./...` reported no test packages.
