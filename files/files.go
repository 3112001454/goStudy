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

func WriteBytes(){
	f, err := os.Create("./bytes")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	d := []byte{
		104,
		101,
		108,
		111,
		32,
		119,
		111,
		114,
		108,
		100,
	}
	n, err := f.Write(d)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n, "bytes written successfully")
}


