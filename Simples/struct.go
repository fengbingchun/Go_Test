// 结构体相关测试代码
package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	// 访问结构体成员
	var Book1, Book2 Books

	Book1 = Books{"Go语言", "www.runoob.com", "Go语言教程", 649507}

	Book2.title = "Python 教程"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 6495700

	fmt.Printf("Book 1 title: %s, autor: %s, subject: %s, book_id: %d\n",
		Book1.title, Book1.author, Book1.subject, Book1.book_id)

	fmt.Printf("Book 2 title: %s, autor: %s, subject: %s, book_id: %d\n",
		Book2.title, Book2.author, Book2.subject, Book2.book_id)

	// 结构体作为函数参数
	printBook(Book1)
	printBook(Book2)

	// 结构体指针
	printBook2(&Book1)
	printBook2(&Book2)
}

func printBook(book Books) {
	fmt.Printf("Book 1 title: %s, autor: %s, subject: %s, book_id: %d\n",
		book.title, book.author, book.subject, book.book_id)
}

func printBook2(book *Books) {
	fmt.Printf("Book 1 title: %s, autor: %s, subject: %s, book_id: %d\n",
		book.title, book.author, book.subject, book.book_id)
}
