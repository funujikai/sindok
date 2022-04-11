package middleware

import (
	"encoding/json"
	"fmt"
	// "encoding/json"
	// "pkmapi/global"
	beecontext "github.com/beego/beego/v2/server/web/context"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/dgrijalva/jwt-go"
	"PMM/global"
	"strings"
)

func Jwt(ctx *beecontext.Context) {

	KEY,_ := beego.AppConfig.String("KEY")

	ctx.Output.Header("Content-Type", "application/json")
	var uri string = ctx.Input.URI()

	//karena login tidak menggunakan token
	if strings.Contains(uri,"/v1/user/") == true {
		return
	}	

	if strings.Contains(uri,"/v1/other/") == true {
		return
	}		

	// karena Preflight Request tidak mengirim token
	if ctx.Input.Method() == "OPTIONS" {
        return
    }

	if ctx.Input.Header("Authorization") == "" {
		ctx.Output.SetStatus(403)
		resBody, err := json.Marshal(global.APIResponse{403, "notAllowed"})
		err = ctx.Output.Body(resBody)
		if err != nil {
			panic(err)
		}
	}

	var tokenString string = ctx.Input.Header("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(KEY), nil
	})	

	if err != nil {
		ctx.Output.SetStatus(403)
		var responseBody global.APIResponse = global.APIResponse{403, err.Error()}
		resBytes, err := json.Marshal(responseBody)
		err = ctx.Output.Body(resBytes)
		if err != nil {
			panic(err)
		}
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid && claims != nil {
		return
	} else {
		ctx.Output.SetStatus(403)
		resBody, err := json.Marshal(global.APIResponse{403, ctx.Input.Header("Authorization")})
		err = ctx.Output.Body(resBody)
		if err != nil {
			panic(err)
		}
	}
}