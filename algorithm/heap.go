// heap
package algorithm

type head interface {
	Less(int, int) bool
	Swap(int, int)
	Len() int
}

func AdjustHeap(a head, i int) {
	lchild := 2*i + 1
	rchild := 2*i + 2
	l := a.Len()
	if lchild >= l { // 当前结点没有子结点
		return
	} else if rchild <= l-1 { // 当前有左右子节点
		r1 := a.Less(i, lchild)
		r2 := a.Less(i, rchild)
		if r1 && r2 { //当前结点已经是最小值 无需替换
			return
		}
		r3 := a.Less(lchild, rchild)
		if !r1 && r3 { // 左孩子值最小  i > l, l < r
			a.Swap(i, lchild)
			AdjustHeap(a, lchild)
			return
		}
		if !r2 && !r3 { // 右孩子值最小  i > r, l > r
			a.Swap(i, rchild)
			AdjustHeap(a, rchild)
			return
		}
	} else { // 只有左结点
		r := a.Less(i, lchild)
		if !r {
			a.Swap(i, lchild)
			AdjustHeap(a, lchild)
		}
	}
}

func CreateHeap(a head) {
	l := a.Len()
	lp := (l - 1) / 2
	for i := lp; i >= 0; i-- {
		AdjustHeap(a, i)
	}
}
