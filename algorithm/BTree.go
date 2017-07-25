// BTree
package algorithm

const (
	RED   = true
	BLACK = false
)

type node struct {
	lt    *node
	rt    *node
	p     *node
	color bool
	v     int
	count int
}

func (n *node) GetValue() int {
	return n.v
}
func (n *node) SetValue(v int) {
	n.v = v
}

func (n *node) IsRed() bool {
	return n.color
}

func NewNode() *node {
	return &node{
		color: RED,
		count: 0,
	}
}

type Tree struct {
	root *node
	Less func(a, b int) bool
}

func (t *Tree) insertFlipColor(z *node) {
	for z.p.IsRed() {
		if z.p == z.p.p.lt {
			/*
				        爷爷
					爸爸      y
				z
			*/
			y := z.p.p.rt
			if y.IsRed() {
				// case 1:z的叔叔是红色
				z.p.color = BLACK
				y.color = BLACK
				z.p.p.color = RED
				z = z.p.p
			} else {
				if z == z.p.rt {
					/*
						 case 2:z的叔叔是黑色 z是右孩子场景
						需要先翻转为左孩子场景
						        爷爷
							爸爸      y
						         z
								|
								|
								V
						        爷爷
							z         y
						爸爸
					*/
					z = z.p
					t.rl(z)
				}
				/*
					case 3:
						        爷爷
							爸爸         y
						z
				*/
				z.p.color = BLACK
				z.p.p.color = RED
				t.rr(z.p.p)
			}
		} else { //  分析插入右边情况
			/*
							       爷爷
								y       爸爸
				                               z
			*/
			y := z.p.p.lt
			if y.IsRed() {
				// case 1:z的叔叔是红色
				z.p.color = BLACK
				y.color = BLACK
				z.p.p.color = RED
				z = z.p.p
			} else {
				if z == z.p.lt {
					/*
						 case 2:z的叔叔是黑色 z是左孩子场景
						需要先翻转为左孩子场景
						        爷爷
							y        爸爸
						         z
								|
								|
								V
						        爷爷
							y         z
						                 爸爸
					*/
					z = z.p
					t.rr(z)
				}
				/*
					case 3:
						        爷爷
							y         爸爸
						                   z
				*/
				z.p.color = BLACK
				z.p.p.color = RED
				t.rl(z.p.p)
			}
		}
	}
	t.root.color = BLACK
}

//左旋转
func (t *Tree) rl(x *node) *node {
	y := x.rt
	x.rt = y.lt
	if y.lt != nil {
		y.lt.p = x
	}
	y.p = x.p
	if x.p == nil {
		t.root = y
	} else if x == x.p.lt {
		x.p.lt = y
	} else {
		x.p.rt = y
	}
	y.lt = x
	x.p = y
	return y
}

//右旋转
func (t *Tree) rr(y *node) *node {
	x := y.lt
	y.lt = x.rt
	if x.rt != nil {
		x.rt.p = y
	}
	x.p = y.p
	if y.p == nil {
		t.root = x
	} else if y == y.p.lt {
		y.p.lt = x
	} else {
		y.p.rt = x
	}
	x.rt = y
	y.p = x
	return x
}

func (t *Tree) insert(z *node) {
	var y *node
	x := t.root
	for x != nil {
		y = x
		if t.Less(z.v, x.v) {
			x = x.lt
		} else {
			x = x.rt
		}
	}
	z.p = y
	if y == nil {
		t.root = z
	} else if t.Less(z.v, y.v) {
		y.lt = z
	} else {
		y.rt = z
	}
	z.lt = nil
	z.rt = nil
	t.insertFlipColor(z)
}

func (t *Tree) Insert(v int) {
	n := NewNode()
	n.v = v
	t.insert(n)
}

func NewTree() *Tree {
	return &Tree{}
}
