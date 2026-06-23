package stringx_test

import (
	"testing"

	"github.com/bycigo/stringx"
)

func TestCamelCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", ""},
		{"single word lowercase", "hello", "hello"},
		{"snake_case", "hello_world", "helloWorld"},
		{"kebab-case", "hello-world", "helloWorld"},
		{"space separated", "hello world", "helloWorld"},
		{"mixed separators", "hello_world-foo bar", "helloWorldFooBar"},
		{"PascalCase input", "HelloWorld", "helloWorld"},
		{"already camelCase", "helloWorld", "helloWorld"},
		{"multiple underscores", "hello__world", "helloWorld"},
		{"leading separator", "_hello_world", "helloWorld"},
		{"trailing separator", "hello_world_", "helloWorld"},
		{"unicode characters", "hello_世界", "hello世界"},
		{"with numbers", "hello_2_world", "hello2World"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := stringx.CamelCase(tt.input)
			if result != tt.expected {
				t.Errorf("CamelCase(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestPascalCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", ""},
		{"single word lowercase", "hello", "Hello"},
		{"snake_case", "hello_world", "HelloWorld"},
		{"kebab-case", "hello-world", "HelloWorld"},
		{"space separated", "hello world", "HelloWorld"},
		{"mixed separators", "hello_world-foo bar", "HelloWorldFooBar"},
		{"already PascalCase", "HelloWorld", "HelloWorld"},
		{"camelCase input", "helloWorld", "HelloWorld"},
		{"multiple underscores", "hello__world", "HelloWorld"},
		{"leading separator", "_hello_world", "HelloWorld"},
		{"trailing separator", "hello_world_", "HelloWorld"},
		{"unicode characters", "hello_世界", "Hello世界"},
		{"with numbers", "hello_2_world", "Hello2World"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := stringx.PascalCase(tt.input)
			if result != tt.expected {
				t.Errorf("PascalCase(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestSnakeCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", ""},
		{"single word lowercase", "hello", "hello"},
		{"single word uppercase", "HELLO", "hello"},
		{"camelCase", "helloWorld", "hello_world"},
		{"PascalCase", "HelloWorld", "hello_world"},
		{"snake_case already", "hello_world", "hello_world"},
		{"kebab-case", "hello-world", "hello_world"},
		{"space separated", "hello world", "hello_world"},
		{"mixed separators", "hello_worldU-foo bar", "hello_world_u_foo_bar"},
		{"multiple underscores", "hello__world", "hello_world"},
		{"leading separator", "_hello_world", "hello_world"},
		{"trailing separator", "hello_world_", "hello_world"},
		{"unicode characters", "hello世界", "hello世界"},
		{"with numbers", "hello2World", "hello2_world"},
		{"uppercase acronym", "XMLParser", "xml_parser"},
		{"mixed case complex", "myHTTPServer", "my_http_server"},
		{"acronym at end", "parseXML", "parse_xml"},
		{"multiple acronyms", "XMLToJSON", "xml_to_json"},
		{"single letter words", "aBC", "a_bc"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := stringx.SnakeCase(tt.input)
			if result != tt.expected {
				t.Errorf("SnakeCase(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestKebabCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", ""},
		{"single word lowercase", "hello", "hello"},
		{"single word uppercase", "HELLO", "hello"},
		{"camelCase", "helloWorld", "hello-world"},
		{"PascalCase", "HelloWorld", "hello-world"},
		{"snake_case", "hello_world", "hello-world"},
		{"kebab-case already", "hello-world", "hello-world"},
		{"space separated", "hello world", "hello-world"},
		{"mixed separators", "hello_worldU-foo bar", "hello-world-u-foo-bar"},
		{"multiple dashes", "hello--world", "hello-world"},
		{"leading separator", "-hello-world", "hello-world"},
		{"trailing separator", "hello-world-", "hello-world"},
		{"unicode characters", "hello世界", "hello世界"},
		{"with numbers", "hello2World", "hello2-world"},
		{"uppercase acronym", "XMLParser", "xml-parser"},
		{"mixed case complex", "myHTTPServer", "my-http-server"},
		{"acronym at end", "parseXML", "parse-xml"},
		{"multiple acronyms", "XMLToJSON", "xml-to-json"},
		{"single letter words", "aBC", "a-bc"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := stringx.KebabCase(tt.input)
			if result != tt.expected {
				t.Errorf("KebabCase(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}
