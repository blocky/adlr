package prettyprinter

// The encoding/json package has limitations with marshaling errors.
// This wrapper properly stringifies errors.
// Can be used as a struct field to customized error json tags
type FieldError struct {
	Error error
}

// Wrap an error in a FieldError
func MakeFieldError(err error) FieldError {
	return FieldError{err}
}

// Marshal and stringify an error, with proper nil value representation
// as "null" per json spec https://www.json.org/json-en.html
func (be FieldError) MarshalJSON() ([]byte, error) {
	val := errorToString(be.Error)
	if val == "" {
		return prettyJSON(nil)
	}
	return prettyJSON(val)
}

// A key-value error for basic error standalone representation in json.
// Can be used as a struct field for with predetermined error json tag of
// "err"
type KeyValueError struct {
	Error FieldError `json:"err"`
}

// Wrap an error in a FieldError inside a KeyValueError
func MakeKeyValueError(err error) KeyValueError {
	return KeyValueError{MakeFieldError(err)}
}

// Marshal and stringify an error, with a json key of "err"
func (kve *KeyValueError) MarshalJSON() ([]byte, error) {
	type Alias KeyValueError
	return prettyJSON(&struct {
		*Alias
	}{
		Alias: (*Alias)(kve),
	})
}
