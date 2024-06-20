package fs

import (
	"fmt"
	"strings"
)

type fileSystemNode struct {
	name        string
	isFile      bool
	children    *fileSystemNode
	content     string
	nextBrother *fileSystemNode
}

func (this *fileSystemNode) addChildren(name string, node *fileSystemNode) {

	if this.children == nil {
		this.children = node
		return
	}

	// insert the new node
	if this.children.name > name {
		node.nextBrother = this.children
		this.children = node
		return
	}
	root := this.children
	for ; root.nextBrother != nil; root = root.nextBrother {
		if root.nextBrother.name > name {
			node.nextBrother = root.nextBrother
			root.nextBrother = node
			return
		}
	}
	root.nextBrother = node
}

type FileSystem struct {
	nodes map[string]*fileSystemNode
}

func Constructor() FileSystem {
	nodeMap := make(map[string]*fileSystemNode)
	nodeMap["/"] = &fileSystemNode{name: "/"}
	return FileSystem{
		nodes: nodeMap,
	}
}

func (this *FileSystem) Ls(path string) []string {
	names := []string{}
	node, exists := this.nodes[path]
	if exists {
		if node.isFile {
			names = append(names, node.name)
		} else {
			root := node.children
			for root != nil {
				names = append(names, root.name)
				root = root.nextBrother
			}
		}
	}

	return names
}

func (this *FileSystem) Mkdir(path string) {
	this.makesureDir(path)
}

func (this *FileSystem) AddContentToFile(filePath string, content string) {
	fnode, exists := this.nodes[filePath]
	if !exists {
		fmt.Println("flag", filePath)
		fnode = this.createFile(filePath)
	}
	fnode.content += content
}

func (this *FileSystem) ReadContentFromFile(filePath string) string {
	node, exists := this.nodes[filePath]
	if exists {
		return node.content
	}

	return ""
}

func (this *FileSystem) createFile(filePath string) *fileSystemNode { // consume the input is valid
	var i int
	for i = len(filePath) - 1; i > 0; i-- {
		c := filePath[i]
		if c == '/' {
			break
		}
	}
	path, name := filePath[:i], filePath[i+1:]
	if path == "" {
		path = "/"
	}
	fs := this.makesureDir(path)
	fmt.Println("flag2", fs.name, name)
	newNode := &fileSystemNode{
		name:   name,
		isFile: true,
	}
	// add to the folder children
	fs.addChildren(name, newNode)
	// add to the nodes map
	this.nodes[filePath] = newNode

	return newNode
}

// make sure the given dir path exists
// if not create it recursively
func (this *FileSystem) makesureDir(path string) *fileSystemNode { // assume are the inputs are correct

	fs, exists := this.nodes[path]
	if exists { // dir exits
		return fs
	}

	nodeNames := strings.Split(strings.Trim(path, "/"), "/")
	parentPath := ""
	parentNode := this.nodes["/"]

	for _, name := range nodeNames {
		node, exists := this.nodes[parentPath+"/"+name]
		if !exists {
			node = &fileSystemNode{
				name: name,
			}
			parentNode.addChildren(name, node)
			this.nodes[parentPath+"/"+name] = node

		}
		parentPath += ("/" + name)
		parentNode = node
	}

	return parentNode
}
