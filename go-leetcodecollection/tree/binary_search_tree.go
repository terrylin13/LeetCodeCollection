package tree

import "errors"

/**
 * 如果要插入的数据比节点的数据大，并且节点的右子树为空，就将新数据直接插到右子节点的位置；
 * 如果不为空，就再递归遍历右子树，查找插入位置。
 * 同理，如果要插入的数据比节点数值小，并且节点的左子树为空，就将新数据插入到左子节点的位置；
 * 如果不为空，就再递归遍历左子树，查找插入位置。
 */
func (t *TreeNode) Insert(value int) error {
	if t == nil {
		return errors.New("tree is nil")
	}

	p := t
	for p != nil {
		if p.Val > value {
			if p.Left == nil {
				p.Left = &TreeNode{Val: value}
				return nil
			}
			p = p.Left
		} else if p.Val < value {
			if p.Right == nil {
				p.Right = &TreeNode{Val: value}
				return nil
			}
			p = p.Right
		} else {
			return errors.New("this node value already exists")
		}
	}
	return errors.New("insert error")
}

func (t *TreeNode) Find(value int) *TreeNode {
	// if t == nil {
	// 	return nil
	// }

	// if t.Val > value {
	// 	return t.Left.Find(value)
	// } else if t.Val < value {
	// 	return t.Right.Find(value)
	// } else {
	// 	return t
	// }

	p := t
	for p != nil {
		if p.Val < value {
			p = p.Right
		} else if p.Val > value {
			p = p.Left
		} else {
			return p
		}
	}
	return nil
}

/**
 * 第一种情况是，如果要删除的节点没有子节点，我们只需要直接将父节点中，指向要删除节点的指针置为 null。
 * 第二种情况是，如果要删除的节点只有一个子节点（只有左子节点或者右子节点），我们只需要更新父节点中，指向要删除节点的指针，让它指向要删除节点的子节点就可以了。
 * 第三种情况是，如果要删除的节点有两个子节点，这就比较复杂了。我们需要找到这个节点的右子树中的最小节点，把它替换到要删除的节点上。
 * 然后再删除掉这个最小节点，因为最小节点肯定没有左子节点（如果有左子结点，那就不是最小节点了），所以，我们可以应用上面两条规则来删除这个最小节点。
 */
func (t *TreeNode) Remove(value int) error {
	p := t
	var pp *TreeNode

	for p != nil && p.Val != value {
		pp = p
		if p.Val < value {
			p = p.Right
		} else {
			p = p.Left
		}
	}

	if p == nil {
		return errors.New("not found value")
	}

	if p.Left != nil && p.Right != nil {
		minP := p.Right
		minPP := p // minPP 表示 minP 的父节点
		for minP.Left != nil {
			minPP = minP
			minP = minP.Left
		}
		p.Val = minP.Val // 将 minP 的数据替换到 p 中
		p = minP         // 下面就变成了删除 minP 了
		pp = minPP
	}

	// 删除节点是叶子节点或者仅有一个子节点
	var child *TreeNode // p 的子节点
	if p.Left != nil {
		child = p.Left
	} else if p.Right != nil {
		child = p.Right
	} else {
		child = nil
	}

	if pp.Left == p {
		p.Left = child
	} else if pp.Right == p {
		p.Right = child
	} else {
		return errors.New("you try to remove root")
	}
	return nil
}
