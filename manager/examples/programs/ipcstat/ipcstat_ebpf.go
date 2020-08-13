// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// ebpf_prog/ipcstat.elf
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

var _ipcstatElf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x52\xb1\x6e\xd4\x40\x10\x7d\xbe\xb3\xe3\x03\x22\x01\x91\x40\x81\x2a\x05\x48\x88\xc2\x81\x7c\x41\x74\x12\xd0\xb8\x38\x4a\x0a\x64\x39\x9b\x05\x2c\xf9\x7c\x96\x77\x2f\xc2\xa2\x08\x0d\x82\x8e\x0a\x3a\x1a\x68\x52\xd2\x25\xdd\xb5\x7c\x02\x25\x9f\x10\x2a\x28\x22\x19\xed\x7a\x8c\x97\x3d\x9b\x84\x92\x29\x3c\xf7\xde\xec\x9b\x37\x3b\x7b\xfb\xf7\xc2\xfb\x03\xc7\x41\x13\x0e\x7e\xa0\x45\x6d\x7c\x19\xb6\xbf\xb7\xe9\x7b\x01\x0e\x16\x57\x6b\xee\x90\x44\xec\xfa\x49\xa5\xf2\xe2\x63\x8d\xfd\x01\x70\x52\x55\xd5\xba\xd5\xf4\x95\xf6\x02\xae\x60\x55\xe3\x92\xea\xe5\xce\x81\xce\x17\xd7\x6a\xfc\xe2\xf6\x71\x67\xbf\xc5\x27\xc2\x43\xe0\xb8\xa3\xff\xa1\xdb\xfa\x0c\x14\x26\xfe\x1d\xe5\xff\x75\x6e\xc5\xa9\x23\x23\x9a\xa3\x2f\xce\x7a\xee\xc1\x24\xfc\x4b\x15\x98\xeb\x3e\xdf\x2b\x9b\x7f\xaf\xbf\x43\x1c\x58\xfc\x5b\xfd\xf5\x96\xf8\x09\x80\xcb\x7a\x9a\x3a\xe8\x9a\xb8\xa6\x79\x1f\xb7\x08\x37\x79\x95\x78\x58\xfc\x33\x00\x6b\xf0\x7e\xf3\x47\x94\x37\x35\x3f\x5c\xe2\x1b\x9d\xda\xc1\x0a\x80\x37\x16\x36\xeb\x9e\x55\x57\x78\xe4\xaf\x78\x2e\x02\xc9\x9f\x4b\x4c\xe3\x5c\x20\xc9\x84\x2c\xe6\x4c\x26\xb3\x4c\x80\x95\x2c\xe5\x02\x41\xc1\xd3\x9c\x17\x4f\x22\xbe\xc7\x33\xb9\x39\xcb\x22\x96\xcf\x23\xe3\x24\x82\x34\xdd\x9b\x46\xf1\xee\x6e\x21\x92\xa7\x88\xd2\x84\xf1\x4c\xf0\x1e\xa1\xee\x8a\x24\x67\x42\xc6\x32\x60\x08\x84\x2c\x64\xbc\x83\x40\x94\x53\x95\xc3\xf1\xf8\x6e\xb4\xa5\xd2\x9d\x68\xab\xff\xe9\xfe\x29\xf6\xf5\x6b\x2e\xc7\x4d\x22\x3f\x58\xbc\xfd\x9f\x72\x8c\x9d\x9a\xb1\xdd\xe3\xe7\x5a\xf8\xc6\x29\xfa\x23\x8b\x1f\x59\x78\x03\xc0\xb9\x0e\x9f\x6f\x83\xb6\xae\xe2\x3c\xdd\xb3\xd1\x5f\xa2\xfc\xb8\xc7\xff\xe7\x19\xfd\x1f\xf5\xf8\x37\x4b\x35\xfd\xbd\x0e\x7f\x9f\xfc\xed\x37\xf8\x4c\x8b\x9e\x9c\xe2\xff\xb0\x47\x0f\xba\xbf\xbd\x6f\xfb\xfd\xc6\x4a\x1b\x56\x33\x00\x2f\x4d\x7e\x83\x1a\x7a\xc6\xfc\x5d\xfa\xd7\xad\xd5\x9f\x73\x12\xf9\xd5\xd0\xb9\xc6\xfc\xeb\x94\x7f\x05\x00\x00\xff\xff\x5a\x41\x68\x71\x88\x06\x00\x00")

func ipcstatElfBytes() ([]byte, error) {
	return bindataRead(
		_ipcstatElf,
		"ipcstat.elf",
	)
}

func ipcstatElf() (*asset, error) {
	bytes, err := ipcstatElfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ipcstat.elf", size: 1672, mode: os.FileMode(420), modTime: time.Unix(1597301297, 0)}
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
	"ipcstat.elf": ipcstatElf,
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
	"ipcstat.elf": &bintree{ipcstatElf, map[string]*bintree{}},
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
