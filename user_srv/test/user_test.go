package test

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"mxshop_srvs/user_srv/proto"
	"testing"
)

func TestGetUserList(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		t.Errorf("did not connect: %v", err)
	}
	defer conn.Close()
	ctx := context.Background()
	user := proto.NewUserClient(conn)
	userList, err := user.GetUserList(ctx, &proto.PageInfo{
		Pn:    1,
		PSize: 5,
	})
	if userList == nil {
		t.Log("userList is nil")
		return
	}
	if err != nil {
		t.Errorf("GetUserList error: %v", err)
	}
	// 校验密码
	for _, data := range userList.Data {
		t.Logf("nick_name:%s - mobile:%s - password: %s \n", data.NickName, data.Mobile, data.PassWord)
		UserPwd := data.PassWord
		check, err := user.CheckPassWord(ctx, &proto.PasswordCheckInfo{
			Password:          "admin@123!",
			EncryptedPassword: UserPwd,
		})
		if err != nil {
			t.Errorf("CheckPassWord error: %v", err)
		}
		t.Logf("check: %v", check.Success)
	}
}

func TestCreateUser(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		t.Errorf("did not connect: %v", err)
	}
	defer conn.Close()
	ctx := context.Background()
	user := proto.NewUserClient(conn)
	res, err := user.CreateUser(ctx, &proto.CreateUserInfo{
		NickName: "admin",
		Mobile:   "15270914973",
		PassWord: "123456",
	})
	if err != nil {
		t.Errorf("CreateUser error: %v", err)
	}
	fmt.Println(res.Id)
}

func TestUpdateUser(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		t.Errorf("did not connect: %v", err)
	}
	defer conn.Close()
	ctx := context.Background()
	user := proto.NewUserClient(conn)
	res, err := user.UpdateUser(ctx, &proto.UpdateUserInfo{
		Id:       10,
		NickName: "admin123",
		Gender:   "female",
		BirthDay: 896112000,
	})
	if err != nil {
		t.Errorf("UpdateUser error: %v", err)
	}
	fmt.Println(res)
}

func TestGetUserByMobile(t *testing.T) {

}
