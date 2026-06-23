# stringx

[![Go Reference](https://pkg.go.dev/badge/github.com/bycigo/stringx.svg)](https://pkg.go.dev/github.com/bycigo/stringx)

[English](README.md) | 简体中文

为 Go 提供支持 Unicode 的扩展字符串函数。

## 安装

```bash
go get github.com/bycigo/stringx
```

## 使用

更多示例请参见 [example_test.go](example_test.go)。

## API 参考

### 字符串函数

| 函数 | 说明 |
|----------|-------------|
| `Len(s string) int` | 返回字符串的字符（rune）数量，而非字节数 |
| `Reverse(s string) string` | 反转字符串 |
| `PadLeft(s string, length int, padding string) string` | 在左侧填充字符串 |
| `PadRight(s string, length int, padding string) string` | 在右侧填充字符串 |
| `PadBoth(s string, length int, padding string) string` | 在两侧填充字符串 |

### 大小写转换函数

| 函数 | 说明 |
|----------|-------------|
| `CamelCase(s string) string` | 转换为驼峰式（camelCase） |
| `PascalCase(s string) string` | 转换为帕斯卡式（PascalCase） |
| `SnakeCase(s string) string` | 转换为蛇形式（snake_case） |
| `KebabCase(s string) string` | 转换为短横线式（kebab-case） |

### 随机字符串函数

| 函数 | 说明 |
|----------|-------------|
| `Random(length int) string` | 生成指定长度的随机字符串 |

### 字节转换函数

| 函数 | 说明 |
|----------|-------------|
| `FromBytes(bytes []byte) string` | 零拷贝地将字节切片转换为字符串 |
| `ToBytes(s string) []byte` | 零拷贝地将字符串转换为字节切片 |

### 密码生成函数

| 函数 | 说明 |
|----------|-------------|
| `Password(length int, opt ...PasswordOption) string` | 生成加密安全的密码（默认使用 `crypto/rand`） |
| `NewPasswordFactory() PasswordFactory` | 创建基于 `math/rand/v2` 的工厂（速度快，非加密安全） |
| `NewSecuredPasswordFactory() PasswordFactory` | 创建基于 `crypto/rand` 的工厂（加密安全） |
| `NewPasswordFactoryWithSource(src PasswordSource) PasswordFactory` | 创建带自定义字符源的非加密安全工厂 |
| `NewSecuredPasswordFactoryWithSource(src PasswordSource) PasswordFactory` | 创建带自定义字符源的加密安全工厂 |

通过 `PasswordIncludes` 选择密码可包含的字符类型（`PasswordLetter` 字母、`PasswordNumber` 数字、`PasswordSymbol` 符号、`PasswordSpace` 空格）。生成结果保证每种被选中的类型至少出现一个字符。默认包含字母、数字和符号。


## 许可证

详情请参见 [LICENSE](LICENSE)。
