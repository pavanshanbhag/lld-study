package main

import "templatemethod"

func main() {
	csvProcessor := &templatemethod.CSVDataProcessor{}
	baseCSV := &templatemethod.BaseDataProcessor{Processor: csvProcessor}
	baseCSV.Process()

	xmlProcessor := &templatemethod.XMLDataProcessor{}
	baseXML := &templatemethod.BaseDataProcessor{Processor: xmlProcessor}
	baseXML.Process()
}
