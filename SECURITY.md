# Security Policy

`migraine-postgresql-grammar` is a Go library providing ANTLR4-generated
bindings for a PostgreSQL grammar. It is consumed by the migraine CLI's lint
rules and by other tooling that needs to parse PostgreSQL statements. **This
library does not execute SQL** — it only parses it for lint and analysis
purposes.

## Supported Versions

| Version  | Supported          |
| -------- | ------------------ |
| v1.x.x   | :white_check_mark: |
| < v1.0   | :x:                |

## Reporting a Vulnerability

Please report suspected vulnerabilities privately to
**security@thegorangers.com**.

- We aim to acknowledge reports within **48 hours**.
- Once a fix is available, we coordinate public disclosure after a **30-day
  grace period** to give downstream consumers time to upgrade.
- Please do not open public GitHub issues for security-sensitive reports.

## Scope

In scope for security reports:

- Parser crashes or panics on well-formed PostgreSQL syntax.
- Hangs or unbounded memory/CPU consumption triggered by adversarial input
  (e.g. quadratic blowup, stack exhaustion on pathological nesting).
- Security-relevant grammar mismatches — for example, the lexer failing to
  recognize a keyword in a way that causes downstream tools (lint, policy
  enforcement) to silently miss its meaning.

## Out of Scope

- **PostgreSQL server security.** Issues in the database engine itself
  belong upstream at https://www.postgresql.org/support/security/.
- **Caller-side SQL injection.** This library does not execute SQL; safe
  handling of user-supplied SQL is the responsibility of the calling
  application.
- **`antlr4-go/antlr/v4` runtime issues.** Bugs in the generated ANTLR
  runtime should be reported to that project upstream.
- **Missing grammar coverage.** Unsupported syntax is a feature gap, not a
  vulnerability — please file a regular bug report (see future
  `docs/coverage.md` for the supported surface).

## Threat Model Summary

The primary attack surface is **malformed or adversarial SQL input** leading
to denial-of-service in the parser. The library handles no secrets, no
network I/O, and no database connections, so the classic service-side
threats (auth bypass, data exfiltration, privilege escalation) do not apply.
