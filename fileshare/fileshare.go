package fileshare

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/cbrapplication/base64"
)

// Upload - function of upload file
func Upload(bytes string) error {
	bytes = strings.TrimSpace(bytes)
	fmt.Println(bytes)
	decodeBytes, err := base64.Base64Decode(bytes)
	if err != nil {
		return err
	}

	folder, err := os.Open("./store")
	if err != nil {
		return err
	}
	names, err := folder.Readdirnames(0)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("./store/downloadfile"+fmt.Sprint(len(names))+".xlsx", decodeBytes, 0644)
	if err != nil {
		return err
	}
	return nil
}

// Download - function of download file
func Download(name string) ([]byte, error) {
	file, err := os.Open("./store/" + name)
	if err != nil {
		return nil, err
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	bytes = base64.Base64Encode(bytes)

	return bytes, nil
}
