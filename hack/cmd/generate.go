package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"

	templates "leetcoding/hack/template"
	"leetcoding/utils"

	"github.com/ymichaelson/klog"
)

var (
	configPath string
)

type TemplateConfig struct {
	PackageName string  `json:"packageName"`
	FuncName    string  `json:"funcName"`
	Args        string  `json:"args"`
	Result      string  `json:"result"`
	Return      string  `json:"return"`
	Subject     Subject `json:"subject"`
}

type Subject struct {
	Order    int      `json:"order"`
	Name     string   `json:"name"`
	Topic    string   `json:"topic"`
	Examples []string `json:"examples"`
}

func main() {
	utils.InitKlog()

	// init path
	path, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		klog.Error(err)
		os.Exit(1)
	}

	configFilePath := fmt.Sprintf("%s/../config/config.json", path)
	if len(configPath) > 0 {
		configFilePath = configPath
	}

	file, err := os.Open(configFilePath)
	if err != nil {
		klog.Error(err)
		os.Exit(1)
	}

	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		klog.Error(err)
		os.Exit(1)
	}

	var templateConfig TemplateConfig
	err = json.Unmarshal(content, &templateConfig)
	if err != nil {
		klog.Error(err)
		os.Exit(1)
	}

	dirPath := fmt.Sprintf("%s/../../pkg/%s", path, templateConfig.PackageName)
	codingFilePath := fmt.Sprintf("%s/%s.go", dirPath, templateConfig.PackageName)
	testFilePath := fmt.Sprintf("%s/%s_test.go", dirPath, templateConfig.PackageName)

	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		klog.Error(err)
		os.Exit(1)
	}

	codingFile, err := os.Create(codingFilePath)
	if err != nil {
		klog.Error(err)
		os.Exit(1)
	}

	defer codingFile.Close()

	testFile, err := os.Create(testFilePath)
	if err != nil {
		klog.Error(err)
		os.Exit(1)
	}

	defer testFile.Close()

	t2, err := codingTemplate(&templateConfig)
	if err != nil {
		klog.Error(err)
		os.Exit(1)
	}
	_, err = codingFile.Write([]byte(t2))
	if err != nil {
		klog.Error(err)
		os.Exit(1)
	}

	t1, err := testTemplate(&templateConfig)
	if err != nil {
		klog.Error(err)
		os.Exit(1)
	}
	_, err = testFile.Write([]byte(t1))
	if err != nil {
		klog.Error(err)
		os.Exit(1)
	}
}

func testTemplate(temp *TemplateConfig) (string, error) {
	var w bytes.Buffer

	t := template.Must(template.New("test").Funcs(template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
	}).Parse(templates.TestTemplate))
	err := t.Execute(&w, temp)
	if err != nil {
		return "", err
	}
	return w.String(), nil
}

func codingTemplate(temp *TemplateConfig) (string, error) {
	var w bytes.Buffer

	t := template.Must(template.New("conding").Funcs(template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
	}).Parse(templates.CodingTemplate))
	err := t.Execute(&w, temp)
	if err != nil {
		return "", err
	}
	return w.String(), nil
}

func init() {
	flag.StringVar(&configPath, "config", "", "your config.json path")
}
