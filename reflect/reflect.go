/*
* @Author: grant
* @Date: 2023/2/6 18:13
* @功能:
 */

package reflect

import (
	"fmt"
	"reflect"
)

// 反射第一定律:反射可以将interface类型变量转换成反射对象

func Demo1() {
	var x float64 = 3.1415
	t := reflect.TypeOf(x)
	fmt.Println(t)
	v := reflect.ValueOf(x)
	fmt.Println(v)
}

// 反射第二定律: 反射可以将将反射对象还原为interface对象

func Demo2() {
	var x float64 = 3.1415
	v := reflect.ValueOf(x)
	var y float64 = v.Interface().(float64)
	fmt.Println(y)
}

// 反射第三定律:反射对象可修改，value值必须是可设置的

func Demo3() {
	var x float64 = 3.1415
	v := reflect.ValueOf(&x)
	v.Elem().SetFloat(7.1)
	fmt.Println(v.Elem().Interface())
}