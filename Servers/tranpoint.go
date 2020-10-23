package Servers

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func DecodeRequest(c context.Context,r *http.Request)(interface{},error){
	vars:=mux.Vars(r)
	uid,ok:=vars["uid"]
	id,_:=strconv.Atoi(uid)
	if ok{
		return UserRequest{UserId: id,Methods: r.Method},nil
	}
	return nil,errors.New("参数错误")
}

func EncodeResponse(c context.Context,w http.ResponseWriter,response interface{})error{
	return json.NewEncoder(w).Encode(response)
}
