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
	con_url := "http://52.199.217.23/app/v0/login"
	//api_ket := "e33d36e16e80413abfaf5caa66d49e66" 
	form := url.Values{}
	
	form.Add( "account", "jphacks" )
	form.Add( "password", "jphacks" )
	
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
		err = json.Unmarshal( []byte( json_token ), re_data )

		if err != nil {
			fmt.Println( err )
			return
		}
	} else {
		fmt.Println( err.Error() )
		return
	}

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

}
