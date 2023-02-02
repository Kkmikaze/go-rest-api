package entity

type ReqBodyRegister struct {
	Name     string `form:"name" json:"name" xml:"name" binding:"required"`
	Email    string `form:"email" json:"email" xml:"email" binding:"required,email"`
	Password string `form:"password" json:"password" xml:"password" binding:"required,gte=6"`
}

type ReqBodyLogin struct {
	Email    string `form:"email" json:"email" xml:"email" binding:"required,email"`
	Password string `form:"password" json:"password" xml:"password" binding:"required,gte=6"`
}
