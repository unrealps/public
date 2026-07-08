# Copilot Instructions for Go Fundamentals Workspace

This repository contains multiple independent Go exercises.

## Project Structure

- Each root folder is a separate, standalone Go program.
- Each program entry point is `main.go` inside its folder.
- Do not assume shared packages, shared state, or cross-folder dependencies.

Current exercise folders:

- `01.04-Hello_World/`
- `01.05-Stdout_Values/`
- `02-Var_Types/`
- `03-Math_Operators/`

## Editing Rules

- Make changes only in the folder requested by the user.
- Keep examples beginner-friendly and easy to read.
- Prefer straightforward Go syntax over advanced patterns.
- Preserve the existing lesson intent for each exercise.

## Run and Verify

When validating changes, run from the specific exercise folder:

```bash
go run main.go
```

If formatting is needed:

```bash
gofmt -w main.go
```

## Output Expectations

- Many exercises are output-based; match expected stdout exactly.
- Avoid adding extra logging, prompts, or decorative text unless requested.

## Dependency Guidance

- Prefer Go standard library only for these fundamentals exercises.
- Do not add external dependencies or new module-level setup unless explicitly requested.
