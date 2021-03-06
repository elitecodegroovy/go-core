// Code generated by go-bindata.
// sources:
// templates/aboutme.tmpl
// templates/index.tmpl
// templates/layouts/base.tmpl
// templates/layouts/js.tmpl
// templates/layouts/style.tmpl
// templates/skills.tmpl
// DO NOT EDIT!

package tmpl

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templatesAboutmeTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\x4e\x49\x4d\xcb\xcc\x4b\x55\x50\x2a\xc9\x2c\xc9\x49\x55\xaa\xad\x7d\xd6\x31\xf1\xf9\xac\x96\xea\xea\xd4\xbc\x94\xda\x5a\x2e\x84\x7c\x72\x7e\x5e\x49\x6a\x5e\x89\x52\x6d\x2d\x97\x4d\x81\x9d\x4d\xa2\x42\x46\x51\x6a\x9a\xad\x92\xbe\x92\xdd\xcb\x65\xd3\x5e\x2e\xdc\x6a\xa3\x9f\x68\x67\xa3\x5f\x60\xc7\x65\x93\x61\x6c\xf7\x64\xef\x9c\x17\x0d\xad\x4f\xf6\x2f\x7c\xd6\xb8\xde\x46\x3f\xc3\xd8\x8e\xcb\xa6\x34\xc7\x8e\x4b\x41\x41\x41\xc1\x26\x27\xd3\xee\xf9\x94\x15\xcf\x3a\xb6\x3f\x5d\x3e\xf9\xe9\x84\xde\xf7\x7b\x66\x55\x57\x2b\xe8\xf9\x25\xe6\xa6\x2a\xd4\xd6\x2a\xd8\xe8\xe7\x64\x22\x14\x3e\xeb\x6c\x78\x3a\x67\xc5\xd3\xf9\x7d\x4f\x77\x34\xbd\xdf\x33\x4b\x01\xa4\xd2\x39\xb3\xa4\x52\xa1\xb6\x16\x55\xe1\xd3\xd9\x7b\x9f\x6f\xec\x85\x29\xf1\x4b\x2c\xc9\xcc\xcf\x4b\xcc\x41\x56\x69\xa3\x0f\x72\x01\x17\xd4\x63\x80\x00\x00\x00\xff\xff\x7a\xba\xce\x23\xf8\x00\x00\x00")

func templatesAboutmeTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesAboutmeTmpl,
		"templates/aboutme.tmpl",
	)
}

func templatesAboutmeTmpl() (*asset, error) {
	bytes, err := templatesAboutmeTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/aboutme.tmpl", size: 248, mode: os.FileMode(438), modTime: time.Unix(1519440068, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesIndexTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\x4e\x49\x4d\xcb\xcc\x4b\x55\x50\x2a\xc9\x2c\xc9\x49\x55\xaa\xad\x7d\xb9\x6c\xda\xcb\x85\x5b\xab\xab\x53\xf3\x52\x6a\x6b\xb9\x10\xf2\xc9\xf9\x79\x25\xa9\x79\x25\x4a\xb5\xb5\x5c\x0a\x0a\x0a\x0a\x36\xa5\x39\x76\x60\x06\x98\x93\x93\x69\x67\x93\xa8\x90\x51\x94\x9a\x66\xab\xa4\x5f\x9c\x9d\x99\x93\x53\xac\x64\xf7\xac\xab\xe1\x45\xf3\xde\xa7\x1d\xd3\x5f\x2c\x5c\x61\xa3\x9f\x68\x67\xa3\x9f\x93\x89\x4b\x4f\x62\x52\x7e\x69\x49\x6e\xaa\x92\xdd\xb3\x8e\x89\xcf\x67\xb5\xa0\x2a\xb7\xd1\x07\x59\x06\x75\x12\x20\x00\x00\xff\xff\xd3\x61\x50\x09\xb2\x00\x00\x00")

func templatesIndexTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesIndexTmpl,
		"templates/index.tmpl",
	)
}

func templatesIndexTmpl() (*asset, error) {
	bytes, err := templatesIndexTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/index.tmpl", size: 178, mode: os.FileMode(438), modTime: time.Unix(1519440953, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesLayoutsBaseTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x90\xc1\x0d\x83\x30\x0c\x45\xef\x99\xe2\x2b\x03\x34\x0b\x58\xec\x02\xc4\x08\xda\x90\x54\xc5\x17\x64\x65\xf7\xaa\x18\x95\x36\xb7\x58\x7a\x7e\x7a\xb1\x2a\x22\x4f\x4b\x66\xf8\xa1\xdf\xd8\xa3\x56\x47\xb3\xac\xa9\x73\x34\x73\x1f\x3b\x07\x00\x24\x8b\x24\xee\x54\x87\x54\xc6\x07\xfc\x31\x7a\xdc\x6a\x85\x2a\xe7\x58\x2b\x05\x43\x0e\xfc\xcb\x6d\xb2\x37\x9c\xa3\x60\x5a\x1a\x4a\xdc\x4f\xfb\xf5\xb4\x6d\xe1\xf5\x99\x7a\x61\xf8\xb1\x64\xe1\x2c\x87\xc2\xd8\xf0\xb3\x37\x95\x22\xfc\xba\xb2\x6c\x6e\xba\x4e\xe8\x3f\xec\xbe\xb5\x55\xa6\xa5\x60\x7f\x77\xaa\xe0\x1c\x3f\xd7\x78\x07\x00\x00\xff\xff\x4c\xb8\x05\x15\x22\x01\x00\x00")

func templatesLayoutsBaseTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesLayoutsBaseTmpl,
		"templates/layouts/base.tmpl",
	)
}

func templatesLayoutsBaseTmpl() (*asset, error) {
	bytes, err := templatesLayoutsBaseTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/layouts/base.tmpl", size: 290, mode: os.FileMode(438), modTime: time.Unix(1474253314, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesLayoutsJsTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x24\xca\x41\x0a\xc2\x40\x0c\x05\xd0\x7d\x4e\xf1\xc9\x4a\x37\x7a\x00\x6b\xef\x12\xda\x28\x1d\xe2\xcc\x60\x06\x41\x42\xee\xde\xc5\x6c\x1f\x2f\x02\xbb\xbe\x8e\xaa\xe0\xe2\x8c\x4c\x5a\x7c\xfb\x1e\x7d\x60\xfc\xbb\x3e\xb9\xc8\x4f\x26\xf0\x4a\x00\xb0\xb5\xea\xcd\xf4\x66\xed\x7d\x61\x81\x7f\xc4\x0c\xc5\xf9\xfa\xa0\xe5\x3e\xe7\x4a\x11\xd0\xba\x23\xf3\x0c\x00\x00\xff\xff\x98\x26\x12\xf6\x5f\x00\x00\x00")

func templatesLayoutsJsTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesLayoutsJsTmpl,
		"templates/layouts/js.tmpl",
	)
}

func templatesLayoutsJsTmpl() (*asset, error) {
	bytes, err := templatesLayoutsJsTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/layouts/js.tmpl", size: 95, mode: os.FileMode(438), modTime: time.Unix(1474253314, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesLayoutsStyleTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\xcc\x4b\x0e\x83\x30\x0c\x04\xd0\xbd\x4f\x61\xb1\x47\xfd\x2c\x01\xf5\x2e\x81\x18\x88\x9a\xda\x55\x9a\x7e\x22\xcb\x77\xaf\x88\xb2\x1b\xcd\x3c\x8d\xaa\xa7\x35\x30\x61\xf7\xca\x25\x52\x67\x06\x53\x4d\x37\x98\xc5\x17\x54\x40\x44\x9c\xdd\x72\xdf\x92\xbc\xd9\xf7\x8b\x44\x49\x03\x7e\xf7\x90\x69\x04\x03\xd8\x2f\x0d\xb5\x85\xdd\xa7\x8c\xb5\x78\xb8\xb4\x05\xee\x23\xad\x79\xc0\xeb\xf9\xf9\x3b\xfc\x74\x6a\xf7\xaa\xc4\xde\xec\x1f\x00\x00\xff\xff\x00\x68\x12\xa3\x81\x00\x00\x00")

func templatesLayoutsStyleTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesLayoutsStyleTmpl,
		"templates/layouts/style.tmpl",
	)
}

func templatesLayoutsStyleTmpl() (*asset, error) {
	bytes, err := templatesLayoutsStyleTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/layouts/style.tmpl", size: 129, mode: os.FileMode(438), modTime: time.Unix(1519625293, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesSkillsTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x90\x41\x4a\x33\x41\x10\x85\xf7\x7d\x8a\xa2\x0f\x90\x5e\x64\x5b\xa9\x13\xfc\x97\xe8\xa4\x2b\xd3\x81\xa1\x13\x26\xfd\x0b\x52\x14\x44\xdc\x04\x0f\xa0\xeb\xd9\x68\x40\x24\x2b\x37\x82\x78\x19\x13\x27\xb7\x10\x33\x13\x07\x09\xba\xab\x57\xf5\xf1\xea\xf1\x44\x02\x4f\x67\x89\xc1\xe6\x59\x2e\xd9\xaa\xee\x6f\x56\xcd\xf5\xeb\xfb\x5b\xbd\xbf\xda\x8a\x70\x0a\xaa\xa6\xa7\x26\xf3\x94\x39\x65\xab\x6a\x70\x41\x80\x1e\x62\xc5\xd3\x91\x75\x96\x0e\xf7\xb7\x87\xfa\x19\x9d\x27\x74\x0b\x32\x18\x87\xd4\x7a\xed\xd6\x77\x4d\xbd\x41\x17\x87\x64\x30\xfb\x71\xc9\x30\x29\xfd\x72\x39\xb2\x47\x61\xc9\x00\x00\x60\x8e\xec\x43\x3b\xb7\xba\xea\x45\x07\x50\xb3\x7d\x6a\x36\x2b\x74\x39\x9e\xdf\x3e\x5e\x1e\x76\xeb\xc7\x9f\x37\x74\x27\x97\xaf\xfd\xb7\x3f\xe6\xf1\x3c\x5c\xb6\xb3\x08\x54\x3e\x15\x0c\x03\x50\xfd\xeb\x7b\x20\x11\x18\xfc\xf3\xa9\xf8\xef\x0b\x06\x55\x74\x39\xfc\x02\xf1\x05\x97\x67\x44\x1f\x46\x04\x38\x85\xd3\x3f\x74\x5d\x1c\x74\xc7\x42\xc8\x98\xae\xf8\xcf\x00\x00\x00\xff\xff\xa0\x87\x96\xd6\x9e\x01\x00\x00")

func templatesSkillsTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesSkillsTmpl,
		"templates/skills.tmpl",
	)
}

func templatesSkillsTmpl() (*asset, error) {
	bytes, err := templatesSkillsTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/skills.tmpl", size: 414, mode: os.FileMode(438), modTime: time.Unix(1519626443, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"templates/aboutme.tmpl": templatesAboutmeTmpl,
	"templates/index.tmpl": templatesIndexTmpl,
	"templates/layouts/base.tmpl": templatesLayoutsBaseTmpl,
	"templates/layouts/js.tmpl": templatesLayoutsJsTmpl,
	"templates/layouts/style.tmpl": templatesLayoutsStyleTmpl,
	"templates/skills.tmpl": templatesSkillsTmpl,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"templates": &bintree{nil, map[string]*bintree{
		"aboutme.tmpl": &bintree{templatesAboutmeTmpl, map[string]*bintree{}},
		"index.tmpl": &bintree{templatesIndexTmpl, map[string]*bintree{}},
		"layouts": &bintree{nil, map[string]*bintree{
			"base.tmpl": &bintree{templatesLayoutsBaseTmpl, map[string]*bintree{}},
			"js.tmpl": &bintree{templatesLayoutsJsTmpl, map[string]*bintree{}},
			"style.tmpl": &bintree{templatesLayoutsStyleTmpl, map[string]*bintree{}},
		}},
		"skills.tmpl": &bintree{templatesSkillsTmpl, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

