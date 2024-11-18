// Trong 1 file go bắt buộc phải có package clause, để go biết file đó thuộc về package nào.
// package clause
package main

// File scope: các khai báo này chỉ có thể dùng được ở file này không được dùng ở file khác , mặc dùng chung 1 package
import "fmt"

// Package scope: Khi chữ cái đầu của func là viết thường thì chỉ func đó chỉ dùng được qua lại ở các file trong cùng 1 package
// còn chữ cái đầu viết hoa thì ở trong package nào cũng dùng được.

// Có 2 loại package đó là executable package and library package
// Executable:
// +) tên bắt buộc phải là package main
// +) func cũng bắt buộc phải là main
// +) được tạo để go runtime tìm đến để bắt đầu chạy

// Library
// +)Được tạo ra để tái sử dụng
// +)Không bắt buộc phải thực thi
// +)Có nhiều tên
// Không được là func main

func main() {
	fmt.Println(testScope)
}
