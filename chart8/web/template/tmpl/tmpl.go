package tmpl

import (
	"fmt"
	"github.com/oxtoacart/bpool"
	"html/template"
	"log"
	"net/http"
	"io"
	"strings"
)

//声明全局模板变量
var templates map[string]*Tmpl
//声明缓存池类型变量
var bufpool *bpool.BufferPool

//声明主模板
var mainTmpl = `{{define "main" }} {{ template "base" . }} {{ end }}`

// 初始化缓存池
func init() {
	bufpool = bpool.NewBufferPool(64)
	log.Println("缓存bufpool分配成功")
}

type TemplateError struct {
	s string
}

func (e *TemplateError) Error() string {
	return e.s
}

func NewError(text string) error {
	return &TemplateError{text}
}

func LoadTemplates() (err error) {
	if templates == nil {
		templates = make(map[string]*Tmpl)
	}

	mainTemplate := template.New("main")
	mainTemplate, err = mainTemplate.Parse(mainTmpl)
	if err != nil {
		log.Fatal(err)
	}
	fileNames := AssetNames()
	var layoutFiles []string
	for _, fileName := range fileNames {
		if strings.Contains(fileName, "layouts") {
			layoutFiles = append(layoutFiles, fileName)
		}
	}

	for _, fileName := range fileNames {
		if !strings.Contains(fileName, "layouts") {
			files := append(layoutFiles, fileName)
			mainTemplateClone, err := mainTemplate.Clone()
			if err != nil {
				return err
			}
			templates[fileName] = &Tmpl{Asset, mainTemplateClone}
			templates[fileName] = Must(templates[fileName].ParseFiles(files...))
		}
	}

	log.Println("模板加载成功")
	return nil
}

func RenderTemplate(w http.ResponseWriter, name string, data interface{}) error {
	tmpl, ok := templates[name]
	if !ok {
		http.Error(w, fmt.Sprintf("模板%s不存在.", name),
			http.StatusInternalServerError)
		err := NewError("模板不存在")
		return err
	}

	buf := bufpool.Get()
	defer bufpool.Put(buf)

	err := tmpl.Execute(buf, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		err := NewError("模板执行异常")
		return err
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	buf.WriteTo(w)
	return nil
}


// FuncMap 是一个map类型
type FuncMap template.FuncMap

// HTML 字符类型
type HTML string

//声明函数类型，便于通过名称查询到对应文件内容的切片字节数组
type AssetFunc func(string) ([]byte, error)

// 与template.Must(template.New("name").Parse("templates/my.tmpl"))使用类似，
// 只是输入参数为自定义类型Tmpl，返回类型也是Tmpl指针。
func Must(t *Tmpl, err error) *Tmpl {
	if err != nil {
		panic(fmt.Sprintf("template error: %s", err))
	}
	if t == nil {
		panic(fmt.Sprintf("template was nil"))
	}
	return t
}

// 自定义结构Tmpl，是对template的二次组合封装
type Tmpl struct {
	AssetFunc AssetFunc
	tmpl      *template.Template
}

// 创建一个Tmpl结构对象
//  tmpl := tmpl.New("mytmpl", Asset) //Asset is the function that go-bindata generated for you
//
func New(name string, fn AssetFunc) *Tmpl {
	return &Tmpl{fn, template.New(name)}
}

// 获取模板的名称
func (t *Tmpl) Name() string {
	return t.tmpl.Name()
}

// 代理内部的Funcs模板函数
func (t *Tmpl) Funcs(funcMap FuncMap) *Tmpl {
	return t.replaceTmpl(t.tmpl.Funcs(template.FuncMap(funcMap)))
}

//代理内部的Delims模版函数
func (t *Tmpl) Delims(left, right string) *Tmpl {
	return t.replaceTmpl(t.tmpl.Delims(left, right))
}

// 查找文件名对应的资源，如果没有发现或者分析失败，返回的错误信息error不为nil
func (t *Tmpl) Parse(filename string) (*Tmpl, error) {
	tmplBytes, err := t.file(filename)
	if err != nil {
		return nil, err
	}
	newTmpl, err := t.tmpl.Parse(string(tmplBytes))
	if err != nil {
		return nil, err
	}
	return t.replaceTmpl(newTmpl), nil
}

// 类似于template包中的ParseFiles，区别在于这个方法是从文件tmpl_data.go中
//读取数据
func (t *Tmpl) ParseFiles(filenames ...string) (*Tmpl, error) {
	fileBytes := []byte{}
	for _, filename := range filenames {
		tmplBytes, err := t.file(filename)
		if err != nil {
			return nil, err
		}
		fileBytes = append(fileBytes, tmplBytes...)
	}
	newTmpl, err := t.tmpl.Parse(string(fileBytes))
	if err != nil {
		return nil, err
	}
	return t.replaceTmpl(newTmpl), nil
}

// 代理执行模板函数Execute
func (t *Tmpl) Execute(w io.Writer, data interface{}) error {
	return t.tmpl.Execute(w, data)
}

// 代理执行模板函数ExecuteTemplate
func (t *Tmpl) ExecuteTemplate(wr io.Writer, name string, data interface{}) error {
	return t.tmpl.ExecuteTemplate(wr, name, data)
}

func (t *Tmpl) replaceTmpl(tmpl *template.Template) *Tmpl {
	t.tmpl = tmpl
	return t
}

// 通过文件名查看文件的数据
func (t *Tmpl) file(fileName string) ([]byte, error) {
	tmplBytes, err := t.AssetFunc(fileName)
	if err != nil {
		return nil, err
	}
	return tmplBytes, nil
}