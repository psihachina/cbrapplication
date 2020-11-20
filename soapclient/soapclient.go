package soapclient

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/cbrapplication/base64"
	"github.com/cbrapplication/model"
)

// GetCursOnDate -
func GetCursOnDate(dateString string) (model.Envelope, error) {

	var currentTime string
	/* if dateString == "default" {
		timeNow := time.Now()
		day, mounth, year := timeNow.Date()

		currentTime = fmt.Sprintf("%v-%v-%vT%vZ", day, int(mounth), year,
			fmt.Sprintf("%v:%v:%v", timeNow.Hour(), timeNow.Minute(), timeNow.Second()))
		fmt.Println(currentTime)
	} else {
		date := strings.Fields(dateString)
		fmt.Println(date)
		currentTime = fmt.Sprintf("%v-%v-%vT12:10:10Z", date[2], date[1], date[0])
		fmt.Println(currentTime)
	} */
	currentTime = dateString
	url := "http://www.cbr.ru/DailyInfoWebServ/DailyInfo.asmx"
	var payload = []byte(strings.TrimSpace(fmt.Sprintf(`
		<?xml version="1.0" encoding="utf-8"?>
		<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
		  <soap:Body>
			<GetCursOnDateXML xmlns="http://web.cbr.ru/">
			  <On_date>%v</On_date>
			</GetCursOnDateXML>
		  </soap:Body>
		</soap:Envelope>`, currentTime),
	))
	var soapAction = "http://web.cbr.ru/GetCursOnDateXML"
	var httpMethod = "POST"
	req, err := http.NewRequest(httpMethod, url, bytes.NewReader(payload))
	if err != nil {
		return model.Envelope{}, err
	}

	req.Header.Set("Content-type", "text/xml")
	req.Header.Set("SOAPAction", soapAction)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal("Error on dispatching request. ", err.Error())
		return model.Envelope{}, err
	}

	result := new(model.Envelope)
	err = xml.NewDecoder(res.Body).Decode(result)
	if err != nil {
		return model.Envelope{}, err
	}

	return *result, nil
}

func Upload(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	byt, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	str := base64.Base64Encode(byt)
	fmt.Println(string(str))
	url := "http://localhost:8080/WebService.asmx"
	var payload = []byte(strings.TrimSpace(fmt.Sprintf(`
		<?xml version="1.0" encoding="utf-8"?>
		<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
		  <soap:Body>
			<Content>
			  %v
			</Content>
		  </soap:Body>
		</soap:Envelope>`, string(str)),
	))
	var soapAction = "UploadFile"
	var httpMethod = "POST"
	req, err := http.NewRequest(httpMethod, url, bytes.NewReader(payload))
	if err != nil {
		return err
	}

	req.Header.Set("Content-type", "text/xml")
	req.Header.Set("SOAPAction", soapAction)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal("Error on dispatching request. ", err.Error())
		return err
	}

	result := new(model.EnvelopeUploadFile)
	err = xml.NewDecoder(res.Body).Decode(result)
	if err != nil {
		return err
	}

	return nil
}
