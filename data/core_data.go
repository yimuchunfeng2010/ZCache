package zdata

import (
	"ZCache/global"
	"ZCache/tool"
	"ZCache/tool/logrus"
	"ZCache/types"
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func CoreAdd(key string, value string) (*types.Node, error) {
	hashIndex, err := tool.GetHashIndex(key)
	if err != nil {
		return nil, err
	}
	if nil == global.GlobalVar.GRoot {
		return nil, errors.New("GRoot nil")
	}
	tmpNode, err := Add(global.GlobalVar.GRoot[hashIndex], key, value)
	if err != nil {
		return nil, err
	}

	global.GlobalVar.GRoot[hashIndex] = tmpNode

	// 修改全局信息
	global.GlobalVar.GCoreInfo.KeyNum += 1
	return global.GlobalVar.GRoot[hashIndex], nil
}

func CoreDelete(key string) (*types.Node, error) {
	hashIndex, err := tool.GetHashIndex(key)
	if err != nil {
		return nil, err
	}
	if nil == global.GlobalVar.GRoot {
		return nil, errors.New("GRoot nil")
	}

	tmpNode, err := Delete(global.GlobalVar.GRoot[hashIndex], key)
	if err != nil {
		return nil, err
	}
	global.GlobalVar.GRoot[hashIndex] = tmpNode

	// 修改全局信息
	global.GlobalVar.GCoreInfo.KeyNum -= 1
	return global.GlobalVar.GRoot[hashIndex], nil
}

//查找并返回节点
func CoreUpdate(key string, Value string) (*types.Node, error) {
	hashIndex, err := tool.GetHashIndex(key)
	if err != nil {
		return nil, err
	}

	if nil == global.GlobalVar.GRoot {
		return nil, errors.New("GRoot nil")
	}

	tmpNode, err := Update(global.GlobalVar.GRoot[hashIndex], key, Value)
	if err != nil {
		return nil, err
	}
	global.GlobalVar.GRoot[hashIndex] = tmpNode

	return global.GlobalVar.GRoot[hashIndex], nil
}


func CoreInDecr(key string , step string) (*types.Node, error) {
	hashIndex, err := tool.GetHashIndex(key)
	if err != nil {
		return nil, err
	}

	if nil == global.GlobalVar.GRoot {
		return nil, errors.New("GRoot nil")
	}

	tmpNode, err := InDecr(global.GlobalVar.GRoot[hashIndex], key, step)
	if err != nil {
		return nil, err
	}
	global.GlobalVar.GRoot[hashIndex] = tmpNode

	return global.GlobalVar.GRoot[hashIndex], nil
}


//查找并返回节点
func CoreGet(key string) (*types.Node, error) {
	hashIndex, err := tool.GetHashIndex(key)
	if err != nil {
		return nil, err
	}
	if nil == global.GlobalVar.GRoot {
		return nil, errors.New("GRoot nil")
	}

	node, err := Get(global.GlobalVar.GRoot[hashIndex], key)
	if err != nil {
		logrus.Warningf("%s  CoreGet Failed[Key:%s,Err:%s]", tool.GetFileNameLine(), key, err.Error())
		return nil, err
	}
	return node, nil
}

func CoreGetAll() (*types.DataNode, error) {
	if nil == global.GlobalVar.GRoot {
		return nil, errors.New("GRoot nil")
	}

	var index int64
	var head *types.DataNode = nil
	var tail *types.DataNode = nil
	for index = 0; index < global.Config.MaxLen; index++ {
		err := GetAll(global.GlobalVar.GRoot[index], index, &head, &tail)
		if err != nil {
			return nil, err
		}
	}
	return head, nil
}

func CoreFlush() error {
	if nil == global.GlobalVar.GRoot {
		return errors.New("GRoot nil")
	}

	var index int64
	var rspRoot *types.DataNode = nil
	for index = 0; index < global.Config.MaxLen; index++ {
		err := GetAll(global.GlobalVar.GRoot[index], index, &rspRoot, &rspRoot)
		if err != nil {
			return err
		}
	}
	// 写文件
	file, err := os.OpenFile(tool.GetDataLogFileName(), os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	fileWrite := bufio.NewWriter(file)
	curNode := rspRoot
	for nil != curNode {
		msg := fmt.Sprintf("%s  %s\n", curNode.Key, curNode.Value)
		fileWrite.WriteString(msg)
		curNode = curNode.Next

	}
	fileWrite.Flush()

	return nil
}

func CoreImport() error {
	file, err := tool.GetNewestFile(tool.GetDataLogDir())
	if err != nil {
		return err
	}

	fi, err := os.Open(fmt.Sprintf("%s%s", tool.GetDataLogDir(), file))
	if err != nil {
		logrus.Warningf("%s  Open File Failed! [Err:%s]\n", tool.GetFileNameLine(), err.Error())
		return err
	}
	defer fi.Close()
	br := bufio.NewReader(fi)

	for {
		data, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}

		array := strings.Split(string(data), "  ")
		if len(array) != 2 {
			logrus.Warningf("%s  Invaild Data! [Data: %s]\n", tool.GetFileNameLine(), string(data))
			continue
		}
		key := array[0]
		value := array[1]
		_, err := CoreAdd(key, value)
		if err != nil {
			logrus.Warningf("%s  CoreAdd Data[Key: %s, Value: %s]\n", tool.GetFileNameLine(), key, value)
			continue
		}

	}
	return nil
}

func CoreDeleteAll() (error) {

	if nil == global.GlobalVar.GRoot {
		return errors.New("GRoot nil")
	}

	var index int64
	for index = 0; index < global.Config.MaxLen; index++ {
		err := DeleteAll(global.GlobalVar.GRoot[index])
		if err != nil {
			return err
		}
	}

	for index = 0; index < global.Config.MaxLen; index++ {
		global.GlobalVar.GRoot[index] = nil
	}
	// 修改全局信息
	global.GlobalVar.GCoreInfo.KeyNum = 0
	return nil
}

func CoreExpension(size int64)(error){
	if nil == global.GlobalVar.GRoot {
		return errors.New("GRoot nil")
	}

	newSize := global.Config.MaxLen + size
	if newSize <= 0{
		return errors.New("Tree size less than 0")
	}
	// 获取全部节点
	var index int64
	var head *types.DataNode = nil
	var tail *types.DataNode = nil
	for index = 0; index < global.Config.MaxLen; index++ {
		err := GetAll(global.GlobalVar.GRoot[index], index, &head, &tail)
		if err != nil {
			return err
		}
	}

	// 解析全部节点，并添加到新树上
	global.GlobalVar.GRootTmp = make([]*types.Node,newSize)
	curNode := head
	for curNode != nil {
		expensionAdd(curNode.Key,curNode.Value, newSize)
		curNode = curNode.Next
	}

	// 更新根节点
	global.GlobalVar.GRoot = global.GlobalVar.GRootTmp
	global.GlobalVar.GRootTmp = nil
	global.Config.MaxLen = newSize
	return nil
}

func expensionAdd(key string, value string, newSize int64) (*types.Node, error) {
	hashIndex, err := tool.GetHashIndex(key, newSize)
	if err != nil {
		return nil, err
	}
	if nil == global.GlobalVar.GRootTmp {
		return nil, errors.New("GRoot nil")
	}
	tmpNode, err := Add(global.GlobalVar.GRootTmp[hashIndex], key, value)
	if err != nil {
		return nil, err
	}

	global.GlobalVar.GRootTmp[hashIndex] = tmpNode

	// 修改全局信息
	return global.GlobalVar.GRootTmp[hashIndex], nil
}

func CoreGetKeyNum() (int, error) {
	return global.GlobalVar.GCoreInfo.KeyNum, nil
}
