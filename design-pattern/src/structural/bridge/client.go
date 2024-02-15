package bridge

import "fmt"

func Main() {
	var macComputer, winComputer Computer
	var hpPrinter, epsonPrinter Printer

	hpPrinter, epsonPrinter = &Hp{}, &Epson{}
	macComputer, winComputer = &Mac{}, &Windows{}

	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()

	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()
	fmt.Println()

	winComputer.SetPrinter(hpPrinter)
	winComputer.Print()
	fmt.Println()

	winComputer.SetPrinter(epsonPrinter)
	winComputer.Print()
	fmt.Println()
}
