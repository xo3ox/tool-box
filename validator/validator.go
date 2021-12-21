package validator

import (
	"errors"
	"reflect"

	"github.com/go-playground/locales/zh_Hans_CN"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/zh"

	ut "github.com/go-playground/universal-translator"
)

// 初始化翻译器
func Validator(data interface{}) error {
	// 初始化校验器
	validate := validator.New()

	// 初始化中文翻译器,万能翻译器，保存所有的语言环境和翻译数据
	uni := ut.New(zh_Hans_CN.New())
	// 获取所需要的语言
	trans, _ := uni.GetTranslator("zh_Hans_CN")

	if err := zh.RegisterDefaultTranslations(validate, trans); err != nil {
		return err
	}

	// 注册 根据json名称返回字段名 的函数
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		var tag string
		if tag == "" {
			tag = field.Tag.Get("label")
			if tag == "" {
				tag = field.Tag.Get("json")
				if tag == "-" {
					return field.Name
				}
			}
		}

		return tag
	})

	if err := validate.Struct(data); err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			return errors.New(v.Translate(trans))
		}
	}
	return nil
}
