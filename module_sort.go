/**
* @File:          module_sort.go
* @Aim:           根据某天数据，各类型进行自定义排序
* @Author:        foyon
* @Created Time:  2020-09-08
 */
package noticeservice

import (
	"fmt"
	"sort"
)

type ModuleInfo struct {
	Name string
	Num  float64
}

type Modules []ModuleInfo

// Len()方法和Swap()方法不用变化
// 获取此 slice 的长度
func (p Modules) Len() int { return len(p) }

// 交换数据
func (p Modules) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// 嵌套结构体  将继承 Person 的所有属性和方法
// 所以相当于SortByName 也实现了 Len() 和 Swap() 方法
type SortByName struct{ Modules }

// 根据元素的姓名长度降序排序 （此处按照自己的业务逻辑写）
func (p SortByName) Less(i, j int) bool {
	return len(p.Modules[i].Name) > len(p.Modules[j].Name)
}

type SortByNum struct{ Modules }

// 根据元素的年龄降序排序 （此处按照自己的业务逻辑写）
func (p SortByNum) Less(i, j int) bool {
	return p.Modules[i].Num > p.Modules[j].Num
}

func DemoTest() { // {{{
	persons := Modules{
		{
			Name: "test123",
			Num:  20,
		},
		{
			Name: "test1",
			Num:  22,
		},
		{
			Name: "test12",
			Num:  21,
		},
	}

	fmt.Println("排序前")
	for _, person := range persons {
		fmt.Println(person.Name, ":", person.Num)
	}
	sort.Sort(SortByName{persons})

	fmt.Println("排序后")
	for _, person := range persons {
		fmt.Println(person.Name, ":", person.Num)
	}
} // }}}
