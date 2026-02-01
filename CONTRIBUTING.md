# Contributing to brodot

Thank you for your interest in contributing to brodot.

## Development Setup

1. Clone the repository:

   ```bash
   git clone https://github.com/cboone/brodot.git
   cd brodot
   ```

2. Install development tools:

   ```bash
   make tools
   ```

3. Run all checks:

   ```bash
   make all
   ```

## Code Style

- Use descriptive variable names; avoid abbreviations
- All public types and functions require doc comments
- Doc comments must end with a period
- Zero external dependencies in core packages (standard library only)
- See [.claude/CLAUDE.md](.claude/CLAUDE.md) for detailed conventions

## Commit Messages

This project uses [Conventional Commits](https://www.conventionalcommits.org/):

```
<type>: <short description>
```

Types:

- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation only
- `refactor`: Code change that neither fixes a bug nor adds a feature
- `test`: Adding or updating tests
- `build`: Changes to build system or dependencies
- `ci`: Changes to CI configuration
- `chore`: Other changes that do not modify source or test files

Examples:

```
feat: add circle drawing with Bresenham algorithm
fix: correct off-by-one error in line rendering
docs: update installation instructions
```

## Testing

Run the test suite:

```bash
make test
```

Run tests with visual output (for debugging):

```bash
make test-visual
```

Generate coverage report:

```bash
make coverage
```

## Pull Request Guidelines

1. Create a feature branch from `main`
2. Make your changes following the code style guidelines
3. Ensure all tests pass: `make test`
4. Ensure code is formatted: `make fmt`
5. Ensure linting passes: `make lint`
6. Write clear commit messages following Conventional Commits
7. Open a pull request with a clear description of your changes

## Questions

If you have questions, open an issue on GitHub.
