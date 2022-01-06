package test

import (
	"fmt"
	"testing"

	"github.com/71010068/tool-box/validator"
)

func TestValidator(t *testing.T) {

	type User struct {
		ID   int      `validate:"required,gt=0" json:"id"`
		Name []string `validate:"required,gt=1,lt=3" label:"姓名"` // 姓名
		Age  int      `validate:"required,gt=0"  json:"-"`
		Addr string   `validate:"max=100"  json:"addr" label:"地址"`
	}

	var me = User{
		ID:   1,
		Age:  14,
		Name: []string{"小哥哥", "小美眉"},
		Addr: "       - w -     ",
	}

	if err := validator.Validator(me); err != nil {
		fmt.Println(err)
	}

	fmt.Println(me)

}
