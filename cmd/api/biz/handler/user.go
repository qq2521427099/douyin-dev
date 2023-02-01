package handler

import (
	"context"
	"douyin/cmd/api/biz/handler/vo"
	"douyin/cmd/api/biz/rpc"
	"douyin/cmd/api/conn"
	"douyin/kitex_gen/user"
	"douyin/pkg/middleware"
	"github.com/cloudwego/hertz/pkg/app"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts

type UserLoginResponse struct {
	vo.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	vo.Response
	User vo.User `json:"user"`
}

func Register(ctx context.Context, c *app.RequestContext) {
	username := c.Query("username")
	password := c.Query("password")
	//检查用户输入格式是否正确
	if len(username) == 0 {
		c.JSON(http.StatusOK, vo.Response{
			StatusCode: 1,
			StatusMsg:  "用户名不能为空",
		})
		return
	}
	if len(password) <= 5 {
		c.JSON(http.StatusOK, vo.Response{
			StatusCode: 1,
			StatusMsg:  "密码长度需要大于5",
		})
		return
	}
	//检查是否存在敏感词
	flag := conn.WordsFilter.Contains(username, conn.Root)
	if flag {
		c.JSON(http.StatusOK, vo.Response{
			StatusCode: 1,
			StatusMsg:  "存在敏感词，请修改",
		})
		return
	}
	//检查用户名是否重复
	containsName, err := rpc.ContainsName(context.Background(), &user.ContainsNameReq{Username: username})
	if err != nil {
		c.JSON(http.StatusOK, vo.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	if containsName {
		c.JSON(http.StatusOK, vo.Response{
			StatusCode: 1,
			StatusMsg:  "User already existed",
		})
		return
	}
	hashpw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) //加密处理
	if err != nil {
		c.JSON(http.StatusOK, vo.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	password = string(hashpw)
	//插入记录
	userId, err := rpc.CreateUser(context.Background(), &user.CreateUserReq{
		Username: username,
		Password: password,
	})
	if err != nil {
		c.JSON(http.StatusOK, vo.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	//合成token
	token := middleware.GenerateToken(username, password)
	err = middleware.SaveUserId(token, userId)

	if err != nil {
		c.JSON(http.StatusOK, vo.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: vo.Response{StatusMsg: "Register success!"},
		UserId:   userId,
		Token:    token,
	})
	return
}

func Login(ctx context.Context, c *app.RequestContext) {
	username := c.Query("username")
	password := c.Query("password")
	//检查用户名是否存在
	u, err := rpc.GetUserByName(context.Background(), &user.GetUserByNameReq{Username: username})
	if err != nil {
		c.JSON(http.StatusOK, vo.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	//检查密码是否正确
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		c.JSON(http.StatusOK, vo.Response{
			StatusCode: 1,
			StatusMsg:  "Wrong password!",
		})
		return
	}
	token := middleware.GenerateToken(username, password)
	err = middleware.SaveUserId(token, u.Id)
	if err != nil {
		c.JSON(http.StatusOK, vo.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: vo.Response{StatusMsg: "Login Success!"},
		UserId:   u.Id,
		Token:    token,
	})
}

//func UserInfo(ctx context.Context, c *app.RequestContext) {
//	token := c.Query("token")
//	userIdData := c.Query("user_id")
//	userId, err := strconv.ParseInt(userIdData, 10, 64)
//	if err != nil {
//		fmt.Println(err)
//		c.JSON(http.StatusOK, vo.Response{
//			StatusCode: 1,
//			StatusMsg:  "Invalid user_id",
//		})
//		return
//	}
//	//鉴权并获取登录用户信息
//	uId, err := middleware.GetUserIdByToken(token)
//	if err != nil {
//		fmt.Println(err)
//		c.JSON(http.StatusOK, vo.Response{
//			StatusCode: 1,
//			StatusMsg:  err.Error(),
//		})
//	}
//	//获取目标用户信息
//	userTarget, err := rpc.GetUserById(context.Background(), &user.GetUserByIdReq{Id: userId})
//	if err != nil {
//		fmt.Println(err)
//		c.JSON(http.StatusOK, vo.Response{
//			StatusCode: 1,
//			StatusMsg:  err.Error(),
//		})
//	}
//	us, err := utils.PackageUser(uId, userTarget.Id, userTarget.Name)
//	if err != nil {
//		fmt.Println(err)
//		c.JSON(http.StatusOK, vo.Response{StatusCode: 1, StatusMsg: "Invalid Token"})
//		return
//	}
//	//转换
//	c.JSON(http.StatusOK, UserResponse{
//		Response: vo.Response{StatusCode: 0},
//		User:     *us,
//	})
//}
