package stringx_test

import (
	"fmt"

	"github.com/bycigo/stringx"
)

func ExampleLen() {
	fmt.Println(stringx.Len("hello"))
	fmt.Println(stringx.Len("你好"))
	fmt.Println(stringx.Len("🎉🎊"))
	// Output:
	// 5
	// 2
	// 2
}

func ExampleReverse() {
	fmt.Println(stringx.Reverse("hello"))
	fmt.Println(stringx.Reverse("你好世界"))
	fmt.Println(stringx.Reverse("abc123"))
	// Output:
	// olleh
	// 界世好你
	// 321cba
}

func ExamplePadLeft() {
	fmt.Println(stringx.PadLeft("go", 5, "x"))
	fmt.Println(stringx.PadLeft("hello", 10, "0"))
	fmt.Println(stringx.PadLeft("test", 8, "ab"))
	// Output:
	// xxxgo
	// 00000hello
	// ababtest
}

func ExamplePadRight() {
	fmt.Println(stringx.PadRight("go", 5, "x"))
	fmt.Println(stringx.PadRight("hello", 10, "0"))
	fmt.Println(stringx.PadRight("test", 8, "ab"))
	// Output:
	// goxxx
	// hello00000
	// testabab
}

func ExamplePadBoth() {
	fmt.Println(stringx.PadBoth("go", 6, "x"))
	fmt.Println(stringx.PadBoth("hi", 8, "ab"))
	fmt.Println(stringx.PadBoth("test", 10, "="))
	// Output:
	// xxgoxx
	// abahiaba
	// ===test===
}

func ExampleCamelCase() {
	fmt.Println(stringx.CamelCase("hello_world"))
	fmt.Println(stringx.CamelCase("hello-world"))
	fmt.Println(stringx.CamelCase("hello world"))
	fmt.Println(stringx.CamelCase("HelloWorld"))
	// Output:
	// helloWorld
	// helloWorld
	// helloWorld
	// helloWorld
}

func ExamplePascalCase() {
	fmt.Println(stringx.PascalCase("hello_world"))
	fmt.Println(stringx.PascalCase("hello-world"))
	fmt.Println(stringx.PascalCase("hello world"))
	fmt.Println(stringx.PascalCase("helloWorld"))
	// Output:
	// HelloWorld
	// HelloWorld
	// HelloWorld
	// HelloWorld
}

func ExampleSnakeCase() {
	fmt.Println(stringx.SnakeCase("helloWorld"))
	fmt.Println(stringx.SnakeCase("HelloWorld"))
	fmt.Println(stringx.SnakeCase("hello-world"))
	fmt.Println(stringx.SnakeCase("HTTPServer"))
	// Output:
	// hello_world
	// hello_world
	// hello_world
	// http_server
}

func ExampleKebabCase() {
	fmt.Println(stringx.KebabCase("helloWorld"))
	fmt.Println(stringx.KebabCase("HelloWorld"))
	fmt.Println(stringx.KebabCase("hello_world"))
	fmt.Println(stringx.KebabCase("HTTPServer"))
	// Output:
	// hello-world
	// hello-world
	// hello-world
	// http-server
}

func ExampleRandom() {
	// Generate a random string of length 10
	s := stringx.Random(10)
	fmt.Printf("Random string length: %d\n", len(s))
	// Output:
	// Random string length: 10
}

func ExampleFromBytes() {
	bytes := []byte("hello world")
	s := stringx.FromBytes(bytes)
	fmt.Println(s)
	// Output:
	// hello world
}

func ExampleToBytes() {
	s := "hello world"
	bytes := stringx.ToBytes(s)
	fmt.Println(string(bytes))
	// Output:
	// hello world
}

func ExamplePassword() {
	// Generate a secure password of length 16 using the default character sets
	// (letters, numbers and symbols).
	pwd := stringx.Password(16)
	fmt.Printf("Password length: %d\n", len(pwd))
	// Output:
	// Password length: 16
}

func ExamplePassword_includes() {
	// Restrict the password to letters and numbers only.
	pwd := stringx.Password(12, stringx.PasswordIncludes{
		stringx.PasswordLetter,
		stringx.PasswordNumber,
	})
	fmt.Printf("Password length: %d\n", len(pwd))
	// Output:
	// Password length: 12
}

func ExampleNewPasswordFactory() {
	// A non-cryptographic factory is faster and suitable for non-sensitive
	// scenarios such as test fixtures.
	factory := stringx.NewPasswordFactory()
	pwd := factory.MakePassword(20)
	fmt.Printf("Password length: %d\n", len(pwd))
	// Output:
	// Password length: 20
}

func ExampleNewSecuredPasswordFactory() {
	// A cryptographically secure factory backed by crypto/rand.
	factory := stringx.NewSecuredPasswordFactory()
	pwd := factory.MakePassword(20)
	fmt.Printf("Password length: %d\n", len(pwd))
	// Output:
	// Password length: 20
}

func ExampleNewPasswordFactoryWithSource() {
	// Provide a custom character source to fully control the alphabet.
	src := stringx.PasswordSource{
		stringx.PasswordLetter: []byte("abcdef"),
		stringx.PasswordNumber: []byte("0123"),
	}
	factory := stringx.NewPasswordFactoryWithSource(src)
	pwd := factory.MakePassword(16, stringx.PasswordIncludes{
		stringx.PasswordLetter,
		stringx.PasswordNumber,
	})
	fmt.Printf("Password length: %d\n", len(pwd))
	// Output:
	// Password length: 16
}
