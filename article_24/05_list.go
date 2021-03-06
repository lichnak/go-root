// Seriál "Programovací jazyk Go"
//
// Dvacátá čtvrtá část
//
// Demonstrační příklad číslo 5:
//     	Operace nad seznamem

package main

import (
	"container/list"
	"fmt"
)

func printList(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func main() {
	l := list.New()
	l.PushBack("foo")
	l.PushBack("bar")
	l.PushBack("baz")
	printList(l)
}
