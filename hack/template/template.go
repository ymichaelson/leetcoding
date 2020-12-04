package template

const (
	TestTemplate = `package {{.PackageName}}

import (
	"leetcoding/utils"
	"testing"

	"github.com/ymichaelson/klog"
)

var (
	input1  = 
	expect1 = 

	input2  = 
	expect2 = 
)

func Test{{.FuncName}}(t *testing.T) {
	utils.InitKlog()

	output1 := {{.FuncName}}(input1)
	if output1 != expect1 {
		t.Errorf("test failed, expect %v, but got %v", expect1, output1)
		return
	}

	output2 := {{.FuncName}}(input2)
	if output2 != expect2 {
		t.Errorf("test failed, expect %v, but got %v", expect2, output2)
		return
	}

	klog.Infof("input: %v; expect: %v; output: %v", input1, expect1, output1)
	klog.Infof("input: %v; expect: %v; output: %v", input2, expect2, output2)
}`

	CodingTemplate = `package {{.PackageName}}

/*
{{.Subject.Order}}.{{.Subject.Name}}
	题目：
		{{.Subject.Topic}}
	{{range $i, $v := .Subject.Examples}}
	示例{{add $i 1}}:
		{{$v}}
	{{end}}
	解题思路：
		
*/

func {{.FuncName}}({{.Args}}) {{.Result}} {
	return {{.Return}}
}
`
)
