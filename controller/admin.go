package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"go_gateway_demo/dao"
	"go_gateway_demo/dto"
	"go_gateway_demo/lib"
	"go_gateway_demo/middleware"
	"go_gateway_demo/public"
)

type AdminController struct{}

func AdminRegister(Router *gin.RouterGroup) {
	adminLogin := &AdminController{}
	Router.GET("/admin_info", adminLogin.AdminInfo)
	Router.POST("/change_pwd", adminLogin.ChangePwd)
}

// AdminInfo godoc
//// @Summary 管理员信息
//// @Description 管理员信息
//// @Tags 管理员接口
//// @ID /admin/admin_info
//// @Accept  json
//// @Produce  json
//// @Success 200 {object} middleware.Response{data=dto.AdminInfoOutput} "success"
//// @Router /admin/admin_info [get]
func (adminlogin *AdminController) AdminInfo(c *gin.Context) {
	sess := sessions.Default(c)
	sessInfo := sess.Get(public.AdminSessionInfoKey)
	adminSessionInfo := &dto.AdminSessionInfo{} //获取信息
	if err := json.Unmarshal([]byte(fmt.Sprint(sessInfo)), adminSessionInfo); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	//1.读取sessionKey对应json 转换为结构体
	//读出数据然后封装输出结构体

	out := &dto.AdminInfoOutput{
		ID:           adminSessionInfo.ID,
		Name:         adminSessionInfo.UserName,
		LoginTime:    adminSessionInfo.LoginTime,
		Avatr:        "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
		Introduction: "I am a super administrator",
		Roles:        []string{"admin"},
	}
	middleware.ResponseSuccess(c, out)
}

// ChangePwd godoc
// @Summary 修改密码
// @Description 修改密码
// @Tags 管理员接口
// @ID /admin/change_pwd
// @Accept  json
// @Produce  json
// @Param body body dto.ChangePwdInput true "body"
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /admin/change_pwd [post]
func (adminlogin *AdminController) ChangePwd(c *gin.Context) {
	params := &dto.ChangePwdInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	//1.读取session用户信息到结构体 sessInfo.ID
	//2.sessInfo.ID 读取数据库信息 adminInfo
	//3.params.password+adminInfo.salt sha256 saltPassword
	//4.saltPassword ==> adminInfo.password 执行数据库保存

	sess := sessions.Default(c)
	sessInfo := sess.Get(public.AdminSessionInfoKey)
	adminSessionInfo := &dto.AdminSessionInfo{}
	if err := json.Unmarshal([]byte(fmt.Sprint(sessInfo)), adminSessionInfo); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	//DB数据库链接
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	//查找数据库对应信息
	adminInfo := &dao.Admin{}
	adminInfo, err = adminInfo.Find(c, tx, (&dao.Admin{UserName: adminSessionInfo.UserName}))
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}

	//对新密码进行盐加密
	saltPassword := public.GenSaltPassword(adminInfo.Salt, params.Password)
	//数据库密保保存
	adminInfo.Password = saltPassword
	if err := adminInfo.Save(c, tx); err != nil {
		middleware.ResponseError(c, 2003, err)
		return
	}
	middleware.ResponseSuccess(c, "")
}
