# Pretty Printer
[![Build Status](https://www.travis-ci.com/blocky/prettyprinter.svg?branch=main)](https://www.travis-ci.com/blocky/prettyprinter)
[![Go Report Card](https://goreportcard.com/badge/github.com/blocky/prettyprinter)](https://goreportcard.com/report/github.com/blocky/prettyprinter)

Pretty printer is a simple printer designed to prettify everything. It can process simple to complex structs into readable json output.

To get started:
```
go get -u github.com/blocky/prettyprinter
```

To create a pretty printer with default outputs to `stdout` and `stderr`, use:
```golang
p := prettyprinter.NewPrettyPrinter()
```

Add the thing you want to print out with `Add()`:
```golang
var str string = "string"
p.Add(struct{Str string}{str})
```
```json
{
 "Str": "string"
}
```

Pretty printer allows method chaining. To add, print, and check for printing error:
```golang
var integer int = 1234
err := p.Add(struct{Int int}{integer}).
	StdoutDump().
	Error()
```
```json
{
 "Int": 1234
}
```

Pretty printer will have purged the thing added with `Add()` after attempting to print via `Flush()`, and can be continued to be reused to print more.


Pretty printer can also attempt to write its own printing error to `stderr`:
```golang
var str string = "string"
err := p.Add(struct{Str string}{str}).
	StdoutDump().
	StderrDumpOnError().
	Error()
```

Want to write to something other than `stdout` or `stderr`? Use `Dump()` to write to anything that implements the `io.Writer` interface, such as files. Pretty printer will either write all or nothing:
```golang
file, err := os.Open("./folder/file.txt")
defer file.Close()

var string string = "\"{msg: writing some json}\""
err := p.Add(integer).
	Dump(file).
	Error()
```
```text
"{msg: writing some json}"
```

To print more complicated things like multi-fielded structs, just pass the struct to pretty printer. Make sure the fields you want to print are public. Uncapitalized fields will not be printed:
```golang
type thing struct {
	Str string
	Num int
	secret string
}
t := thing{"string", 1234, "shhh"}

p := prettyprinter.NewPrettyPrinter()
err := p.Add(t).
	StdoutDump().
	Error()
```
```json
{
 "Str": "string",
 "Num": 1234,
}
```

Pretty printer will utilize struct json tags:
```golang
p := prettyprinter.NewPrettyPrinter()
err := p.Add(struct{
			F1 string `json:"field_1"`
			F2 string `json:"field_2"`
		}{
			"apples",
			"oranges",
		}).
		StdoutDump().
		Error()
```
```json
{
 "field_1": "apples",
 "field_2": "oranges",
}
```

Due to marshaling limitations in golang's `encoding/json` package, errors to pretty printer should be converted to strings via `error.Error()` or wrapped in pretty printer's custom error types.

To print a simple error, wrap the error in the pretty printer type `KeyValueError`. It will produce a key-value format: `"err": "error message"`:
```golang
var existErr error = errors.New("unicorns exist!")

p := prettyprinter.NewPrettyPrinter()
err := p.Add(prettyprinter.MakeKeyValueError(existErr)).
	StdErrDump().
	Error()
```
```json
{
 "err": "unicorns exist!"
}
```
```golang
var nilErr error = nil
err := p.Add(prettyprinter.MakeKeyValueError(nilErr)).
	StdErrDump().
	Error()
```
```json
{
 "err": "null"
}
```

If you want a custom json key name for an error in a struct, wrap the error in the `FieldError` type with the desired key name as a struct tag. It may be useful to create custom constructors for existing structs to wrap the error type fields for pretty printing. Or, to implement a custom `MarshalJSON()` method for existing structs that stringifies or wraps the error fields for pretty printing
```golang
type Struct1 struct {
	Error error
	Str string
}

type Struct1Output struct {
	Error prettyprinter.KeyValueError `json:"not-used-tag"`
	Str string `json:"string"`
}

func MakeStruct1Output(s Struct1) Struct1Output {
	err := prettyprinter.MakeKeyValueError(s.Error)
	return Struct1Output{err, s.Str}
}

var s Struct1 = Struct1{errors.New("error!"), "grapes"}

p := prettyprinter.NewPrettyPrinter()
err := p.Add(MakeStruct1Output(s)).
	StdOutDump().
	Error()
```
```json
{
 "err": "error!",
 "string": "grapes"
}
```
```golang
type Struct2 struct {
	Error error
	Str string
}

func (s Struct2) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct{
		Error prettyprinter.FieldError `json:"custom_error"`
		Str string `json:"string"`
	}{
		prettyprinter.MakeFieldError(s.Error),
		s.Str,
	})
}

var s Struct2 = Struct2{errors.New("error!"), "grapes"}

p := prettyprinter.NewPrettyPrinter()
err := p.Add(s).
	StdOutDump().
	Error()
```
```json
{
 "custom_error": "error!",
 "string": "grapes"
}
```

# Dependencies for testing
Mockery - mockery v1 is used to autogenerate code for golang interfaces. Mocked interfaces are automatically outputted to the mocks/ folder. The golang binary tool can be downloaded from https://github.com/vektra/mockery