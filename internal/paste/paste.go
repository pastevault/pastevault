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
	Expiration string `form:"expiration" json:"expiration"`
	Content    string `form:"content" json:"content" binding:"required"`
}

func init() {
	if err := db.DB.AutoMigrate(&Paste{}); err != nil {
		panic(err)
	}
}

func NewPasteHandler(c *gin.Context) {
	var p Paste
	if err := c.ShouldBind(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//	Generate a UUID for the paste
	p.UUID = uuid.New().String()

	// Set expiration to 1 day from now
	p.Expiration = time.Now().Add(24 * time.Hour).Format(time.RFC3339)

	//	Store the paste in the database
	db.DB.Create(&p)

	//	Return the UUID of the paste
	c.JSON(201, gin.H{"uuid": p.UUID, "expiration": p.Expiration})
}
