// Code generated by go-bindata. DO NOT EDIT.
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
	info  fileInfoEx
}

type fileInfoEx interface {
	os.FileInfo
	MD5Checksum() string
}

type bindataFileInfo struct {
	name        string
	size        int64
	mode        os.FileMode
	modTime     time.Time
	md5checksum string
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
func (fi bindataFileInfo) MD5Checksum() string {
	return fi.md5checksum
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _bindataProbeo = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x56\x3f\x68\xd4\x50\x18\xff\xd2\xeb\xf5\xce\x56\xe4\xaa\x20\xd7\x70" +
	"\x43\x44\x84\x16\x21\xb5\xa2\x22\x82\x58\x2a\xd6\xa5\x83\x48\x87\x6e\x47\x7a\x4d\x6d\xc8\xe5\x0f\x49\xac\xb6\x29" +
	"\x28\x88\x83\x9b\x8b\x83\x2e\xe2\x9f\xcd\x41\xb7\xe2\x74\x37\x76\xec\x22\x38\xba\x08\x45\x17\x75\xb1\x93\x91\xef" +
	"\xe5\x25\xf9\xfa\x2e\xb1\xd7\x4d\xc1\x0f\xae\x79\xdf\xef\xbd\xef\x7d\xdf\xf7\x7b\xbf\xbc\xf4\xde\xb5\xb9\xd9\x01" +
	"\x49\x82\xc4\x24\xf8\x09\x99\x97\xd9\x74\x95\x8c\xf9\xdf\x51\x90\xa0\x73\x3c\xc6\xea\xa5\xbd\xeb\xc3\x4b\xbb\x11" +
	"\x3e\x37\x25\x80\xa5\x61\x00\x53\xde\x61\x7e\x5d\x02\x08\x34\x3b\xc0\xb1\x72\x59\x39\x15\xca\x9f\x52\x5c\xb1\xd6" +
	"\x9a\x38\x6e\x39\xb6\x1f\xca\x1f\x53\xdc\x75\x0c\x1b\xc7\xc1\x84\xb2\x11\xca\xdb\x29\xbe\xec\x2b\x2b\x38\x76\x1c" +
	"\x53\x09\xe5\xad\x14\xb7\xcc\x25\x03\xc7\x9e\x32\xbe\x1a\xca\xdd\xb4\x0e\x34\x5f\xfe\xc2\xfc\xce\x8b\xd8\xaf\x48" +
	"\x00\xdd\x28\x8a\x36\x07\x00\x4e\x03\xc0\x43\x00\x18\xe2\x4f\x6c\xf9\x16\xd6\x09\x00\x77\xf8\xb3\xf3\x9a\xc7\x0d" +
	"\x02\xec\x46\x51\xd4\xd1\x78\xff\x03\x7b\xfb\xef\x70\x3e\x36\xcb\xf1\x3e\xb8\xdf\x18\xfa\x7c\xfe\x49\xb2\xee\x2f" +
	"\xe4\xcf\xb3\xfe\x1d\xfe\xb0\xae\x09\xc6\xcf\xb7\xb4\x7e\xc7\x54\x5c\xd6\x97\x61\x07\x21\xe1\xcd\x6f\x69\x6d\x1c" +
	"\xb7\x95\x15\x87\xf2\x16\xdc\xf6\x18\x3f\xca\xb8\xbf\x46\x79\xcb\x74\xe4\xe9\x09\x6f\x19\x0f\x3f\x7a\x78\xd8\xe6" +
	"\x3c\x9c\x24\x3c\x88\xf5\x0e\xc2\x7f\xfb\x93\x21\xe7\xa8\x3b\xd4\x26\xea\x12\xb5\x8b\xba\x45\x6d\xa3\xae\x51\xfb" +
	"\xa8\x7b\x7c\x27\xd0\x50\xab\x07\x59\x1f\xef\xef\xe9\x78\xe6\x78\xde\xa8\x89\x36\x0d\x1d\x86\xeb\x37\xe6\x00\x7e" +
	"\x45\x51\xf4\xfc\xab\x04\x75\xa1\x3e\x76\xfc\x74\x62\x90\xfc\x50\xa8\xb5\x18\xae\x25\xeb\xd7\x6f\x42\x75\x63\x44" +
	"\x3a\x8c\x9a\xe2\xbf\xc4\xa6\xc9\x45\xdb\x00\x80\x73\x05\x73\xc9\xfc\x3c\xf1\xdf\xf5\xc9\xe7\x7b\x56\xda\xf7\x28" +
	"\x6f\xae\x04\xa5\xdc\x98\x12\x94\x0b\xf0\x4a\x0f\xf6\x12\x00\x46\x61\x38\xf5\x13\x7d\x3f\x65\xf8\x48\x0f\x8e\x3d" +
	"\x1c\x25\x79\x93\x3e\x55\x86\x97\x7b\xf0\x75\x86\x67\x79\x93\xbe\x6b\x84\x63\x6a\x27\x58\xde\xec\x63\xd5\xe0\xfb" +
	"\x24\x08\xba\x58\x6d\x55\xca\xfc\x91\x3e\xe6\x1b\x64\x1e\xaf\xaa\x8b\xc4\xc7\x6e\x16\x88\xcf\x7a\x55\x03\xfd\x6e" +
	"\x00\xea\xcc\xfc\xac\x8a\x03\x22\x47\xb0\x34\xd7\x9f\x24\x80\x0f\xa6\xeb\x39\x8b\x7a\x73\x75\xd9\x6f\x32\x41\x83" +
	"\xea\xe9\xed\x18\x9c\xcc\x40\xb2\x8a\xc9\x58\x5c\x15\x83\xa6\xa7\x07\x2e\xae\x23\x2e\x5b\x12\xfb\xcd\x55\xdd\xf3" +
	"\x0d\xc7\x86\x66\xdb\x68\xe9\xb6\xaf\xb3\x5d\x54\x7d\xa5\xb9\xec\x69\x96\x0e\x96\x66\xd8\x6a\x0b\x54\x3f\xf0\x02" +
	"\x6d\x11\x54\x7f\xcd\x62\xcf\x99\xf9\x59\x50\x3d\x67\x49\x0b\x34\x9c\x9b\x52\xa7\x2e\xe4\xd0\x7f\x60\xeb\x72\xfe" +
	"\x44\x5b\xe0\x47\xfe\x59\xc0\xc5\xff\x4d\x24\xfe\x1b\x12\xf0\xe9\x82\x7c\xe2\xfd\x6b\xec\x13\x2f\xbe\x87\x55\x61" +
	"\x9d\x0e\x00\x87\x72\xf2\xec\xf0\x0d\x15\xee\xd7\x78\x9f\x55\xe2\xa3\x5d\x2d\xc8\x7f\x5f\xea\x2f\xff\x95\x82\xfc" +
	"\xd5\x4a\x6f\xfe\x72\x4e\xfe\x47\x05\xf9\xbb\xfc\x7b\x2c\xde\x33\x62\xfe\x06\x79\x07\xa8\xb9\x1c\x68\x08\xf5\x8b" +
	"\xfc\x7f\xe0\xf1\x67\x05\xfc\x31\x5f\xf8\x40\xc0\xd3\x7b\x95\x3f\x5f\x15\xe4\x3f\x56\xce\xcf\x27\xea\xe7\x59\x41" +
	"\x7c\xbd\x20\x5e\xf4\xb7\x73\xf6\x44\x6b\xf0\xf8\xb1\x7d\xf2\x57\x0a\xe2\xcf\xf3\x78\x65\x9f\xf8\xb7\xe4\x4e\xa2" +
	"\xb6\xc0\xe3\x5d\x01\x17\xcf\xef\x4d\x81\x7e\xc6\xb9\x7e\xce\x70\x1f\x79\x3f\x92\xa3\x9f\xad\x9c\xdc\x0c\x4f\xea" +
	"\x27\xf7\x28\xd5\x5f\xf2\x1d\xfc\x1d\x00\x00\xff\xff\x58\x41\x60\x89\x80\x0c\x00\x00")

func bindataProbeoBytes() ([]byte, error) {
	return bindataRead(
		_bindataProbeo,
		"/probe.o",
	)
}



func bindataProbeo() (*asset, error) {
	bytes, err := bindataProbeoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "/probe.o",
		size: 3200,
		md5checksum: "",
		mode: os.FileMode(436),
		modTime: time.Unix(1594298237, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}


//
// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
//
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
// nolint: deadcode
//
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

//
// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or could not be loaded.
//
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// AssetNames returns the names of the assets.
// nolint: deadcode
//
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

//
// _bindata is a table, holding each asset generator, mapped to its name.
//
var _bindata = map[string]func() (*asset, error){
	"/probe.o": bindataProbeo,
}

//
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
//
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, &os.PathError{
					Op: "open",
					Path: name,
					Err: os.ErrNotExist,
				}
			}
		}
	}
	if node.Func != nil {
		return nil, &os.PathError{
			Op: "open",
			Path: name,
			Err: os.ErrNotExist,
		}
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

var _bintree = &bintree{Func: nil, Children: map[string]*bintree{
	"": {Func: nil, Children: map[string]*bintree{
		"probe.o": {Func: bindataProbeo, Children: map[string]*bintree{}},
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
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
