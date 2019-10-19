package api

import (
	"../util"
	"../logger"
	"../database"
	"../config"
	"../jwt"
	"encoding/json"
	"net/http"
	"bytes"
	"fmt"
	//"os"
)

type Plan_Data struct {
	name string `json:name`
	key string `json:key`
}

/*
type Form_Day struct {
	Day string `json:"day"`
	Check int `json:check`
}

type Form struct {
	Area string `json:"area"`
	Genre string `json:"genre"`
	Free string `json:"free"`
	Day []Form_Day `json:"select_day"`
}
*/

func Plan_Check( conf config.Connect_data, keys *jwt.JWTKeys) http.HandlerFunc {
	return func( w http.ResponseWriter, req *http.Request ) {		
		logger.Write_log( "plan check start " + req.RemoteAddr, 1 )

		if req.Method == "OPTIONS" {
			logger.Write_log( "OPTIONS return", 1 )
			util.CORSforOptions( &w )
			return
		}
		
		account := req.FormValue( "account" )

		if len( account ) == 0 {
			logger.Write_log( "not set account " + req.RemoteAddr, 1 )
			fmt.Fprintf( w, "false" )
			return
		}

		db, err := database.Connect( conf.DB )
		defer database.Disconnect( db )
		
		if err != nil {
			logger.Write_log( "database not connect", 4 )
			logger.Write_log( err.Error(), 4 )
			fmt.Fprintf( w, "false" )
			return
		}

		user_id, err := database.Account_ID( db.Sess, account )

		if err != nil {
			logger.Write_log( "fail get ID " + req.RemoteAddr, 1 )
			fmt.Fprintf( w, "false" )
			return
		}

		key_list, err := database.Plan_Key( db.Sess, user_id )

		var name_key_list []Plan_Data

		for i := 0; i < len( key_list ); i++ {
			file_name := key_list[i] + "_form.json"
			bytes, err := util.FileDownload( file_name )
			instance := Plan{}

			err = json.Unmarshal( bytes, &instance )

			if err != nil {
				logger.Write_log( "fail json change", 4 )
				logger.Write_log( err.Error(), 4 )
				fmt.Fprintf( w, "false" )
				return
			}

			check := Plan_Data{}
			check.key = key_list[i]
			check.name = instance.Plan_Name

			name_key_list = append( name_key_list, check )
		}

		res_bytes, err := json.Marshal( name_key_list )

		if err != nil {
			logger.Write_log( "fail json change", 1 )
			logger.Write_log( err.Error(), 1 )
			return
		}
		
		var buf bytes.Buffer
	
		_ = json.Indent( &buf, res_bytes, "", "  " )
		
		responseResult := ResponseResult{
			Status:    "OK",
			Data:      map[string]interface{}{ "json": buf.String() },
			ErrorText: "",
		}

		res, _ := json.Marshal( responseResult )

		logger.Write_log( "plan check success" + req.RemoteAddr, 1 )
		util.Respond( res, w )
	}
}
