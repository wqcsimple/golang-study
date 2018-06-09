package main

import (
	"fmt"
    "github.com/360EntSecGroup-Skylar/excelize"
)

func main()  {
	xlsx, err := excelize.OpenFile("file/test.xlsx")
	if err != nil {
        fmt.Println(err)
        return
	}
	
	// Get value from cell by given worksheet name and axis.
	rows := xlsx.GetRows("战魂")

	for _, row := range rows {

		for _, colCell := range row {
            fmt.Print(colCell, "\t")
		}
		
		fmt.Println()
	}
}