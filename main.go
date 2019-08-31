package main

import (
	"begin/model"
	"begin/pkg/setting"
)

func main() {
	setting.SetUp("conf/app.ini")
	model.SetUp()
}
