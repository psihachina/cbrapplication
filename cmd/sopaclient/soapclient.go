package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/c-bata/go-prompt"
	"github.com/cbrapplication/excel"
	"github.com/cbrapplication/model"
	"github.com/cbrapplication/soapclient"
)

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "getCurs", Description: "getCurs"},
		{Text: "exit", Description: "Closes the command line interface"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func main() {
	for {
		cmdString := prompt.Input("> ", completer)
		err := runCommand(cmdString)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
func runCommand(commandStr string) error {
	var err error
	commandStr = strings.TrimSuffix(commandStr, "\n")
	arrCommandStr := strings.Fields(commandStr)
	var arrValutes []model.ValuteCursOnDate
	var arrTime []string
	switch arrCommandStr[0] {
	case "exit":
		os.Exit(0)
	case "getCurs":
		if len(arrCommandStr) < 2 {
			return errors.New("Please enter the symbolic currency code!")
		}
		var result model.Envelope
		var currentTime string
		for i := 0; i < 5; i++ {
			timeNow := time.Now().AddDate(0, 0, -i)
			day, mounth, year := timeNow.Date()
			hour := fmt.Sprintf("%v", timeNow.Hour())
			if timeNow.Hour()/10 == 0 {
				hour = "0" + hour
			}
			minute := fmt.Sprintf("%v", timeNow.Minute())
			if timeNow.Minute()/10 == 0 {
				minute = "0" + minute
			}
			second := fmt.Sprintf("%v", timeNow.Second())
			if timeNow.Second()/10 == 0 {
				second = "0" + second
			}
			currentTime = fmt.Sprintf("%v-%v-%vT%vZ", day, int(mounth), year,
				fmt.Sprintf("%v:%v:%v", hour, minute, second))
			fmt.Println(currentTime)
			fmt.Println(arrCommandStr[1])
			result, err = soapclient.GetCursOnDate(currentTime)
			if err != nil {
				return err
			}
			for _, v := range result.ValuteCursOnDate {
				if v.VchCode == arrCommandStr[1] {
					fmt.Println(v)
					arrTime = append(arrTime, currentTime)
					arrValutes = append(arrValutes, model.ValuteCursOnDate{VchCode: v.VchCode, Vcode: v.Vcode, Vcurs: v.Vcurs, Vname: strings.TrimSpace(v.Vname), Vnom: v.Vnom})
				}
			}

		}
		fmt.Println(arrValutes)
		path := currentTime + ".xlsx"
		f := excelize.NewFile()
		f.SaveAs(path)
		err := excel.CreateExcelWithValutes(arrValutes, arrTime, path)
		if err != nil {
			return nil
		}
		err = soapclient.Upload(path)
		if err != nil {
			fmt.Printf("Error: %v", err)
		}
		fmt.Fprintln(os.Stdout, "XML file created successfully")
		return nil
		// add another case here for custom commands.
	}
	cmd := exec.Command(arrCommandStr[0], arrCommandStr[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func sum(numbers ...int64) int64 {
	res := int64(0)
	for _, num := range numbers {
		res += num
	}
	return res
}
