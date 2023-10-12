package main

import (
	"choice/organization"
	"fmt"

	"github.com/xuri/excelize/v2"
)

func main() {
	f, err := excelize.OpenFile("/home/star/Desktop/ap.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	//items.SelectFromItemSheet(f, 10, 1)
	//doctor.SelectFromItemSheet(f, 10)
	organization.SelectFromItemSheet(f, 10, 1)
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

}
