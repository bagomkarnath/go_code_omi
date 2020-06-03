package main

import (
	"fmt"
	"strings"
)

func main() {
	var s CaseConverter = &sentence{sent: "Good Morning"}
	fmt.Println(s.UpperCaseConvert())
	fmt.Println(s.LowerCaseConvert())
}

type LowerCaseConverter interface {
	LowerCaseConvert() string
}

type UpperCaseConverter interface {
	UpperCaseConvert() string
}

type CaseConverter interface {
	UpperCaseConverter
	LowerCaseConverter
}

type sentence struct {
	sent string
}

func (s *sentence) LowerCaseConvert() string {
	return strings.ToLower(s.sent)
}

func (s *sentence) UpperCaseConvert() string {
	return strings.ToUpper(s.sent)
}
