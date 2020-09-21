package main

import (
	"fmt"
	"regexp"
	//"strings"
)

func main() {
	FormatStr("test one 苹果")
}

func FormatStr(str string) string {

	fmt.Printf("%v", str)
	mark := []map[string]interface{}{
		{`苹果[\S]*`: "apple iphone"},
		{`one`: "one plus"},
	}
	for _, k := range mark {
		for key, _ := range k {
			//fmt.Printf("%v", k[key])
			valre := k[key]

			//fmt.Printf("%v", k[key])
			//return str
			re, _ := regexp.Compile(key)
			str = re.ReplaceAllString(str, valre.(string))
			//fmt.Printf("%v", str)
		}
	}

	fmt.Printf("%v", str)
	return str
}
