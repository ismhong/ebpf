// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// ebpf_prog/bio_tracepoint.elf
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

var _bio_tracepointElf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x58\x4d\x4c\x5c\xd5\x17\x3f\xef\x3d\xe6\xe3\x4f\xf9\x0b\xa5\x41\xc9\x83\xda\x69\x2b\x09\x31\x32\x65\x06\xb4\x7e\x24\x0d\xd2\xa8\x8d\x21\x01\x9b\x60\x1b\x37\xc3\x63\x66\x24\x23\x03\xc3\xcc\x3c\x85\x61\x16\xa5\x8b\x26\xb5\x0b\x83\x0b\x13\x77\x32\x74\x21\x3b\xd9\x41\xdc\x0c\x2b\xc3\xc2\x18\x56\x86\x85\x31\x2c\x8c\xc1\xc4\x28\x26\x1a\x59\xd4\x3c\x73\xee\x3d\x6f\xee\x7b\xa7\x33\x85\xa6\xac\xbc\x8b\x9e\x39\xbf\xfe\xee\xf9\xf8\xdd\x77\xcf\x7b\xe1\xd6\x1b\x23\x6f\xea\x9a\x06\xee\xd2\xe0\x6f\x50\x9e\x5a\xfb\x6d\xea\xf7\x10\xfd\xdb\x02\x1a\x54\x9f\x96\xd8\x06\x6d\x2a\x9b\x7b\x8e\xb4\xbb\x64\x77\xc8\x6e\x0b\x5b\x4d\x4a\x5e\xc8\x00\x08\xa3\x5f\x21\x5f\x03\xd8\x75\x1c\x67\x43\x07\x68\x02\x80\x3b\x20\xad\x97\xdf\xc6\xf8\x7b\xc4\x0f\x37\xe0\x77\x33\xfe\x83\x3a\xf1\x43\x41\x80\x5e\xc6\xfb\x95\x78\x6d\x9e\x78\x2e\xdf\xaa\xc8\xbe\x3a\x34\x29\x48\xa9\x22\xfb\x1d\xd7\x74\x70\x70\x1f\xd3\x01\xf7\x05\xd0\x6f\x96\xfd\x5b\x95\x07\x3e\x5d\xaa\xab\x94\x57\x97\xfd\x57\xef\xab\xfa\xb7\x1d\xc7\xe9\x64\x87\xb1\xd1\xa4\xea\xd1\xd1\x27\xfc\x33\xb2\xb5\xf3\x08\x51\x1d\x8b\x87\x32\xdf\xe2\xc1\x91\xfa\x1f\x3c\xa6\xfe\x87\x75\xf4\x47\x3d\x23\x8c\xb7\xef\xe1\x71\x3d\x4b\x95\x43\xd2\x4f\x13\xfa\xb9\xf5\x7a\x75\x39\xa8\xa3\xc3\x1d\xf1\xac\x02\x54\xa9\xcf\x8e\xd0\x9f\xe0\xd5\xbb\x1a\x94\x78\x29\x2f\xed\x39\xd2\xe5\x72\x10\x60\xdf\x00\xa8\x4e\x4a\x7f\x5e\x97\xf5\x76\xe8\x5d\x52\x37\x43\xc6\x45\x9d\xd1\xf6\xe8\x58\x97\xd2\x7d\xaa\x49\xd6\x7d\x6b\x98\x74\xd6\x3c\x3c\xcf\x3d\x98\xd2\x00\x0c\xe4\xb5\x4b\x7f\xf4\x75\x75\x7e\x2e\xbf\x15\xfc\x71\x75\x0f\xcf\x8d\x8f\x3c\x83\xea\x12\x3c\xaa\x6f\x34\x46\xbc\xb8\xea\x43\xe0\xe7\x95\xee\x11\xa1\x45\xb7\x7c\x2e\xa8\xdf\x29\xea\xd7\xed\xfb\xbf\xd4\xaf\x46\x75\x88\xfe\xc8\xe2\x73\x67\x02\x40\x5f\xbb\x56\xab\xdf\xf4\xdc\x4f\xab\x42\xf7\xa2\x5d\xe9\xd3\x02\x00\x9d\x22\xbf\xfc\x3f\xa7\x15\x20\x41\x79\xab\x67\xa4\xbd\x61\x80\xe8\x1f\xf3\x20\xff\x86\x26\x89\x6e\x9d\x6e\x5d\x98\x1f\x57\xf9\x79\x39\x37\x92\x34\x17\x4b\xf3\x61\xc2\xb7\x09\x97\x76\x3e\x04\xd0\x8c\x79\xf2\x4a\x57\xde\x8f\xef\xbc\xb4\x06\xe7\xe5\xea\x6f\x78\x78\x4c\x57\x71\x5e\x6e\xbd\x57\xd9\x79\x69\x0d\xce\x8b\x78\x6e\x7c\xe4\x19\x9e\x3e\xa7\xdc\x73\xb9\x48\x3c\xd2\x61\x9e\xce\x67\xf4\x8c\x7a\x8e\x30\x7e\x5f\x4c\xab\xd5\x89\x7e\xf9\x55\x39\x17\xf1\xff\x31\x64\x7e\x75\x5f\xf8\x38\x07\xbe\x76\x1c\xa7\x47\x6f\x82\xd3\x9e\xbe\x72\x17\xd5\x79\x68\x11\x1d\xc6\x0d\xa9\x6b\x7e\xf5\x97\xc7\xdc\xa7\xf9\x9e\xef\xa2\xf9\xad\xd8\x5f\x34\x7f\x90\xf5\x84\x25\x5e\xbe\xb7\x55\x77\x6e\x37\x9a\x4f\x1d\x20\x05\x2f\x69\xfe\xe7\xb4\x6c\xd6\x8f\xe3\x9d\xff\x5b\xc7\x98\xff\xe5\x7b\x9b\x0f\xc5\xa9\xf7\xde\x38\xba\x9e\xfa\x71\xbc\xf5\x6c\x1e\xe7\x7d\xa4\x01\x74\x79\xee\xd7\x49\xe9\xd6\x6a\x49\x7b\x52\xba\x61\x9d\x6d\xde\xef\x93\x27\xd4\xb1\x35\xef\xd6\x77\x72\x3a\xe2\x1c\x98\x36\xd7\x45\x3c\xe4\xa7\x67\x93\x25\xfc\xdd\x77\xa5\x27\x55\x36\xd7\x6a\x78\x6a\x3a\x23\xa6\x7b\x24\x6b\xd9\x65\x73\xa5\x86\x47\x8a\x99\x45\xfc\x9d\xee\xbb\xd2\x53\x36\x3f\x57\x71\x22\xb9\x39\x15\x67\xb9\x86\x27\x92\xb9\x19\xfc\x3d\x97\x4d\xdb\x65\xf3\x6e\x0d\x9f\xcc\xe6\xc4\x5b\x7b\x3a\x51\xc8\x97\xcd\x25\x81\xe7\xef\xcb\xfb\xe1\x7d\xcf\x2f\xd1\x7b\x1e\x47\x65\xd5\x96\x78\x35\xad\xfa\x0a\x1e\xf3\xbd\x6e\xd4\xf9\xae\x41\x1a\x9e\x17\xde\xee\x9f\xba\xa0\xe1\xf2\xf2\x7e\x6f\x3b\x9a\x27\x28\xbd\x8d\x79\xd8\x3b\xf6\x8d\xda\xa0\x2e\xa8\x1d\xea\x86\xda\xa2\xae\xa8\x3d\xea\x8e\x67\x83\x6e\x33\xbc\x35\x36\xd2\x38\x1c\x7c\x22\xbe\x7b\xfe\x70\x38\xfe\x82\xd0\xc1\x80\x08\xd3\xa3\x83\xf0\x25\x86\x7f\x2c\xfc\x00\x6c\x32\xfc\x3d\xc2\xf7\x18\xfe\x1a\xe1\x9d\xba\x1f\xff\x3f\xe1\x63\x0c\xff\x19\x24\xbe\xc0\xf0\x0d\xc2\xd7\x18\x7e\x95\xe2\x6c\xb3\xbe\x3e\x20\x7c\x87\xf1\xaf\x13\x7e\xc8\xf0\xe7\x08\xef\x35\xfc\xf8\x3f\x94\x17\x9a\xfc\xf8\xf7\x84\xf7\x32\xfc\x36\xc5\xb9\xc9\xf0\xb7\x09\x5f\x62\x78\x37\xe1\x5b\x0c\xff\x8b\xe2\x87\x03\x7e\xfc\x3b\xc2\x87\x18\xfe\x15\xe1\xd7\x18\x3e\x49\xf1\xd7\x19\x3e\x40\xf8\x50\xd0\x8f\x07\x08\xdf\x67\xf8\x8f\x14\xff\x6e\xc8\x8f\x7f\x43\xf8\x4a\xd8\x8f\xff\x46\xf8\x04\x7b\x1e\xaa\x6e\x1c\x86\x2f\x00\xc0\x69\xf8\x5f\xcd\x77\xe5\x78\x1f\x00\xda\x41\x15\xbf\x4b\x79\x86\x05\xae\x0e\x6b\x99\xe2\x69\x22\x8e\x2a\xd2\xbd\x62\x3d\x84\xf7\x32\xfc\x36\xe1\x63\x0c\x9f\xf0\xc4\xbb\x80\x79\xc1\xef\x2f\x04\x94\x8f\xdf\x47\x3b\xcc\x0f\x07\x95\x8f\x9f\x22\x13\xcc\x5f\x0f\x32\x7e\xc8\xef\x8f\x85\xfc\xfc\x35\xe6\xaf\x87\xfd\xf5\x9c\x7b\x36\x72\xfe\xc2\x59\xc8\xe4\x12\x34\x14\x20\x6a\xa7\x17\x6c\x98\xb1\xe6\x8a\x10\xcd\x66\x3f\x9a\x49\x58\xa9\x54\xa1\x98\x99\x42\x0e\x0e\x11\x88\x16\xd2\x59\xbb\x60\x25\xd3\x73\xb9\xcc\xac\x7d\x49\x4c\x9c\x4b\xb5\xb9\x93\x29\x16\x3f\x3c\x8a\xe3\xce\x26\x48\x64\x33\xc9\xf4\x6c\x31\x0d\x45\xdb\x2a\xd8\x89\x42\x3a\x99\x2b\xa4\x60\x32\x93\x4b\xa8\xcd\xd1\x24\x44\x8b\x76\xc1\xb6\x26\x21\x5a\x2c\xcd\xa0\x1d\x19\x1e\x8e\x25\x06\x5e\x91\x36\x46\xf6\x65\x69\x06\xc9\x0e\x90\x8d\xbb\xf8\x65\xa2\x93\x7d\x89\x68\x64\xe3\x64\x07\x5f\x24\x9a\xb0\xfd\x89\x41\xa2\x91\x8d\x0f\x4a\x78\x80\xd8\x64\x63\x64\xe3\xc4\x26\x1b\x27\x1b\x23\x3b\x18\x23\x3c\x06\xd1\x42\x2e\x65\xd9\x16\x76\x16\x8b\x12\x3c\xd0\x4f\xf4\x7e\x78\xe2\xf5\x05\xbd\x97\xf8\x7a\xb7\x55\xda\x4f\xd9\x3d\xe2\x7f\x43\x69\x21\x8c\x5d\x67\xfa\x5b\xca\xc3\x8b\x8d\x21\x88\x1d\xb1\x7f\x99\x25\x64\x63\x00\xfa\x00\x3c\xb7\x5a\xad\x95\xa7\xa4\x75\x5f\x95\xa7\xa8\x4f\x77\xbf\x8b\x8f\x37\xc8\xbf\x43\x79\x77\xc3\x8f\xce\xff\x4e\x83\xfc\xeb\x94\x7f\xc5\x93\x3f\x50\x27\x7f\x3b\xe5\xe7\x67\xb0\xd6\x2c\xed\x02\x3c\x3a\x7f\x81\xbe\x73\xe3\x0c\x8f\x9c\x92\x96\xe3\x1a\xb3\xa5\x06\xf9\xaf\xd3\x7e\x7e\x5e\xfc\xfc\x9f\xc1\xbd\x23\x4e\x0e\x00\x96\xbc\xf8\x18\x3d\x3f\xae\xae\xa7\x1a\xec\xff\x92\xbe\x0b\xf9\xba\x49\x1b\xae\x19\x6a\xdf\x59\x4f\xff\x9d\x64\xff\x0d\x00\x00\xff\xff\x0b\x4e\x93\x15\xe8\x13\x00\x00")

func bio_tracepointElfBytes() ([]byte, error) {
	return bindataRead(
		_bio_tracepointElf,
		"bio_tracepoint.elf",
	)
}

func bio_tracepointElf() (*asset, error) {
	bytes, err := bio_tracepointElfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bio_tracepoint.elf", size: 5096, mode: os.FileMode(420), modTime: time.Unix(1596590992, 0)}
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
	"bio_tracepoint.elf": bio_tracepointElf,
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
	"bio_tracepoint.elf": &bintree{bio_tracepointElf, map[string]*bintree{}},
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
