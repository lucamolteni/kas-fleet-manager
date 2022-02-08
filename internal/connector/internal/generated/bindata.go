// Code generated by go-bindata. (@generated) DO NOT EDIT.

//Package generated generated by go-bindata.// sources:
// .generate/openapi/connector_mgmt.yaml
package generated

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

var _connector_mgmtYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x7d\xff\x73\xdb\x38\xae\xf8\xef\xf9\x2b\xf8\x71\x3f\x37\xb9\x7b\x2f\x76\x6c\xc7\xf9\xe6\x79\xbd\x99\x34\x49\xbb\xd9\x4d\xd3\x6e\x92\x6e\xb7\x77\x73\xe3\xd0\x12\x6c\xb3\x91\x48\x85\xa4\x92\xba\x77\xef\x7f\x7f\x43\x52\xb2\x28\x89\xb2\x65\x27\xfd\xb6\x6b\xcd\xec\x36\x96\x48\x10\x00\x01\x10\x00\x21\x8a\x45\x40\x71\x44\xfa\x68\xa7\xd5\x6e\xb5\xd1\x33\x44\x01\x7c\x24\x27\x44\x20\x2c\xd0\x88\x70\x21\x51\x40\x28\x20\xc9\x10\x0e\x02\xf6\x80\x04\x0b\x01\x9d\x9d\x9c\x0a\x75\xeb\x96\xb2\x07\xd3\x5a\x75\xa0\x28\x01\x87\x7c\xe6\xc5\x21\x50\xd9\xda\x78\x86\x8e\x82\x00\x01\xf5\x23\x46\xa8\x14\xc8\x87\x11\xa1\xe0\xa3\x09\x70\x40\x0f\x24\x08\xd0\x10\x90\x4f\x84\xc7\xee\x81\xe3\x61\x00\x68\x38\x55\x23\xa1\x58\x00\x17\x2d\x74\x36\x42\x52\xb7\x55\x03\x24\xd8\x31\x74\x0b\x10\x19\x4c\x66\x90\x37\x9e\xa1\x46\xc4\xc9\x3d\x96\xd0\xd8\x42\xd8\x57\x54\x40\xa8\x1a\xcb\x09\xa0\x86\xc7\x28\x05\x4f\x32\x3e\x08\xc7\xa1\x6c\x26\x2d\x5b\x53\x1c\x06\x0d\x34\x22\x01\x6c\x10\x3a\x62\xfd\x0d\x84\x24\x91\x01\xf4\xd1\x71\xda\x01\x5d\x01\xbf\x27\x1e\xa0\x97\x01\x80\x44\xaf\x31\xc5\x63\xe0\x1b\x08\xdd\x03\x17\x84\xd1\x3e\x6a\xb7\x3a\xad\xf6\x06\x42\x3e\x08\x8f\x93\x48\xea\x9b\x0b\xfa\x1b\x7a\x2e\x41\x48\x74\xf4\xf6\x4c\xa1\x19\xea\x07\x68\x86\xa8\x68\x6d\x08\xe0\x6a\x10\x85\x55\x13\xc5\x3c\xe8\xa3\x89\x94\x91\xe8\x6f\x6f\xe3\x88\xb4\x14\xb3\xc5\x84\x8c\x64\xcb\x63\xe1\x06\x42\x05\x04\x5e\x63\x42\xd1\x5f\x23\xce\xfc\xd8\x53\x77\xfe\x86\x0c\x38\x37\x30\x21\xf1\x18\x16\x81\xbc\x92\x78\x4c\xe8\xd8\x09\xa8\xbf\xbd\x1d\x30\x0f\x07\x13\x26\x64\xff\xa0\xdd\x6e\x97\xbb\xcf\x9e\x67\x3d\xb7\xcb\xad\xbc\x98\x73\xa0\x12\xf9\x2c\xc4\x84\x6e\x48\x3c\x4e\x18\x40\x71\x98\x9b\x97\xeb\x69\x04\xa2\xdc\xbf\xd1\x70\xb5\xae\xdd\x10\x1d\x07\xb1\x90\xb0\x44\x87\x64\x7e\x9d\xed\x37\x22\x2c\x27\x1a\xff\x67\xea\x3f\xe4\xec\xf6\x6c\x63\x03\xa1\x86\x9a\x86\xed\xbc\x98\x6e\xdf\x77\x1a\x7d\x0d\x77\x0c\xd2\xfc\x81\x50\xca\x10\x73\x35\x2b\x10\x41\x4a\x17\x39\x56\x88\x9c\xf9\x7d\xd5\xff\x37\x23\xae\xaf\x41\x62\x1f\x4b\x9c\xb4\x12\x71\x18\x62\x3e\xed\xa3\x4b\x90\x31\xa7\x42\x6b\x4b\x22\xd9\x28\xcc\xb7\xcd\x11\x57\xa7\x03\x07\x11\x31\x2a\xc0\xc2\xb7\xd1\x6d\xb7\x1b\xd9\x4f\xa4\xe4\x5d\x02\x95\xf6\x2d\x84\x70\x14\x05\xc4\xd3\xd8\x6f\x7f\x14\x8c\xe6\x9f\x22\x24\xbc\x09\x84\xb8\x78\x17\xa1\xff\xcf\x61\xd4\x47\x9b\xcf\xb6\x3d\x16\x46\x8c\x02\x95\x62\xdb\xb4\x15\xdb\x05\xfa\x37\xad\xce\x39\xc2\x7e\x2b\xd2\x32\x9b\xbc\xb2\xe8\xcd\x9b\xb9\xed\x5b\x3c\xba\xc5\x83\xec\xbe\x54\x9d\xb6\xff\x9d\xbf\x31\x20\xfe\xff\x26\xfc\x88\x30\xc7\x21\xc8\x44\xe1\xcd\xe4\x1a\x59\x2b\x75\xd9\x70\x62\x7e\x3d\x01\x44\x7c\xc4\xb4\xc9\xcc\x3a\x21\xd5\x69\xa3\x9a\x75\xea\x71\x1f\x09\xc9\x09\x1d\xcf\x6e\x13\xda\x47\x4a\x76\x67\x37\x38\xdc\xc5\x84\x83\xdf\x47\x92\xc7\x50\x5f\x28\x33\x2d\x45\x48\x80\x17\x73\x22\xa7\x76\xcb\x17\x80\x39\xf0\x3e\xfa\x27\xfa\x57\x85\xe0\xce\x60\x29\x50\x2f\xa6\x67\x27\x45\xd1\x7d\x05\x12\xe1\x02\xbd\x6a\x19\x99\xf1\x29\x2f\xb8\x0b\x9b\x7f\x23\xb1\x6d\x38\xc5\x36\x47\x7d\xa3\xd0\x15\x3e\xe1\x30\x0a\x6c\x44\xd3\x2b\xd7\xed\xd4\x34\x2b\xb7\x72\x0f\x9d\x42\xdd\x76\x01\x69\x54\xe9\xcd\x75\x49\xe6\x50\x88\xa5\x37\x51\x0b\x86\x92\x47\x25\x40\xa0\x6d\x7f\xc2\xd2\x5e\xbb\xf3\x6d\x58\x7a\xca\x39\xe3\xf5\x59\xd9\x6b\x77\x56\x65\x60\xd6\xb5\x92\x6d\x47\xb1\x9c\x20\xc9\x6e\x81\x2a\x97\x80\xd0\x7b\x1c\x58\xfa\xdd\xe8\xb5\x7b\x3f\x08\x93\x7a\xab\x33\xa9\xb7\x88\x49\x17\x2c\x93\xa5\x82\x8c\xc1\x27\x22\xa4\xc8\x18\xb6\xfb\xad\x14\x75\x49\x86\xed\xb6\xdb\xab\x32\x2c\xeb\x5a\xc9\xb0\x77\x14\x3e\x45\xe0\x49\xf0\x11\x28\xbc\x10\xf3\xb4\x5f\xe5\x2f\xbd\x60\x2d\xe3\x80\x3c\xb1\xad\x17\x55\x3e\x0a\x46\x01\x11\x52\x2d\x74\x79\x61\x10\x2e\x7b\x5f\xb7\x53\x79\xf9\x55\x28\xbb\x26\x22\x6b\xb9\x1d\xe1\xb1\x35\x09\x0b\x9b\x0b\xf2\x79\x99\xe6\x8c\xfb\xc0\x5f\x4c\x97\x19\x00\x30\xf7\x26\x8d\xef\x7e\x21\x3b\x27\x42\x56\x9b\xc4\x05\x33\xb5\x5e\x3b\xea\xad\x1d\x6b\x53\xb8\xd0\x14\x16\x1c\xfb\x25\x5d\xfa\xd4\x38\x46\x2a\xe6\x5d\x64\x1d\x1f\x61\x18\x3d\x0e\x58\x82\x8d\x25\xb2\xcd\xe2\xb1\x7e\xac\xd3\x23\x0f\x99\xca\xb8\x6c\xe1\xdc\x96\x6e\x03\xa8\x02\x81\xbb\x18\xf8\xd4\xe2\xaf\x89\x4a\xb0\x98\x52\xaf\x8a\xeb\x6f\x81\x8f\x18\x0f\xb5\xe7\x87\x75\xfe\x01\x11\x8a\x30\x35\xbd\x26\x9c\x51\x16\x0b\x14\x62\x4a\x81\x6f\xcc\x97\x36\x13\x9f\x0c\x19\x0b\x00\x53\xeb\x89\x23\x22\x41\xa9\x97\xf9\x82\xf9\x16\x83\x2b\x12\x33\x56\xa4\xea\x54\x8e\xf9\xaa\xe1\x56\x8c\x5a\x16\xf0\xd2\x20\x99\xd7\x90\x2a\xfd\x98\xf5\x32\x93\x57\xa9\x29\xf5\x3c\xf9\x1c\x90\xc6\xbc\xe8\xae\x6a\xf9\xe8\x7e\xe3\xe5\xa3\xda\x1a\x7a\x1e\x44\x12\x72\xce\xf3\x8f\x61\x00\x7b\xed\xb6\x9e\x17\xc2\xe8\xea\xab\x45\x11\x44\x25\x9f\x7e\x53\xab\x84\x6e\x69\x0c\xa2\xc8\x2c\xa2\xc5\xb9\xf5\xfa\xba\x8e\xcd\x6a\xc5\x66\xd7\x59\x6c\x0f\xbe\xb2\x19\x2c\xe6\x1e\x20\x9f\x81\xa0\x9b\xd2\xc4\x67\x6b\x9f\xa4\x20\x58\x14\xc5\x55\x6e\x89\x59\xed\xd3\xac\x49\x7e\x91\xae\x13\x85\x3d\xc2\xcf\x50\x6e\x77\x19\xce\x1f\x2c\xfa\xfa\xae\x63\xa3\x65\xe3\xa2\x75\x48\xb4\x0e\x89\xbe\x4d\x76\x48\x6c\xff\x7b\xfe\xd6\xc5\x02\x65\x24\x7e\xe3\x6b\x98\x34\x3b\xa7\x54\x34\x68\x85\x8d\x00\x97\xf9\x72\x37\xf9\x3e\x6d\x47\xcd\xcc\xfc\x3a\x29\xbf\x76\xfc\xea\x30\x69\x9d\x94\x5f\x8a\x61\x8f\x32\xbb\xa6\x69\x00\x12\xbe\xa4\x2d\x34\x23\x54\x9a\xc3\x13\xfd\x78\x91\x45\xac\x6c\xe5\x36\x8a\xdf\x8b\xa2\x38\x68\x58\x87\xbb\x7f\x58\xab\x67\x26\xf8\x11\xb6\x2f\x07\x60\x9e\x05\xd4\x5e\x51\xba\x8c\xa2\x07\x22\x27\x48\x44\xe0\x91\x11\x01\x1f\x9d\x9d\xfc\xc8\x96\xf0\x71\x4c\x2c\x02\x58\xd1\x2a\x46\x6a\x85\xf9\x92\x46\x51\x0f\x50\x69\x13\xdf\xaa\xa7\x8b\x4c\x62\x55\xa3\xc5\xb9\xe8\x13\x2c\x31\x92\xcc\x20\x51\xa8\xda\x51\xb2\xb4\x31\x47\x4c\x6c\x21\x09\x81\x8f\xa1\xa9\xa1\xfc\x77\xdd\x4c\xb5\x49\xab\xb3\xe1\x47\xf0\x64\x05\x58\x05\x6a\x49\xa8\x85\x80\xf5\xe7\xab\x37\x17\x86\x3f\x5b\xe8\xf2\xe5\x31\xda\x3b\x6c\x77\x51\x73\x56\x79\x28\x19\x0b\x44\x8b\x80\x1c\xb5\x18\x1f\x6f\x4f\x64\x18\x6c\xf3\x91\xa7\x5a\xad\x86\xed\x97\x48\xd1\xcf\xba\xff\x11\x92\xe4\xeb\x58\xe0\xcf\xbb\x2a\xae\x63\x81\x1f\x21\x16\x70\x94\x9b\x26\x25\xc9\xcb\x16\x9c\x7a\x49\x25\xf3\x32\xbb\xd4\xf9\xf2\xe7\xf9\xfb\xd0\x19\x5a\xa8\xf6\xda\xbb\x60\xd3\x1a\x79\x39\x98\x35\x36\xaf\x0b\x3d\xfe\x74\x9b\xd8\x09\xf9\xdf\x6e\x33\x3b\x91\x82\x15\xf7\xb4\x4d\xe7\xa7\xd9\xda\x76\xc0\xfa\x21\x77\xb8\x13\x42\xd6\x1b\xdd\xeb\x8d\x6e\x8b\x73\x6b\x1f\xa7\x06\x93\xd6\x1b\xdd\xdf\x97\x9b\xb3\xc2\x46\x77\x6e\x41\xaf\x55\x76\x5c\x70\x59\x1e\xbb\xf1\x5d\x04\x57\x67\xff\xdb\xcb\xf7\xa9\xbd\x05\x5e\xe8\xf7\xb5\x6b\x90\xbf\xcf\x8d\xac\x64\x02\x96\xae\x11\x2e\x30\x73\x6d\xdd\xd7\x7b\xe2\x5f\xf9\x8d\x89\x54\x02\xed\xb7\xfc\x92\x7b\x4b\xbe\xe8\x97\xf5\x72\x07\x00\x55\xef\xfa\xe5\xa3\xa1\xaf\xff\xba\xdf\xe3\x6d\xb1\xbd\x63\x5f\x88\x30\xab\x5e\xf8\x9b\x13\x34\xce\x6f\xfa\x5d\xdb\xbf\x9a\x39\xbc\x34\x00\x5c\xe7\xf2\xd6\x7e\x6e\x1d\x26\xad\x98\xcb\x4b\xc5\x6c\x9d\xd3\x5b\x75\x27\x2b\xfe\x2a\xe6\x33\x8e\x7c\x47\x8e\xee\xc5\xf4\xcc\x2f\x5a\xd1\xd8\x8f\x70\x7e\x27\x7f\x9e\x21\x5d\xd8\xba\xfe\x6e\x97\x41\xd1\x5f\x71\xaf\xeb\xab\x24\xaf\x96\xc8\x16\xe5\x4d\x46\x3e\x4b\x97\xe8\x8c\x90\x58\xc6\xfa\x8c\x94\x84\xf4\xb5\x5d\x5e\xdb\xe5\x27\xb6\xcb\x6b\x93\xec\xe2\xd9\x93\x94\x5c\x3d\x81\x55\x2e\x94\x5e\x55\xf8\xb5\xe5\xda\xaa\x79\x16\x79\x61\xeb\x75\x45\xd6\xda\x2e\x7e\x27\x4c\xfa\x8a\x15\x59\xb3\xc4\xec\xba\x18\xeb\x29\x8b\xb1\x9e\x2e\x0b\xb2\x8d\x7d\x9f\xd1\x41\x96\x05\x59\xa7\x45\x56\x4b\x8b\x1c\x29\x3e\xbe\x9d\x71\xad\xb8\x9a\x54\xa4\x3e\x36\x05\xd2\x13\x60\xf1\xdb\xb5\xba\x2c\xdd\xfb\xbb\xca\xa5\xe4\x59\x33\x37\x93\xac\x44\x26\x23\x06\xc9\x09\x96\x48\x4c\x58\x1c\xf8\x68\x08\x28\x16\xe6\xc4\x41\x8f\xd1\x11\x19\xc7\x1c\xb4\x60\x99\xb3\xfa\xec\x08\xc6\x30\x85\x51\x23\x77\x86\x57\xad\xf5\x72\xf6\x47\x5d\xce\xd6\xe9\x97\x1f\xc9\xd7\xdf\xc8\x20\xaa\x81\x13\xec\xfb\xe6\x55\xd0\x67\xe6\xff\xe8\x98\x85\x21\xa3\xc9\x2d\xfd\x8f\x32\x1b\xfd\x99\x75\x4b\x0c\xbf\x65\xb1\x6f\x09\xf5\xad\x9f\x11\x1e\x83\xf5\x53\x90\xcf\xf6\x4f\xc9\x24\x0e\xac\xdf\x44\x42\x98\x4e\xa1\xa3\xb8\x35\xe2\xca\xfa\x4b\x62\xb3\x51\x8d\xb7\x70\xc9\x52\x58\x94\x1b\x11\x2a\x61\x6c\xaf\x7f\xe4\x73\x8d\x56\x1a\xe7\xea\x66\xfa\x81\x16\x81\xb4\x0d\x0e\x82\x37\x23\x9b\x45\xf3\x84\xe7\x8d\xa6\xf7\x12\x46\xc0\x81\x7a\xb9\x2d\xcc\x8a\x6a\x5f\x17\x53\x90\x96\x77\xbf\x24\x51\x4e\xe6\x20\x3d\x93\xd8\x21\xfd\x95\xcd\x67\x8b\xf0\x80\xf8\x73\x3b\xe9\x67\x05\x9a\xfa\xcb\x4d\x30\x59\x3c\xbd\xb5\x64\x60\xa2\xb8\x5e\xd5\xc8\xc2\xf3\x35\x48\xbc\x24\x8a\xec\x81\x02\x5f\x88\x80\x29\x14\xf4\x07\x38\x67\x83\x46\x8c\x87\x58\xf6\x91\x8f\x25\x34\x25\x09\x61\x11\x98\x90\xf9\xda\x77\x5f\x15\x8e\xbe\x9f\x9c\x8a\x9a\x38\x4f\x84\xd1\x2b\x90\x92\xd0\xcc\x53\x73\xa9\x36\xb1\x15\x3b\xe6\xc1\xe3\x26\x2d\xe6\x0e\x2d\x72\xe0\x78\xe4\x79\x2c\xa6\x73\x6d\x8e\x17\x10\xa0\x72\x90\xc3\x2f\xb9\x27\xc0\xe3\x30\x6f\xee\x66\x7d\x17\xcf\x9f\x0d\x71\x3e\xea\x27\x10\x05\x6c\x1a\x02\x95\xe7\xcc\xac\x2e\x69\x7b\x9f\x28\xe3\x1c\x12\x8a\x25\xb3\x44\x26\xc1\x6c\x7a\xa1\x5d\xfb\x9c\x0d\x0d\x71\x14\x11\x3a\xb6\x07\x2c\xfa\xbc\x75\x33\xba\xd7\x98\x8f\x61\xe6\xf4\x31\x0a\xf5\xed\x52\x15\xa8\x0d\x17\x3e\xe6\x61\xdf\xe5\x41\x37\xcc\x33\x81\x1e\x18\xbf\x0d\x18\xf6\xf5\x91\xd9\x98\x26\xbe\xa2\x97\xdf\xe5\x73\xe8\xdf\x82\x35\x67\xe5\x25\x22\x0b\xa2\xe6\x4f\x6d\xe1\xc4\xdc\xaf\x64\xe4\xc1\xe5\x21\x68\xba\x50\xe3\xe8\xed\x59\x82\x54\xde\xe9\x20\xea\xe1\x7d\x27\x7f\x73\x62\xd0\xaa\x38\x57\xb9\xb0\x80\x04\x81\x31\x0e\x25\xaf\xa5\x69\x80\xeb\x18\x57\x14\x5d\x9d\x05\x83\x94\xcf\x0b\x2b\xf5\x4f\x08\xab\x3c\x00\xa2\x7a\xc9\xab\xc4\xd8\xf0\x15\x73\x8e\xa7\x85\x27\xda\xe7\x28\xbb\x5e\x85\x09\xb5\x69\x5f\x6a\x6a\x73\xee\x54\x62\xd2\x84\xed\x50\xfd\xa2\xd8\x51\x6d\x88\x73\xda\xf3\x13\x0b\x7c\x91\x46\xf1\x3a\xf4\x32\x95\x9c\x26\x16\x53\x10\xb4\x36\x19\x98\xe8\x8c\x0a\x89\xa9\x07\xad\x55\x64\xb4\x72\x85\xc8\x26\xe2\x59\xf2\xa2\x5f\x92\x4e\xf2\xac\x79\xc9\xda\x54\x88\xf4\xb3\xfc\x2c\x1a\x83\xaf\x87\xbe\x84\x31\x11\x92\x4f\x9f\x98\x25\x1a\x38\x4a\x81\x7f\x05\xde\x98\xc6\x88\xa7\x23\x3e\x15\x97\x52\x59\xd2\xd1\x7c\x4e\x92\xf2\xf1\xbd\xdb\xfc\x1e\x15\x33\x15\x73\x4c\xed\x8a\x0b\xfb\x3d\x0e\x62\x87\x1f\x6d\x1b\xd1\x72\x26\xa2\x0a\xdb\xb4\xa0\xad\x98\x5f\xc9\xa3\x6d\xeb\x75\x41\x9f\xeb\x27\x44\x1a\xc5\xd0\xa7\xfc\xa6\xc9\x8c\xd5\xc5\x15\xef\x4a\x62\x59\x70\x6c\x73\x5c\x01\x1a\x87\xb6\x74\x29\x37\xc0\x80\x00\xdb\x69\xe1\x80\xfd\xa9\x7b\x84\x64\x37\xd6\xf6\x4e\x5d\xf3\xa3\x73\x83\x73\x79\x5f\x01\xd8\x3d\x01\x46\x25\x95\x73\x69\x57\xd0\x64\xfb\xd5\x08\xeb\x77\x2c\x50\x14\x60\x0a\x56\x3a\xcc\x6c\xee\x36\x56\x51\xae\x39\x84\x57\xb8\x1b\x36\x4f\x56\x58\x87\x0d\xe4\x2f\x85\xdc\x95\xe6\xc4\xbc\x29\x13\xb9\x16\x15\xba\x58\xd5\x39\x05\x50\x0a\xf5\x96\xa1\x42\x4b\x6f\xe1\x25\x15\x3b\x82\xad\x2f\x4c\x4f\xed\x0e\x2d\x43\xc5\x63\xe6\xf1\x2a\x91\x57\x27\x51\xb6\x7d\x5a\x8a\xb0\xbc\xdf\xb2\x74\x04\xef\xf4\x4c\x96\x76\x64\x96\xab\xae\x73\x9b\x40\xeb\xee\xf1\x04\x53\x0a\xc1\x1c\x5b\xe7\xc3\x08\xc7\x81\x54\x77\xf1\x30\x80\x0a\x0b\x98\x3c\xcc\x33\xfc\x04\x84\xf2\xed\x97\xb5\xa6\xc6\x6c\xda\xb0\x59\x14\xe5\x0c\xab\x9f\x6c\xa5\xe6\x87\x5b\x76\x1c\x2c\x04\x19\x53\x7b\xad\x5b\x66\xec\xec\x4e\xc4\xd9\x3d\x51\x0e\x7b\x1e\x94\x0f\xb9\x07\x05\x59\xb4\xbd\x9b\xb9\x09\x38\xe5\x93\xd9\x08\x19\x67\x65\x80\x4d\x1c\x6d\x3d\x29\x9d\x75\xeb\x8a\x9d\x14\x34\x5b\xa8\xe6\x89\x53\x85\x2b\x9b\xa9\x40\x01\x97\x32\x5c\xf7\x57\x4a\xf2\x99\x80\xec\x23\x25\xe6\xf9\x20\x75\xb1\xea\xa2\xb9\xc8\xcf\xcc\xf0\x9d\x71\xc8\x06\xfd\xcc\xca\x40\xcf\x73\xea\x54\x4b\xbd\x38\x8a\x09\x8e\x20\x77\x3b\xe2\xcc\x03\x21\xec\xa3\xea\xd4\x6d\x93\xa1\x9d\x60\xea\x07\xf9\x8c\x5b\xce\x70\xe4\xe5\xc2\xe1\x17\xb8\xa4\x42\xf9\x05\xae\xa9\x2f\x7d\x3c\x45\x8b\x61\x92\xbc\x18\x04\x49\xf6\x22\xf7\x54\xab\xe8\x40\x2f\x3a\xab\x3a\x22\x25\xfe\xa6\x68\x2c\xee\x91\x37\x3f\x8b\xa6\x3a\xb1\x56\xd9\x8c\x3a\x88\xab\x0b\xab\x9c\xd4\xb1\xc1\x5a\x5c\xa9\x8d\x9c\xcb\xec\x15\xd7\xa0\x82\x7b\xb6\x9a\x2b\x95\x73\x53\x96\xec\x9b\x33\x3c\x45\xec\xbe\x85\xeb\x55\x41\xcc\x92\x8b\x6b\x5a\x11\x31\x48\xbf\x58\xe6\x5c\x67\x8b\xc9\x7f\x73\xa5\xb9\x56\x42\xe5\x5e\xcf\xb1\xa6\x7c\xb7\xfe\xde\x13\x38\x7a\xdf\xc4\xc3\x7b\x0a\xc1\x5d\xb2\xb7\xdb\x23\xfc\x13\xb8\x82\x0d\xb7\x0b\x68\x7d\xc7\xa4\x18\x03\xab\x27\xce\xf0\xf1\xef\xcd\x19\x0a\x97\x10\x71\x10\x6a\xc4\xf2\x67\xa6\x44\x1c\x45\x8c\x4b\xf0\xd1\x70\xaa\xc3\xcc\xa3\xb7\x67\x49\xc7\x52\x8e\xba\xbc\xb6\xa1\xf2\xfa\x66\x6e\x25\x8a\x5d\xb8\x6b\xe8\x7d\x4a\x88\x1f\x05\xa3\x83\x1c\xd8\x6f\xb4\xe3\x57\x5c\x72\x4b\xf3\x71\x81\x43\x70\x7f\x09\xad\x35\xcf\x00\xd8\x0f\x2a\xac\xa5\xf3\x63\x71\x8f\x1b\x29\x59\xe9\x4b\x62\x9c\x2f\x26\x4f\x1a\x2d\x33\xd6\x53\x29\x4c\xd1\xb5\x28\x22\x37\x0f\xef\x23\xfb\x67\x09\xf9\xda\x3c\x22\x1e\xa3\x83\xe2\xc6\x66\x69\xb0\x77\x97\xe7\xc9\x26\x8b\x6a\xbf\xfa\x68\x01\x1e\x2e\x9a\x8f\x73\xdd\x24\x2b\x10\xc2\x12\xc6\x8c\x93\xcf\xe0\x38\xb7\xfb\x11\xf3\x52\x2d\x34\x38\xc2\x43\x12\x90\xb2\x72\xb8\x5e\x17\xb3\x1a\x97\x8d\x90\xa7\xe6\xfb\x8b\x22\x5b\xe3\xb0\x30\xcb\x82\xa6\xd7\x91\x36\x38\x69\x7a\x59\x57\x66\x79\x98\xda\x65\x59\xf7\xe6\x88\x07\x40\xb8\xe4\x46\x96\xa0\x65\x0a\x33\x22\x10\xf8\x6e\x59\x28\x59\x20\x64\x1b\xbd\x1f\x87\x80\xf2\xb2\xf5\x27\x58\xcf\xcd\x27\x12\x37\xca\xa5\xa4\x59\xb4\x65\x0a\x4a\xdd\x5f\x86\x54\x8a\x72\x76\xa2\x8c\x06\x07\x8f\xf1\xd9\x0b\x53\x85\xa9\x77\x08\x79\xa1\x4e\xd4\x51\x25\x6a\x97\xe5\x18\x1c\xac\x72\xa1\xe2\xf1\x44\x85\xd3\x05\xc7\x80\x08\xf5\xe1\x53\x09\xfa\x08\x07\x02\xea\x63\x59\x2e\xcc\x2a\x16\x0b\x99\xfd\x0c\xd4\x48\xb6\x47\xed\x2a\x21\x83\xb4\x55\xd4\x34\x17\xe9\x8b\x38\x1c\x02\x57\xac\xd4\xf3\x89\x08\x45\x80\xbd\x89\x4d\xf4\x13\x92\x51\xac\x66\x9a\x91\xd1\x6e\x1b\x42\x92\x0f\xa9\x39\x3d\xb7\xff\x64\x6a\x7b\x95\x14\x8b\x9b\x4d\x36\xdd\x49\x99\x48\x8f\x13\x09\x9c\xe0\x96\x96\x10\x31\xa5\x12\x7f\x32\x4b\x0b\x11\x99\xa8\x21\x22\x2c\x84\x42\x12\x60\x9e\x7e\x7d\xda\xee\x02\xe8\x26\x05\x7c\x83\xbc\x00\xc7\x42\xfb\x29\x98\xa2\xab\x5f\xcf\x4d\xbc\x63\xbe\x9c\x9d\xc2\x3a\x55\x7c\xd3\x8c\x4e\x6d\x87\xee\x6f\xac\x37\xa6\xd3\x19\xd8\x9c\x1a\xdc\x18\x1b\x21\x32\x38\x2f\x19\x4f\x59\xb7\xa5\x10\xe3\xfa\x58\x0a\xfd\xad\xed\xe3\x9c\x2b\x21\xec\x01\xe4\x04\x08\xd7\x93\xbf\xa5\x6c\x96\x1e\x69\xc4\x82\x80\x3d\xe8\x0f\x41\x6b\xc2\xfa\x1b\xb3\x41\x6e\x6e\x6e\xc4\x5d\x56\xe6\xa6\xfa\x21\x2c\x3c\xfb\x79\xd6\xf8\x7a\x79\x24\xd0\x00\x53\x7f\x90\xba\x66\x8f\x41\x69\x6b\xf6\xd1\xe2\x4a\xfc\xcc\xb7\xc7\x73\x33\x4c\x37\xa5\x49\x69\xfa\xe0\x6f\x21\xc6\x11\x31\x6d\xb4\xc4\x21\x22\x10\x84\x91\x9c\x6e\xa9\x7b\x99\xef\x6c\xb6\x93\x44\x1c\xa8\x88\x80\xe7\xe6\x4f\x61\xd3\x9a\xc9\x75\x14\x30\x1f\x72\xaf\x1c\x96\x65\xbd\x20\xca\xb6\xb8\xa7\xa4\x35\x2a\x34\xd4\xa8\x70\x02\xe0\xb1\x5a\x28\xe4\x34\x80\xbe\xce\x0f\x18\x5b\xa1\xbf\x3c\xe8\xd6\xb0\x4c\xc1\x74\xa3\x4c\xa1\x2c\x59\x98\xaf\x59\x0b\x34\xea\x61\x02\x1c\x72\xea\x94\x0d\x99\xd3\x2a\x74\xa4\xe4\x04\xfc\x44\x3b\x94\x5d\xd2\xe0\x0c\x5e\x6a\x72\x6e\x14\x97\x6e\xb6\xd0\x8d\x45\x82\xfa\x99\x48\x8b\xfa\x53\x3b\x87\x37\x5b\x08\x53\x1f\xdd\x24\xbe\xfb\x4d\xa6\x68\xe9\x10\xa6\x72\x90\x71\x33\xe9\x37\xff\xf3\x77\xd5\xf7\xf9\x8d\x16\x9b\x9b\xf3\xb3\x5f\x4e\x1d\x7d\x3c\x46\x3f\xc6\xd4\x93\xe4\x1e\x8a\xfd\x8f\x2e\x4e\x6e\xcc\x90\x6f\x2e\x6f\x5a\xe8\x27\xf6\x00\xf7\xc0\xb7\xd0\x94\xc5\xda\x30\x28\xca\x31\x0a\xf1\x27\x12\xc6\xa1\xe2\x41\xa7\x9d\x81\x63\x54\xd3\x8a\x53\x4a\xb5\x58\x58\xec\x3f\x9d\xc9\x99\x4b\x3b\x0b\xa1\xb1\x79\xb3\x46\xf1\x4d\x4b\xdc\x0d\x7e\x10\x4d\x71\x27\x9a\x26\xcb\x64\x90\xd4\x6e\xa5\x61\x0d\xba\x31\x1b\x20\x37\x75\xd5\x35\xaf\xab\xcf\x51\x1e\xbe\x06\x9f\x82\x7e\x9e\xdf\x79\xd1\xdd\xff\x19\x35\xff\xe5\x26\xc3\xd4\x8a\x90\xa4\x1e\xc2\x90\x81\xcd\x28\xa6\xec\x5f\x62\x2e\x85\xb9\xaf\xa8\x5a\x11\xe3\x80\xdc\x82\x42\xfa\x2f\xdd\xdd\x2f\x62\x58\xb4\xb9\x54\x0f\xf3\xd3\x62\xd9\x1b\x2c\xf5\xf3\x58\x00\x47\x13\x2c\x50\x04\x3c\x24\x42\x24\xc5\x22\x02\x40\x8b\x94\xe1\x0b\xf8\x96\x1c\x5c\x30\x09\xad\x14\x3f\xb3\xe8\x64\x95\xfa\x4a\xe2\x93\xc4\x3d\x11\x56\xef\x6a\xf3\x95\x38\x0d\x5a\xe6\x2a\x8c\x92\xdb\x00\x39\xd6\xf8\x9c\x7d\x41\x45\xb3\x57\x4b\x4a\x1a\xab\x99\xb7\x8d\xec\x85\x1d\xbd\x9f\x92\xa2\x95\xbc\xb1\x63\x03\x85\x3e\x1a\xea\xbb\xc9\x4d\xf3\xe3\x65\x92\x44\xfd\xf9\xfd\xf5\x86\x3d\xe2\x44\xca\x48\x41\xcf\x53\x5b\xeb\x70\xc4\x42\xe5\x89\x61\x74\xe3\xf5\x34\x77\x76\xc8\xbc\x8f\x80\x17\x00\x10\xbf\x8f\x02\x36\x1e\x08\x42\x6f\x07\xed\x56\x67\xf6\xc0\x14\xa8\xe5\x20\xcd\x9e\x2d\x55\xfc\x96\x7c\xff\xde\x1e\xa4\x51\xc0\xff\x9c\x8d\xd1\x15\xa1\xb7\xb3\xdb\x69\x0a\x06\x35\x72\xad\x5d\xf9\x92\x66\xd1\x12\xe4\x83\xf5\x22\xe4\x2c\x9d\xb0\x22\xfe\xad\x88\x8e\x33\x8c\xca\xf9\x82\x26\x12\xf6\x78\x55\xd1\x7a\x53\xef\x9b\x0d\x8a\xfb\x66\x4d\xd7\xbe\x59\x39\x06\xad\xae\x0e\x0c\xc3\x72\x5a\x26\x53\xb5\x7f\xfe\xab\x18\x8f\x11\x19\x98\x09\xa8\x1b\x15\x57\x0f\xae\xae\x30\x0e\x24\x19\x04\x84\x3a\xdf\xf4\x98\x6d\x9b\xdb\x2a\x9f\x6f\x60\xcd\xdd\x6b\x05\x0b\x9d\x13\xea\x6a\x99\x20\x3e\xbf\x4d\xc5\xd9\xab\xe9\xf5\xa9\x39\xe6\x2c\x8e\xfa\xa8\x01\xd4\x8f\x18\xa1\xb2\x5c\xa7\x29\x26\xec\x61\x80\x83\xe0\xf1\xe4\x5c\x4d\xd8\x83\x5a\xef\xab\x89\x99\xd7\xe2\x91\xa4\x48\x16\x11\x6f\x41\x9e\x91\x85\xa1\xf2\x13\xd4\xea\x24\xc1\x9f\x95\xa5\x99\xc5\x53\x03\xd0\xea\x2a\xdc\x22\x74\x5d\xdd\xa0\x2a\x39\x64\xa3\xad\x95\x2e\x8f\xb3\x90\x10\x3d\x3e\x81\x50\x48\xaf\x67\x57\x73\xae\x20\x27\x30\xa9\x00\x2e\x07\xda\x69\xac\x6a\x53\x1d\x56\x96\xaf\x23\xdf\xd7\x9b\x03\xb1\x90\x2c\x34\xbe\x68\xea\x8d\x78\x4c\xbb\x27\x32\x59\xf9\x13\x7f\x37\x04\x21\x4c\x1e\x00\x49\x8e\xa9\x20\xb2\x98\xfd\xc9\xae\xc5\xe4\xa8\x6b\x01\x2d\x25\x7a\xae\x53\x77\x2f\xf1\xb9\x0d\xd2\x92\xa9\x80\x14\xfb\xbe\x55\xf4\xe1\xba\x12\xe1\x78\xa9\x3a\xcd\x6f\x58\x2d\x24\xf6\x55\x2a\xbc\xac\x81\xbd\x61\xa8\x8d\x7e\x1d\x94\x7f\x53\xbd\x1e\x8f\xb2\x7b\xeb\x25\x7f\x35\x17\x62\xd5\x34\x44\x54\xb6\x48\x70\x3e\xd3\xe2\x6a\xb8\x8d\x8e\x3c\x59\xdc\xc8\x29\x63\xef\x34\xf0\xf5\x30\x6f\xe6\xb4\xc3\xd9\x68\xc1\x18\x75\x34\x10\x3e\x49\x8e\xbd\xe5\x54\xf0\xd4\xf4\x41\x38\x11\xd6\x11\x67\xe6\x2c\xef\x21\xf3\x8b\x56\x23\xbb\xfe\xf8\xea\xf3\x14\xb2\x98\x60\x94\xb2\xf8\x6b\x89\x5a\x4e\x0c\xbe\x94\xac\x4d\xb0\x18\x4c\x00\xfb\xc0\x07\x23\x12\x48\x28\xd5\x0c\x64\x57\x6e\x8e\x5f\xea\xc6\x68\x88\x85\x8a\xfe\x4d\x66\xc1\x6c\x05\x7b\x7a\xde\x19\x05\x64\xe0\x3e\x52\xf8\x5c\xbb\x9f\x73\xf0\x52\xb2\x67\xc6\x4d\x62\x5d\x86\x40\xd9\x91\xac\xaa\xc9\x7d\xa5\x6f\xe2\x24\x9d\x2f\x8a\xbb\xc4\xc5\x2b\x91\x89\x9f\xcc\x50\x8b\x9b\x3f\x9d\xac\x96\x36\xb0\x5d\x68\x61\x91\xa2\x96\x4c\xd4\x97\x17\xd7\x92\x24\xd5\x13\xd9\xf9\x9f\x7c\x77\x86\x7e\xaf\xa7\xe7\x6c\x6c\x97\xef\x2c\xa8\xfe\x2a\xbc\x77\xe4\x38\x72\xd2\x7a\x4b\x0c\x35\x0e\x87\xe2\xbe\x2d\xf6\x25\x85\xfd\x71\xbb\x3b\x9e\xec\x8e\x7b\x56\xf4\x53\xaa\x9c\xb4\xfa\xec\x0d\xf9\x88\xb7\xdb\xdd\x68\x44\x6f\x27\xed\xfc\x00\xe9\xeb\x88\xa8\x21\xf8\xbd\xd7\xc4\x9e\x27\x9b\x9d\xbd\x2e\x8c\xba\xfe\x41\xb3\xdd\x6d\x1f\x36\x7b\x9d\xce\x7e\xf3\xa0\xb7\xd7\x6d\xfa\xa3\xbd\x1d\xaf\xdb\xee\xee\x7a\xdd\x3d\x07\x94\xe4\x55\x45\xd4\x18\x76\x7a\x3d\xff\xf0\xb0\xd3\x6c\x1f\xc0\xb0\xd9\xeb\xed\x77\x9b\x07\xe0\x75\x9a\x30\x6c\xef\xf4\xbc\xbd\xc3\xee\x4e\x67\x68\xf7\x8f\x79\xd0\x47\x8d\x11\x63\x4d\x17\xbe\xad\x5b\x2c\x5a\xd8\x0b\xa1\xe5\xb1\xb0\xdf\xeb\xed\x34\x0a\xd1\x98\xb3\x22\xd3\x22\xbf\x7d\x7b\x10\xd0\x71\x7b\xa7\x23\xe0\xf0\xae\x06\xf9\xd0\xee\xee\x76\xf7\x76\xa1\x89\x0f\x0e\x70\xb3\xd7\x1b\x0d\x9b\x07\xbd\xdd\x76\x13\xfc\x76\xa7\x0d\xc3\xbd\xa1\xb7\xeb\xcd\x23\xdf\xf7\x76\xf1\x41\xf7\xf0\xa0\x39\x04\x7f\xbf\xd9\xeb\x76\xa1\x79\x70\xd8\xdb\x6f\x8e\xf6\x46\x3e\xde\x3b\xec\x1e\x76\x47\xa3\x32\xf9\x43\xcc\x13\xf2\xbb\xe1\xc8\xc3\xed\x76\x57\x1e\xde\xed\x8b\x71\x4b\xf0\x2a\xf2\xd3\xea\xc4\x62\xd8\x5d\xae\x73\x44\x0d\x77\xcc\xef\xac\x38\x75\x45\xae\xb3\xd8\xcb\x4e\x2d\x99\x2b\x8b\x33\x45\xe9\x69\x12\xeb\xe8\xc9\xdd\x1a\xe2\xfc\xe1\x53\xb3\xa0\x3b\x3f\xd4\x2d\x4c\x8b\xda\x9c\x56\xc0\x35\xae\xae\x2f\xcf\x2e\x5e\xe5\x63\x13\xa7\x1f\x3a\xeb\xf1\xf3\xd5\x9b\x8b\xc2\xcb\x7c\x49\x4c\x5f\xac\xc0\x99\x1f\x60\x24\xd9\x1d\xfd\x54\x99\xd5\x72\x78\x9a\xe6\xc2\x74\x13\xed\xb2\x56\x15\x6c\x16\xaa\xb8\x75\x3a\x6f\x90\xd6\xe1\xda\x43\xfb\x80\xfd\x41\x00\x52\xd9\x80\xbb\x18\x8a\x64\x6a\xee\x2a\x81\x0b\xee\xcc\x50\xd5\xdf\x76\x70\xa4\x9a\x1a\x9d\xb6\x25\x4b\x89\x31\x2a\x9c\x26\x31\x3f\x3b\x63\xbe\xe8\xb0\x9d\x83\xa3\xcf\x01\x40\x8d\xe3\x37\x17\x17\xa7\xc7\xd7\x6f\x2e\x9b\xaf\x5f\xbd\xbe\x6e\xe6\x9a\x24\x6f\xff\xa3\xc6\x95\xf5\x05\x97\xf4\xdb\x2e\x02\x51\x26\xb3\xf2\x08\x93\xfd\xd5\xdf\x7a\x79\xae\x64\xab\xfc\x26\x59\xe1\x78\x00\xd4\xe8\x90\xf7\x67\x24\xbc\x7b\xe5\xf1\x93\xf8\x7c\xaf\x83\xdf\x7d\x3a\xfb\xc7\xdd\x8b\xeb\xbb\x8b\x4b\x3c\xe3\xd2\x99\xc9\xa6\xfe\x1a\x03\x9f\xd6\xe0\x54\xf7\x89\x38\xd5\x5d\xc8\xa8\xae\x83\x4f\xff\xb1\x26\xfd\x25\x26\x81\xa9\x5a\x88\x30\x17\x90\xdb\x4b\xe8\xa3\x77\x54\xd9\x01\xf5\x54\x67\x0c\x7e\xb1\x3f\x8c\x28\xf4\x6b\x55\x38\x22\x03\x93\x54\x4b\x8a\xe3\xfb\xa8\x84\x41\x7f\x89\xf1\xb2\x3a\x16\x8f\x05\x71\x48\x8d\x77\xa3\x46\x4a\x92\xc5\x68\x93\xf8\x9b\x2d\x74\xe5\x6a\xa7\x77\x55\xec\xd1\x94\x21\x67\x74\x2b\xd9\xeb\xf4\x02\x16\xfb\x83\x24\x23\xcf\xd3\xbb\xa6\x9e\xb5\x85\x7e\x35\x99\x71\x33\x91\x7d\x44\x7c\xf4\x1c\x75\xba\x3b\x95\x52\x11\xbc\x3f\x79\x15\x4f\x87\x67\xfc\x94\x7e\xe2\x47\x10\xee\x77\x7b\xe3\xbb\xdb\x5b\x72\x72\x9f\x4a\x45\xf1\x44\x19\x97\x24\xf4\xda\xbd\x27\x91\x84\xfd\x45\x82\xb0\xef\xd0\x97\x3a\x5f\xbf\x98\x11\xe3\x3c\xad\xcc\x45\xd2\xfe\xb7\x23\xe8\x38\x77\xfa\x2c\x22\xfe\xf3\xcd\x0e\xf9\x65\xc7\x8f\x7f\xfb\x70\x76\x7f\xbf\xfb\xe1\xfe\x3c\x98\x7e\xee\x84\xaf\x2e\x77\x7e\x9e\xde\x5d\x6c\x6a\xd3\x30\x62\x31\xf5\xe7\x28\xff\x87\x37\xfb\xe3\xee\x78\xef\xa7\x6b\xff\xdd\x2f\xef\x70\xf7\x56\xfc\x74\xd0\xbd\xfd\xf5\x64\x67\x9a\x72\xa6\x78\xba\x92\xd3\x34\x76\x9e\xc6\x32\x76\x16\x1a\xc6\x8e\x83\x2d\x99\x1a\xdf\x03\x27\xa3\x29\xfa\xf9\xfd\xb5\x39\xb3\xa9\x8f\x2e\x13\x87\x17\xe1\x58\x4e\x18\x27\x9f\xd3\xf7\x8f\x6f\x81\xd6\xe3\xcf\xce\xbb\xc9\xe9\xe4\x21\xfc\xfd\x45\xf4\xfe\xed\xe8\xac\x1b\x5c\xc0\x6d\xe4\xf7\xfe\x71\x92\xf2\xe7\x50\x2d\x6f\xc7\x8c\x8e\x02\xe2\xc9\x1a\xbc\xda\xd9\x7b\x12\x5e\xd9\x60\xdc\xbc\xb2\x5b\xd8\x22\x64\xea\xe6\x8c\xe5\x21\x02\xe1\x40\x2f\xaf\xba\xbc\xab\x92\x0f\x7b\xb7\x1f\xda\xef\xc8\xe9\xed\xe7\xdb\xdf\x8f\x3f\xbf\x7f\x0b\x67\x5d\xf6\x01\x26\xfe\xce\x69\xc2\x86\xf2\x31\x49\x2e\xd2\x0f\x9f\x84\xf2\xc3\x45\x84\x1f\x3a\x65\x24\x3b\x56\x11\xf2\x83\x96\xa6\x1c\x4e\xcf\xef\x5f\x1e\x7e\x7c\xfd\xeb\x87\xbd\x0f\xe3\xc9\xe8\xf5\xe1\xf8\xd5\xa5\xf8\xe9\xfe\xf4\xfd\x8c\xd6\xda\xc6\xe2\xdb\x51\x6c\xaf\x82\x7a\x4c\xf3\xa6\x5b\xb2\x1d\xec\x09\xe5\x7a\xbf\x39\x7e\xdd\x3c\xfd\xbd\x79\xd8\x47\xa3\xd9\xba\x65\xde\x87\xcb\xda\xc0\x27\xd9\x4c\xd6\x3e\x1c\x91\x66\x87\x7c\x6a\xef\x04\xd4\x0f\xc2\xbb\xf6\xdd\xc8\xdb\x17\x44\xe2\x5d\x11\x7c\xbc\x3f\xb0\xfd\xd8\x91\x75\xfc\x97\xe2\x43\x67\xbc\xeb\x1f\x1c\xdc\xb5\x03\xee\xf9\xf7\xbd\xf1\x3e\x0e\x86\xfb\x22\x18\x8d\xe9\xc7\x1d\x7f\x32\x14\x1f\xff\xf2\xff\xfe\x7a\xfa\xfb\xf5\xe5\x11\xfa\x2f\x43\x71\x4b\x63\xfc\x9c\xf8\x40\xa5\x9a\x33\x3b\x08\x25\x02\x6d\xf6\xda\xbd\xcd\x2d\xcd\x0b\xfd\xf3\xf8\xfc\xdd\xd5\xf5\xe9\xe5\x95\x61\x86\x7a\xa8\xf7\x52\x67\x13\x8b\x32\x40\xba\x7d\x67\xbc\xcb\xf8\x6e\xfb\x9e\xc4\xed\x7d\x06\x6a\xda\x26\xfc\xd6\xeb\xee\xf9\xe3\x91\xfc\xd8\xc1\xde\xa6\xbd\xc8\xa6\x27\x62\x6f\x2e\x22\xc2\xb2\xb7\x7f\x9b\x63\x4f\xae\xc5\x7b\x3e\xdd\xa3\xe2\x6e\xd8\x15\x17\xe1\xcb\x8f\xbb\xc3\xdf\xa3\x93\xfd\x63\xdc\xd8\xf8\xbf\x00\x00\x00\xff\xff\x95\xb3\x40\x78\xda\xa0\x00\x00")

func connector_mgmtYamlBytes() ([]byte, error) {
	return bindataRead(
		_connector_mgmtYaml,
		"connector_mgmt.yaml",
	)
}

func connector_mgmtYaml() (*asset, error) {
	bytes, err := connector_mgmtYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "connector_mgmt.yaml", size: 41178, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
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
	"connector_mgmt.yaml": connector_mgmtYaml,
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
	"connector_mgmt.yaml": &bintree{connector_mgmtYaml, map[string]*bintree{}},
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
