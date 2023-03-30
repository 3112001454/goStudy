/*
* @Author: 龚国宁
* @Date: 2023/3/30 18:17
* @功能:
 */

package files

import (
	"fmt"
	"os"
)

func WriteString(){
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	l, err := f.WriteString("Hello, world!")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(l, "string written successfully")
}


