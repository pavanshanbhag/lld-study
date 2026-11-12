package templatemethod

func Demo() {

	csvProcessor := &CSVDataProcessor{}
	baseCSV := &BaseDataProcessor{Processor: csvProcessor}
	baseCSV.Process()

	xmlProcessor := &XMLDataProcessor{}
	baseXML := &BaseDataProcessor{Processor: xmlProcessor}
	baseXML.Process()

}
