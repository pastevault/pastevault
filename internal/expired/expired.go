package expired

import (
	"time"

	"pastevault.com/internal/db"
	pb "pastevault.com/internal/proto"
)

func Delete() {
	for {
		var pastes []pb.Paste
		db.DB.Where("expiration < ?", time.Now()).Delete(&pastes)
		time.Sleep(5 * time.Minute)
	}
}
