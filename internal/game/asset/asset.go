package asset

import (
	"bytes"
	"embed"
	"path"
	"sync"

	"github.com/charmbracelet/log"
)

var (
	//go:embed image/*.txt
	images       embed.FS
	loadedImages = map[string][]string{}
	imageMux     sync.RWMutex
)

// Image loads a named image. If the image has been loaded previously, it is returned from a global cache.
func Image(logger *log.Logger, name string) []string {
	logger = logger.With("name", name)

	imageMux.RLock()
	if img, ok := loadedImages[name]; ok {
		imageMux.RUnlock()
		return img
	}
	imageMux.RUnlock()

	p := path.Join("image", name)
	f, err := images.ReadFile(p)
	if err != nil {
		logger.With("error", err).Error("failed to load image")
		return nil
	}

	lines := bytes.Split(f, []byte("\n"))
	img := make([]string, 0)
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		img = append(img, string(line))
	}

	logger.Debug("loaded image")
	imageMux.Lock()
	loadedImages[name] = img
	imageMux.Unlock()

	return img
}
