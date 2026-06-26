---
title: 'Hello World CLI'
type: 'feature'
created: '2026-06-26'
status: 'done'
route: 'one-shot'
---

# Hello World CLI

## Intent

**Problem:** No runnable CLI program exists in this project.

**Approach:** Add `cmd/hello/main.go` — a minimal Go main that prints `hello <name>` from the first CLI argument, with a usage error when the argument is missing.

## Suggested Review Order

**Entry point**

- Arg check + output: the whole program is here
  [`main.go:1`](../../cmd/hello/main.go#L1)
