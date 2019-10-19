package main

import (
	"encoding/json"
	"bytes"
	"fmt"
)

type Plan struct {
	Plan_Name string `json:"plan_name"`
	Memo string `json:"memo"`
	Day [] string `json:"day"`
}


func main() {
	test_plan := Plan{}

	test_plan.Plan_Name = "飲み会"
	test_plan.Memo = "いっぱい飲むぞ"
	test_plan.Day = append( test_plan.Day, "2019:10/5" )
	test_plan.Day = append( test_plan.Day, "2019:10/7" )
	test_plan.Day = append( test_plan.Day, "2019:10/9" )

	jsonBytes, err := json.Marshal( test_plan )
	
    if err != nil {
		fmt.Println( "1" )
        fmt.Println( err )
        return
    }

	var buf bytes.Buffer
	
	err = json.Indent( &buf, jsonBytes, "", "  " )

	new_plan := new( Plan )
	err = json.Unmarshal( []byte( buf.String() ), new_plan )

	if err != nil {
		fmt.Println( "2" )
        fmt.Println( err )
        return
    }

	fmt.Println( new_plan )
}

