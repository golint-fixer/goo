package goo

import (
	"bytes"
	"compress/gzip"
	"sync"
)

var (
	resourceData        = map[string][]byte{}
	resourceCompression = map[string]bool{}
	resourceMutex       = &sync.Mutex{}
)

func AddResource(name string, file []byte, compressed bool) {
	resourceMutex.Lock()

	defer resourceMutex.Unlock()

	resourceData[name] = file
	resourceCompression[name] = compressed
}

func GetResource(name string) ([]byte, bool) {
	resourceMutex.Lock()

	defer resourceMutex.Unlock()

	var d, ok = resourceData[name]

	if !ok {
		return nil, false
	}

	if resourceCompression[name] {
		var b bytes.Buffer
		var r, err = gzip.NewReader(&b)

		if err != nil {
			panic(err)
		}

		if _, err := r.Read(d); err != nil {
			panic(err)
		}

		if err := r.Close(); err != nil {
			panic(err)
		}

		d = b.Bytes()
	}

	return d, true
}

func RemoveResource(name string) bool {
	resourceMutex.Lock()

	defer resourceMutex.Unlock()

	var _, ok = resourceData[name]

	delete(resourceData, name)
	delete(resourceCompression, name)

	return ok
}
