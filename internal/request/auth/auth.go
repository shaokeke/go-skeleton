package auth

type IndexReq struct {
	Name string `form:"name" query:"name" json:"name" xml:"name" validate:"required"`
	Id   int    `form:"id" query:"id" xml:"id" validate:"required"`
}

type LoginReq struct {
	UserName  string `form:"username" json:"username" validate:"required"`
	Password  string `form:"password" json:"password" validate:"required"`
	Captcha   string `form:"captcha" json:"captcha" validate:"required"`
	CaptchaId string `form:"captchaId" json:"captchaId" validate:"required"`
}

type RigsterReq struct {
	UserName string `form:"username" json:"username" validate:"required"`
	Password string `form:"password" json:"password" validate:"required"`
}
