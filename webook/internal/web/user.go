package web

import (
	"errors"
	"github.com/Teresajw/go_project/webook/internal/domain"
	"github.com/Teresajw/go_project/webook/internal/service"
	_ "github.com/dlclark/regexp2"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

const (
	Biz = "login"
)

// 确保UserHandler 实现了 handler 接口, 否则编译的时候会报错,以下两种方法，任选其一，第二种没有初始化任何对象，减少内存开销
// var _ handler = &UserHandler{}
var _ handler = (*UserHandler)(nil)

type UserHandler struct {
	svc         service.UserService
	codeSvc     service.CodeService
	emailExp    *regexp.Regexp
	passwordExp *regexp.Regexp
}

type UserClaims struct {
	jwt.RegisteredClaims
	// 声明你要放入token中的数据
	Uid int64
}

func NewUserHandler(svc service.UserService, codeSvc service.CodeService) *UserHandler {
	const (
		emailRegexPattern    = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
		passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,72}$`
	)
	// 编译正则表达式
	emailExp := regexp.MustCompile(emailRegexPattern, regexp.None)
	passwordExp := regexp.MustCompile(passwordRegexPattern, regexp.None)
	return &UserHandler{
		svc:         svc,
		codeSvc:     codeSvc,
		emailExp:    emailExp,
		passwordExp: passwordExp,
	}
}

func (u *UserHandler) RegisterRouters(server *gin.Engine) {
	ug := server.Group("/users")
	ug.POST("/signup", u.SignUp)
	//ug.POST("/login", u.Login)
	ug.POST("/login", u.LoginJWT)
	ug.POST("/login_sms/code/send", u.SendSMSLoginCode)
	ug.POST("/login_sms", u.LoginSMS)
	ug.POST("/edit", u.Edit)
	ug.GET("/profile", u.Profile)
	ug.POST("/delete", u.Delete)

}

func (u *UserHandler) SignUp(ctx *gin.Context) {
	type SignUpReq struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
		Nickname        string `json:"nickname"`
		Phone           string `json:"phone"`
		Birthday        string `json:"birthday"`
		Profile         string `json:"profile"`
	}
	var req SignUpReq
	// Bind 方法会根据 Content-Type 来解析你的数据到req中
	// 解析错了就会直接写回一个4xx错误
	if err := ctx.Bind(&req); err != nil {
		return
	}

	ok, err := u.emailExp.MatchString(req.Email)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "系统错误")
	}
	if !ok {
		ctx.String(http.StatusOK, "邮箱格式错误")
		return
	}
	if req.Password != req.ConfirmPassword {
		ctx.String(http.StatusOK, "两次密码不一致")
		return
	}

	ok, err = u.passwordExp.MatchString(req.Password)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "系统错误")
		return
	}
	if !ok {
		ctx.String(http.StatusOK, "密码必须大于8位、包含数字、特殊字符")
		return
	}
	// 调用service层
	//u.svc.SignUp(ctx.Request.Context(), req)
	err = u.svc.SignUp(ctx, domain.User{
		Email:    req.Email,
		Password: req.Password,
		Nickname: req.Nickname,
		Phone:    req.Phone,
		Birthday: req.Birthday,
		Profile:  req.Profile,
	})
	if errors.Is(err, service.ErrDuplicateEmail) {
		ctx.String(http.StatusOK, "邮箱已经存在")
		return
	}
	if err != nil {
		ctx.String(http.StatusInternalServerError, "系统异常")
		return
	}
	//数据库操作
	ctx.JSON(http.StatusOK, Result{
		Code: 200,
		Msg:  "注册成功！",
		Data: gin.H{
			"email": req.Email,
		},
	})
}

func (u *UserHandler) LoginJWT(ctx *gin.Context) {
	type LoginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req LoginReq
	if err := ctx.Bind(&req); err != nil {
		return
	}
	user, err := u.svc.Login(ctx, req.Email, req.Password)
	if errors.Is(err, service.ErrInvalidUserOrPassword) {
		ctx.JSON(http.StatusOK, Result{
			Code: 200,
			Msg:  "邮箱或密码错误",
		})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusOK, Result{
			Code: 500,
			Msg:  "系统异常",
		})
		return
	}
	if err = u.setJwtToken(ctx, user.Id); err != nil {
		ctx.JSON(http.StatusOK, Result{
			Code: 500,
			Msg:  "系统异常",
		})
		return
	}

	ctx.JSON(http.StatusOK, Result{
		Code: 200,
		Msg:  "登录成功",
	})
}

func (u *UserHandler) LoginJWTOut(ctx *gin.Context) {

}

func (u *UserHandler) Login(ctx *gin.Context) {
	type LoginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req LoginReq
	if err := ctx.Bind(&req); err != nil {
		return
	}
	user, err := u.svc.Login(ctx, req.Email, req.Password)
	if errors.Is(err, service.ErrInvalidUserOrPassword) {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "邮箱或密码错误",
		})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "系统异常",
		})
		return
	}
	sess := sessions.Default(ctx)
	//sess.Options(sessions.Options{
	//	//Secure: true,
	//	//HttpOnly: true,
	//	MaxAge: 30,
	//})
	//随便设置值，你要放在session里面的值
	sess.Set("userid", user.Id)
	err = sess.Save()

	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登录成功",
	})
}

func (u *UserHandler) LoginOut(ctx *gin.Context) {
	sess := sessions.Default(ctx)
	sess.Options(sessions.Options{
		// 一分钟过期
		MaxAge: 60,
	})
	sess.Save()
	ctx.JSON(http.StatusOK, Result{
		Code: 200,
		Msg:  "退出登录成功！",
	})
}

func (u *UserHandler) Edit(ctx *gin.Context) {
	type EditReq struct {
		//Email    string `json:"email"`
		//Password string `json:"password"`
		Nickname string `json:"nickname"`
		//Phone    string `json:"phone"`
		Birthday string `json:"birthday"`
		Profile  string `json:"profile"`
	}
	var req EditReq
	if err := ctx.Bind(&req); err != nil {
		return
	}
	if req.Nickname == "" {
		ctx.JSON(http.StatusOK, Result{
			Code: 200,
			Msg:  "昵称不能为空",
		})
		return
	}
	if len(req.Profile) > 1024 {
		ctx.JSON(http.StatusOK, Result{
			Code: 200,
			Msg:  "简介不能超过1024个字符",
		})
		return
	}

	claims, _ := ctx.Get("userClaims")
	userClaims, ok := claims.(*UserClaims)
	if !ok {
		ctx.JSON(http.StatusOK, Result{
			Code: 500,
			Msg:  "系统异常",
		})
	}

	err := u.svc.UpdateNoSensitiveInfo(ctx, domain.User{
		Id:       userClaims.Uid,
		Nickname: req.Nickname,
		Birthday: req.Birthday,
		Profile:  req.Profile,
	})
	switch {
	case errors.Is(err, service.ErrUserNotFound):
		ctx.JSON(http.StatusOK, Result{
			Code: 200,
			Msg:  "修改失败,用户不存在",
		})
		return
	case err == nil:
		ctx.JSON(http.StatusOK, Result{
			Code: 500,
			Msg:  "修改成功",
		})
	default:
		ctx.JSON(http.StatusOK, Result{
			Code: 200,
			Msg:  "系统异常",
		})
		return
	}
}

func (u *UserHandler) Profile(ctx *gin.Context) {
	//sess := sessions.Default(ctx)
	//id := sess.Get("userid")
	//if id == nil {
	//	ctx.JSON(http.StatusOK, gin.H{
	//		"code": 200,
	//		"msg":  "请先登录",
	//	})
	//	return
	//}
	//uid, ok := id.(int64)
	claims, _ := ctx.Get("userClaims")
	userClaims, ok := claims.(*UserClaims)
	if !ok {
		ctx.JSON(http.StatusOK, Result{
			Code: 500,
			Msg:  "系统异常",
		})
	}

	user, err := u.svc.Profile(ctx, userClaims.Uid)
	if errors.Is(err, service.ErrUserNotFound) {
		ctx.JSON(http.StatusOK, Result{
			Code: 200,
			Msg:  "用户不存在",
		})
		return
	}
	ctx.JSON(200, Result{
		Code: 200,
		Msg:  "success",
		Data: domain.User{
			Id:       user.Id,
			Email:    user.Email,
			Nickname: user.Nickname,
			Phone:    user.Phone,
			Birthday: user.Birthday,
			Profile:  user.Profile,
			Ctime:    user.Ctime,
			Utime:    user.Utime,
		},
	})
}

func (u *UserHandler) Delete(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": "这是删除",
	})
}

func (u *UserHandler) SendSMSLoginCode(ctx *gin.Context) {
	type Req struct {
		Phone string `json:"phone"`
	}
	var req Req
	if err := ctx.Bind(&req); err != nil {
		return
	}
	if req.Phone == "" {
		ctx.JSON(http.StatusOK, Result{
			Code: 500,
			Msg:  "手机号不能为空",
		})
		return
	}
	err := u.codeSvc.Send(ctx, Biz, req.Phone)
	if errors.Is(err, service.ErrCodeSendTooMany) {
		ctx.JSON(http.StatusOK, Result{
			Code: 500,
			Msg:  "发送短信过于频繁",
		})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusOK, Result{
			Code: 500,
			Msg:  "系统异常",
		})
		return
	}
	ctx.JSON(http.StatusOK, Result{
		Code: 200,
		Msg:  "发送成功!",
	})
}

func (u *UserHandler) LoginSMS(ctx *gin.Context) {
	type Req struct {
		Phone string `json:"phone"`
		Code  string `json:"code"`
	}
	var req Req
	if err := ctx.Bind(&req); err != nil {
		return
	}
	ok, err := u.codeSvc.Verify(ctx, Biz, req.Phone, req.Code)
	if errors.Is(err, service.ErrCodeVerifyTooMany) {
		ctx.JSON(http.StatusOK, Result{
			Code: 500,
			Msg:  "验证码验证次数过多",
		})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusOK, Result{
			Code: 500,
			Msg:  "系统异常",
		})
		return
	}
	if !ok {
		ctx.JSON(http.StatusOK, Result{
			Code: 500,
			Msg:  "验证码错误",
		})
		return
	}
	//验证成功后，生成token
	user, err := u.svc.FindOrCreate(ctx, req.Phone)
	if err != nil && !errors.Is(err, service.ErrUserNotFound) && !errors.Is(err, service.ErrDuplicateEmail) {
		ctx.JSON(http.StatusOK, Result{
			Code: 500,
			Msg:  "系统异常",
		})
		return
	}
	err = u.setJwtToken(ctx, user.Id)
	if err != nil {
		ctx.JSON(http.StatusOK, Result{
			Code: 500,
			Msg:  "系统异常",
		})
	}
	ctx.JSON(http.StatusOK, Result{
		Code: 200,
		Msg:  "注册并登录成功！",
	})
}

func (u *UserHandler) setJwtToken(ctx *gin.Context, uid int64) error {
	// 使用JWT
	// 生成jwt
	claims := UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		Uid: uid,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenStr, err := token.SignedString([]byte("SUwcr3HfInY49a4XVQ03lV4u1AgcQkynTkf5dPbEAknqr8K7zh5WFNLLPgpUocHi"))
	if err != nil {
		return err
	}
	ctx.Header("x-jwt-token", tokenStr)
	return nil
}
