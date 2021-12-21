package test

import (
	"fmt"
	"testing"

	"github.com/71010068/tool-box/validator"
)

func TestValidator(t *testing.T) {

	type User struct {
		ID   int      `validate:"required,gt=0" json:"id"` // 姓名
		Name []string `validate:"required,gt=1,lt=3"`
		Age  int      `validate:"required,gt=0"  json:"-"`
	}

	var me = User{
		ID:  1,
		Age: 14,
	}

	if err := validator.Validator(me); err != nil {
		fmt.Println(err)
	}

	fmt.Println(me)

}
