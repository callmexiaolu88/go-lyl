package handler

import (
	"fmt"
	"lyl/learn/restful/model"

	restful "github.com/emicklei/go-restful"
)

type UserResource struct {
	Users map[string]model.User
}

var users UserResource = UserResource{map[string]model.User{}}

func GetUser(req *restful.Request, rsp *restful.Response) {
	id := req.PathParameter("id")
	fmt.Println("Get user by id:" + id)
}
