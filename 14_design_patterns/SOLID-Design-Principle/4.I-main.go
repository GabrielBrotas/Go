package main

// Interface Segregation Principle
// you should not have a large interfaces, you should have chunks of interfaces that build a bigger type

type Document struct {}

type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

type NewPrinter struct {}

func (n *NewPrinter) Print() {} // Ok
func (n *NewPrinter) Fax() {} // Ok
func (n *NewPrinter) Scan() {} // Ok

type OldPrinter struct {}


func (o *OldPrinter) Print() {} // Ok
// Deprecated: ...
func (o *OldPrinter) Fax() {} // Not Ok - incapable of
// Deprecated: ...
func (o *OldPrinter) Scan() {} // Not Ok - incapable of

// ISP - Break interface in smaller piece, eg:

type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

type Faxer interface {
	Fax(d Document)
}

type MyPrinter struct {}
func (m *MyPrinter) Print() {} // Ok

type MyMultiFuncDevice interface {
	Printer
	Scanner
	// Fax
}


func main() {
}