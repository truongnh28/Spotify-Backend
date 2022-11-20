package helper

import (
	"fmt"
	"regexp"
	"strings"
)

var domainName = "@vng.com.vn"
var HOST = "https://gitlab.zalopay.vn/"

func ConvDomainsToEmails(domains []string) []string {
	emails := []string{}
	for i := 0; i < len(domains); i++ {
		emails = append(emails, fmt.Sprintf("%s%s", domains[i], domainName))
	}
	return emails
}

func ConvAccountToEmail(domain string) string {
	return fmt.Sprintf("%s%s", domain, domainName)
}

// change "zlp:top:cicd:everest" to "top / cicd / everest"
func GitlabProjectName(name string) string {
	m1 := regexp.MustCompile("^zlp:")
	m2 := regexp.MustCompile(":")
	name = m1.ReplaceAllString(name, "")
	return m2.ReplaceAllString(name, " / ")
}

func GitlabURL(name string) string {
	m1 := regexp.MustCompile("^zlp:")
	m2 := regexp.MustCompile(":")
	name = m1.ReplaceAllString(name, "")
	return HOST + m2.ReplaceAllString(name, "/")
}

func ConvEmailToAccount(email string) string {
	if strings.Index(email, "@") == -1 {
		return email
	}
	return email[0:strings.Index(email, "@")]
}
