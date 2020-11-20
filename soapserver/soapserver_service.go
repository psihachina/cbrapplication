package soapserver

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/cbrapplication/fileshare"
	"github.com/cbrapplication/model"
)

func licenseServer(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.Header.Get("SOAPAction"))
	switch r.Method {
	case "POST":
		{
			switch r.Header.Get("SOAPAction") {
			case "UploadFile":
				{
					var res interface{}
					result := new(model.EnvelopeUploadFile)
					xml.NewDecoder(r.Body).Decode(&result)
					fmt.Println(result.Body.Content)
					err := fileshare.Upload(result.Body.Content)
					if err != nil {
						fmt.Println(err)
					}

					res = "Success upload file"
					v := model.SOAPEnvelope{
						Xsi:  "http://www.w3.org/2001/XMLSchema-instance",
						Xsd:  "http://www.w3.org/2001/XMLSchema",
						Soap: "http://schemas.xmlsoap.org/soap/envelope/",
						Body: model.SOAPBody{
							Content: res,
						},
					}
					w.WriteHeader(http.StatusOK)
					w.Header().Set("Content-Type", "text/xml")
					x, err := xml.MarshalIndent(v, "", "  ")
					if err != nil {
						w.WriteHeader(http.StatusInternalServerError)
						return
					}
					w.Write([]byte(xml.Header + string(x)))

					return
				}
			case "DownloadFile":
				{
					var res interface{}
					result := new(model.EnvelopeDownloadFile)
					xml.NewDecoder(r.Body).Decode(&result)
					bytes, err := fileshare.Download(result.Body.Content.FileName)
					if err != nil {
						fmt.Println(err)
					}

					res = bytes

					v := model.SOAPEnvelope{
						Xsi:  "http://www.w3.org/2001/XMLSchema-instance",
						Xsd:  "http://www.w3.org/2001/XMLSchema",
						Soap: "http://schemas.xmlsoap.org/soap/envelope/",
						Body: model.SOAPBody{
							Content: res,
						},
					}
					w.WriteHeader(http.StatusOK)
					w.Header().Set("Content-Type", "text/xml")
					x, err := xml.MarshalIndent(v, "", "  ")
					if err != nil {
						w.WriteHeader(http.StatusInternalServerError)
						return
					}
					w.Write([]byte(xml.Header + string(x)))

					return
				}
			case "FileList":
				{
					folder, err := os.Open("./store")
					if err != nil {
						fmt.Println(err)
					}
					var res interface{}
					res, err = folder.Readdirnames(0)
					if err != nil {
						fmt.Println(err)
					}

					v := model.SOAPEnvelope{
						Xsi:  "http://www.w3.org/2001/XMLSchema-instance",
						Xsd:  "http://www.w3.org/2001/XMLSchema",
						Soap: "http://schemas.xmlsoap.org/soap/envelope/",
						Body: model.SOAPBody{
							Content: res,
						},
					}
					w.WriteHeader(http.StatusOK)
					w.Header().Set("Content-Type", "text/xml")
					x, err := xml.MarshalIndent(v, "", "  ")
					if err != nil {
						w.WriteHeader(http.StatusInternalServerError)
						return
					}
					w.Write([]byte(xml.Header + string(x)))

					return
				}
			}

		}
	default:
		fmt.Fprintf(w, "Sorry, only POST methods are supported.")
	}

}
