package tmpl_test

import (
	"testing"
	"github.com/elitecodegroovy/go-core/chart8/web/template/tmpl"
	"testing/quick"
	"errors"
)

func TestA(t *testing.T)  {
	quick.Check(func(name string, bytes []byte) bool {
		assertFunc := func(name string) ([]byte, error) {
			return bytes, nil
		}
		tmpl := tmpl.New(name, assertFunc)
		return tmpl.Name() == name
	}, nil)
}

func createValidAssetFunc(name string, bytes []byte, notFound error) tmpl.AssetFunc {
	return tmpl.AssetFunc(func(n string) ([]byte, error) {
		if n == name {
			return bytes, nil
		}
		return []byte{}, notFound
	})
}

func TestParse(t *testing.T) {
	tmplText := `
    <html>
      <head>
        <title>hello {{.name}}</title>
      </head>
      <body>
        {{.greeting}}
      </body>
    </html>
  `
	fileName := "mytmpl.tmpl"
	tmplBytes := []byte(tmplText)
	expectedErr := errors.New("template not found")
	assetFunc := createValidAssetFunc(fileName, tmplBytes, expectedErr)

	tmpl1, err1 := tmpl.New("test", assetFunc).Parse(fileName)
	t.Log(tmpl1)
	t.Log(err1)
}

