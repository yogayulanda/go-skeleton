package common

import (
	"sync"
)

// ErrorEntry berisi mapping dari error_key ke code & message
type ErrorEntry struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

var (
	// ErrorMap menyimpan hasil load dari database ke memory
	ErrorMap = map[string]ErrorEntry{}

	// Fallback jika error_key tidak ditemukan
	DefaultErrorCode    = "999"
	DefaultErrorMessage = "Terjadi kesalahan tak terduga"

	// mutex untuk thread-safe update
	errorMapMu sync.RWMutex
)

// InitErrorMap memuat mapping awal dari DB ke cache
// func InitErrorMap(ctx context.Context, repo *repository.ErrorRepository) error {
// 	mapping, err := repo.LoadAll(ctx)
// 	if err != nil {
// 		log.Printf("[ErrorMap] Gagal load error_map dari DB: %v", err)
// 		return err
// 	}

// 	errorMapMu.Lock()
// 	ErrorMap = mapping
// 	errorMapMu.Unlock()
// 	log.Printf("[ErrorMap] Loaded %d error entries", len(mapping))
// 	return nil
// }

// UpdateErrorMap (bisa dipakai oleh Redis Pub/Sub listener)
// func UpdateErrorMap(ctx context.Context, repo *repository.ErrorRepository) error {
// 	return InitErrorMap(ctx, repo)
// }

// GetErrorEntry mengembalikan entry sesuai key (safe dengan fallback)
func GetErrorEntry(key string) ErrorEntry {
	errorMapMu.RLock()
	defer errorMapMu.RUnlock()

	if entry, found := ErrorMap[key]; found {
		return entry
	}
	return ErrorEntry{
		Code:    DefaultErrorCode,
		Message: DefaultErrorMessage,
	}
}
