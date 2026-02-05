# AGENTS.md

This file provides guidance for AI coding assistants working with the OpenHue Go codebase.

## Project Overview

OpenHue Go is a Go library for interacting with Philips Hue smart lighting systems. It is based on the [OpenHue API](https://github.com/openhue/openhue-api) specification and uses [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen) to generate most of its code automatically.

**Key Point**: Most of the code in this repository is auto-generated. Manual edits should be limited to specific files.

## File Structure

### Auto-Generated Files (DO NOT EDIT)
- `openhue.gen.go` - Generated client and type definitions from the OpenAPI spec
  - **NEVER modify this file manually**
  - Regenerate using `make generate` when the OpenAPI spec changes

### Manually Maintained Files (CAN EDIT)
- `openhue.go` - Core library wrapper and high-level interfaces
- `auth.go` - Authentication helpers
- `discovery.go` - Bridge discovery functionality (mDNS and URL-based)
- `helpers.go` - Utility functions
- `error.go` - Error handling utilities
- `testing.go` - Testing utilities
- `doc.go` - Package documentation
- `openhue_test.go` - Tests

## Development Workflow

### Code Generation
```bash
# Generate from the latest OpenHue API spec (default)
make generate

# Generate from a local spec file
make generate spec=/path/to/openhue.yaml
```

**Important**: After regenerating `openhue.gen.go`, always verify that existing functionality still works.

### Testing
```bash
# Run all tests with coverage
make test

# Run tests manually
go test -cover ./...
```

### Dependency Management
```bash
# Tidy go.mod
make tidy
```

## Key Dependencies

- `github.com/grandcat/zeroconf` - mDNS discovery
- `github.com/oapi-codegen/runtime` - OpenAPI code generation runtime
- `github.com/stretchr/testify` - Testing framework
- `gopkg.in/yaml.v3` - YAML parsing for configuration

## Architecture

### Core Components

1. **Home** - Main entry point representing a Philips Hue bridge
   - Created via `openhue.NewHome()`
   - Provides high-level methods for interacting with devices

2. **BridgeDiscovery** - Discovers bridges on the local network
   - Uses mDNS first, falls back to discovery.meethue.com
   - Configurable timeout and options

3. **Authenticator** - Handles bridge authentication
   - Implements the link button authentication flow
   - Returns API keys for subsequent requests

4. **Generated Client** - Low-level API client (in `openhue.gen.go`)
   - Auto-generated from OpenAPI spec
   - Provides raw API access

## Important Considerations for AI Agents

### When Making Changes

1. **Check if it's generated code**
   - If the change affects `openhue.gen.go`, you must modify the upstream OpenAPI spec instead
   - Never directly edit `openhue.gen.go`

2. **Prefer high-level wrappers**
   - Add new functionality to manually maintained files (`openhue.go`, `helpers.go`, etc.)
   - Wrap generated client methods with ergonomic interfaces

3. **Configuration**
   - The library uses a well-known configuration file for bridge credentials
   - See https://www.openhue.io/cli/setup#manual-configuration for format details

4. **Error handling**
   - Use `openhue.CheckErr()` helper for simple error checking
   - Provide detailed error messages for authentication and network failures

### Testing Guidelines

- Test files should focus on manually maintained code
- Mock external dependencies (mDNS, HTTP calls) where possible
- Bridge discovery and authentication require real hardware to fully test

### Code Style

- Follow standard Go conventions
- Use Go 1.23+ features (as specified in go.mod)
- Keep functions focused and composable
- Document all exported types and functions

### Git Commit Guidelines

- **NEVER include co-author attribution** in commit messages
- Do NOT add `Co-Authored-By:` lines to commits
- Keep commit messages clear and focused on the changes made
- Follow conventional commits format: `type: description`
- Examples: `fix:`, `feat:`, `docs:`, `refactor:`, `test:`

## Common Tasks

### Adding a New Helper Method
1. Add the method to `openhue.go` or `helpers.go`
2. Ensure it wraps generated client methods appropriately
3. Add tests in `openhue_test.go`
4. Update README.md with usage examples if needed

### Updating the OpenAPI Spec Version
1. Run `make generate` to regenerate from latest spec
2. Review changes in `openhue.gen.go`
3. Update any breaking changes in wrapper code
4. Run `make test` to ensure compatibility
5. Update `go.mod` if needed with `make tidy`

### Adding New Discovery Methods
1. Extend `discovery.go` with new discovery mechanism
2. Maintain backward compatibility with existing methods
3. Add appropriate timeout and option handling
4. Test with real bridges if possible

## Resources

- [OpenHue API Specification](https://github.com/openhue/openhue-api)
- [oapi-codegen Documentation](https://github.com/oapi-codegen/oapi-codegen)
- [Philips Hue API Documentation](https://developers.meethue.com/)
- [OpenHue Website](https://www.openhue.io/)

## License

This project is licensed under Apache License 2.0. See LICENSE file for details.
