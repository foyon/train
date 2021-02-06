/**
* @File:          option.go
* @Aim:
* @Author:        foyon
* @Created Time:  2020-10-27
 */

/*
Option模式的优缺点

优点
支持传递多个参数，并且在参数个数、类型发生变化时保持兼容性
任意顺序传递参数
支持默认值
方便拓展

缺点
增加许多function，成本增大
参数不太复杂时，尽量少用
*/

package main

import (
	"fmt"
)

type Message struct {
	id      int
	name    string
	address string
	phone   int
}

func (msg Message) String() {
	fmt.Printf("ID:%d \n- Name:%s \n- Address:%s \n- phone:%d\n", msg.id, msg.name, msg.address, msg.phone)
}

func New(id, phone int, name, addr string) Message {
	return Message{
		id:      id,
		name:    name,
		address: addr,
		phone:   phone,
	}
}

type Option func(msg *Message)

var DEFAULT_MESSAGE = Message{id: -1, name: "-1", address: "-1", phone: -1}

func WithID(id int) Option {
	return func(m *Message) {
		m.id = id
	}
}

func WithName(name string) Option {
	return func(m *Message) {
		m.name = name
	}
}

// 对参数统一追加 改写
func WithAddress(addr string) Option {
	return func(m *Message) {
		m.address = "foyon" + addr
	}
}

func WithPhone(phone int) Option {
	return func(m *Message) {
		m.phone = phone
	}
}

func NewByOption(opts ...Option) Message {
	msg := DEFAULT_MESSAGE
	for _, o := range opts {
		fmt.Printf("%+v\n", msg)
		//o == WithID 等函数
		o(&msg)
		fmt.Printf("%+v\n", msg)
	}
	return msg
}

func NewByOptionWithoutID(id int, opts ...Option) Message {
	msg := DEFAULT_MESSAGE
	msg.id = id
	for _, o := range opts {
		o(&msg)
	}
	return msg
}

func main() { // {{{
	message1 := New(1, 123, "message1", "cache1")
	message1.String()
	message2 := NewByOption(WithID(2), WithName("message2"), WithAddress("cache2"), WithPhone(456))
	message2.String()
	message3 := NewByOptionWithoutID(3, WithAddress("cache3"), WithPhone(789), WithName("message3"))
	message3.String()
} // }}}
