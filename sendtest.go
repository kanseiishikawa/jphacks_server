package main

import (
	"fmt"
	"strings"
	"encoding/json"
	"io/ioutil"
	"net/url"
	"net/http"
)

// ResponseResult レスポンス結果に関する構造体
type ResponseResult struct {
	Status    string                 `json:"status"`
	Data      map[string]interface{} `json:"data"`
	ErrorText string                 `json:"errorText"`
}

func main() {
	//con_url := "http://52.199.217.23/app/v0/login"
	//con_url := "http://52.199.217.23/app/v0/plan_form"
	con_url := "http://52.199.217.23/app/v0/plan_check"
	//api_ket := "e33d36e16e80413abfaf5caa66d49e66"
/*
	json_data := `{
  "area": "神戸",
  "genre": "高級",
  "free": "お酒飲みたい",
  "select_day": [
    {
      "day": "2019:10/5",
      "Check": 1
    },
    {
      "day": "2019:10/7",
      "Check": 2
    },
    {
      "day": "2019:10/9",
      "Check": 3
    }
  ]
}`
*/	
	form := url.Values{}
	
	form.Add( "account", "jphacks" )
	//form.Add( "password", "jphacks" )

	//form.Add( "plan_key", "wxErdYhFZVEzOaluDVWi" )
	//form.Add( "form_data", json_data )
	
	req, err := http.NewRequest( "POST", con_url, strings.NewReader( form.Encode() ) )

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := new( http.Client )
	resp, err := client.Do( req )

	var json_token string
	var re_data = new( ResponseResult )

	if resp != nil {
		defer resp.Body.Close()
		var byteArray, _ = ioutil.ReadAll( resp.Body )
		json_token = string( byteArray )
		fmt.Println( json_token )
		err = json.Unmarshal( []byte( json_token ), re_data )

		if err != nil {
			fmt.Println( err )
			return
		}

		fmt.Println( re_data.Data )
	} else {
		fmt.Println( err.Error() )
		return
	}
	/*
	con_url = "http://52.199.217.23/app/v0/plan_generate"
	plan_data := `{"plan_name": "飲み会","memo": "いっぱい飲むぞ","day": ["2019:10/5","2019:10/7","2019:10/9"]}`
	form = url.Values{}
	
	form.Add( "account", "jphacks" )
	form.Add( "password", "jphacks" )
	form.Add( "plan_data", plan_data )
	
	req, err = http.NewRequest( "POST", con_url, strings.NewReader( form.Encode() ) )

	req.Header.Set( "Authorization", re_data.Data["token"].( string) )
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client = new( http.Client )
	resp, err = client.Do( req )

	if resp != nil {
		defer resp.Body.Close()
		var byteArray, _ = ioutil.ReadAll( resp.Body )
		fmt.Println( string( byteArray ) )
	} else {
		fmt.Println( err.Error() )
		return
	}
*/

}
