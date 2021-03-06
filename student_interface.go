package main
import (
	"fmt"
)

// 父结构
type human struct {
	name string
	age  int
}

type student struct {
	human //匿名结构
	id int
}

type employ struct {
	human
	job string
}

// 定义human的方法
func (h *human) sleep(do string) string {
	fmt.Println(do)
	return  do
}

func (s *student) paly(do string) string {
	fmt.Println(do)
	return  do
}

func (e *employ) work(do string) string {
	fmt.Println(do)
	return  do
}

// 定有接口
type man interface {
	sleep(do string)
	paly(do string)
	work(do string)
}

func main() {
	var DaTouSon man
	var UncleWang man

	DaTouSon = student{human:human{name:"大头儿子", age:5},id:7}
	UncleWang = employ{human:human{name:"隔壁王叔叔", age:35},job:"develop"}

	DaTouSon.sleep("呼呼呼")
	DaTouSon.play("看漫画")
	DaTouSon.work("写作业")
	UncleWang.sleep("啪啪啪")
	UncleWang.play("玩游戏")
	UncleWang.work("写代码")
}
