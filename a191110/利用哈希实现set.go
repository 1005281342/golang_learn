package main

type Set struct {
	length int
	buf    []interface{}
	hash   map[interface{}]bool
}

func main() {
	obj := NewSet()
	obj.Add(1)
	obj.Add(1)
	obj.Add(4)
}

func NewSet() *Set {
	return new(Set)
}

func (self *Set) Add(v interface{}) bool {
	if self.IsExist(v) {
		return false
	} else {
		self.length++
		self.buf = append(self.buf, v)
		self.hash[v] = true
		return true
	}
}

func (self *Set) IsExist(v interface{}) bool {
	return self.hash[v]
}
