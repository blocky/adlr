package prettyprinter

import "os"

// PrettyPrinter holds data for printing, interfaces for writing,
// and the internal error that may occur from printing attempts
type PrettyPrinter struct {
	spool  interface{}
	stderr Writer
	stdout Writer
	err    error
}

// Returns a new PrettyPrinter with default stdout and stderr files
func NewPrettyPrinter() *PrettyPrinter {
	return NewPrettyPrinterFromRaw(os.Stderr, os.Stdout)
}

// Construct a PrettyPrinter from raw Writers
func NewPrettyPrinterFromRaw(
	stderr, stdout Writer,
) *PrettyPrinter {
	return &PrettyPrinter{nil, stderr, stdout, nil}
}

// Add any one thing to the PrettyPrinter to store internally for printing
func (p *PrettyPrinter) Add(
	value interface{},
) Printer {
	p.spool = value
	return p
}

// Pretty print the spool to the stderr Writer
func (p *PrettyPrinter) StderrDump() Printer {
	return p.Dump(p.stderr)
}

// Pretty print the spool to the stdout Writer
func (p *PrettyPrinter) StdoutDump() Printer {
	return p.Dump(p.stdout)
}

// Pretty print the spool to a provided Writer, such as os.File
func (p *PrettyPrinter) Dump(w Writer) Printer {
	bytes, err := prettyJSON(p.spool)
	p.Flush()
	if err != nil {
		p.err = err
		return p
	}
	p.err = print(bytes, w)
	return p
}

// Clear the spool of anything saved for printing
func (p *PrettyPrinter) Flush() {
	p.spool = nil
}

// Return the internal error which may result from a printing attempt
func (p *PrettyPrinter) Error() error {
	return p.err
}

// If there is an internal error from printing, clean the spool and attempt
// to print the internal error in pretty json to the stderr Writer
func (p *PrettyPrinter) StderrDumpOnError() error {
	if p.Error() != nil {
		p.Flush()
		kve := MakeKeyValueError(p.Error())
		p.Add(kve)
		return p.StderrDump().Error()
	}
	return nil
}
