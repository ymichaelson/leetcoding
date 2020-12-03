package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"

	templates "leetcoding/hack/template"
	"leetcoding/utils"

	"github.com/ymichaelson/klog"
)

type TemplateConfig struct {
	PackageName string  `json:"packageName"`
	FuncName    string  `json:"funcName"`
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

	file, err := os.Open("../config/config.json")
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

	dir := fmt.Sprintf("../../pkg/%s", templateConfig.PackageName)

	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		klog.Error(err)
		os.Exit(1)
	}

	codingFile, err := os.Create(fmt.Sprintf("%s/%s.go", dir, templateConfig.PackageName))
	if err != nil {
		klog.Error(err)
		os.Exit(1)
	}

	defer codingFile.Close()

	testFile, err := os.Create(fmt.Sprintf("%s/%s_test.go", dir, templateConfig.PackageName))
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

	t := template.Must(template.New("test").Parse(templates.TestTemplate))
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
