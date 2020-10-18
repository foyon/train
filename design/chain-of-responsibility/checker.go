/**
* @File:          checker.go
* @Aim:           审批者
* @Author:        foyon
* @Created Time:  2020-10-18
 */
package main

import (
	"fmt"
)

type CheckOpt interface {
	isPass(day int) bool
}

type Checker struct {
	Uname string
	Lv    int
	Day   int
	Next  *Checker
}

func (v1 *Checker) isPass(day int) bool {

}

func (v2 *Checker) isPass(day int) bool {

}
