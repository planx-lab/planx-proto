# AI RULES — PLANX PROTO (MANDATORY)

## Authority Documents

Before modifying protocol definitions:
1. [planx-architecture.md](../planx-architecture.md)
2. [planx-ai-guardrails.md](../planx-ai-guardrails.md)
3. [AI_CONTRACT.md](../AI_CONTRACT.md)

---

## SCOPE

This repo is the single source of truth for Planx Plugin Protocol v4.

- **Status**: ARCHITECTURE-FROZEN
- **No business logic** inside proto files
- **No runtime assumptions**
- **Backward compatibility**: NOT required before v4.0 GA

---

## MODIFICATION RULES

AI MUST NOT:
- Redesign protocol semantics
- Merge services
- Add fields without explicit justification

AI MAY:
- Generate code from existing proto
- Add comments for clarification ONLY
