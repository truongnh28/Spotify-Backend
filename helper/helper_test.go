package helper

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvAccountToEmail(t *testing.T) {
	testcases := []struct {
		name     string
		domain   string
		expected string
	}{
		{
			name:     "halpm",
			domain:   "halpm",
			expected: "halpm@vng.com.vn",
		},
		{
			name:     "random_user",
			domain:   "random_user",
			expected: "random_user@vng.com.vn",
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			got := ConvAccountToEmail(tc.domain)
			assert.Equal(t, tc.expected, got)
		})
	}
}

func TestConvDomainsToEmails(t *testing.T) {
	testcases := []struct {
		name     string
		domains  []string
		expected []string
	}{
		{
			name:     "Empty list",
			domains:  []string{},
			expected: []string{},
		},
		{
			name: "Random domains list",
			domains: []string{
				"halpm",
				"test",
				"anonymous",
				"random_user",
			},
			expected: []string{
				"halpm@vng.com.vn",
				"test@vng.com.vn",
				"anonymous@vng.com.vn",
				"random_user@vng.com.vn",
			},
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			got := ConvDomainsToEmails(tc.domains)
			assert.Equal(t, tc.expected, got)
		})
	}
}

func TestConvEmailToAccount(t *testing.T) {
	testcases := []struct {
		name     string
		email    string
		expected string
	}{
		{
			name:     "Full email",
			email:    "halpm@vng.vn.com",
			expected: "halpm",
		},
		{
			name:     "Empty email",
			email:    "",
			expected: "",
		},
		{
			name:     "Email without @",
			email:    "halpm#vng.com.vn",
			expected: "halpm#vng.com.vn",
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			got := ConvEmailToAccount(tc.email)
			assert.Equal(t, tc.expected, got)
		})
	}
}

func TestGitlabProjectName(t *testing.T) {
	testcases := []struct {
		name        string
		projectName string
		expected    string
	}{
		{
			name:        "Full project name",
			projectName: "zlp:top:cicd:everest",
			expected:    "top / cicd / everest",
		},
		{
			name:        "Project name without zlp",
			projectName: "top:cicd:everest",
			expected:    "top / cicd / everest",
		},
		{
			name:        "Wrong format project name",
			projectName: "zlp-top-cicd-everest",
			expected:    "zlp-top-cicd-everest",
		},
		{
			name:        "Wrong format project name with some : symbol",
			projectName: "zlp:top-cicd:everest",
			expected:    "top-cicd / everest",
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			got := GitlabProjectName(tc.projectName)
			assert.Equal(t, tc.expected, got)
		})
	}
}

func TestGitlabURL(t *testing.T) {
	testcases := []struct {
		name        string
		projectName string
		expected    string
	}{
		{
			name:        "Full project name",
			projectName: "zlp:top:cicd:everest",
			expected:    "https://gitlab.zalopay.vn/top/cicd/everest",
		},
		{
			name:        "Project name without zlp",
			projectName: "top:cicd:everest",
			expected:    "https://gitlab.zalopay.vn/top/cicd/everest",
		},
		{
			name:        "Wrong format project name",
			projectName: "zlp-top-cicd-everest",
			expected:    "https://gitlab.zalopay.vn/zlp-top-cicd-everest",
		},
		{
			name:        "Wrong format project name with some : symbol",
			projectName: "zlp:top-cicd:everest",
			expected:    "https://gitlab.zalopay.vn/top-cicd/everest",
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			got := GitlabURL(tc.projectName)
			assert.Equal(t, tc.expected, got)
		})
	}
}
