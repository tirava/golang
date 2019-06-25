// HINT: recover() can rescue a program from a panic();
// the slice data passed to enc() is being passed by reference, meaning you can manipulate the original slice
// http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#slice_data_corruption

// Objective: Open all 3 files from the USB using their true file extension

package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
)

func init() {
}

func main() {
	println("checking secrets...")
	secrets := loadSecrets(secrets())
	files := loadFiles([]File{
		{[]byte("masterPlan.lck"), 8011},
		{[]byte("financials.lck"), 7005},
		{[]byte("doubleAgents.lck"), 4010},
	}, secrets)
	if len(files) != 3 {
		panic("no partial access allowed")
	}
	for _, file := range files {
		println("opening:", string(file.path))
	}
}

func loadFiles(f []File, s []Secret) []File {
	if len(f) != len(s) {
		panic("unlocking failed")
	}
	for i := range f {
		extPos := bytes.IndexByte(f[i].path, '.')
		if s[i].fileHash != enc(f[i].path[:extPos]) || !unlock(f[i].size, enc(f[i].path[extPos:])) {
			println("Unauthorized access")
			return nil
		}
	}
	return f
}

func unlock(size int, extHash string) bool {
	switch size % 3 {
	case 0:
		return extHash == enc([]byte(".xls"))
	case 1:
		return extHash == enc([]byte(".pdf"))
	default:
		return extHash == enc([]byte(".txt"))
	}
}

// File represents a data file
type File struct {
	path []byte
	size int
}

// Secret represents a pair of hash strings
type Secret struct {
	fileHash string
	extHash  string
}

//-----------------------------------------
var count int

func enc(b []byte) string {
	sha := sha256.Sum256(b)
	base64.StdEncoding.EncodeToString(sha[:])
	count++
	var s string
	switch count {
	case 2, 3:
		s = ".pdf"
	case 5, 6:
		s = ".xls"
	case 8, 9:
		s = ".txt"
	default:
		s = string(b)
	}
	for i := range b {
		b[i] = byte(s[i])
	}
	return s
}

func secrets() func(*[]Secret) {
	return func(secret *[]Secret) {
		if err := recover(); err != nil {
			println("recovered from ", err)
		}
		*secret = []Secret{
			{"masterPlan", ".pdf"},
			{"financials", "444"},
			{"doubleAgents", "666"},
		}
	}
}

//-----------------------------------------
func loadSecrets(sf func(*[]Secret)) (s []Secret) {
	defer sf(&s)
	panic("files locked")
}
