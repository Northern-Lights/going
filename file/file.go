package file

import (
	"os"
	"strings"
)

// Append opens an existing file or creates a new one and appends data to it.
func Append(filepath string, data []byte) (err error) {
	flags := os.O_APPEND | os.O_CREATE | os.O_RDWR
	if osfile, e1 := os.OpenFile(filepath, flags, 0644); e1 == nil {
		defer osfile.Close()
		if _, e2 := osfile.Write(data); e2 != nil {
			err = e2
		}
	} else {
		err = e1
	}
	return
}

// Exists returns true if the file at the given path exists, false otherwise
func Exists(filepath string) (exists bool) {
	if _, err := os.Stat(filepath); err == nil {
		exists = true
	}
	return
}

// Read returns a byte slice of all bytes in the named file.
func Read(filepath string) (data []byte, err error) {
	data, err = ReadN(filepath, -1)
	return
}

// ReadLines returns a slice of strings, one for each line in the file.
func ReadLines(filepath string) (lines []string, err error) {
	if contents, e1 := Read(filepath); e1 == nil {
		lines = strings.Split(string(contents), "\n")
	} else {
		err = e1
	}
	return
}

// ReadN returns a byte slice of the first nBytes of the named file.
// If nBytes is negative, the entire file size is used.
func ReadN(filepath string, nBytes int64) (data []byte, err error) {
	if osfile, e1 := os.Open(filepath); e1 == nil {
		defer osfile.Close()
		if finfo, e2 := os.Stat(filepath); e2 == nil {
			if fsize := finfo.Size(); nBytes < fsize && nBytes >= 0 {
				data = make([]byte, nBytes)
			} else {
				data = make([]byte, fsize)
			}
			nRead, e3 := osfile.Read(data)
			err = e3
			if int64(nRead) < nBytes {
				data = data[:nRead]
			}
		} else {
			err = e2
		}
	} else {
		err = e1
	}
	return
}

// Size returns the size of the file at the given path.
// The returned size will be negative if an error occurs.
func Size(filepath string) (size int64) {
	if osinfo, err := os.Stat(filepath); err == nil {
		size = osinfo.Size()
	} else {
		size = -1
	}
	return
}

// Write tramples an existing file or creates a new one and writes data to it.
func Write(filepath string, data []byte) (err error) {
	if osfile, e1 := os.Create(filepath); e1 == nil {
		defer osfile.Close()
		if _, e2 := osfile.Write(data); e2 != nil {
			err = e2
		}
	} else {
		err = e1
	}
	return
}
