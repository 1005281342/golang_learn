package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	sklMaxLevel int = 18
)

var sklProbs = [sklMaxLevel - 1]int64{ // 层级分布阈值
	3393088950634442752, 1248247667004394496, 459204654181133312, 168931951563480736,
	62146591937174464, 22862453512557408, 8410626622007697, 3094096621605848,
	1138254536086807, 418740442646473, 154046000036667, 56670356408185,
	20847859046429, 7669498735621, 2821450908925, 1037953783668, 381841857897,
}

type sklKeyType = int
type sklElemType = int
type sklKeyCmpCb = func(sklKeyType, sklKeyType) bool

type SkipListNode struct {
	Value sklElemType     // 元素计数
	key   sklKeyType      // 节点元素
	next  []*SkipListNode // 下一个节点(make([]*SkipListNode, sklMaxLevel))
}

type SkipList struct {
	eqCb      sklKeyCmpCb     // a == b
	lessCb    sklKeyCmpCb     // a < b
	root      SkipListNode    // 头节点
	length    int             // 链表长度
	randSrc   rand.Source     // 随机数
	prevNodes []*SkipListNode //  上一个节点(make([]*SkipListNode, sklMaxLevel))
}

func (l *SkipList) Init(eqCb, lessCb sklKeyCmpCb) *SkipList {
	l.eqCb, l.lessCb = eqCb, lessCb
	l.root.next = make([]*SkipListNode, sklMaxLevel)
	l.randSrc = rand.NewSource(time.Now().UnixNano())
	l.prevNodes = make([]*SkipListNode, sklMaxLevel)
	return l
}

func (l *SkipList) Get(key sklKeyType) *SkipListNode {
	p := &l.root // 从头节点开始查询
	var n *SkipListNode
	for i := sklMaxLevel - 1; i >= 0; i-- { // 降级查找
		n = p.next[i]
		for n != nil && l.lessCb(n.key, key) {
			p, n = n, n.next[i]
		}
	}
	if n != nil && l.eqCb(n.key, key) {
		return n
	}
	return nil
} // 查询

func (l *SkipList) getPrevNodes(key sklKeyType) []*SkipListNode {
	p := &l.root
	prevs := l.prevNodes                    // 上一个节点
	for i := sklMaxLevel - 1; i >= 0; i-- { // 从根节点遍历，查找当前节点的上一节点
		n := p.next[i]
		for n != nil && l.lessCb(n.key, key) {
			p, n = n, n.next[i]
		}
		prevs[i] = p
	}
	return prevs
} // 查询上一节点

func (l *SkipList) Add(key sklKeyType, value sklElemType) (*SkipListNode, bool) {
	prevs := l.getPrevNodes(key)
	if e := prevs[0].next[0]; e != nil && l.eqCb(e.key, key) {
		return e, false // 该key已存在
	}
	r := l.randSrc.Int63()
	level := 1
	for ; level < sklMaxLevel && r < sklProbs[level-1]; level++ {
	} // 获取层级
	node := &SkipListNode{value, key, make([]*SkipListNode, level)} // 初始化当前节点
	for i := 0; i < level; i++ {
		node.next[i] = prevs[i].next[i] // 当前节点指向"上一节点"的next
		prevs[i].next[i] = node         // 更新"上一节点"的next指向当前节点node
	}
	l.length++ // 跳表长度加1
	return node, true
} // 插入节点

func (l *SkipList) Remove(key sklKeyType) *SkipListNode {
	prevs := l.getPrevNodes(key) // 上一个节点
	if e := prevs[0].next[0]; e != nil && l.eqCb(e.key, key) {
		for i, n := range e.next {
			prevs[i].next[i] = n // 将当前节点的上一节点的next指向当前节点的下一节点
		}
		l.length-- // 更新跳表长度
		return e
	}
	return nil
} // 移除节点

type Skiplist struct {
	l SkipList
}

func Constructor() Skiplist {
	sl := Skiplist{SkipList{}}
	sl.l.Init(
		func(a, b int) bool { return a == b },
		func(a, b int) bool { return a < b })
	return sl
}

func (sl *Skiplist) Search(target int) bool {
	return sl.l.Get(target) != nil
}

func (sl *Skiplist) Add(num int) {
	if n, ok := sl.l.Add(num, 1); !ok {
		n.Value++ // 说明该key[num]已经存在了，将计数Value+1即可
	}
}

func (sl *Skiplist) Erase(num int) bool {
	if n := sl.l.Get(num); nil != n {
		n.Value--         // 查找到时，将其计数减1
		if 0 == n.Value { // 当计数为0时，将其从跳表中移除并调整结构
			sl.l.Remove(num)
		}
		return true
	}
	return false
}

func main() {
	ct := Constructor()

	for i := 0; i < 20000; i++ {
		ct.Add(i)
	}

	fmt.Printf("%+v", ct)
	fmt.Println()
	b := 0
	a := ct.Search(b)
	fmt.Println(b, a)
	b = 100
	a = ct.Search(b)
	fmt.Println(b, a)
	a = ct.Erase(b)
	fmt.Println(b, a)
}
