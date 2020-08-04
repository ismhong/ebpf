// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// ebpf/bin/probe.o
package main

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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _probeO = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x93\xbd\x6e\xdb\x30\x14\x85\x0f\x65\xd7\x76\xbb\xb5\x93\x21\x78\xf0\xd0\xa1\x5d\x64\xb8\x4f\xe0\xa5\x2d\xd0\x7a\x28\xd0\x07\x10\x68\x49\x48\x14\xeb\x27\x11\x09\xc3\x82\x87\x24\x43\x00\x8f\x59\xf2\x00\x79\x0a\x8f\x59\xf3\x18\x1e\x32\x78\x74\xa6\x64\x0a\x03\x2a\xd4\x4f\xe8\x28\xc9\x05\xec\xcb\xfb\x49\xe7\xde\x43\x82\x3a\xfe\x39\xfe\x65\x10\x82\x3c\x08\xee\x50\x56\x95\x30\xca\xe5\x48\xfd\x7f\x04\xc1\x8a\x00\xfc\xfb\x27\x38\xe6\xbd\x90\xb4\x4b\x80\x84\x3a\x9e\x5c\x1f\xc6\x7e\xb4\x30\xb7\x05\xf7\x22\x9e\xf1\xa4\xff\x8d\x2f\xcc\x4d\xc1\xc3\xa9\xeb\x67\x9c\xf2\xfe\xc2\x5c\x67\xfc\xea\xf2\x69\x56\x9b\x00\x6b\x21\xc4\xca\x00\x7a\x00\xce\x00\xb4\x00\xac\x94\x97\x0b\x95\x65\x0f\xa9\x97\x33\x64\x7f\xe9\x41\xce\x97\xde\x7e\xff\x1b\xe3\x41\x08\xf1\xd2\xb6\xf2\x38\x00\xd0\xc4\xed\xce\x3b\x14\xc0\x67\x7c\x28\xea\xa6\xca\x7f\x32\xde\xda\xe1\x6d\x00\x5f\xd0\x28\xf8\x32\x7f\x6e\x34\x60\x71\x6f\xce\x51\x5a\xb3\x59\xca\xec\xcc\xb0\xad\xec\x57\x1e\x0e\x58\xca\x1c\x1a\x04\x6c\xb0\xfb\x96\x3d\xf3\x12\xe6\xc7\x11\xac\x20\x98\x85\x36\x75\xdd\x84\xf9\x7b\xb0\x03\xdf\xf1\x22\xe6\x21\xa4\x7e\x64\x39\xb0\x18\x4f\x38\x9d\xc0\x62\x69\x98\xe5\x24\x76\x29\xa7\x12\x0f\xad\xe1\x6b\xa7\xf1\xfe\x38\x02\x2a\xbb\x2d\x63\xaa\x2e\xd1\xb9\xc6\xf5\xbb\x45\xd4\xaf\xa5\xf1\x51\xcd\xbc\xa6\x56\x7f\x7d\x43\xbf\xd4\x78\x47\xab\x4f\x95\xfe\x87\xc6\xaf\x55\xee\xd5\xf8\xcf\xf3\x44\xad\xf5\x33\xb8\xa9\xf1\xab\xef\xff\x6f\x8d\x7e\x53\xa3\xd7\xeb\xff\x52\x3b\x16\x31\x80\x93\x2a\xdf\x57\x83\xf2\xbe\x9d\x9a\xf9\xe9\xf3\x4f\xbb\x88\xad\xca\xf3\x8a\xce\xa8\xf4\xe9\xaa\xfc\x18\x00\x00\xff\xff\x72\x7c\x6f\xd5\x40\x04\x00\x00")

func probeOBytes() ([]byte, error) {
	return bindataRead(
		_probeO,
		"probe.o",
	)
}

func probeO() (*asset, error) {
	bytes, err := probeOBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "probe.o", size: 1088, mode: os.FileMode(420), modTime: time.Unix(1596189328, 0)}
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
	"probe.o": probeO,
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
	"probe.o": &bintree{probeO, map[string]*bintree{}},
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
