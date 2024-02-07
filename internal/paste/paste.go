package paste

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"pastevault.com/internal/db"
	"time"
)

type Paste struct {
	UUID       string
	Expiration time.Time
	Content    string
	Encrypted  bool `gorm:"default:false"`
}

type pasteRequest struct {
	Expiration string `json:"expiration"`
	Content    string `json:"content" binding:"required"`
	Encrypted  bool   `json:"encrypted"`
}

func init() {
	if err := db.DB.AutoMigrate(&Paste{}); err != nil {
		panic(err)
	}
}

func NewPasteHandler(c *gin.Context) {
	var r pasteRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p := Paste{
		Content: r.Content,
	}

	//	Generate a UUID for the paste
	p.UUID = uuid.New().String()

	//	Set the expiration date
	now := time.Now()
	switch r.Expiration {
	case "5m":
		p.Expiration = now.Add(time.Minute * 5)
	case "1h":
		p.Expiration = now.Add(time.Hour)
	case "1d":
		p.Expiration = now.Add(time.Hour * 24)
	case "1w":
		p.Expiration = now.Add(time.Hour * 24 * 7)
	case "1m":
		p.Expiration = now.Add(time.Hour * 24 * 30)
	case "1y":
		p.Expiration = now.Add(time.Hour * 24 * 365)
	default:
		p.Expiration = now.Add(time.Hour * 24)
	}

	//	Store the paste in the database
	db.DB.Create(&p)

	//	Return the UUID of the paste
	c.JSON(201, gin.H{"uuid": p.UUID, "expiration": p.Expiration})
}
