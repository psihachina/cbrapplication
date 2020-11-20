package soapserver

import (
	"net/http"
)

// NewSOAPMux - return SOAP server mux
func NewSOAPMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/WebService.asmx", licenseServer)
	return mux
}

// NewSOAPServer create i2c mock server
func NewSOAPServer(port string) *http.Server {
	mux := NewSOAPMux()
	server := &http.Server{
		Handler: mux,
		Addr:    ":" + port,
	}
	return server
}
