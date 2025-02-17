// Code generated by go-bindata. (@generated) DO NOT EDIT.

//Package jsonscheme generated by go-bindata.// sources:
// ../../../../../../../../resources/component-descriptor-ocm-v3-schema.yaml
package jsonscheme

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// ModTime return file modify time
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

var _ResourcesComponentDescriptorOcmV3SchemaYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x1a\x5d\x6f\xdb\x38\xf2\xdd\xbf\x62\xb0\x29\x20\xa7\x89\xe2\x24\x45\x0f\xa8\x5f\x82\x5c\x8b\x03\x8a\xbb\xdd\x2c\xda\xde\x3d\x5c\x9a\x2b\x68\x69\x64\xb3\x4b\x91\x3e\x92\x76\xe2\xed\xf6\xbf\x1f\x48\x8a\x12\x25\x4b\xb6\x65\xbb\xbd\x3b\xec\xf6\xa1\x31\xc9\xf9\xe2\x70\x66\x38\x33\xe2\x33\x9a\x8e\x21\x9a\x69\x3d\x57\xe3\xd1\x68\x4a\x64\x8a\x1c\xe5\x45\xc2\xc4\x22\x1d\xa9\x64\x86\x39\x51\xa3\x44\xe4\x73\xc1\x91\xeb\x38\x45\x95\x48\x3a\xd7\x42\xc6\x22\xc9\xe3\xe5\x0b\xc2\xe6\x33\x72\x15\x0d\x9e\x39\xd8\x80\xd6\x67\x25\x78\xec\x66\x2f\x84\x9c\x8e\x52\x49\x32\x3d\xba\xbe\xbc\xbe\x8c\xaf\xae\x0b\xd2\xd1\xc0\x13\xa4\x82\x8f\x21\xba\x7b\xfd\x23\xbc\xf6\xcc\xe0\x4d\xc9\x0c\x96\x2f\xa0\xc2\xc8\x28\xa7\x06\x41\x8d\x07\x00\x39\x6a\x62\xfe\x02\xe8\xd5\x1c\xc7\x10\x89\xc9\x67\x4c\x74\x64\xa7\xea\xd4\xcb\x6d\xc0\x12\xa5\xa2\x82\x5b\xe4\x94\x68\xe2\xa0\x25\xfe\x7b\x41\x25\xa6\x8e\x1c\x40\x0c\x11\x27\x39\x46\xd5\xb0\xc0\x73\x33\x24\x4d\xad\x18\x84\xfd\x2c\xc5\x1c\xa5\xa6\xa8\xc6\x90\x11\xa6\xd0\xae\xcf\xab\xd9\x82\x82\xa1\xe6\x7f\x03\x3c\x93\x98\x8d\x21\x3a\x19\x05\x3b\xaa\x54\xfd\x53\xc0\xb9\x60\xbb\x05\x55\x22\x23\x4f\x98\xbe\xc7\x7c\x89\xd2\xa3\x32\x32\x41\xa6\xb6\x60\x3a\x20\x8f\x32\x97\x62\x49\x53\x94\x5b\x90\x3c\x98\x47\x4b\x24\x12\xb3\xf2\x81\x86\x9b\x74\x87\xa2\xb4\xa4\x7c\x5a\x4e\x66\x42\xe6\x44\x8f\x21\x25\x1a\x63\x4d\x73\x1c\xd8\x83\x94\x53\xec\x3c\xc9\x75\x65\x12\x36\x15\x92\xea\x59\x5e\x31\x9b\x13\xad\x51\x9a\xa3\xfe\xd7\x3d\x89\x7f\x7d\x30\xff\x5d\xc6\xaf\x46\x9f\xe2\x87\xb3\x67\xa5\x9c\x82\x67\x74\x3a\x86\x2f\xf0\x75\x87\x63\x0c\xf5\x57\x88\x45\xa4\x24\x2b\x47\x8d\x6a\xcc\x4b\x81\xba\x55\x1b\x79\x42\x9d\xdb\xdb\xc1\xf4\x08\x5b\x60\x97\x2e\xea\x86\xd5\xa2\x73\x8b\x3d\x86\x2f\x5f\xbb\x2c\x2a\x50\xdd\xf2\xfe\x32\x7e\x15\x28\x4c\xd1\x29\xa7\x7c\xda\xa4\x1f\x4d\x84\x60\x48\xb8\x07\x0b\xce\xaf\x53\x1b\x16\x66\xbb\xf7\x0c\xcc\x29\x05\x7e\x50\x53\x9b\xdb\x97\x23\x92\x93\xa7\xbf\x21\x9f\xea\xd9\x18\xae\x5f\xbe\x1c\xb4\xda\x40\xec\x8c\xe0\xe1\xf9\xf0\xfe\xe2\xa1\x31\x75\xfa\xdc\xcf\x7d\xb9\x3e\xff\x3a\x1c\xd5\x96\x3f\xb5\xa0\x7c\x32\x38\xa7\x46\x37\x03\x00\x9a\x22\xd7\x54\xaf\x6e\xb5\x96\x74\xb2\xd0\xf8\x57\x5c\x39\x51\x73\xca\x4b\xb9\xda\xa4\x32\xcc\x87\xf7\xf1\xa7\x33\x2f\x88\x9f\x3c\xbd\x71\xa4\x6b\xbe\xec\x68\x9e\x80\x26\xbf\x20\x87\x4c\x8a\x1c\x94\x5d\x30\x71\x15\x08\x4f\x81\xa4\x9f\x17\x4a\x63\x0a\x5a\x00\x61\x4c\x3c\x02\xe1\x20\xe6\x4e\xbf\xc0\x90\xa4\x94\x4f\x21\x5a\x46\xe7\x90\x93\xcf\x26\x78\x73\xb6\x3a\xb7\xa8\x76\x7c\x91\x53\x5e\xcc\x7a\x5e\x33\xaa\x20\x47\xc2\x15\xe8\x19\x42\x26\x0c\x55\x43\xc4\xa9\x5f\x01\x91\x68\x58\x19\xcb\xa2\x69\x5d\x5e\xe5\x05\xbe\xba\xb8\xbe\x78\x11\xfe\x8e\x33\x21\xce\x26\x44\x16\x73\xcb\x10\x60\xd9\x06\x71\x75\x71\xed\x7f\x95\x60\x01\x7c\xf9\xb3\x86\x16\x2a\x7b\xf9\x70\x33\xbc\xfc\xed\xfe\x2a\x7e\xf5\xf0\x31\x7d\x7e\x3a\xbc\x19\x7f\xbc\x08\x27\x4e\x6f\xda\xa7\xe2\xe1\xf0\x66\x5c\x4d\xfe\xf6\x31\xb5\x67\x74\x1b\xff\x33\x7e\x30\xfe\xe1\x7f\x7b\x92\x3b\x02\x9f\x7a\x8e\x67\xc3\x70\xe1\xcc\x12\xa9\xcd\x58\xc8\xc2\x07\x1b\x96\xdf\x66\x7a\xdb\x02\xe7\xca\xf8\x91\x32\x51\xaf\xd5\x31\xdb\x4c\x39\x82\xaf\xce\x14\xe7\x42\x51\x2d\xe4\xea\xb5\xe0\x1a\x9f\x74\x9f\x20\x66\xa0\xba\x82\x96\xa5\xd0\x0c\x2a\xc1\x1e\x45\x42\xdf\xb5\xf3\x26\x8c\xdd\x65\x15\x97\x8e\xdb\xb0\x81\x5a\xc5\xd2\xa6\x9c\x85\xac\x13\xa2\xf0\xef\x92\x45\x55\x4c\x5c\x13\xd9\xfc\x2b\xc0\xc2\xa9\xd6\xd8\xe4\x6f\x9b\x20\x8e\xfd\x48\xe6\xf3\x5a\x24\xdd\x88\x0a\x80\x7c\x91\x8f\xe1\x3e\x5a\x48\xf6\x33\xd1\xb3\xe8\x1c\x22\x35\x23\xd7\x2f\xff\x14\xa7\x74\x8a\x4a\x47\x0f\x83\x06\x9d\xbe\x94\xad\x8e\xa7\x54\x69\xb9\x32\xd4\xef\x5e\xbf\x2d\x87\x0f\xe6\x0c\x48\x92\xa0\x52\x3b\xe6\x57\x46\x33\x16\xca\x5c\xf0\x05\x2a\x2a\x18\x9a\x11\x3e\x69\xe4\xe6\xca\x51\xa7\x5b\x8c\x65\x00\x30\xa5\x7a\xb6\x98\xdc\x6e\xe6\xbd\xd1\xda\xec\xd0\x98\x40\x70\xa0\x76\x26\xdb\xcb\x1a\x9b\x6a\x73\x02\x96\xea\x2f\x18\x6d\x41\x37\x56\xba\x19\x22\x11\x79\x4e\xf5\x26\x9f\xe0\x82\xe3\x21\x7a\x39\x70\xdf\x3f\x09\x8e\xce\x30\x94\x58\xc8\x04\xdf\x94\x0e\xd7\x43\x1c\x93\xac\x94\x83\x22\x11\x29\xc7\x86\x42\x39\x70\x26\xd4\x23\xe7\x59\x13\x7c\xf7\x60\x57\xa0\xe0\x93\x96\xe4\x6d\x01\xb0\x25\x9b\x59\xa3\x73\x84\x7c\x7d\x87\xe3\xd8\x23\xa5\x0f\xdd\xd8\x8e\xf9\xea\x2e\xab\x87\xbf\x56\x2a\x0e\x2f\xda\x0e\x18\x7a\xec\x0e\xe0\xa6\x46\xf4\xc0\x03\x00\x17\xcd\xde\xcf\x31\xe9\x61\x46\x33\xa2\x66\xb7\xbe\x06\xa8\x8c\xcb\x94\x16\x8c\x2a\x5b\x8a\xac\x2f\xdb\x3c\x78\x87\xb4\xbf\xcd\xe0\x6a\x0c\x37\x66\xdb\xed\x42\xec\x90\xa0\xb7\x43\x0c\x5c\x0e\x4e\xf4\x42\x62\x4f\x25\x91\x0d\x1a\x30\xa3\x1c\x53\x4a\x3e\x78\x9f\xeb\xaf\x93\x96\x22\xac\xe7\xe6\xdc\x54\x29\x47\x05\x55\xbf\x5b\x3e\xcc\xd0\x01\xb9\x0b\x46\x64\x36\x2d\x2d\xd5\x02\x41\x79\xb4\x51\x7f\xfb\xc6\x29\x67\xa2\xe5\xb0\xa4\xb7\xa7\xde\xb6\x16\x6c\x8e\xdf\x16\x27\xaf\xfc\x26\xac\xd5\x82\x7d\x76\x62\xd6\xec\xc9\xdd\x2c\x68\x0a\x88\x37\xfb\x78\x62\x97\xa2\xf6\xa9\x53\x9b\x91\xb3\x05\xe6\x28\x21\xba\xb7\x7e\x4b\x15\x95\xfd\x29\xa7\xab\x3e\x17\xf1\xc6\x9b\xaf\x61\x61\x12\x8b\x2b\xd6\x71\x39\x8a\x9d\x7d\x97\x8e\xd3\x9e\x96\xdb\xd8\xef\xda\x1d\x18\x74\x5d\xa0\xd9\x79\xe9\x64\xd3\x34\x6a\x7b\x8c\x4a\x26\xef\x7c\x2a\xb6\x35\xa7\x25\x26\x6d\x43\x89\x3c\x41\x5b\x5c\xc3\xb0\x6a\x86\x32\x91\x10\x76\x5a\xa4\x42\x5d\xf9\x95\xb7\xc0\xf7\xc8\x30\xd1\x62\x5b\x57\xad\xd3\x60\x7b\xdd\xfa\xb6\x6c\x2b\xc4\xde\x77\xa3\xe5\x3e\x77\xed\x50\xb5\x1a\xd2\xe1\x4d\xd3\x96\x4e\x50\x3f\x5b\xde\x94\x28\xc2\x09\x90\x44\x2f\x08\x63\xab\x71\xc5\x29\xb6\x77\xcc\xe3\x08\xd4\x1c\x13\x4a\x98\x31\x4e\x2d\x69\x62\x99\xfc\xff\xe6\x96\x7b\x24\x8e\x4d\x67\x16\x1c\x9b\x89\x63\xa1\x50\xbe\x60\x6c\x87\xcc\xaf\x11\x52\xbd\xd7\x57\x57\x7f\xcf\x2a\xd3\x13\x50\x7d\x5b\xf8\x70\x62\xf1\xad\x0f\x57\x54\xce\x8b\xc6\xd7\x42\x69\xc8\x89\x4e\x66\x81\x1b\xa8\xb5\x90\xbd\x5e\x70\x32\x9b\xf2\x05\x53\x61\x06\xfd\x47\x0d\x53\xee\xca\xc5\xe0\x23\x45\x79\x47\xac\xba\x48\xdc\x21\xec\x5c\xd4\x5a\x13\x88\xce\x21\xc2\x27\x8d\x92\x13\x56\xd6\xf5\xff\x8b\x95\x96\x48\xe8\x9f\x99\xd8\xbd\xd4\xb2\xbb\xfb\x0b\x65\xa8\x56\x4a\x63\xde\x1f\xf7\xae\x8d\xe1\xb7\x8e\x0b\x22\xa1\x6f\x73\x32\x3d\xa8\xd7\x61\x87\xd4\x50\x79\xe7\x6f\xb6\xa3\x34\x41\xc2\x9e\x99\xb7\x94\x3a\x9b\x2d\x5d\xcd\x4a\x9d\x07\x6c\x8c\x91\x95\xf7\xb8\xc3\xf6\x03\x51\x21\x52\x04\x55\x3f\x2b\xeb\xaa\xc3\x6e\xcd\x06\xea\xa9\x82\x29\xc4\x72\xc2\x69\x86\x4a\x37\x2b\xb0\x06\xd3\x3d\xcb\x3c\xa7\x19\x17\x9a\x9d\xa3\x38\x09\x14\x68\xb1\x85\x63\xd3\x50\xd7\xd9\x39\x08\xcf\x4a\x13\x39\x45\x8d\x29\x24\x82\xeb\x32\xf9\xe9\x24\xaf\xe8\xaf\x1b\xf7\x62\xd6\x81\x72\x98\xac\x34\x2a\xcf\x63\x62\x94\xdd\xa4\xcb\x17\xf9\xc4\x1c\xe8\x00\xa0\xd3\x65\x0f\x30\x97\x8c\x32\xac\x6e\xc2\x43\x2d\xa6\x45\xc2\xca\x7a\x3c\xab\x2e\xbd\xf8\xf5\x50\x1d\xa0\x67\x44\x03\x55\x76\xef\x46\xfd\x94\xdb\xb5\x1f\xcc\xa2\xfa\x01\x52\x2a\x6d\xf6\xbc\xea\x3c\x0f\xaf\xb7\xbb\x23\xf9\xd7\x37\x50\xd8\x5d\xd3\xcf\x36\x1b\x67\xdd\x30\xad\xbf\xc3\x23\xd5\xb3\x42\x35\xc9\x42\x4a\xe4\xba\x4a\x50\xa0\x7a\x9c\xb1\x49\x4b\x3e\xb4\xbe\x2b\x72\x9e\x3e\x3a\xda\xf4\x1c\xa2\x4d\x89\x7f\x64\x3f\xdb\xef\x12\x7b\x18\xc7\x4c\x39\xba\xd2\x86\xe0\x42\xfd\x3e\xd7\xf8\x00\xa0\x6a\xf4\x1e\xe0\x8a\x0b\xff\x0d\xe7\xc0\x8b\xdb\x08\x53\x2a\x7a\xb1\xe1\x7b\xcd\x00\x60\x8a\x1c\x25\x4d\xfe\x8b\xdf\x5a\x0a\x09\xdc\xe7\x96\x62\xf0\xbd\x7d\xf6\x38\x8d\xcd\xdf\x99\x4f\x57\x07\xe7\xe6\xbf\x95\x4b\xd7\x4c\xf4\x7b\x25\xe6\xf5\xc7\x61\x7d\x2d\xf0\x9b\xd8\x53\xdf\xce\x98\xda\xd4\xdd\xae\x5f\xc1\xb6\xff\x93\xd1\xc4\x16\x94\xfe\x26\x2e\x32\x43\x33\x0c\xba\x64\xde\xbc\xf4\xbe\x3b\x2d\x3a\x10\x47\x2a\x89\x1b\x9f\x67\x83\x6f\xd0\x2e\x71\x3f\x12\x1f\x59\xaf\xac\xaa\x86\x4e\x7f\xfa\x6b\x95\xf2\x86\xa7\x1d\x55\xd3\x28\xda\x05\xa1\x99\xf2\xec\x84\xd4\x08\xb9\xd1\x60\xd0\x30\x97\xd0\xd2\x4d\xdc\x9c\xd3\x7f\x54\xb1\x35\x86\xe8\x17\xca\xd3\xe2\x67\xf8\xce\x34\x76\x66\x15\x0d\xea\x26\x50\xa1\xd7\x6c\x33\x34\xf5\xa0\x60\xcb\x2f\x1a\x4f\x75\xcb\x97\xb8\xe7\x6e\x59\x89\x4c\x3f\x12\x89\xd5\x82\xcd\x3a\x8d\x4c\x9d\xf4\x13\xc1\x95\x1e\x43\x54\x7e\xe1\x08\xf6\xe3\x77\xe0\x90\x3b\x1e\xf7\xb9\x0d\xae\xbd\xbc\xd9\xed\xf9\x64\xe3\xfc\xbb\x8f\x72\xed\x51\x50\x04\x27\x3e\x1b\x66\xab\x73\x78\x44\x10\x9c\xad\x8a\x87\x70\xb6\x68\x14\x1c\x6b\x8e\xdf\xee\x33\xc5\xd7\x88\xf2\xdb\xd8\x01\xcf\x3e\x4b\x1a\x51\xe3\xd3\xda\x01\x34\xdb\x3f\x3f\x45\xff\x09\x00\x00\xff\xff\x74\x46\x86\xd2\xc2\x2d\x00\x00")

func ResourcesComponentDescriptorOcmV3SchemaYamlBytes() ([]byte, error) {
	return bindataRead(
		_ResourcesComponentDescriptorOcmV3SchemaYaml,
		"../../../../../../../../resources/component-descriptor-ocm-v3-schema.yaml",
	)
}

func ResourcesComponentDescriptorOcmV3SchemaYaml() (*asset, error) {
	bytes, err := ResourcesComponentDescriptorOcmV3SchemaYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../../../../../../../../resources/component-descriptor-ocm-v3-schema.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"../../../../../../../../resources/component-descriptor-ocm-v3-schema.yaml": ResourcesComponentDescriptorOcmV3SchemaYaml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//
//	data/
//	  foo.txt
//	  img/
//	    a.png
//	    b.png
//
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
	"..": {nil, map[string]*bintree{
		"..": {nil, map[string]*bintree{
			"..": {nil, map[string]*bintree{
				"..": {nil, map[string]*bintree{
					"..": {nil, map[string]*bintree{
						"..": {nil, map[string]*bintree{
							"..": {nil, map[string]*bintree{
								"..": {nil, map[string]*bintree{
									"resources": {nil, map[string]*bintree{
										"component-descriptor-ocm-v3-schema.yaml": {ResourcesComponentDescriptorOcmV3SchemaYaml, map[string]*bintree{}},
									}},
								}},
							}},
						}},
					}},
				}},
			}},
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
