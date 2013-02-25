package parser

// A Match represents the result of a pattern matching
// process on the content of an document.
// It indicates whether the pattern was found and if yet,
// the lines in which it was located and the matched text.
type Match struct {
	Found   bool
	Matches []string
}

// Found create a new Match which represents a successful match.
func Found(matches []string) Match {
	return Match{
		Found:   true,
		Matches: matches,
	}
}

// NotFound create a new Match which represents
// an unsuccessful match.
func NotFound() Match {
	return Match{
		Found: false,
	}
}
