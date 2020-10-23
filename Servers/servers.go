package Servers

type IUserServer interface {
	GetName(id int) string
	DelUser(id int) string
}

type UserServer struct {

}

func (user UserServer) GetName(id int)string{
	if id ==101{
		return "asd"
	}else{
		return "zxc"
	}
}

func (user UserServer) DelUser(id int)string{
	if id==101{
		return "Failed"
	}else{
		return "Successful"
	}
}
