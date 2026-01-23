# Planx Plugin Protocol

## Versioning

Planx Plugin Protocol starts at **v4**.

Earlier versions (v1–v3) were internal drafts and never released.
They are intentionally removed.

There is NO backward compatibility requirement.

## Design Principles

- SDK-managed runtime
- Explicit session lifecycle
- Batch-only data model
- Opaque payloads
- Polyglot-friendly (gRPC)

## Specification Authority

The authoritative specification for Planx 4.0 lives in the [planx-spec](github.com/planx-lab/planx-spec) repository.
All behavior, formats, and contracts in this repository MUST conform to it. Local documentation must not redefine system contracts.
