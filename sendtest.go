package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"net/url"
	"net/http"
)

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

	if resp != nil {
		defer resp.Body.Close()
		var byteArray, _ = ioutil.ReadAll( resp.Body )
		json_token = string( byteArray )
		//fmt.Println( "ログイン成功!!!" )
		fmt.Println( json_token )
	} else {
		fmt.Println( err.Error() )
		return
	}
}
