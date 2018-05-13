package main

import (
	"fmt"
	"gitlab.xunlei.cn/pub/fields-transfer/pkg/field"
)

func main() {
	type Person struct {
		Name string `form:"name"`
		Age  uint8  `form:"age"`
	}

	binder := &field.Parser{}
	person := &Person{"zhangsan", 18}

	params := make(map[string][]string)
	params["name"] = []string{"张三"}
	params["age"] = []string{"李四"}

	binder.Bind(person, params, "form")
	fmt.Println(person)
}
