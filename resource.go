package goo

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"sync"
)

var (
	resourceCompression = map[string]map[string]bool{}
	resourceData        = map[string]map[string][]byte{}
	resourceMutex       = &sync.Mutex{}
)

func AddResource(package_, name string, data []byte, compressed bool) {
	resourceMutex.Lock()
	defer resourceMutex.Unlock()

	if nameData, ok := resourceData[package_]; ok {
		if _, ok := nameData[name]; ok {
			panic(ok)
		}

		nameData[name] = data
		resourceCompression[package_][name] = compressed
	} else {
		resourceData[package_] = map[string][]byte{name: data}
		resourceCompression[package_] = map[string]bool{name: compressed}
	}
}

func Resource(package_, name string) []byte {
	resourceMutex.Lock()
	defer resourceMutex.Unlock()

	var nameData, ok = resourceData[package_]

	if !ok {
		panic(ok)
	}

	data, ok := nameData[name]

	if !ok {
		panic(ok)
	}

	if resourceCompression[package_][name] {
		var r, err = gzip.NewReader(bytes.NewBuffer(data))

		if err != nil {
			panic(err)
		}

		if data, err = ioutil.ReadAll(r); err != nil {
			panic(err)
		}

		if err = r.Close(); err != nil {
			panic(err)
		}
	}

	return data
}
