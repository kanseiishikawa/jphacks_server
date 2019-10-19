package api

import (
	"../util"
	"../logger"
	//"../database"
	"../config"
	"../jwt"
	"encoding/json"
	"net/http"
	"bytes"
	"fmt"
	//"os"
)

type Store_Data struct {
	Store_Name string `json:"store_name"`
	Category string `json:"category"`
	URL string `json:"url"`
	Store_Image string `json:"store_image"`
	Rest_Day string `json:"rest_day"`
	Area string `json:"area"`
	Average_Money int `json:"average_money"`
}

func Store_Search( conf config.Connect_data, keys *jwt.JWTKeys) http.HandlerFunc {
	return func( w http.ResponseWriter, req *http.Request ) {
		logger.Write_log( "store search start " + req.RemoteAddr, 1 )

		if req.Method == "OPTIONS" {
			logger.Write_log( "OPTIONS return", 1 )
			util.CORSforOptions( &w )
			return
		}

		api_par := map[string]string{}

		if len( req.FormValue( "freeword" ) ) != 0 {
			api_par["freeword"] = req.FormValue( "freeword" )
		} else {
			api_par["freeword"] = "0"
		}
		
		if len( req.FormValue( "no_smorking" ) ) != 0 { 
			api_par["no_smorking"] = req.FormValue( "no_smorking" )
		} else {
			api_par["no_smorking"] = "0"
		}

		if len( req.FormValue( "card" ) ) != 0  {
			api_par["card"] = req.FormValue( "card" )
		} else {
			api_par["card"] = "0"
		}

		if len( req.FormValue( "bottomless_cup" ) ) != 0 {
			api_par["bottomless_cup"] = req.FormValue( "bottomless_cup" )
		} else {
			api_par["bottomless_cup"] = "0"
		}

		if len( req.FormValue( "buffet" ) ) != 0 {
			api_par["buffet"] = req.FormValue( "buffet" )
		} else {
			api_par["buffet"] = "0"
		}

		if len( req.FormValue( "private_room" ) ) != 0 {
			api_par["private_room"] = req.FormValue( "private_room" )
		} else {
			api_par["private_room"] = "0"
		}

		if len( req.FormValue( "midnight") ) != 0 {
			api_par["midnight"] = req.FormValue( "midnight" )
		} else {
			api_par["midnight"] = "0"
		}

		if len( req.FormValue( "wifi") ) != 0 {
			api_par["wifi"] = req.FormValue( "wifi" )
		} else {
			api_par["wifi"] = "0"
		}

		if len( req.FormValue( "projecter_screen" ) ) != 0 {
			api_par["projecter_screen"] = req.FormValue( "projecter_screen" )
		} else {
			api_par["projecter_screen"] = "0"
		}

		if len( req.FormValue( "web_reserve" ) ) != 0 {
			api_par["web_reserve"] = req.FormValue( "web_reserve" )
		} else {
			api_par["web_reserve"] = "0"
		}

		url_option := ""

		for k, v := range api_par {
			if v != "0" {
				url_option += "&" + k + "=" + v
			}
		}

		result, err := ReturnStruct( url_option )

		if err != nil {
			logger.Write_log( "fail tap api", 1 )
			logger.Write_log( err.Error(), 1 )
			fmt.Fprintf( w, "false")
			return
		}

		var res_store_data []Store_Data

		for i := 0; i < len( result.Rest ); i++ {
			instance := Store_Data{}
			instance.Store_Name = result.Rest[i].Name
			instance.Category = result.Rest[i].Category
			instance.URL = result.Rest[i].URL
			instance.Store_Image = result.Rest[i].ImageURL.ShopImage1
			instance.Rest_Day = result.Rest[i].Holiday
			instance.Area = result.Rest[i].Code.AreanameS
			instance.Average_Money = result.Rest[i].Budget

			res_store_data = append( res_store_data, instance )
		}

		json_byte, err := json.Marshal( res_store_data )

		if err != nil {
			logger.Write_log( "fail change json", 1 )
			logger.Write_log( err.Error(), 1 )
			fmt.Fprintf( w, "false")
			return
		}

		var buf bytes.Buffer
	
		_ = json.Indent( &buf, json_byte, "", "  " )

		responseResult := ResponseResult{
			Status:    "OK",
			Data:      map[string]interface{}{ "json": buf.String() },
			ErrorText: "",
		}

		res, _ := json.Marshal( responseResult )

		logger.Write_log( "store search success " + req.RemoteAddr, 1 )

		util.Respond( res, w )
 	}
}