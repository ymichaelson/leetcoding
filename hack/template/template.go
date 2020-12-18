package template

const (
	TestTemplate = `package {{.PackageName}}
{{ $funcName := .FuncName }}
import (
	"leetcoding/utils"
	"testing"

	"github.com/hex-techs/klog"
)

var (
	{{range $i, $v := .Subject.Examples}}
	input{{add $i 1}} = 
	expect{{add $i 1}} = 
	{{end}}
)

func Test{{.FuncName}}(t *testing.T) {
	utils.InitKlog()
	{{range $i, $v := .Subject.Examples}}
	output{{add $i 1}} := {{$funcName}}(input{{add $i 1}})
	if output{{add $i 1}} != expect{{add $i 1}} {
		t.Errorf("test failed, expect %v, but got %v", expect{{add $i 1}}, output{{add $i 1}})
		return
	}
	klog.Infof("input: %v; expect: %v; output: %v", input{{add $i 1}}, expect{{add $i 1}}, output{{add $i 1}})
	{{end}}
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
