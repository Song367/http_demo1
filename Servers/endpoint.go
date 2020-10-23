package Servers

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/time/rate"
	"http_demo1/uti"
	"strconv"
)

type UserRequest struct {
	UserId int `json:"user_id"`
	Methods string
}

type UserResponse struct {
	Result string `json:"result"`
}


func Ratelimit(limit *rate.Limiter) endpoint.Middleware{
	return func(EP endpoint.Endpoint)endpoint.Endpoint{
		return func(ctx context.Context, request interface{}) (response interface{}, err error){
			if !limit.Allow(){
				return nil,errors.New("too many request")
			}
			return EP(ctx,request)
		}
	}
}

func GetUserInfo(server IUserServer)endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error){
		r:=request.(UserRequest)
		res:="Nothing"
		fmt.Println(r.Methods)
		if r.Methods=="GET" {
			res = server.GetName(r.UserId)+strconv.Itoa(uti.ServerPort)
		}else if r.Methods=="DELETE"{
			res = server.DelUser(r.UserId)
		}
		return UserResponse{Result: res},nil
	}
}
