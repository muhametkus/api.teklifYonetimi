package utils

import "sync"

type PDFCache struct {
	mu    sync.RWMutex
	cache map[string][]byte
}

var pdfCache = &PDFCache{
	cache: make(map[string][]byte),
}

// GetPDF
func GetPDF(key string) ([]byte, bool) {
	pdfCache.mu.RLock()
	defer pdfCache.mu.RUnlock()

	data, ok := pdfCache.cache[key]
	return data, ok
}

// SetPDF
func SetPDF(key string, data []byte) {
	pdfCache.mu.Lock()
	defer pdfCache.mu.Unlock()

	pdfCache.cache[key] = data
}
