package model

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type Sex string

func (s Sex) IsValid() bool {
	switch s {
	case "男", "女":
		return true
	}
	return false
}

func SexRegisterTranslation(trans ut.Translator) (string, ut.Translator, validator.RegisterTranslationsFunc, validator.TranslationFunc) {
	return "enum", trans, func(ut ut.Translator) error {
			return ut.Add("enum", "{0}必须是(男|女)中的任何一个，当前是: {1}", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("enum", fe.Namespace(), string(fe.Value().(Sex)))
			return t
		}
}

type Student struct {
	// 注意只能使用等号. 注意中间不能有空格
	Name string `json:"name" validate:"required,gte=5,lte=5,len=5,contains=a|contains=b,excludes=us,startswith=cn,endswith=cn"`

	// 注意这个不等于的大小写。 necsfield field不同
	Name2 string `json:"name2" validate:"required,gte=5,lte=5,len=5,contains=cn,excludes=us,startswith=cn,endswith=cn,necsfield=Name"`

	// 自定义注册 enum
	Sex Sex `json:"sex" validate:"enum"`

	// 常用于校验密码是否相等, 长度是否符合要求。eqfield field 相同
	Password  string `json:"password" validate:"required,min=5,max=10"`
	Password2 string `json:"password2" validate:"required,eqfield=Password"`

	// 校验 Ip 问题. ip ipv4 ipv6 uri url
	Ip string `json:"ip" validate:"required,ip"`

	Storage string `json:"storage" validate:"base64"`

	// dive 用来在深层次进行验证
	Address   []string `json:"address" validate:"required,min=1,max=1,dive,len=5,contains=cn,excludes=us,startswith=cn,endswith=cn"`
	Privilege string   `json:"privilege" validate:"required,oneof=boss admin user"`
}
