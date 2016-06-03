// Code generated by go-bindata.
// sources:
// in/a/test.asset
// in/b/test.asset
// in/c/test.asset
// in/test.asset
// DO NOT EDIT!

package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)
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

var _inATestAsset = []byte(`// sample file
`)

func inATestAssetBytes() ([]byte, error) {
	return _inATestAsset, nil
}

func inATestAsset() (*asset, error) {
	bytes, err := inATestAssetBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "in/a/test.asset", size: 15, mode: os.FileMode(436), modTime: time.Unix(1464924728, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _inBTestAsset = []byte(`// sample file
`)

func inBTestAssetBytes() ([]byte, error) {
	return _inBTestAsset, nil
}

func inBTestAsset() (*asset, error) {
	bytes, err := inBTestAssetBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "in/b/test.asset", size: 15, mode: os.FileMode(436), modTime: time.Unix(1464924728, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _inCTestAsset = []byte(`// sample file
`)

func inCTestAssetBytes() ([]byte, error) {
	return _inCTestAsset, nil
}

func inCTestAsset() (*asset, error) {
	bytes, err := inCTestAssetBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "in/c/test.asset", size: 15, mode: os.FileMode(436), modTime: time.Unix(1464924728, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _inTestAsset = []byte(`// sample file
`)

func inTestAssetBytes() ([]byte, error) {
	return _inTestAsset, nil
}

func inTestAsset() (*asset, error) {
	bytes, err := inTestAssetBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "in/test.asset", size: 15, mode: os.FileMode(436), modTime: time.Unix(1464924728, 0)}
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
	"in/a/test.asset": inATestAsset,
	"in/b/test.asset": inBTestAsset,
	"in/c/test.asset": inCTestAsset,
	"in/test.asset": inTestAsset,
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
	"in": &bintree{nil, map[string]*bintree{
		"a": &bintree{nil, map[string]*bintree{
			"test.asset": &bintree{inATestAsset, map[string]*bintree{}},
		}},
		"b": &bintree{nil, map[string]*bintree{
			"test.asset": &bintree{inBTestAsset, map[string]*bintree{}},
		}},
		"c": &bintree{nil, map[string]*bintree{
			"test.asset": &bintree{inCTestAsset, map[string]*bintree{}},
		}},
		"test.asset": &bintree{inTestAsset, map[string]*bintree{}},
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
	return save(_filePath(dir, name), data, info.Mode(), info.ModTime())
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

// save saves the given data to the file at filename. If an existing file at
// that filename already exists, this simply chmods the existing file to match
// the given fileMode and otherwise leaves it alone.
func save(filename string, data []byte, fileMode os.FileMode, modTime time.Time) error {
	path := filepath.Dir(filename)
	if err := os.MkdirAll(path, os.FileMode(0755)); err != nil {
		return err
	}

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_EXCL, fileMode)
	if err != nil {
		if !os.IsExist(err) {
			return err
		}

		if dataMatches(filename, data) {
			err2 := chmod(filename, fileMode)
			if err2 != nil {
				return err2
			}
			return os.Chtimes(filename, modTime, modTime)
		}

		file, err = openAndTruncate(filename, fileMode, true)
		if err != nil {
			return err
		}
	}

	if _, err = file.Write(data); err != nil {
		return err
	}
	if err := file.Sync(); err != nil {
		return err
	}
	if err := file.Close(); err != nil {
		return err
	}

	return os.Chtimes(filename, modTime, modTime)
}

func openAndTruncate(filename string, fileMode os.FileMode, removeIfNecessary bool) (*os.File, error) {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, fileMode)
	if err != nil && os.IsPermission(err) && removeIfNecessary {
		if err = os.Remove(filename); err != nil {
			return nil, err
		}
		return openAndTruncate(filename, fileMode, false)
	}

	return file, err
}

// dataMatches compares the file at filename byte for byte with the given data
func dataMatches(filename string, data []byte) bool {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0)
	if err != nil {
		return false
	}
	fileInfo, err := file.Stat()
	if err != nil {
		return false
	}
	if fileInfo.Size() != int64(len(data)) {
		return false
	}
	b := make([]byte, 65536)
	i := 0
	for {
		n, err := file.Read(b)
		if err != nil && err != io.EOF {
			return false
		}
		for j := 0; j < n; j++ {
			if b[j] != data[i] {
				return false
			}
			i = i + 1
		}
		if err == io.EOF {
			break
		}
	}
	return true
}

func chmod(filename string, fileMode os.FileMode) error {
	fi, err := os.Stat(filename)
	if err != nil || fi.Mode() != fileMode {
		return os.Chmod(filename, fileMode)
	}
	return err
}

