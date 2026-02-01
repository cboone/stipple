# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/),
and this project adheres to [Semantic Versioning](https://semver.org/).

## [Unreleased]

## [0.2.0] - 2026-02-01

### Added

- Bresenham's line drawing algorithm in `draw` package
- `draw.Line()` function for drawing lines between two points
- Handles all octants (steep/shallow, positive/negative slopes)
- Updated demo program with line demonstrations (horizontal, vertical, diagonal)

## [0.1.0] - 2026-01-25

### Added

- Core braille canvas with pixel-level control (`canvas` package)
- `canvas.New()` constructor with functional options
- `canvas.Set()`, `canvas.Unset()`, `canvas.Toggle()`, `canvas.Get()` for pixel operations
- `canvas.Frame()` for rendering to string
- `canvas.Clear()` for resetting the canvas
- `WithInvertedY()` option for Y-axis inversion
- Braille encoding constants and pixel mapping
- Demo program (`examples/demo/`)
- Go module initialization (`go.mod`)
- GitHub Actions CI workflow (test, lint, build, format)
- GitHub Actions release workflow with goreleaser
- Makefile with development targets
- golangci-lint configuration
- goreleaser configuration for library releases
- EditorConfig for consistent formatting
- Dependabot configuration for dependency updates
- Contributing guidelines (`CONTRIBUTING.md`)
