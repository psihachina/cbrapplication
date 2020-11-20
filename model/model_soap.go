package model

import "encoding/xml"

// SOAPEnvelope envelope
type SOAPEnvelope struct {
	XMLName xml.Name    `xml:"soap:Envelope"`
	Header  *SOAPHeader `xml:",omitempty"`
	Body    SOAPBody    `xml:",omitempty"`
	Soap    string      `xml:"xmlns:soap,attr"`
	Xsi     string      `xml:"xmlns:xsi,attr"`
	Xsd     string      `xml:"xmlns:xsd,attr"`
}

// SOAPHeader header
type SOAPHeader struct {
	XMLName xml.Name    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`
	Content interface{} `xml:",omitempty"`
}

// SOAPBody body
type SOAPBody struct {
	XMLName xml.Name    `xml:"http://schemas.xmlsoap.org/soap/envelope/ soap:Body"`
	Content interface{} `xml:",omitempty"`
}

// EnvelopeUploadFile - evenlope for upload file action
type EnvelopeUploadFile struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Soap    string   `xml:"soap,attr"`
	Xsi     string   `xml:"xsi,attr"`
	Xsd     string   `xml:"xsd,attr"`
	Body    struct {
		Text    string `xml:",chardata"`
		Xmlns   string `xml:"xmlns,attr"`
		Content string `xml:"Content"`
	} `xml:"Body"`
}

// EnvelopeDownloadFile - evenlope for download file action
type EnvelopeDownloadFile struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Soap    string   `xml:"soap,attr"`
	Xsi     string   `xml:"xsi,attr"`
	Xsd     string   `xml:"xsd,attr"`
	Body    struct {
		Text    string `xml:",chardata"`
		Xmlns   string `xml:"xmlns,attr"`
		Content struct {
			Text     string `xml:",chardata"`
			FileName string `xml:"FileName"`
		}
	} `xml:"Body"`
}
