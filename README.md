# stringx

[![Go Reference](https://pkg.go.dev/badge/github.com/bycigo/stringx.svg)](https://pkg.go.dev/github.com/bycigo/stringx)

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

## License

See [LICENSE](LICENSE) for details.
