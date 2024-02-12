package paste

import (
	"pastevault.com/internal/db"
	pb "pastevault.com/internal/proto"
)

func Migrate() error {
	if err := db.DB.AutoMigrate(&pb.Paste{}); err != nil {
		return err
	}

	return nil
}
