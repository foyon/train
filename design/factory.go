/**
* @File:          factory.go
* @Aim:           工厂设计模式
*                 父类提供创建对象方法，允许子类自定义标准接口
* @Author:        foyon
* @Created Time:  2020-09-19
 */

/**
php interface 与abstract class 区别
1、接口只约定方法,不能私有，没有变量及属性，没有代码实现，不能被实例化
2、接口的方法必须全部实现
3、HasOne
============================
1、抽象类不能被实例化，只能被继承，可以有变量及属性，约定方法可有代码实现
2、final func 不能被重写覆盖，final class 不能被继承
3、IsOne
*/

package main

import (
	"fmt"
)

type Op interface {
	getName() string
	//getNum() int
}

type A struct {
}

type B struct {
}

type Factory struct {
}

func (a *A) getName() string {
	return "A"
}

func (b *B) getName() string {
	return "B"
}

func (f *Factory) create(name string) Op {
	switch name {

	case `a`:
		return new(A)
	case `b`:
		return new(B)
	default:
		panic("name not exist")
	}

	return nil
}

func testInterface() {
	interfaceA := new(A)
	fmt.Printf("%+v", interfaceA.getName())
}

func main() {
	f := new(Factory)
	inf := f.create("b")
	fmt.Printf("%+v", inf)
	fmt.Printf(inf.getName())

	testInterface()
}
