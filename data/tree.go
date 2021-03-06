package zdata

import (
	"zCache/types"
	"errors"
	"strconv"
)

var (
	errNotExist       = errors.New("Key Not Existed")
	errTreeNil        = errors.New("tree is null")
	errTreeIndexExist = errors.New("tree Index is existed")
)

func max(data1 int, data2 int) int {
	if data1 > data2 {
		return data1
	}
	return data2
}

func getHeight(node *types.Node) int {
	if node == nil {
		return 0
	}
	return node.Height
}

// 左旋转
//
//    node  BF = 2
//       \
//         pRchild     ----->       pRchild    BF = 1
//           \                        /   \
//           ppRchild               node  ppRchild
func llRotation(node *types.Node) *types.Node {
	pRchild := node.Rchild
	node.Rchild = pRchild.Lchild
	pRchild.Lchild = node
	//更新节点 node 的高度
	node.Height = max(getHeight(node.Lchild), getHeight(node.Rchild)) + 1
	//更新新父节点高度
	pRchild.Height = max(getHeight(pRchild.Lchild), getHeight(pRchild.Rchild)) + 1
	return pRchild
}

// 右旋转
//             node  BF = -2
//              /
//         pLchild     ----->       pLchild    BF = 1
//            /                        /   \
//        ppLchild                Lchild   node

func rrRotation(node *types.Node) *types.Node {
	pLchild := node.Lchild
	node.Lchild = pLchild.Rchild
	pLchild.Rchild = node
	node.Height = max(getHeight(node.Lchild), getHeight(node.Rchild)) + 1
	pLchild.Height = max(getHeight(node), getHeight(pLchild.Lchild)) + 1
	return pLchild
}

// 先左转再右转
//          node                  node
//         /            左          /     右
//      node1         ---->    node2     --->         node2
//          \                   /                     /   \
//          node2s           node1                 node1  node
func lrRotation(node *types.Node) *types.Node {
	pLchild := llRotation(node.Lchild) //左旋转
	node.Lchild = pLchild
	return rrRotation(node)

}

// 先右转再左转
//       node                  node
//          \          右         \         左
//          node1    ---->       node2     --->      node2
//          /                       \                /   \
//        node2                    node1           node  node1
func rlRotation(node *types.Node) *types.Node {
	pRchild := rrRotation(node.Rchild)
	node.Rchild = pRchild
	node.Rchild = pRchild
	return llRotation(node)
}

//处理节点高度问题
func handleBF(node *types.Node) *types.Node {
	if getHeight(node.Lchild)-getHeight(node.Rchild) == 2 {
		if getHeight(node.Lchild.Lchild)-getHeight(node.Lchild.Rchild) > 0 { //RR
			node = rrRotation(node)
		} else {
			node = lrRotation(node)
		}
	} else if getHeight(node.Lchild)-getHeight(node.Rchild) == -2 {
		if getHeight(node.Rchild.Lchild)-getHeight(node.Rchild.Rchild) < 0 { //LL
			node = llRotation(node)
		} else {
			node = rlRotation(node)
		}
	}
	return node
}

//中序遍历树，并根据钩子函数处理数据
func Midtraverse(node *types.Node, handle func(interface{}) error) error {
	if node == nil {
		return nil
	} else {
		if err := handle(node); err != nil {
			return err
		}
		if err := Midtraverse(node.Lchild, handle); err != nil { //处理左子树
			return err
		}
		if err := Midtraverse(node.Rchild, handle); err != nil { //处理右子树
			return err
		}
	}
	return nil
}

//插入节点 ---> 依次向上递归，调整树平衡
func Add(node *types.Node, Key string, Value string) (*types.Node, error) {
	if node == nil {
		return &types.Node{Lchild: nil, Rchild: nil, Key: Key, Value: Value, Height: 1}, nil
	}
	if node.Key > Key {
		node.Lchild, _ = Add(node.Lchild, Key, Value)
		node = handleBF(node)
	} else if node.Key < Key {
		node.Rchild, _ = Add(node.Rchild, Key, Value)
		node = handleBF(node)
	} else {
		node.Value = Value
		return node, nil
	}
	node.Height = max(getHeight(node.Lchild), getHeight(node.Rchild)) + 1
	return node, nil
}

//删除指定Index节点
//查找节点 ---> 删除节点 ----> 调整树结构
//删除节点时既要遵循二叉搜索树的定义又要符合二叉平衡树的要求   ---> 重点处理删除节点的拥有左右子树的情况
func Delete(node *types.Node, Key string) (*types.Node, error) {
	if node == nil {
		return nil, errNotExist
	}
	if node.Key == Key { //找到对应节点
		//如果没有左子树或者右子树 --->直接返回nil
		if node.Lchild == nil && node.Rchild == nil {
			return nil, nil
		} else if node.Lchild == nil || node.Rchild == nil { //若只存在左子树或者右子树
			if node.Lchild != nil {
				return node.Lchild, nil
			} else {
				return node.Rchild, nil
			}
		} else { //左右子树都存在
			//查找前驱，替换当前节点,然后再进行依次删除  ---> 节点删除后，前驱替换当前节点 ---> 需遍历到最后，调整平衡度
			var n *types.Node
			//前驱
			n = node.Lchild
			for {
				if n.Rchild == nil {
					break
				}
				n = n.Rchild
			}
			//
			n.Value, node.Value = node.Value, n.Value
			n.Key, node.Key = node.Key, n.Key
			node.Lchild, _ = Delete(node.Lchild, n.Key)
		}
	} else if node.Key > Key {
		node.Lchild, _ = Delete(node.Lchild, Key)
	} else { //node.Index < Index
		node.Rchild, _ = Delete(node.Rchild, Key)
	}
	//删除节点后节点高度
	node.Height = max(getHeight(node.Lchild), getHeight(node.Rchild)) + 1
	//调整树的平衡度
	node = handleBF(node)
	return node, nil
}

func DeleteAll(node *types.Node) (err error) {
	if node == nil {
		return nil
	}
	if nil != node.Lchild{
		if err = DeleteAll(node.Lchild);err != nil {
			return err
		}

	}
	if nil != node.Rchild{
		if err = DeleteAll(node.Rchild);err != nil {
			return err
		}

	}
	node = nil
	return nil
}
//查找并返回节点
func Update(node *types.Node, Key string, Value string) (*types.Node, error) {
	for {
		if node == nil {
			return nil, errNotExist
		}
		if Key == node.Key { //查找到Index节点
			node.Value = Value
			return node, nil
		} else if Key > node.Key {
			node = node.Rchild
		} else {
			node = node.Lchild
		}
	}
}

func InDecr(node *types.Node, Key string, Step string)(*types.Node, error){
	for {
		if node == nil {
			return nil, errNotExist
		}
		if Key == node.Key { //查找到Index节点
			f, err := strconv.ParseInt(node.Value, 10, 64)
			if err != nil {
				return nil, err
			} else {
				tmp, _ := strconv.ParseInt(Step, 10, 64)
				f += tmp
				node.Value = strconv.FormatInt(f, 10)
			}
			return node, nil
		} else if Key > node.Key {
			node = node.Rchild
		} else {
			node = node.Lchild
		}
	}
}
//查找并返回节点
func Get(node *types.Node, Index string) (*types.Node, error) {
	for {
		if node == nil {
			return nil, errNotExist
		}
		if Index == node.Key { //查找到Index节点
			return node, nil
		} else if Index > node.Key {
			node = node.Rchild
		} else {
			node = node.Lchild
		}
	}
}

// 深度优先搜索二叉树
func GetAll(node *types.Node, treeIndex int64, head **types.DataNode, tail **types.DataNode) error {
	if nil == node {
		return nil
	}

	newNode := new(types.DataNode)
	newNode.Key = node.Key
	newNode.Value = node.Value
	newNode.Index = treeIndex
	newNode.Next = nil

	if nil == *head {
		*head = newNode
		*tail = newNode
	} else {
		for nil != (*tail).Next {
			(*tail) = (*tail).Next
		}
		(*tail).Next = newNode
		(*tail) = newNode

	}
	if nil != node.Lchild {
		GetAll(node.Lchild, treeIndex, head, tail)
	}

	if nil != node.Rchild {
		GetAll(node.Rchild, treeIndex, head, tail)
	}
	return nil
}
