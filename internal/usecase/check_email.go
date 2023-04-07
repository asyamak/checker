package usecase

import (
	"regexp"
	"strings"
)

type EmailCheckerUsecase interface {
	CheckEmail(s []string) []string
}

type EmailChecker struct{}

func NewEmailChecker() *EmailChecker {
	return &EmailChecker{}
}

func (e *EmailChecker) CheckEmail(s []string) []string {
	var results []string
	res := strings.Join(s, "")
	emailRegex := regexp.MustCompile(`\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}\b`)
	emails := emailRegex.FindAllString(res, -1)
	results = append(results, emails...)

	return results
}
