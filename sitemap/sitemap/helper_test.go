package sitemap

import "testing"

func TestParseLink(t *testing.T) {
	tests := []struct {
		link     string
		expected string
	}{
		{
			link:     "https://mrg.sh/",
			expected: "https://mrg.sh",
		},
		{
			link:     "http://mrg.sh/test/",
			expected: "http://mrg.sh/test",
		},
		{
			link:     "https://mrg.sh?query=test&n=2",
			expected: "https://mrg.sh?query=test&n=2",
		},
	}

	for _, test := range tests {
		result := ParseLink(test.link)

		if result.String() != test.expected {
			t.Errorf("Expected [%s], got [%s]", test.expected, result.String())
		}
	}
}

func TestIsValidHostUrl(t *testing.T) {
	tests := []struct {
		url1     string
		url2     string
		expected bool
	}{
		{
			url1:     "https://mrg.sh",
			url2:     "http://mrg.sh",
			expected: true,
		},
		{
			url1:     "https://test.sh",
			url2:     "http://mrg.sh",
			expected: false,
		},
		{
			url1:     "https://mrg.sh",
			url2:     "mailto:me@mrg.sh",
			expected: false,
		},
	}

	for _, test := range tests {
		result := IsValidHostUrl(test.url1, test.url2)

		if result != test.expected {
			t.Errorf("Expected [%t], got [%t]", test.expected, result)
			t.Errorf("Url1: [%s], Url2: [%s]", test.url1, test.url2)
		}
	}
}

func TestIsValidUrl(t *testing.T) {
	tests := []struct {
		link     string
		expected bool
	}{
		{
			link:     "https://mrg.sh",
			expected: true,
		},
		{
			link:     "test",
			expected: false,
		},
		{
			link:     "mailto:me@mrg.sh",
			expected: false,
		},
	}

	for _, test := range tests {
		result := IsValidUrl(test.link)

		if result != test.expected {
      t.Errorf("Expected [%t], got [%t], link: [%s]", test.expected, result, test.link)
		}
	}
}
