package forms

import (
	"fmt"
	"net/url"
	"strings"
	"unicode/utf8"
)

// Form anonymously embeds url.Values to hold form data and validation errors.
type Form struct {
	url.Values
	Errors errors
}

// New returns a new Form with form values and empty errors.
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Required checks that fields are present and not blank.
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

// MaxLength check specific field to contain max d characters.
func (f *Form) MaxLength(field string, d int) {
	value := f.Get(field)
	if value == "" {
		return
	}
	if utf8.RuneCountInString(value) > d {
		f.Errors.Add(field, fmt.Sprintf("This field is too long (max %d characters)", d))
	}
}

// PermittedValues checks that field has one of opts as value.
func (f *Form) PermittedValues(field string, opts ...string) {
	value := f.Get(field)
	if value == "" {
		return
	}
	for _, opt := range opts {
		if value == opt {
			return
		}
	}
	f.Errors.Add(field, "This field is invalid")
}

// Valid returns true if there are no errors.
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
