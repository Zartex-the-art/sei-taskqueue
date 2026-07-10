# Contributing

## Branching Strategy

- The `main` branch is the protected branch and should remain stable.
- Do not commit directly to `main`.
- Create feature branches using the naming convention:

```text
day<N>/<name>-<short-topic>
```

Examples:

```text
day1/madhu-job-schema-postgres
day1/hari-postgres-compose
day1/gayathri-test-conventions
day1/vishnu-readme-adr008
```

## Pull Requests

- Create one pull request per major task.
- Include a clear summary of the changes.
- Explain how the changes were tested.
- Wait for review before merging whenever possible.

## Code Quality

Before opening a pull request, run:

```bash
go build ./...
go vet ./...
go test ./...
```

Update documentation whenever your changes affect project behavior.