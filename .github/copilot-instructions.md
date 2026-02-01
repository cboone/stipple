# GitHub Copilot Instructions

## PR Review Checklist (CRITICAL)

<!-- KEEP THIS SECTION UNDER 4000 CHARS - Copilot only reads first ~4000 -->

### Documentation Files

- **Markdown tables use single pipes**: All tables in this repo use standard single `|` separators, NOT double `||`. Do not flag table formatting.
- **Code examples are pseudocode**: Examples in planning docs (PLAN.md, ROADMAP.md) are illustrative pseudocode, not production code. Do not flag missing bounds checks, error handling, or allocation details.
- **ROADMAP vs PLAN scope differs intentionally**: ROADMAP.md defines the full v1 API vision. PLAN.md is a focused implementation plan that may defer features (like text rendering). This is intentional, not an inconsistency.

### Code Patterns

- **No io/ package exists**: This codebase does not have an `io/` package. Do not suggest renaming it.
- **Thread safety is a non-goal for v1**: Per ROADMAP.md, thread-safety guarantees are explicitly out of scope for v1.0.
- **Constructor validation is intentionally minimal**: `New()` does not validate dimensions. Invalid constructor args (negative/zero dimensions) are programmer errors that panic on allocation - this is idiomatic Go. Runtime bounds checking uses silent ignore pattern (see below).
- **Out-of-bounds coordinates are silently ignored**: Per project conventions, pixel operations on out-of-bounds coordinates return early without error. This enables partial rendering of shapes that extend beyond canvas edges. Do not flag missing bounds validation.
