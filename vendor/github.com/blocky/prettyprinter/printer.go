package prettyprinter

// Pretty printer fullfills the Printer interface for easier mocking in your testing.
// Use the Printer interface in your code in place of the PrettyPrinter struct
type Printer interface {
	Add(interface{}) Printer
	Dump(Writer) Printer
	StderrDump() Printer
	StdoutDump() Printer
	Error() error
	StderrDumpOnError() error
}

// The Writer interface from https://golang.org/pkg/io/#Writer
type Writer interface {
	Write([]byte) (int, error)
}
