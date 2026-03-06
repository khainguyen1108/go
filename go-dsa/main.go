package main

import (
	hashtable "GO/go-dsa/hashtable"
)

func main() {
	obj := hashtable.Constructor()
	obj.Inc("a")
	obj.Inc("b")
	obj.Inc("c")
	obj.Inc("d")
	obj.Inc("e")
	obj.Inc("f")
	obj.Inc("g")
	obj.Inc("h")
	obj.Inc("i")
	obj.Inc("j")
	obj.Inc("k")
	obj.Inc("l")
	obj.Dec("a")
	obj.Dec("b")
	obj.Dec("c")
	obj.Dec("d")
	obj.Dec("e")
	obj.Dec("f")
	obj.Inc("g")
	obj.Inc("h")
	obj.Inc("i")
	obj.Inc("j")
	obj.GetMaxKey()
	obj.GetMinKey()
}
