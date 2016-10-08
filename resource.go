package goo

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"sync"
)

var (
	resourceData        = map[string][]byte{}
	resourceCompression = map[string]bool{}
	resourceMutex       = &sync.Mutex{}
)

func DeleteResource(name string) bool {
	resourceMutex.Lock()

	defer resourceMutex.Unlock()

	var _, ok = resourceData[name]

	delete(resourceData, name)
	delete(resourceCompression, name)

	return ok
}

func Resource(name string) ([]byte, bool) {
	resourceMutex.Lock()

	defer resourceMutex.Unlock()

	var d, ok = resourceData[name]

	if !ok {
		return nil, false
	}

	if resourceCompression[name] {
		var r, err = gzip.NewReader(bytes.NewBuffer(d))

		if err != nil {
			panic(err)
		}

		if d, err = ioutil.ReadAll(r); err != nil {
			panic(err)
		}

		if err = r.Close(); err != nil {
			panic(err)
		}
	}

	return d, true
}

func SetResource(name string, file []byte, compressed bool) {
	resourceMutex.Lock()

	defer resourceMutex.Unlock()

	resourceData[name] = file
	resourceCompression[name] = compressed
}
