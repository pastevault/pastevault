package expired

import (
	"pastevault.com/internal/db"
	"pastevault.com/internal/paste"
	"time"
)

func Delete() {
	for {
		var pastes []paste.Paste
		db.DB.Where("expiration < ?", time.Now()).Delete(&pastes)
		time.Sleep(5 * time.Minute)
	}
}
