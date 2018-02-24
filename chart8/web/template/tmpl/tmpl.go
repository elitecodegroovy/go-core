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

var templates map[string]*Tmpl
var bufpool *bpool.BufferPool
var mainTmpl = `{{define "main" }} {{ template "base" . }} {{ end }}`

// create a buffer pool
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


// FuncMap is a convenience type that mirrors the FuncMap type in html/template
type FuncMap template.FuncMap

// HTML is another convenience type that mirrors the HTML type in html/template
// (http://golang.org/src/html/template/content.go?h=HTML#L120)
type HTML string

// AssetFunc is the function that go-bindata generates to look up a file
// by name
type AssetFunc func(string) ([]byte, error)

// Must is a helper that wraps a call to a function returning
// (*Template, error) and panics if the error is non-nil. It is intended for
// use in variable initializations such as
//	var t = template.Must(template.New("name").Parse("templates/my.tmpl"))
func Must(t *Tmpl, err error) *Tmpl {
	if err != nil {
		panic(fmt.Sprintf("template error: %s", err))
	}
	if t == nil {
		panic(fmt.Sprintf("template was nil"))
	}
	return t
}

// Template is a wrapper around a Template (from html/template). It reads
// template file contents from a function instead of the filesystem.
type Tmpl struct {
	AssetFunc AssetFunc
	tmpl      *template.Template
}

// New creates a new Template with the given name. It stores
// the given Asset() function for use later.
// Example usage:
//  tmpl := template.New("mytmpl", Asset) //Asset is the function that go-bindata generated for you
//
func New(name string, fn AssetFunc) *Tmpl {
	return &Tmpl{fn, template.New(name)}
}

// Name gets the name that was passed in the New function
func (t *Tmpl) Name() string {
	return t.tmpl.Name()
}

// Funcs is a proxy to the underlying template's Funcs function
func (t *Tmpl) Funcs(funcMap FuncMap) *Tmpl {
	return t.replaceTmpl(t.tmpl.Funcs(template.FuncMap(funcMap)))
}

//Delims is a proxy to the underlying template's Delims function
func (t *Tmpl) Delims(left, right string) *Tmpl {
	return t.replaceTmpl(t.tmpl.Delims(left, right))
}

// Parse looks up the filename in the underlying Asset store,
// then calls the underlying template's Parse function with the result.
// returns an error if the file wasn't found or the Parse call failed
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

// ParseFiles looks up all of the filenames in the underlying Asset store,
// concatenates the file contents together, then calls the underlying template's
// Parse function with the result. returns an error if any of the files
// don't exist or the underlying Parse call failed.
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

// Execute is a proxy to the underlying template's Execute function
func (t *Tmpl) Execute(w io.Writer, data interface{}) error {
	return t.tmpl.Execute(w, data)
}

// ExecuteTemplate is a proxy to the underlying template's ExecuteTemplate function
func (t *Tmpl) ExecuteTemplate(wr io.Writer, name string, data interface{}) error {
	return t.tmpl.ExecuteTemplate(wr, name, data)
}

// replaceTmpl is a convenience function to replace t.tmpl with the given tmpl
func (t *Tmpl) replaceTmpl(tmpl *template.Template) *Tmpl {
	t.tmpl = tmpl
	return t
}

// file is a convenience function to look up fileName using t.AssetFunc, then
// return the contents or an error if the file doesn't exist
func (t *Tmpl) file(fileName string) ([]byte, error) {
	tmplBytes, err := t.AssetFunc(fileName)
	if err != nil {
		return nil, err
	}
	return tmplBytes, nil
}