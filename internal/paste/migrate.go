package paste

import "pastevault.com/internal/db"

func Migrate() error {
	if err := db.DB.AutoMigrate(&Paste{}); err != nil {
		return err
	}

	return nil
}
