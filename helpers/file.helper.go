// helpers/file_helper.go
package helpers

import (
	"path/filepath"
	"strings"
	"strconv"
	"time"
)

// GenerateUniqueFilename generates a unique filename based on the original filename
func GenerateUniqueFilename(originalFilename string) string {
	ext := filepath.Ext(originalFilename)
	filename := strings.TrimSuffix(originalFilename, ext)

	uniqueID := time.Now().UnixNano()
	uniqueFilename := filename + "_" + strconv.FormatInt(uniqueID, 10) + ext

	return uniqueFilename
}
