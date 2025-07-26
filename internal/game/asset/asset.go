package asset

import (
	"bytes"
	"embed"
	"path"
	"sync"
)

var (
	//go:embed image/logo.txt
	images       embed.FS
	loadedImages = map[string][]string{}
	imageMux     sync.RWMutex
)

// Image loads a named image. If the image has been loaded previously, it is returned from a global cache.
func Image(name string) []string {
	imageMux.RLock()
	if img, ok := loadedImages[name]; ok {
		imageMux.RUnlock()
		return img
	}
	imageMux.RUnlock()

	p := path.Join("image", name)
	f, err := images.ReadFile(p)
	if err != nil {
		// TODO(davidsbond): something here when the image doesn't exist. Maybe just panic?
		return nil
	}

	lines := bytes.Split(f, []byte("\n"))
	img := make([]string, len(lines))
	for i, line := range lines {
		img[i] = string(line)
	}

	imageMux.Lock()
	loadedImages[name] = img
	imageMux.Unlock()
	
	return img
}
