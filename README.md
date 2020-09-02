# React Gin Starter

Web starter kit using: React at frontend and golang, gin at backend.

## Development

Frontend pages run on port 1234:

```bash
$ make frontend-install
$ make frontend-run
```

Backend pages run on port 5678:

```bash
$ make backend-run
```

**NOTE:** Browser should enable CORS when local develpment.

## Build

```bash
$ make build-local
# For Linux
$ make build-cross
```

Then run ./build/main, website will be served on port 5678.
