package controller

import (
	"encoding/json"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"go_gateway_demo/dao"
	"go_gateway_demo/dto"
	"go_gateway_demo/lib"
	"go_gateway_demo/middleware"
	"go_gateway_demo/public"
	"time"
)

type AdminLoginController struct{}

//管理员路由注册
func AdminLoginRegister(Router *gin.RouterGroup) {
	adminLogin := &AdminLoginController{}
	Router.POST("/login", adminLogin.AdminLogin)    // 管理员登陆登陆
	Router.GET("/logout", adminLogin.AdminLoginOut) //管理员登出
}

// AdminLogin godoc
// @Summary 管理员登陆登陆
// @Description 管理员登陆登陆
// @Tags 管理员登陆接口
// @ID /admin_login/login
// @Accept  json
// @Produce  json
// @Param body body dto.AdminLoginInput true "body"
// @Success 200 {object} middleware.Response{data=dto.AdminLoginOutput} "success"
// @Router /admin_login/login [post]
func (adminlogin *AdminLoginController) AdminLogin(c *gin.Context) {
	//绑定输入结构体测试
	//xorm gorm 基本形式一样 通过BindShould绑定tag获取URL传递过来的信息，挺方便
	params := &dto.AdminLoginInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	//登陆检查，检验gorm数据库链接
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	//引用dao层结构体，传入接受到的额数据
	admin := &dao.Admin{}
	//登陆信息检查
	admin, err = admin.LoginCheck(c, tx, params)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}

	//设置session
	sessInfo := &dto.AdminSessionInfo{
		ID:        admin.Id,
		UserName:  admin.UserName,
		LoginTime: time.Now(),
	}
	sessBts, err := json.Marshal(sessInfo)
	if err != nil {
		middleware.ResponseError(c, 2003, err)
		return
	}
	sess := sessions.Default(c)
	sess.Set(public.AdminSessionInfoKey, string(sessBts))
	sess.Save()

	//1.params.username 取得管理员信息,AdminInfo
	//2.adminInfo.salt + params.Password sha256=> saltPassword
	//3.saltPassword==adminInfoPassword
	out := &dto.AdminLoginOutput{Token: admin.UserName}
	middleware.ResponseSuccess(c, out)
}

// AdminLogin godoc
// @Summary 管理员登出
// @Description 管理员登出
// @Tags 管理员登出接口
// @ID /admin_login/logout
// @Accept  json
// @Produce  json
// @Param body body dto.AdminLoginInput true "body"
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /admin_login/logout [get]
func (adminlogin *AdminLoginController) AdminLoginOut(c *gin.Context) {

	sess := sessions.Default(c)
	sess.Delete(public.AdminSessionInfoKey)
	sess.Save()
	//1.params.username 取得管理员信息,AdminInfo
	//2.adminInfo.salt + params.Password sha256=> saltPassword
	//3.saltPassword==adminInfoPassword
	middleware.ResponseSuccess(c, "")
}
