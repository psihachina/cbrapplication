package model

// Envelope - CBR response with rate
type Envelope struct {
	ValuteCursOnDate []struct {
		Vname   string `xml:"Vname"`
		Vnom    string `xml:"Vnom"`
		Vcurs   string `xml:"Vcurs"`
		Vcode   string `xml:"Vcode"`
		VchCode string `xml:"VchCode"`
	} `xml:"Body>GetCursOnDateXMLResponse>GetCursOnDateXMLResult>ValuteData>ValuteCursOnDate"`
}

// ValuteCursOnDate - element of envelope response
type ValuteCursOnDate struct {
	Text    string `xml:",chardata"`
	Vname   string `xml:"Vname"`
	Vnom    string `xml:"Vnom"`
	Vcurs   string `xml:"Vcurs"`
	Vcode   string `xml:"Vcode"`
	VchCode string `xml:"VchCode"`
}
