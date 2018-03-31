// Code generated by go-bindata.
// sources:
// static/migrations/1_initial.sql
// static/scripts/passport.min.js
// DO NOT EDIT!

package static

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

var _staticMigrations1_initialSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xd0\x31\x6f\xb3\x30\x10\x06\xe0\xdd\xbf\xe2\xe4\x09\xf4\x85\xe1\x93\xaa\x2e\x4c\x26\x5c\xd2\x53\xc1\x50\x63\x57\xc9\x54\xa1\xe2\x46\xa8\x0d\x20\x07\x9a\xbf\x5f\x91\xc4\x55\x23\x71\xf3\xf3\xbe\xf6\x5d\x14\xc1\xbf\x63\x7b\x70\xf5\x68\xc1\x0c\x8c\xad\x15\x0a\x8d\x80\x3b\x8d\xb2\xa2\x42\x02\x6d\x40\x16\x1a\x70\x47\x95\xae\x80\x4f\x53\xdb\x44\xfd\xe9\x34\xf0\xf8\x17\x6b\x91\x64\x08\xfc\xa3\xed\x0e\xd6\x0d\xae\xed\x46\x0e\x01\x03\xe0\x6d\xc3\xc1\x4f\x42\xdb\x0a\x15\x89\x0c\x4a\x45\xb9\x50\x7b\x78\xc6\xfd\x6a\x56\xc7\xda\x7d\x5a\x77\x95\xc6\x50\xea\x13\xf3\xb3\xd2\x64\xd9\x05\x7d\xd7\x5f\x93\xbd\xb5\xbd\x0a\xb5\x7e\x12\x2a\x78\x7c\x08\xef\xd1\x7b\x3f\x75\xe3\xad\x2a\xa1\x2d\x49\x7d\xdf\x04\x29\x6e\x84\xc9\x34\xfc\xbf\x72\x67\xeb\xd1\x36\x6f\xf5\xc8\x41\x53\x8e\x95\x16\x79\xb9\xc4\xbb\xfe\x1c\x84\x97\xc8\x34\x34\xcb\x91\x99\xcf\xe3\x23\xfe\x4f\x46\xd2\x8b\x41\x08\xfc\x96\x2b\xbf\x4a\xc8\xc2\x98\xb1\xbf\xe7\x4f\xfb\x73\xc7\x58\xaa\x8a\x72\xe9\xa2\x31\xfb\x09\x00\x00\xff\xff\xc9\xeb\xcc\xee\xab\x01\x00\x00")

func staticMigrations1_initialSqlBytes() ([]byte, error) {
	return bindataRead(
		_staticMigrations1_initialSql,
		"static/migrations/1_initial.sql",
	)
}

func staticMigrations1_initialSql() (*asset, error) {
	bytes, err := staticMigrations1_initialSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/migrations/1_initial.sql", size: 427, mode: os.FileMode(420), modTime: time.Unix(1522409989, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _staticScriptsPassportMinJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x54\xc1\x8e\x23\x27\x10\xbd\xe7\x2b\x30\x07\x0b\xe4\x0a\xf1\xe4\x14\xb5\xc5\x5e\x36\x59\x29\x51\x92\xdd\x68\x47\xca\x99\x85\x6a\x9b\x09\x03\x2d\xa8\x1e\x8f\x85\xf9\xf7\xa8\xbb\x6d\x8f\xe7\x90\x9c\xa8\x7a\xa2\xe0\xbd\x57\x14\xab\x7e\x8c\x96\x7c\x8a\x82\x20\x02\x82\x87\x2c\x2b\x1f\x0b\xb2\x42\xd9\x5b\xe2\xbb\x17\x93\x99\xd1\xa4\xbe\x98\x52\x86\x94\xe9\x2e\x3c\x9f\xab\xc3\x6f\xe3\xbe\xcb\x6a\x5e\x1b\x24\x5d\x7b\x1f\xf7\x98\x87\xec\x23\x75\x2f\xc9\x3b\xb6\x85\x67\x24\xe3\x0c\x99\x5b\x6e\xf2\x3f\x98\x3b\x5e\xab\xfa\x25\xda\x7c\x1a\x08\xdd\x1f\x33\xd6\x1a\x6f\x60\xf5\x16\x9c\x5e\x3d\x40\xd1\xab\x87\xdd\x95\x20\x1b\x04\xc9\x6a\x96\x9b\xd6\x6b\x2f\xb2\x1a\x32\xf6\xfe\x75\x43\xb2\xdd\x36\xf5\x93\x0e\x59\x7d\x2f\x56\x6e\xbd\xb6\x1f\xb2\x0a\xfe\xd9\xd3\xbb\xfd\xdc\x66\x4f\xde\x9a\xd0\xb1\xc1\x9c\x42\x32\x8e\x1d\x4d\x61\x31\x11\xfb\x86\xac\x60\x24\x2e\xc1\x9d\xcf\xb7\x6a\x99\x91\xc6\x1c\x99\x0d\x68\xf2\xaf\x91\x30\xbf\x98\x20\x48\xc2\xac\x67\x10\x71\xc3\x99\x2f\xcc\xa5\x88\x5c\xee\xca\xf9\x2c\x8a\x5e\x6d\xc1\x6e\x36\x80\xa2\xd2\x69\xc0\x8e\x7f\xf9\xfc\xf5\x91\xc3\x98\x43\x97\x15\x46\x37\x24\x1f\x09\x66\x53\x7e\xfb\xfa\xf9\x4f\x35\xb9\x1d\xf7\xbe\x3f\x89\x24\xc1\xa6\x48\x18\xe9\x71\x2e\x34\xc3\x10\xbc\x35\x93\xba\x1f\x9e\x4a\x8a\x3b\x66\x0f\x26\x17\x24\x3d\x52\xff\xfd\x4f\x1c\x5e\x0f\xf9\x93\xc7\xe0\x4a\x57\x8f\x9e\x0e\x1f\x33\x3a\x8c\xe4\x4d\x28\xdd\x6a\xdb\xa0\x8c\xd6\x62\x29\xdd\xad\xd1\xb2\xba\x89\x9d\x51\x77\x9d\xd2\xe9\x3e\x03\xa3\xae\x2d\xd3\xe9\x16\xc2\x20\x78\xc1\xe8\x30\xb3\x83\x29\xac\x9c\xa2\x45\xc7\xcc\xd5\x43\x2e\x1b\xd8\xf4\x3c\x04\x24\xbc\xbf\x6c\x6a\x22\x2c\x1e\xcd\x65\x18\x89\x99\xc9\x6c\xdf\x5f\x64\x31\x4a\x8c\x6f\xde\x6c\x91\xad\x49\xd9\xa6\x77\x17\x60\xd4\x05\xe9\x66\xb9\x08\xfa\x01\x44\xc4\x23\x8b\x52\xed\x91\xc4\xfd\xeb\x95\xf5\x9d\x08\x4d\xf0\xc6\x5d\xc7\x26\xe1\x8e\xd4\x7f\x9f\x41\x2b\xfd\xde\x8c\xf5\x5a\xfc\xdf\xb9\x10\xf4\x76\x72\xc6\xa6\x9c\xd1\x52\x5a\xcc\x79\x36\x0e\x99\x61\x17\xd0\xa7\xc8\xa5\x84\xcd\x26\x7c\xd0\x59\xd1\x21\x63\x39\xa4\xe0\xd6\xeb\x5e\x8c\xf0\x56\xc9\x65\x93\x4d\x42\x56\x17\x44\xc2\xfe\x9d\xfa\x3b\xfe\xbd\xd8\x03\x3f\x1a\xb2\x07\x9c\xca\x20\xab\x39\x91\x4d\x1c\x7d\x74\xe9\x08\xcb\xa2\x3e\xbd\xf1\xfe\xf1\x8a\x3d\xfd\x35\x62\x3e\x29\xf3\x64\x5e\xaf\x90\x4d\xb1\xa4\x80\x2a\xa4\x3d\xd4\x6b\x1b\x2e\xf3\xb9\x24\xad\x71\x58\xc6\xa7\xe3\xc3\x65\xfc\x3b\xc6\x61\x9e\x8f\xae\x56\xf5\xfb\x14\xb4\x06\x37\x75\x13\xf8\x78\x4d\xda\xf4\x36\x66\x51\x13\xfc\x71\x09\x5b\x83\x99\xf5\x04\xfd\x3d\x05\xad\xc1\xf2\x9f\xd4\xaa\x7e\x9e\x3f\x94\xd6\xe4\xee\xbb\x7f\x03\x00\x00\xff\xff\x35\xd7\x5b\x40\xa4\x04\x00\x00")

func staticScriptsPassportMinJsBytes() ([]byte, error) {
	return bindataRead(
		_staticScriptsPassportMinJs,
		"static/scripts/passport.min.js",
	)
}

func staticScriptsPassportMinJs() (*asset, error) {
	bytes, err := staticScriptsPassportMinJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/scripts/passport.min.js", size: 1188, mode: os.FileMode(420), modTime: time.Unix(1522524892, 0)}
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
	"static/migrations/1_initial.sql": staticMigrations1_initialSql,
	"static/scripts/passport.min.js": staticScriptsPassportMinJs,
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
	"static": &bintree{nil, map[string]*bintree{
		"migrations": &bintree{nil, map[string]*bintree{
			"1_initial.sql": &bintree{staticMigrations1_initialSql, map[string]*bintree{}},
		}},
		"scripts": &bintree{nil, map[string]*bintree{
			"passport.min.js": &bintree{staticScriptsPassportMinJs, map[string]*bintree{}},
		}},
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

