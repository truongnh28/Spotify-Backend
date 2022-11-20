package cache

import (
	"sync"
)

type MemoryCache interface {
	GetScanningVersion() (project string, mapRevisionsScanned map[string]bool, ok bool)
	SetScanningVersion(project string, mapRevisionsScanned map[string]bool)
	RemoveScanningVersion()
}

type memoryCacheImpl struct {
	project            string
	mapRevisionScanned map[string]bool
	lock               sync.RWMutex
}

func (cache *memoryCacheImpl) GetScanningVersion() (project string, mapRevisionsScanned map[string]bool, ok bool) {
	cache.lock.Lock()
	defer cache.lock.Unlock()
	if cache.project != "" {
		return cache.project, cache.mapRevisionScanned, true
	}
	return "", nil, false
}

func (cache *memoryCacheImpl) SetScanningVersion(project string, mapRevisionsScanned map[string]bool) {
	cache.lock.Lock()
	defer cache.lock.Unlock()
	cache.project = project
	cache.mapRevisionScanned = mapRevisionsScanned
}

func (cache *memoryCacheImpl) RemoveScanningVersion() {
	cache.lock.Lock()
	defer cache.lock.Unlock()
	cache.project = ""
	cache.mapRevisionScanned = nil
}

func NewMemoryCache() MemoryCache {
	return &memoryCacheImpl{
		project:            "",
		mapRevisionScanned: make(map[string]bool, 0),
		lock:               sync.RWMutex{},
	}
}
