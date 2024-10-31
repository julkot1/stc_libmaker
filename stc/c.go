package stc

import (
	"errors"
	"regexp"
)

var cKeywords = map[string]bool{
	"auto": true, "break": true, "case": true, "char": true, "const": true,
	"continue": true, "default": true, "do": true, "double": true, "else": true,
	"enum": true, "extern": true, "float": true, "for": true, "goto": true,
	"if": true, "int": true, "long": true, "register": true, "return": true,
	"short": true, "signed": true, "sizeof": true, "static": true, "struct": true,
	"switch": true, "typedef": true, "union": true, "unsigned": true, "void": true,
	"volatile": true, "while": true,
}

// IsValidCFunctionName checks if a string is a valid C function name
func IsValidCFunctionName(name string) error {

	validName := regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*$`)

	if !validName.MatchString(name) {
		return errors.New("Invalid name: " + name)
	}

	if cKeywords[name] {
		return errors.New("Invalid name (is C keyword): " + name)
	}

	return nil
}
