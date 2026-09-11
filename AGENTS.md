# Planx Plugin Protocol v4

## Collaboration and references

Use [workspace guidance](../AGENTS.md) for task completion, authorization and
proportionate checks. Read relevant clauses of the [canonical contract](../planx-spec/AI_CONTRACT.md),
[architecture](../planx-spec/planx-architecture.md) and accepted
[ADR-017](../planx-spec/adr/017-builtin-typed-data-integration.md) when their
subject changes. Use [repo.lock](repo.lock) for source ownership. Do not reread
the entire specification for an unrelated edit.

## Ownership and changes

This repository is the protocol source of truth. Keep business logic, SDK helpers
and Engine implementation out of the protocol. Do not redesign services or add
fields without explicit justification and an authorized contract change.

ADR-017's approved typed-data and consumer updates may include necessary protocol
changes; synchronize generated code, SDK/Engine consumers and contract tests.
Unrelated frozen-model changes still require approval. The historical
comments-only rule must not veto an already approved contract change.

Generate from `v4/` using the checked-in Buf configuration; do not hand-edit
generated Go files. Before v4.0 GA, follow the repository's v4 compatibility
policy. A generation command alone is not end-to-end protocol verification.
