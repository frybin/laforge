// Code generated by fileb0x at "2018-10-14 21:12:17.604306234 -0700 PDT m=+0.006288647" from config file "assets.toml" DO NOT EDIT.
// modification hash(ff255d370664d8b723095393ceb0023d.e5b9c5ef4c0b7aef8593382d0449dfd6)

package static

import (
	"bytes"
	"compress/gzip"
	"context"
	"io"
	"net/http"
	"os"
	"path"

	"golang.org/x/net/webdav"
)

var (
	// CTX is a context for webdav vfs
	CTX = context.Background()

	// FS is a virtual memory file system
	FS = webdav.NewMemFS()

	// Handler is used to server files through a http handler
	Handler *webdav.Handler

	// HTTP is the http file system
	HTTP http.FileSystem = new(HTTPFS)
)

// HTTPFS implements http.FileSystem
type HTTPFS struct{}

// FileCommandTfTmpl is "command.tf.tmpl"
var FileCommandTfTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xb4\x93\x31\x8f\xd4\x30\x10\x85\x7b\xff\x8a\x91\x45\x01\x12\x44\x14\x88\xee\x0a\xb8\x02\xe8\x10\x14\x14\xe8\x64\xe5\xe2\xd9\xdb\x11\xf1\x4c\xe4\x99\xec\xde\x29\xca\x7f\x47\x93\x90\xb0\xc5\x22\x41\x81\x9b\x64\xf4\x66\x9e\x9f\x3f\xd9\xd3\x04\x19\x0f\xc4\x08\xb1\x93\x52\x5a\xce\x11\xe6\x39\x54\x54\x19\x6b\x87\x10\x79\xec\xfb\xb4\x95\x11\xe2\x50\xe5\x44\x4a\xc2\x69\x9a\xa0\xf9\x80\x06\x71\x53\x13\xb7\x05\x7d\x3c\xa9\xe1\xb0\xcb\x5e\x24\x1e\xcb\x3d\x56\x17\x23\x4c\x01\x20\xe3\x80\x9c\x35\x09\xc3\x0d\x7c\x0f\x00\x00\x91\xee\x4b\xea\xa4\x0c\xa3\x61\x3a\x95\x44\xac\xd6\x72\x87\xcd\x9f\x37\x8a\x01\xe0\x2e\x04\x80\x3d\x15\x56\x6f\x2b\x62\xf8\x0a\x1f\xb1\x5b\x37\x03\x98\x26\xa0\x03\x34\x1f\x45\xad\xf9\xa4\xdf\x88\xb3\x9c\xd5\x0f\x0a\xcb\xea\x84\x19\x3b\x23\xe1\x5f\xfd\xbe\x8e\xa2\xb6\xfc\xdc\x40\x7c\x36\xfd\x7b\xb8\x86\x86\xd3\x9b\xd4\xe6\x5c\x51\x75\x89\xba\x2e\x7b\x1a\x70\xf3\x3d\x13\xd7\xf2\x5b\x1a\x15\xeb\x26\xbd\xcb\x85\x98\xd4\x6a\x6b\x52\x2f\xa6\xa9\xa0\x8c\xb6\xb4\xbc\x7d\x7d\x31\x3b\xb4\xaa\x67\xa9\xd9\x05\x0f\x75\x2b\x65\x40\x23\x3f\x54\xf3\x45\xc4\x3e\x6f\xfa\xbc\x67\x99\x37\x36\xd8\x2b\xfe\x2d\x8d\xff\x08\x64\xb5\x56\x3d\x5e\x21\xb2\x6a\x55\xc4\xae\xb0\x80\x2b\x38\x2a\x9d\x5a\xc3\xf4\x03\x9f\xd6\xbc\x07\xea\xf1\xb9\x93\x21\xce\xf8\x08\xcd\xfb\x91\xfa\xdc\xdc\x0a\x1f\xe8\xc1\xc3\xf6\x49\xf5\x98\x2e\xc6\x92\x4f\x2c\xb7\xec\xc5\x15\x62\xec\x20\xc3\x52\x12\xf7\xfe\x80\xb6\x7b\x0c\x3b\x7f\x7f\x4e\xdb\xf7\xab\x55\xe2\x07\x77\x7b\xb9\x74\xdd\x05\x77\x9b\xc3\xee\xf5\x33\x00\x00\xff\xff\x0d\x32\x43\xcf\x8a\x03\x00\x00")

// FileDNSRecordTfTmpl is "dns_record.tf.tmpl"
var FileDNSRecordTfTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xac\x92\x3f\x4b\x04\x31\x10\xc5\xfb\x7c\x8a\x47\xb0\x3c\xc2\x81\xf5\x75\x82\xd8\x5c\xa1\x60\xa1\x48\x58\x37\x73\x6b\x30\x37\x59\x92\xac\x82\x21\xdf\x5d\xb2\xec\x9f\x3b\x0e\x1b\xb5\xcc\xbc\xc7\x9b\x99\xdf\x24\x67\x18\x3a\x58\x26\x48\xc3\x51\x07\x6a\x7d\x30\x12\xa5\x08\x11\x28\xfa\x21\xb4\x93\xd2\x4c\x9a\x8e\x94\x24\x64\x1f\xfc\x87\x8d\xd6\xb3\xce\x19\xea\x96\x12\xe4\xec\xd7\xdc\x1c\xa9\x46\xe8\x98\xa8\x5f\xe4\xfa\xd0\x3c\x1c\x5f\x29\x54\x51\x22\x0b\xe0\xcb\x33\x61\x07\x59\x5d\x37\xfb\x87\xfb\xb1\x85\x7a\xaa\xd5\x52\x94\x14\x40\x0d\xbb\x74\xec\x6b\xb5\x94\x6a\x68\x8c\x09\x14\x23\x45\xec\xf0\x2c\x00\x20\x67\xd8\xc3\xa9\xfb\x8e\xdf\x28\xd8\x44\xa6\xee\x55\x1d\xf2\x2a\x77\xde\x77\x8e\x74\xeb\x8f\xfd\x90\x48\x5b\x8e\xa9\xe1\x96\xd4\xcf\xeb\x28\xa6\xf4\xe9\xc3\xbb\xb6\x9c\x28\x1c\x9a\x96\xd4\x76\xad\xf5\x45\x6e\xe6\xee\xe4\x22\x2d\xad\xce\x07\x7f\x6c\xdc\x30\x4e\xbe\x9a\x79\x1a\xeb\x45\x00\x29\x39\xec\x70\xbd\xdd\x8a\xb3\x03\xf0\xe0\x9c\x9e\x9f\xff\x47\x3f\x05\xdb\x75\x14\xe2\xf8\x00\x66\x06\xda\x9a\x4a\xfc\x57\x8c\x4e\x32\xc6\xeb\x14\x21\x00\x43\x3d\xb1\x89\xda\xf3\x72\x23\xb9\x64\xac\xa2\x5c\xb9\x5c\x7c\x39\xf5\xf7\x95\x37\x23\xe3\x22\xc4\x0a\xfd\x3b\x00\x00\xff\xff\xf5\xf2\x7c\x23\xfe\x02\x00\x00")

// FileInfraTfTmpl is "infra.tf.tmpl"
var FileInfraTfTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xdc\x5a\x6d\x6f\xdb\x38\x12\xfe\xae\x5f\x41\xa8\x29\xae\x01\x2a\x25\x71\x16\x69\x1b\x6c\x80\x4b\x93\xb4\x1b\x5c\x2f\x39\xa4\xd9\xdb\x0f\xbb\x85\x40\x8b\x63\x87\x17\x89\xd4\x92\x94\x13\xc7\xd0\x7f\x3f\xf0\x4d\x2f\xb6\xec\xc4\x89\x73\x87\x3b\x7f\x48\x6d\x71\xe6\xe1\x70\x86\xf3\xcc\x90\xea\x9b\xe7\x7f\x82\x37\xe8\xdb\xf1\x97\xcb\xab\xaf\x67\xe8\xeb\xd9\xc5\xd9\xd5\xf1\xf5\xd9\x29\xba\x3e\xbb\xba\xd2\x0f\xff\x8e\x4e\x2e\x2f\xbe\x9c\x7f\xfd\xf5\xea\xf8\xfa\xfc\xf2\x22\x78\x83\xa2\x08\xfd\x76\x7c\x75\x71\x7e\xf1\x15\x45\x51\xf0\x06\x5d\xdf\x50\x89\x46\x34\x03\x44\x25\xc2\xa5\xe2\x39\x56\x34\xc5\x59\x36\x45\x63\x60\x20\xb0\x02\x12\xa3\x53\x8e\x18\x57\x08\x08\x55\x88\xaa\xbf\xc8\xe0\x0d\x4a\x39\x53\xc0\x94\x44\x84\x0a\x48\x55\x36\x8d\xd1\xaf\x12\xd0\x37\x3c\xe2\x62\x0c\x08\x33\x82\x04\xa0\x61\x49\x33\x82\x94\x9f\x24\x0e\x5e\xb2\xd2\x40\x81\x10\x1a\x3f\x47\xb3\x00\xa1\x21\x4e\x6f\x81\x11\x14\x82\x4a\xc9\x64\x3f\x34\x0f\x11\x2a\x04\x8c\xe8\x3d\x3a\x42\xa1\x54\x58\xc1\xce\x6c\x86\xb6\xe2\xcf\xda\x8e\xf8\xfc\x14\x55\x95\x7d\x70\x0d\x38\x37\x7f\x2e\xca\x7c\x08\x42\x3f\xaf\xd1\x63\x35\x32\xaa\xa1\xc1\x03\x46\x0a\x4e\xf5\x4a\x8f\xd0\xef\xe6\x09\x42\xe1\x8d\x52\x85\x3c\xdc\xd1\x58\x94\x11\xb8\xaf\xa7\x38\xe1\x6c\x44\xc7\xd6\xa6\x24\xc7\x52\x81\x08\x51\x55\x85\xef\xd7\xd3\x94\x19\x9e\x40\x4b\xf1\x87\xf9\x5b\x4a\x10\x0c\xe7\xa0\x17\xb7\x5a\xdf\x4b\x1a\x08\xeb\x16\x2c\xe5\x1d\x17\xe4\x71\x5d\x2f\xe9\x75\xab\xa0\x0a\x82\x09\x16\x14\x0f\x33\x40\xe1\x24\x97\xf4\x01\xac\xb7\xd5\xb4\x30\xc6\xe4\xb8\xd0\x92\x04\x46\xb8\xcc\x14\x3a\x72\xa1\x08\x65\x8e\xb3\x2c\xd4\x12\x6c\x2f\x92\x0a\x33\x82\x05\x89\xf6\xac\x45\x61\x0e\x84\x96\xf9\xc2\xf0\xc0\x0d\x67\x58\x8c\x61\x61\xf4\x27\x37\x7a\xdf\x3f\xbc\x3f\xe8\x35\x99\xcb\xa7\x19\x5c\x0e\x4b\xa6\xca\xbd\x03\x83\x6b\x7f\x44\x5c\x46\x69\xc6\x4b\xb2\xe3\x7e\xef\x1d\xec\xfe\x14\x65\x4a\x86\x1d\x95\x8f\x2b\x55\x3e\x76\x54\x52\x60\x8a\xcb\x0f\x46\xc3\x7e\x77\xe2\xee\xc7\x87\x8e\xdc\xc1\x52\xb9\x03\x27\x47\x60\x48\x31\x5b\x69\x41\x4e\x19\xcd\x71\x36\x6f\x89\xd5\xfc\x64\x34\xed\x77\xa7\xe6\x7e\x7c\x72\x72\x77\x83\xdb\x7d\x23\x74\x47\x19\xe1\x77\x1e\xdc\xff\x1a\xec\xee\x1d\x44\x29\x17\xd0\x88\x7f\x5c\x29\xbe\xfb\x31\x12\x83\x46\x78\x6f\xb0\x1a\x7c\xd0\x95\x3e\x78\xc4\x94\x7a\x0b\x14\x82\x4f\x28\x01\x81\xc2\x31\xe7\xe3\xcc\xed\xda\x54\x00\x01\xa6\x28\xce\x74\x4e\x87\x5b\x33\x4d\x4d\xef\x96\x67\xc5\x38\x2d\x12\xad\x93\x68\x39\x93\x15\xdb\x26\x31\x0a\xc1\xff\x05\xa9\x5a\x99\x51\x5a\xd7\xc9\xf9\x7c\x12\x30\xa6\x9c\x3d\xaa\x65\xc5\xac\x52\x15\x04\x9a\xb5\x08\x93\x0f\xe8\xf0\x68\x85\x12\x61\x32\x79\xe0\x0c\x12\x6a\xd2\x37\x08\x08\x56\xd8\xaf\xde\x8c\xe6\x98\xe1\x31\x10\x23\x15\x1a\x13\x2c\x6c\x55\x59\xe7\xb4\xf8\xa5\x1e\xd0\xf3\x0b\x90\xbc\x14\x29\xd4\x60\x29\xcf\x8b\x52\x41\xc2\x40\xdd\x71\x71\x1b\xa2\x70\x52\xa4\x8b\x18\xf1\x19\x9b\x50\xc1\x59\x0e\x4c\x59\x02\x8e\xd4\x12\x06\x8e\x34\x40\x80\x4c\x19\xd2\x0e\xc7\x0a\x12\x59\x0e\xdd\x04\x3a\x56\x23\x9c\x49\x58\x69\xcd\x88\x0a\xb8\x33\xa4\x13\xe2\x2c\xe3\x77\x09\x4d\xf3\xe2\x65\x56\x19\x9c\xc8\xe0\x68\x14\x6b\x8d\xdd\x38\xfd\xae\x88\x27\x45\x1a\x4b\xc8\x46\x49\x46\xd9\x6d\x15\x06\x7a\x4d\x1a\xa4\xae\x50\x5c\xf1\x94\x67\x1a\xc3\xc3\x56\x5a\xc8\xae\x29\x11\x98\x8d\xa1\x29\x37\xcb\xb7\xc9\xa4\x48\x93\x94\x92\xa6\xc2\xfc\x58\xcf\x37\x98\xe4\x94\x6d\xc2\x39\x16\xe8\x95\xbd\xb3\x4c\x48\xa5\x8f\xcb\x94\xe4\xf9\x5e\x36\x8b\x4b\x68\xa1\xbd\xbc\xb3\x3f\x70\x15\x79\x36\x43\x06\x01\x6d\x25\xef\xd1\x96\x8e\x82\xce\xcc\xae\xef\x8e\xb5\xe6\xc9\xf9\xe9\x95\xd4\xb9\xe8\x67\xb1\xc2\x75\x69\x9f\xcd\x74\x8b\x61\x05\xd6\x8c\xdf\x84\xd0\x04\xee\xd5\x26\x22\x38\x21\x34\xd2\x50\x9b\x8f\xa1\x0b\x0f\x42\x05\x17\xdd\x26\x6a\x30\x68\xda\xa2\xfd\xfd\x8f\x9f\x5a\xbd\xce\xf3\xd2\x81\xd0\xe4\xee\x86\x2a\xc8\xa8\x54\xad\x9c\xd0\x75\x5f\x37\x0b\x2a\x51\x78\xdc\x05\x5a\x8f\x9e\x08\x7d\x4e\x96\x69\xb3\x28\xdb\x58\x94\x34\xd4\xff\x69\xa6\xf5\xf0\xd9\x7f\x2d\x76\x04\xd8\x34\xc9\xc7\xb9\x7a\x79\xec\x34\x54\xa4\xa1\x5e\x14\x3b\xd3\xae\xb2\xe9\x93\x32\x0c\x85\x9f\x3e\x7d\xd8\x0b\x5f\x90\x4a\xbd\x95\x25\x68\x48\x8f\x81\xd2\xee\x78\x6f\xbe\x2d\x32\xdf\x39\x4b\xb3\x92\x00\xb9\xf0\xb5\xbb\xaa\x96\x7b\xbd\xa9\xf1\xae\x21\x71\xe0\xbd\x3d\xc9\x3a\x8e\x77\x58\x56\x4c\xbb\x88\x16\x66\x55\xd6\x0f\x1e\x51\x4b\x68\x92\x7e\x7e\x7b\xf6\xbc\x88\xae\x45\x22\xaf\xe0\x16\x87\xe5\xfb\x1b\x36\xe4\x25\x23\xff\x63\xd4\xb2\x19\x52\x40\x2f\xf5\xe1\x86\xc9\xaa\x17\xbc\x0a\x9a\x4e\xc1\x9f\x06\x12\xbf\x81\x12\x23\x7b\x78\x84\xc2\xd0\x8f\x37\x89\xfa\x1e\x6d\xdd\x70\xa9\xe4\x62\x9a\xfe\xc2\xa5\xfa\x3c\x75\x49\x6a\x5b\x90\x6e\x5f\xa3\xf5\x8c\x9a\x05\x70\x5d\x8c\x4e\x8c\x11\x82\x3f\xed\x63\xbd\x9a\x90\xc9\xdd\xbd\xd0\x8f\x1b\x89\x1e\xfb\x8e\x50\x21\x28\x53\x23\x14\xbe\x95\x91\x7a\x4b\xa2\xb7\x32\x7a\x2b\xc3\x45\xe7\x2c\x7a\xc5\xd0\x4c\x3d\x5d\x63\x47\xdd\x3a\x35\xdf\xbb\x6e\x7a\xb1\x1b\xb4\x0e\x25\xf7\xab\x9c\xb1\xb5\x10\x86\xcd\xad\xd3\xcc\xb1\x94\x27\x30\x21\x02\xa4\x74\xbc\xd9\x35\xc3\xd3\x84\xfe\xb4\xa9\x62\x41\xca\xc8\x3c\x36\x13\x65\x52\x61\x96\xc2\x06\xa6\x42\x28\xc7\xe9\x0d\x65\x90\xf8\xeb\x98\xad\xd9\x04\x8b\xd8\xde\x2c\xfd\x6e\x34\xad\x07\xdc\x9c\xdf\xe9\x83\xd1\xfe\x51\x03\xe8\xd3\xeb\xa3\x14\x6d\x8f\xb8\x95\x65\x27\xfd\x19\x72\xae\x12\x42\xe5\x6d\x6d\x2c\x42\x94\x51\x45\x71\x46\x1f\x20\x29\xb0\xc0\xb9\x6c\x8d\x21\xa4\x0d\xf2\x8b\x31\x26\x9d\x52\x79\x1b\x7b\x7b\x5a\x82\x7e\x25\x05\x89\xa4\x24\xed\x11\x9a\xe3\x71\x67\x91\x5c\xb6\x96\x78\xf9\xbd\xbb\x30\x1d\x08\xff\xaf\x77\xa8\xdd\x97\xba\x0f\x01\x31\xc2\x29\xb4\x2c\x6c\x4a\x67\x2f\x5b\x37\xc3\xb1\xe3\x14\x54\x55\x1d\xde\xf6\x40\x6e\x23\x59\x14\x5d\x23\xb5\x75\xe6\x4e\xe4\x9d\x77\xf0\xca\xfa\xae\xb1\xb7\xeb\x32\xfa\x1e\xd5\x0b\xfc\x86\xa5\xba\x4c\x95\x99\x79\xbb\x09\x05\x42\x38\x4d\x41\xca\x24\xb5\x01\x6b\x3b\x9d\x61\x95\xd0\xa2\x77\x41\xce\xcc\xb8\x6f\x73\xc5\x6e\x70\x95\x2f\xb5\x5e\x49\x0c\x05\xf8\x75\x19\x2b\xbf\xa7\x82\x16\x4a\x76\x1f\xfe\x13\x0b\x89\xc2\x52\x82\x48\x08\x56\x38\x91\x46\x28\xa1\x24\xdc\xde\xae\x53\x13\xa1\x1c\x14\x36\xd7\x2c\xcd\x1a\xb4\x7a\x3b\x0f\x0c\xdc\x2f\xfe\x61\x55\xb5\xa3\x61\x8a\xc0\x09\xcf\x0b\x50\x54\x51\xce\xe2\xd3\x8b\xef\xf1\x15\xe7\xea\x94\xe7\x98\xb2\xce\x3e\xb3\xc4\xeb\x92\x43\xfe\x66\xaf\xbc\x1a\xd6\x45\xc8\xdf\x82\x49\x85\x85\x2a\x8b\xc8\x9a\x1c\x15\x72\xcf\x5c\x94\x82\x32\x77\xc8\xc8\x1c\x69\xa9\x54\x02\x2b\x2e\xd0\x0e\x4e\x15\x9d\xc0\xe1\x14\x64\x33\x95\x9c\xca\x42\x40\x11\xc9\x02\x52\x9b\x1f\x73\x60\xad\x6b\x33\x9f\x7c\x57\x90\x1d\x4b\x09\xea\x0b\x17\x9a\xd2\xfa\xe8\xad\xeb\x09\x1d\x8b\xf8\x33\x96\xd0\x5c\xaa\xd5\x0b\x85\xcc\x3c\x6e\x0c\x92\x37\x7f\x83\xa9\xde\xa1\x3f\xff\x7c\x76\xf9\x25\x10\x9c\xab\xc3\xad\x59\x7a\xc3\xf3\xe2\xdd\x23\x37\x78\x02\xb2\x44\xca\x9b\xa4\x28\x87\x19\x4d\x93\x5b\x98\xb6\xee\xf2\xb6\x2b\xa4\xb1\xfe\x9a\xd9\x57\x17\x16\xf8\x29\x17\x7a\x1e\xf2\x16\xa6\x1a\xc9\xa2\x18\xd7\xc6\x69\xa1\xd2\xd8\xa5\x5e\xa0\xad\xed\xac\xcc\xd7\x2c\xb7\x35\x83\x76\x70\x19\x2c\xc4\x57\x89\xb2\xe5\x09\xbf\xdd\x12\x17\x63\xb7\x2d\x5f\x39\x24\xad\xa2\xea\x08\xaf\xdd\xe1\xbc\xa4\xcb\x69\xbf\x1d\x41\x75\x47\xde\xf3\x74\x3e\x89\xda\xe3\xcb\x26\x99\xc7\x58\xb4\x6f\x41\xa2\x9d\x8a\xbd\x12\x4d\x71\x9e\x1f\x59\x28\x77\x7e\xfc\x87\xf7\x9a\xb9\x90\x96\x94\x33\x10\x28\xb4\x5b\x70\xf6\xf4\xec\x46\x28\xe5\x8c\x41\xaa\x6d\xeb\x30\xa6\x65\x1c\xf3\x65\x03\xac\x59\xd7\x33\x87\x77\x47\x99\xc8\xbb\xc3\x86\x45\xdc\xf0\x71\x9b\x4c\xe6\x50\x68\x0e\xbc\x54\x46\xec\x60\x77\x0e\x63\xee\x8d\x54\xd7\xf3\x9a\x00\xff\xe1\x05\xba\xa5\xb6\x5a\xc5\x12\x2b\x5c\x84\xc7\xc0\x94\xfb\x7e\x84\x42\x73\x97\x1c\xf6\x3b\xf1\x15\xfc\x68\x21\xa5\xbc\x59\xe2\x48\x3b\xae\x29\x64\x89\x0b\xd1\x12\x2f\x0a\x3a\xc1\x0a\x34\xa7\x3d\xe9\x65\x46\x4d\x85\x8d\xde\xfc\x7b\x8d\x7e\x47\x77\xb2\xdf\x1f\xc0\xf4\x94\xf1\x4e\x5f\x76\xee\x38\x3a\x8d\x8c\xdb\xc3\x75\x36\x39\x01\xa9\x28\xc3\xca\x1d\xc3\x4f\x0e\xff\xf8\x63\x39\xd8\xe2\x06\x98\x53\xdf\xe1\x85\x5a\x61\xcc\x1c\x19\xb7\xba\x05\x5e\xaa\xa2\x54\x4b\x0a\xb8\x77\x1f\x2d\x9a\x0c\x9e\xe0\xac\x84\xde\x8d\xe3\x9b\xe7\xfe\x9d\xb3\xd0\xe0\xc5\xbb\xcd\xb3\xa2\xdb\x9f\xaf\x36\xca\x96\xb7\xd7\xb2\xa9\xd3\xb1\x69\x1b\x4d\x9f\xd6\xb5\xaf\x1d\xde\x6b\x3c\x3e\xfb\xb3\xc4\x99\x34\x57\xb3\x21\x0a\x75\x21\xab\x0f\x89\x0b\x27\x0d\x7b\x5e\x4c\xb9\x20\x89\x04\xf5\xe8\x39\xa3\xfd\x22\xcb\x2e\x50\x57\xc5\x78\xc9\xab\xae\xb8\xf5\x3e\x2b\xd6\x58\x4d\x2b\xfa\xd4\x46\x6d\x59\x8d\x89\xfb\x0b\x4b\xbc\x86\x41\x7a\xd0\x19\xd5\x3d\x4e\x1c\xd7\x0f\x94\xb9\x02\xd9\xdf\xdd\xad\xed\x16\x42\xe3\xcf\x15\xe1\x57\x0c\x73\x53\xcc\x9a\x54\x99\x23\x06\xfb\xb6\x51\x41\x5e\x64\xd8\xdc\x62\x65\x8f\x9f\x17\xbd\xf4\xa6\x1a\x98\xb0\xa9\xb1\x24\x31\x43\xaa\xc8\x6a\x76\xab\xd3\x42\xb4\x0f\x7a\x02\x72\xee\xf8\xfd\x3f\x91\x2b\xfa\x93\xf1\x14\x67\x9b\x9c\x71\x81\x31\x7c\x45\x4b\x6c\x7b\x6f\x6e\xd9\x74\x02\xb6\x56\xdd\xbd\x9c\x09\xd7\xb1\x61\x0d\x46\x7f\x5e\xa1\xef\xe1\x76\x6a\x5e\xdb\x2b\x5b\xb0\x56\x1e\xff\x57\x56\xb9\xa7\xd1\x7f\xc3\x4f\x36\x54\x2b\x36\x73\x52\x6f\x39\xca\xc6\x49\xb7\xad\x73\xff\x4f\xaa\xc5\x50\x9d\xf4\xe8\x8f\xaf\x00\x46\x40\x00\xa9\x6d\xd5\xa2\x3e\x4a\x4b\xea\x6d\xb7\x35\x8d\x5d\xc5\x0b\xeb\x64\xed\xbd\x20\xfb\x77\x00\x00\x00\xff\xff\x41\xc4\x2e\x4d\x79\x26\x00\x00")

// FileProvisionedHostTfTmpl is "provisioned_host.tf.tmpl"
var FileProvisionedHostTfTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x94\x92\x3d\xcb\xdb\x40\x10\x84\x7b\xfd\x8a\xc5\xb8\x76\x61\x30\x24\x85\x8b\x7c\x41\xd2\xa4\x08\x81\x94\xe2\xa2\x5b\x45\x8b\x75\xb7\xe2\x76\x2d\x61\x14\xfd\xf7\x17\xdd\xe9\xc3\xc2\x6a\x5e\x77\x37\x3b\x33\xfb\x78\x51\x13\xb8\x25\x21\xf6\x68\xf3\x8a\x45\xa1\xef\xe1\x78\xfa\xce\xa2\xa7\x1f\x5f\xe1\x3f\x54\x45\x2d\x1a\xc8\xff\x83\x61\x80\x3e\x03\x28\xd8\x35\xa8\xa4\xc4\x3e\x27\x0b\xd7\x14\xf8\xb2\xaa\xaf\xb9\x0c\x00\x7d\x4b\x81\xbd\x43\xaf\x4f\xa9\x6f\xab\xba\x9b\xfa\x7b\xa7\xda\x8e\xfe\xf4\x9b\x52\x9f\x47\x75\xd7\xaf\x68\xdc\x6a\x9f\xfd\xbf\xd1\xb8\x5d\xbb\x47\xed\x38\xdc\xe6\xc4\x64\xff\x99\xd4\xdd\xc4\x78\xa1\xd7\x05\xfb\xd7\xca\x00\x02\xa6\xdb\xee\xf1\xff\x9a\x67\xd1\x69\x0a\xa5\x16\x61\x2d\x3e\x1c\xfb\xb8\x2c\x0d\x86\x43\x6c\x73\xac\x98\x1b\x6b\xc3\xe2\x79\xd2\xa2\xa7\xe6\xc2\xd4\x8b\x25\x79\x56\x6d\xaa\x11\xbe\x87\x02\x73\x6f\x1c\x2e\x35\x4f\xda\x70\xc8\xb2\x0c\x46\x56\x2a\x97\x7f\x27\x7f\xc8\x5b\xee\x24\xe1\x76\xe4\x83\x8b\x9f\xc3\x16\x6b\x97\x09\xa0\xe1\xa0\x70\x85\xcb\xc7\x0f\x97\xf8\xae\x54\x1b\x81\x2b\x94\xa6\x16\x8c\x8a\xdc\xa8\xc9\x5b\x0c\x54\x3e\x36\xfa\x5d\x30\xb6\x7e\xb2\x8e\x3c\x89\x06\xa3\x1c\xa6\x52\x23\xd2\x71\xb0\x69\xe9\xfc\x8a\x1b\x87\x84\x8f\xb5\x60\xe2\x15\xa9\xde\x4d\x7b\x3e\x6f\x08\x02\xb3\xa6\x39\x59\xf4\x4a\xfa\xc8\x4b\xaa\x31\x95\x6c\xa4\x0d\x82\xb7\x23\xc1\xf0\x16\x00\x00\xff\xff\xd3\xb7\x99\x54\x67\x03\x00\x00")

// FileRemoteFileTfTmpl is "remote_file.tf.tmpl"
var FileRemoteFileTfTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xcc\x94\xcf\x6e\xd4\x30\x10\xc6\xef\x79\x8a\x91\xd5\x03\x48\x95\xcb\x89\x03\xd2\x1e\x4a\x11\x7f\x2e\x08\xf5\xc2\x01\x21\xcb\xac\x67\xb3\x23\x12\x4f\xe4\x99\x74\xa9\xa2\xbc\x3b\xb2\xb3\x49\x53\x11\xe0\x42\x25\x7c\x8a\xf5\xf9\x9b\xcc\xfc\xc6\x9e\x61\x80\x80\x07\x8a\x08\x26\x61\xcb\x8a\xee\x40\x0d\x1a\x18\xc7\x2a\xa1\x70\x9f\xf6\x08\x26\xf6\x4d\xe3\xe6\xad\x01\xd3\x25\xbe\x23\x21\x8e\x6e\x18\xc0\xbe\x43\xcd\xe6\x49\x75\xd1\xb7\xc5\xee\x44\xb1\x5b\xe4\xbc\x71\xb1\x6f\xbf\x61\x2a\x62\xdf\x35\xec\x83\x81\xa1\x02\xd0\x44\x75\x8d\x49\xca\x06\x80\xa2\xa8\x8f\x7b\x74\x14\x60\x07\xe6\x62\xa8\x99\xeb\x06\xdd\x9e\xdb\xae\x57\x74\xb3\x6e\x7f\xff\x6f\xbb\x8a\x31\x9a\x0a\x60\xac\x2a\x80\x80\x1d\xc6\x20\x8e\x23\xec\xe0\x4b\xf9\x97\x59\x62\x3c\x88\x39\x40\xf6\x7c\xcd\x9e\x61\x80\x8b\x0c\x24\x47\x86\x57\x3b\xb0\xb7\x05\xd2\x5b\x6a\xd0\x5e\x8b\xa0\x7e\xcc\xc2\x58\xe2\x2f\x54\x30\x81\x99\x28\x4e\x15\x0d\x03\xd0\x01\xec\x7b\x16\xb5\x1f\xe4\x33\xc5\xc0\x27\xc9\x26\x28\x6b\xcf\x31\xe2\x5e\x89\xe3\xf9\x7c\x5e\x47\x16\x2d\x1f\x5b\x08\x7c\x08\x09\x45\xfe\x44\xe0\x7c\xa4\x54\x32\x2d\xbd\xef\x70\x8e\x78\xa2\x98\xda\x07\xa9\x17\x4c\xb3\x74\x1d\x5a\x8a\x24\x9a\xbc\x72\x5a\xb9\xa9\x45\xee\xb5\x1c\x79\xf9\x62\xe5\xed\xbc\xc8\x89\x53\xe9\x55\xce\xe7\x86\xdb\x0e\x95\x72\x39\xf6\x96\x59\x3f\xcd\xfa\xb8\xe4\x32\xce\x54\xb0\x11\xfc\x0b\x07\x5f\x63\xd4\xf3\xf7\x0e\xcc\xc1\x37\x82\xe6\x57\x4a\x4f\x02\x6a\x0a\x2a\x72\xdc\x20\x35\x69\x89\x59\x37\x18\xc1\x06\xa6\x44\x77\x5e\xd1\x7d\xc7\xfb\x29\xd3\x7c\x41\x9e\x65\x62\x14\x03\xfe\x00\xfb\xba\xa7\x26\xd8\x1b\x8e\x07\xaa\x73\x9e\x8d\x13\x39\xba\x95\x6d\x79\x98\xe6\xf9\x06\xc9\x18\x66\x90\xe7\x47\x3b\x67\x61\xed\x95\xb5\x57\xc1\xab\xbf\x7a\x74\x99\xe7\x76\x04\x14\xa5\xe8\x0b\xf6\x73\x0b\x57\x77\xfc\xcd\x4a\xcd\x8e\xf2\x96\xc6\xea\x29\x67\xc3\xff\x3b\x14\x2e\x27\xf9\x51\xc5\xf6\x9f\xcd\xc2\xcb\x32\x73\xc6\xaa\x5a\xda\xf9\x33\x00\x00\xff\xff\xbf\xe6\x1b\x66\x9c\x05\x00\x00")

// FileRootModuleTfTmpl is "root_module.tf.tmpl"
var FileRootModuleTfTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x4c\xca\x3b\x0a\x43\x21\x10\x05\xd0\x7e\x56\x71\x11\xcb\x44\xfb\x80\x4d\xd6\x90\x0d\x48\x1c\x82\xe0\x07\xfc\x54\xc3\xec\x3d\xf8\xaa\x57\x1e\x38\x22\x18\xb1\xfd\x18\x76\xe5\xf4\x80\x5d\x1c\x2b\x5e\x01\xd6\xbd\x77\x2e\xc9\x7d\x38\xd6\x09\x55\xaa\x3d\xed\xc2\x30\x27\x3c\x45\xae\x0f\x55\x03\x21\x60\xf6\x3d\xbe\x8c\x00\xe3\xfc\x09\xd3\xdf\x06\x29\x91\x08\xb8\x1d\xfd\x03\x00\x00\xff\xff\x66\x90\xa0\x30\x70\x00\x00\x00")

// FileScriptTfTmpl is "script.tf.tmpl"
var FileScriptTfTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xec\x57\x4b\x8f\xdb\x36\x10\xbe\xeb\x57\x0c\x88\x14\x48\x50\x5b\xdb\xa0\x68\x0f\x01\xf6\x90\x5d\xf4\x05\x04\xc5\xa2\x3e\xf4\x50\x07\x04\x23\x8d\x65\xa2\x14\xa9\x72\xa8\xd8\x0b\x45\xff\xbd\x18\xca\x92\xe5\x67\xd2\xee\x23\x41\xb0\x3c\x18\xa6\x87\xc3\x99\xf9\xf8\xf1\xe3\xb8\x69\x20\xc7\x85\xb6\x08\x82\x32\xaf\xab\x20\xa0\x6d\x13\x8f\xe4\x6a\x9f\x21\x08\x5b\x1b\x23\xfb\xa9\x00\x51\x79\xf7\x5e\x93\x76\x56\x36\x0d\xa4\xbf\x60\x00\xd1\x5b\xa5\x55\x25\xb2\xbb\xa4\x80\xd5\x60\xe6\x89\xb4\x75\xf9\x0e\x7d\x34\xd6\x95\x71\x2a\x17\xd0\x24\x00\xc1\xeb\xa2\x40\x4f\x71\x02\xa0\x2d\x05\x65\x33\x94\x3a\x87\x4b\x10\xcf\x9a\xc2\xb9\xc2\xa0\xcc\x5c\x59\xd5\x01\x65\x6f\x4f\x4f\xc7\x4e\x47\x7b\xb4\x22\x01\x68\x93\x04\x20\xc7\x0a\x6d\x4e\xd2\x59\xb8\x84\xbf\x62\x2c\x31\xec\xb1\x35\xf2\x06\xec\xf3\x96\x7d\x86\x4a\xd1\x83\x58\x68\x83\x62\x93\x65\xd3\x80\x5e\x40\xfa\xab\xa3\x90\xfe\x46\x7f\x6a\x9b\xbb\x15\x31\x6a\x10\x47\xe6\xac\xc5\x2c\x68\x67\x37\xeb\x79\x2c\x1d\x85\xf8\xe5\x58\x59\x2a\xcf\x3d\x12\x9d\xab\x6a\xb3\x24\x66\xd7\x8d\x70\x5b\x61\xbf\xe3\x4a\x5b\x5f\x6e\x4d\x35\xa1\xef\x4d\xaf\xf3\x52\x5b\x4d\xc1\xab\xe0\xfc\xc8\x5b\x97\xe8\xea\x10\x97\xfc\xf8\xdd\xc8\xb7\x52\x44\x2b\xe7\x23\xfe\x9c\xcf\xb5\x2b\x2b\x0c\x9a\xcb\x49\xff\x70\x2e\xdc\xf4\xf6\x76\xc8\x25\x22\xcc\x63\xc3\x99\x3e\x34\xbb\x47\x90\xf8\x83\x2b\x81\xb6\xbd\x50\x44\x18\xe8\x82\x6d\xb3\xc8\xb7\xf4\x4a\x11\x8e\x76\xeb\xd0\xc5\x7f\x06\xfb\x1b\x65\x8b\x5a\x15\x4c\xd0\x25\x1a\x23\xb6\x48\xf3\xb9\x52\xd0\x56\x45\xb0\xbb\x88\x95\xd7\x36\x2c\x40\x5c\xbf\x9a\xcf\xe7\x73\xa3\x16\xce\x17\x28\x3b\x6a\xcb\x6f\x28\x7d\xa7\x82\x80\xe7\x87\xcc\x7c\xb1\x9b\x02\x9a\x98\xd4\x1d\x02\x55\xf4\xf2\x93\x02\xd9\xbc\x8f\x73\x10\xb6\x69\xe0\x59\x4e\x81\xb9\x07\xaf\x2e\x87\x88\x17\xde\xb9\x70\x71\x18\x91\x96\x27\x03\x9e\xa3\xa6\x2a\xd0\x06\x18\x4e\x6d\xa1\x0c\xa1\x38\x24\xee\x83\x70\xb7\xdb\x94\x68\x79\x84\xbc\x9d\x8d\x6b\x3d\x42\x5b\x38\xc2\x5c\xaf\xdf\xab\x80\xf2\x6f\xbc\xed\x32\x65\xdc\x9e\xf3\x51\x69\x9b\xe3\x1a\xd2\xab\x5a\x9b\x3c\xbd\x76\x76\xa1\x0b\xce\xd3\x48\xa2\xa5\x1c\xb9\xc9\xee\x96\xb7\xad\x78\x31\x22\xf7\x3d\x72\xfb\x90\x45\xc3\xf9\xf6\x8b\xc6\x94\x68\x93\x36\x79\x50\x29\xc6\x35\x66\x9f\x5f\x88\x77\xea\x4a\xef\xed\x81\x99\x1c\x15\x72\x8f\xa5\x0b\x38\xdd\x56\xfe\xa4\xe7\x27\xf5\x5c\x5b\xc3\x8d\x41\x7f\x50\xff\x59\x9d\xef\xac\xc7\x93\x71\xdc\x7d\x49\x16\x95\x5b\xa1\x8f\x51\x61\xfa\xbb\xbb\xf1\x2e\x5e\xa4\xe9\x4f\x6b\xcc\x6a\xae\xef\xc6\x19\x9d\xdd\xc2\xd5\x2d\x63\x00\xd3\x9f\xd9\x3a\xbf\xab\x74\xcf\xc5\x5e\x56\x5b\xfd\x66\xba\x3d\xe9\xf8\x17\xab\xe3\xa7\x49\x2d\xb2\x65\xe9\x72\xf8\x76\x0d\x7b\x8a\x3c\xd9\xa1\xf2\xc8\x32\x92\xf8\xd7\xbe\x98\x05\xaf\x6d\x31\x76\x78\xfb\xf8\x5a\xbe\x52\x3a\x7c\x7d\x5a\x1e\x75\x7a\x50\x72\x3e\x04\xa3\x6d\xbd\x96\x64\x10\x2b\xc9\x34\xf2\x7c\xa5\xbe\xef\x40\x66\xfb\xaa\xd3\xef\xfd\x15\x2f\x7f\xe0\x25\xc9\xa0\xf6\x6f\x5c\xa6\xcc\xa1\xdc\xef\xbc\x15\x86\xd7\xec\x3c\x15\x1f\x7d\x2c\xf8\x6e\x96\xa5\xb2\x11\xea\x59\x50\x3e\x4c\x67\x9c\xc8\xc9\xd4\xce\x76\x9d\x67\xf6\x3a\x84\xe1\x54\x5b\xc9\x0c\x08\xe8\x2b\x8f\x01\x3d\x9f\x94\xb8\x61\xe1\x9c\x45\xb9\x9e\x80\x98\x5e\x77\x51\x44\x47\xda\x0d\x8e\xe3\x54\xee\x15\x14\xba\x17\x38\xe8\x7f\x00\xd1\x57\xd6\xfd\xf2\xb0\xd7\xf1\x71\x6e\xe2\xe4\x31\xaf\x62\x14\x98\xa7\xa6\xea\x73\x37\x55\x00\x22\x47\x03\x0f\xda\x5a\x7d\x7a\x8c\xb3\xff\x72\x9f\x1a\xa5\xaf\xad\x51\xf2\x25\x4c\x17\x7b\x5d\x12\x7c\xf8\x00\xc1\xd7\xf8\x91\xe6\x67\x3b\x4f\xfe\x0d\x00\x00\xff\xff\x3e\x2f\xbd\x02\x74\x14\x00\x00")

func init() {
	err := CTX.Err()
	if err != nil {
		panic(err)
	}

	var f webdav.File

	var rb *bytes.Reader
	var r *gzip.Reader

	rb = bytes.NewReader(FileCommandTfTmpl)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "command.tf.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	rb = bytes.NewReader(FileDNSRecordTfTmpl)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "dns_record.tf.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	rb = bytes.NewReader(FileInfraTfTmpl)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "infra.tf.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	rb = bytes.NewReader(FileProvisionedHostTfTmpl)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "provisioned_host.tf.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	rb = bytes.NewReader(FileRemoteFileTfTmpl)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "remote_file.tf.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	rb = bytes.NewReader(FileRootModuleTfTmpl)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "root_module.tf.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	rb = bytes.NewReader(FileScriptTfTmpl)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "script.tf.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	Handler = &webdav.Handler{
		FileSystem: FS,
		LockSystem: webdav.NewMemLS(),
	}

}

// Open a file
func (hfs *HTTPFS) Open(path string) (http.File, error) {

	f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// ReadFile is adapTed from ioutil
func ReadFile(path string) ([]byte, error) {
	f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(make([]byte, 0, bytes.MinRead))

	// If the buffer overflows, we will get bytes.ErrTooLarge.
	// Return that as an error. Any other panic remains.
	defer func() {
		e := recover()
		if e == nil {
			return
		}
		if panicErr, ok := e.(error); ok && panicErr == bytes.ErrTooLarge {
			err = panicErr
		} else {
			panic(e)
		}
	}()
	_, err = buf.ReadFrom(f)
	return buf.Bytes(), err
}

// WriteFile is adapTed from ioutil
func WriteFile(filename string, data []byte, perm os.FileMode) error {
	f, err := FS.OpenFile(CTX, filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}

// WalkDirs looks for files in the given dir and returns a list of files in it
// usage for all files in the b0x: WalkDirs("", false)
func WalkDirs(name string, includeDirsInList bool, files ...string) ([]string, error) {
	f, err := FS.OpenFile(CTX, name, os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}

	fileInfos, err := f.Readdir(0)
	if err != nil {
		return nil, err
	}

	err = f.Close()
	if err != nil {
		return nil, err
	}

	for _, info := range fileInfos {
		filename := path.Join(name, info.Name())

		if includeDirsInList || !info.IsDir() {
			files = append(files, filename)
		}

		if info.IsDir() {
			files, err = WalkDirs(filename, includeDirsInList, files...)
			if err != nil {
				return nil, err
			}
		}
	}

	return files, nil
}
