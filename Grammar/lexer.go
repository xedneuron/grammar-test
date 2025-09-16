package grammar

/*
This code's sole purpose is to transform a report as a sanitized string
into tokens. Any spacing errors get added to the errors list. Tokens are
either words, numerals, or identifiers.
*/

// Tokens are words, numerals, or identifiers
type Token struct {
	Data string
	Form string
}
