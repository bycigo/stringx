# stringx

[![Go Reference](https://pkg.go.dev/badge/github.com/bycigo/stringx.svg)](https://pkg.go.dev/github.com/bycigo/stringx)

English | [简体中文](README_zh.md)

Extended string functions for Go with Unicode support.

## Installation

```bash
go get github.com/bycigo/stringx
```

## Usage

More examples can be found in the [example_test.go](example_test.go).

## API Reference

### String Functions

| Function | Description |
|----------|-------------|
| `Len(s string) int` | Returns the rune count of a string |
| `Reverse(s string) string` | Reverses a string |
| `PadLeft(s string, length int, padding string) string` | Pads string on the left |
| `PadRight(s string, length int, padding string) string` | Pads string on the right |
| `PadBoth(s string, length int, padding string) string` | Pads string on both sides |

### Case Conversion Functions

| Function | Description |
|----------|-------------|
| `CamelCase(s string) string` | Converts to camelCase |
| `PascalCase(s string) string` | Converts to PascalCase |
| `SnakeCase(s string) string` | Converts to snake_case |
| `KebabCase(s string) string` | Converts to kebab-case |

### Random String Functions

| Function | Description |
|----------|-------------|
| `Random(length int) string` | Generates a random string |

### Byte Conversion Functions

| Function | Description |
|----------|-------------|
| `FromBytes(bytes []byte) string` | Zero-copy conversion from bytes to string |
| `ToBytes(s string) []byte` | Zero-copy conversion from string to bytes |

### Password Functions

| Function | Description |
|----------|-------------|
| `Password(length int, opt ...PasswordOption) string` | Generates a cryptographically secure password (uses `crypto/rand` by default) |
| `NewPasswordFactory() PasswordFactory` | Creates a factory backed by `math/rand/v2` (fast, non-cryptographic) |
| `NewSecuredPasswordFactory() PasswordFactory` | Creates a factory backed by `crypto/rand` (secure) |
| `NewPasswordFactoryWithSource(src PasswordSource) PasswordFactory` | Creates a non-cryptographic factory with a custom character source |
| `NewSecuredPasswordFactoryWithSource(src PasswordSource) PasswordFactory` | Creates a secure factory with a custom character source |

Use `PasswordIncludes` to choose which character kinds (`PasswordLetter`, `PasswordNumber`, `PasswordSymbol`, `PasswordSpace`) the password may contain. The result always includes at least one character from each selected kind. The default includes letters, numbers and symbols.

## License

See [LICENSE](LICENSE) for details.
