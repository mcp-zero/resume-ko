package analyzer

import (
	"reflect"
	"testing"
)

func TestExtractKeywords(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		expected map[string]bool
	}{
		{
			name: "Simple text",
			text: "软件工程师 software engineer",
			expected: map[string]bool{
				"软件":       true,
				"工程":       true,
				"software": true,
				"engineer": true,
			},
		},
		{
			name: "Text with punctuation",
			text: "软件, 工程师; developer.",
			expected: map[string]bool{
				"软件":        true,
				"工程":        true,
				"工程师":       true,
				"developer": true,
			},
		},
		{
			name:     "Empty text",
			text:     "",
			expected: map[string]bool{},
		},
		{
			name: "Text with short words",
			text: "I am a software engineer",
			expected: map[string]bool{
				"am":       true,
				"software": true,
				"engineer": true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := extractKeywords(tt.text)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("extractKeywords() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestCompareKeywords(t *testing.T) {
	tests := []struct {
		name            string
		resumeKeywords  map[string]bool
		jobKeywords     map[string]bool
		expectedMatched []string
		expectedMissing []string
	}{
		{
			name: "Perfect match",
			resumeKeywords: map[string]bool{
				"software":  true,
				"engineer":  true,
				"developer": true,
			},
			jobKeywords: map[string]bool{
				"software":  true,
				"engineer":  true,
				"developer": true,
			},
			expectedMatched: []string{"software", "engineer", "developer"},
			expectedMissing: []string{},
		},
		{
			name: "Partial match",
			resumeKeywords: map[string]bool{
				"software": true,
				"engineer": true,
			},
			jobKeywords: map[string]bool{
				"software":  true,
				"engineer":  true,
				"developer": true,
			},
			expectedMatched: []string{"software", "engineer"},
			expectedMissing: []string{"developer"},
		},
		{
			name: "No match",
			resumeKeywords: map[string]bool{
				"manager": true,
				"sales":   true,
			},
			jobKeywords: map[string]bool{
				"software":  true,
				"engineer":  true,
				"developer": true,
			},
			expectedMatched: []string{},
			expectedMissing: []string{"software", "engineer", "developer"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMatched, gotMissing := compareKeywords(tt.resumeKeywords, tt.jobKeywords)

			// Sort slices for comparison
			if !stringSlicesEqualIgnoreOrder(gotMatched, tt.expectedMatched) {
				t.Errorf("compareKeywords() matched = %v, expected %v", gotMatched, tt.expectedMatched)
			}

			if !stringSlicesEqualIgnoreOrder(gotMissing, tt.expectedMissing) {
				t.Errorf("compareKeywords() missing = %v, expected %v", gotMissing, tt.expectedMissing)
			}
		})
	}
}

func TestAnalyze(t *testing.T) {
	resume := "I am a software engineer with experience in Go programming"
	jobDescription := "We need a software engineer proficient in Go and Python"

	result, err := Analyze(resume, jobDescription)
	if err != nil {
		t.Fatalf("Analyze() error = %v", err)
	}

	// Check that we have results
	if result == nil {
		t.Error("Analyze() returned nil result")
		return
	}

	// Check that we have both matched and missing keywords
	if len(result.MatchedKeywords) == 0 {
		t.Error("Analyze() returned no matched keywords")
	}

	if len(result.MissingKeywords) == 0 {
		t.Error("Analyze() returned no missing keywords")
	}
}

// Helper function to compare string slices ignoring order
func stringSlicesEqualIgnoreOrder(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	// Create maps to count occurrences
	aMap := make(map[string]int)
	bMap := make(map[string]int)

	for _, s := range a {
		aMap[s]++
	}

	for _, s := range b {
		bMap[s]++
	}

	// Compare maps
	for key, count := range aMap {
		if bMap[key] != count {
			return false
		}
	}

	return true
}
