# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/),
and this project adheres to [Semantic Versioning](https://semver.org/).

## [Unreleased]

## [0.5.0] - 2026-02-01

### Added

- Optional per-cell ANSI color support in `canvas` package
- `canvas.Color` type with 8 basic ANSI colors (Black, Red, Green, Yellow, Blue, Magenta, Cyan, White)
- `canvas.SetColor()` method for setting pixels with color
- `WithColor()` option to enable color grid allocation
- `Color.ANSI()` method for escape sequence generation
- `ANSIReset()` function for reset escape sequence
- Color tests with visual flag support
- Updated demo program with color demonstrations (demos 21-24)

## [0.4.0] - 2026-02-01

### Added

- Circle drawing primitives in `draw` package
- `draw.Circle()` function for drawing circle outlines using Bresenham's midpoint algorithm
- `draw.CircleFilled()` function for drawing filled circles
- 8-way symmetry optimization for efficient circle rendering
- Updated demo program with circle demonstrations

## [0.3.0] - 2026-02-01

### Added

- Rectangle drawing primitives in `draw` package
- `draw.Rectangle()` function for drawing rectangle outlines
- `draw.RectangleFilled()` function for drawing filled rectangles
- Updated demo program with rectangle demonstrations

### Changed

- Refactored visual test helpers to shared file for reuse across test files

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
