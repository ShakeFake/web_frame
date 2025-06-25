package utils

import (
	log "github.com/cihub/seelog"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslation "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"sync"
	"wilikidi/gin/model"
)

var VALIDATE *DefaultValidator
var trans ut.Translator

func BindValidator() *DefaultValidator {
	VALIDATE = new(DefaultValidator)

	// 注册翻译器
	trans, _ = ut.New(zh.New()).GetTranslator("zh")
	err := zhTranslation.RegisterDefaultTranslations(VALIDATE.Engine().(*validator.Validate), trans)
	if err != nil {
		log.Infof(err.Error())
	}

	// 自定义注册校验接口，注册自定义 validator 翻译器
	VALIDATE.Engine().(*validator.Validate).RegisterTranslation(model.SexRegisterTranslation(trans))

	return VALIDATE
}

// DefaultValidator 定义主持刷校验器
type DefaultValidator struct {
	once     sync.Once
	validate *validator.Validate
}

func (v *DefaultValidator) ValidateStruct(obj interface{}) error {

	if kindOfData(obj) == reflect.Struct {

		v.lazyinit()

		if err := v.validate.Struct(obj); err != nil {
			return error(err)
		}
	}

	return nil
}

func (v *DefaultValidator) Engine() interface{} {
	v.lazyinit()
	return v.validate
}

func (v *DefaultValidator) lazyinit() {
	v.once.Do(func() {
		v.validate = validator.New()
		// validate 代表开启的校验的标签
		v.validate.SetTagName("validate")

		// add any custom validations etc. here
		v.validate.RegisterValidation("enum", ValidateEnum)
	})
}

func kindOfData(data interface{}) reflect.Kind {

	value := reflect.ValueOf(data)
	valueType := value.Kind()

	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}

// Enum 用户自定义注册器
type Enum interface {
	IsValid() bool
}

func ValidateEnum(fl validator.FieldLevel) bool {
	value := fl.Field().Interface().(Enum)
	return value.IsValid()
}

// ErrorTranslate 将翻译为对应的翻译器
func ErrorTranslate(err error) (message string) {
	if validationErrors, ok := err.(validator.ValidationErrors); !ok {
		return err.Error()
	} else {
		for _, e := range validationErrors {
			message += e.Translate(trans) + ";"
		}
	}
	return message
}
