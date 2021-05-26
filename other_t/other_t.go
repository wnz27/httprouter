/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/5/26 5:13 下午
* Description:
 */
package main

import "fmt"

func main()  {
	a := "asdfasdf/asdfas/asdf"
	for idx, c := range []byte(a) {
		fmt.Println(idx, c)

	}
}


