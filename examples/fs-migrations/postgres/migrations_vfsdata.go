// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package postgres

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// Migrations statically implements the virtual filesystem provided to vfsgen.
var Migrations = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2020, 10, 10, 0, 10, 21, 356686441, time.UTC),
		},
		"/00001_create_users_table.sql": &vfsgen۰CompressedFileInfo{
			name:             "00001_create_users_table.sql",
			modTime:          time.Date(2020, 10, 10, 0, 10, 18, 288000000, time.UTC),
			uncompressedSize: 243,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x5c\x8f\xc1\x4e\x86\x30\x10\x84\xef\xfb\x14\x73\x2b\xc4\x92\xe8\x99\x53\x95\x1e\x88\xb5\x90\x52\x48\x38\x1a\x6d\x14\x8d\xd4\xd0\x82\x3e\xbe\xa1\x98\xf0\xe7\x3f\x6c\xb2\x33\xfb\x65\x27\x53\x14\xb8\x79\xf3\x3e\x38\xf4\xdf\xf4\x60\xa4\xb0\x12\x56\xdc\x2b\x89\x35\xb8\x25\x20\x23\x00\x98\x5e\x31\xcd\x11\xba\xb1\xd0\xbd\x52\x68\x4d\xfd\x24\xcc\x88\x47\x39\xf2\x04\xec\xf0\xfc\xfc\xe5\x10\xdd\x6f\x3c\xac\x2b\x19\xd6\x13\xa0\xbc\x24\xaa\x75\x27\x8d\x45\xad\x6d\xf3\x9f\x35\x08\xd5\xcb\x8e\xb2\x5b\x0e\xb6\x78\x1f\x19\x07\x4b\x93\x73\xca\xee\x38\xd8\xe6\x3f\xa2\x7b\x79\xdf\xa6\xe8\x3e\xf7\xc3\x70\xe8\xb4\x26\x6f\xff\x7b\x36\xaa\xfc\xcf\x4c\x95\x69\xda\xcb\x46\x25\xfd\x05\x00\x00\xff\xff\x86\xed\xe0\x21\xf3\x00\x00\x00"),
		},
		"/00002_rename_root.go": &vfsgen۰CompressedFileInfo{
			name:             "00002_rename_root.go",
			modTime:          time.Date(2020, 10, 10, 1, 10, 49, 134575632, time.UTC),
			uncompressedSize: 451,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xac\xce\x41\x4b\xfc\x30\x10\x05\xf0\x73\xe6\x53\xcc\x3f\x97\xb6\x7f\x96\xb6\x8a\x07\x51\x7a\x58\xd8\x82\x17\x41\xb4\x8b\x47\x49\xdb\x6c\x1d\x6c\x93\xee\x24\xc5\x8a\xec\x77\x97\xc6\x75\x11\xbc\x89\xb7\x97\x79\xe4\xf1\x1b\x55\xf3\xa2\x3a\x8d\x83\x22\x03\x40\xc3\x68\xd9\x63\x0c\x42\xb6\xca\xab\x5a\x39\x9d\xb9\x7d\x2f\x01\x84\xec\xc8\x3f\x4f\x75\xda\xd8\x21\xab\x59\xb5\x6f\x67\x17\xf9\x65\xd6\x59\xeb\xb4\x84\x04\x60\x37\x99\x06\xc9\x90\x8f\x13\x7c\x07\x11\x8a\x74\xdd\xb6\xb7\xd4\xb1\xf2\x64\x4d\xbc\x1d\xf3\x3c\xcf\xcf\x57\xb8\xb1\xaf\x26\xc4\x04\x0e\xc7\x8f\xc7\x2e\xf6\x33\xfe\x77\xfb\x3e\xad\xe6\x04\x35\xb3\xe5\x65\xeb\x69\xb5\x64\xbc\x2a\xd0\xcf\x69\x39\xeb\x26\x96\xdb\xbb\xcd\xba\x2a\x71\x72\x9a\x1d\x3e\x94\x55\x48\x46\x0d\xba\x88\x54\x3b\x90\x89\xf0\xf1\xa6\xbc\x2f\xbf\x9d\xd9\x5a\x1f\x5d\xcb\x04\x04\xed\xc2\xdc\xbf\x02\x0d\xf5\xcb\xbe\x60\xed\x27\x36\xcb\x15\xc4\x01\xbe\x9e\x86\xfa\x93\xef\x44\xfe\x0b\x61\xa0\xfc\x00\x7e\xba\x7f\x21\xfc\x08\x00\x00\xff\xff\x76\x37\xf8\x82\xc3\x01\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/00001_create_users_table.sql"].(os.FileInfo),
		fs["/00002_rename_root.go"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}