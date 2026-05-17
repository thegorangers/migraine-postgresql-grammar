# Stability commitment

This library follows semver. Within v1.x, the following are stable:

## Stable API

- **Go-package public surface** (`pkg/postgresql.*` exported types,
  functions, methods). No renames or removals without a major version bump.
- **ANTLR listener method names** tied to grammar rules
  (e.g., `EnterInsertstmt`, `ExitCreatetrigstmt`). Downstream linters
  embedding `*BasePostgreSQLParserListener` and overriding `EnterX` /
  `ExitX` methods rely on these names. They will not change.
- **Production coverage** documented under `docs/coverage.md` (when
  available). Anything listed as "stable" in coverage docs stays
  parseable across v1.x.

## What may change within v1.x

- **Internal generated code structure** — regenerated parser files
  (`postgresql_parser.go`, etc.) may differ between releases. Only the
  listener method API is the contract; internal helpers, struct field
  ordering, byte-identical generated output are NOT.
- **Grammar `.g4` source** — may be updated to add productions or fix
  bugs. Listener method names tied to existing rules stay stable; new
  rules add new listener methods.
- **Performance characteristics** — we may make the parser faster.
- **Coverage** — new productions move from "incomplete" or "unsupported"
  to "stable" as fixtures land (improvement, not breaking change).

## Breaking changes (v2.0+)

The following would require a v2.0:

- Renaming or removing public Go types/methods.
- Renaming listener methods for existing rules.
- Removing a previously-stable production from the "stable" coverage set.
- Changing the `module` path.

We will not ship v2.0 capriciously. A deprecation period of at least
one minor v1.x release will precede any v2.0 break.

## Reporting stability violations

If you observe a breaking change in a v1.x patch or minor release,
that is a bug — file an issue with `regression:` prefix.
