package paste

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"math/rand"
	"net/http"
	"pastevault.com/internal/db"
	. "pastevault.com/internal/logger"
	pb "pastevault.com/internal/proto"
	"time"
)

type Paste struct {
	Id         string
	Expiration time.Time
	Content    string
	Encrypted  bool `gorm:"default:false"`
}

type pasteRequest struct {
	Expiration string `json:"expiration"`
	Content    string `json:"content" binding:"required"`
	Encrypted  bool   `json:"encrypted"`
}

func generateId(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func GetPasteById(id string) (*pb.Paste, error) {
	start := time.Now()
	var p *pb.Paste
	if err := db.DB.Where("id = ?", id).First(&p).Error; err != nil {
		Logger.Error("Error getting paste by id: ", "paste.GetPasteById", err)
		return nil, err
	}

	Logger.Info("GetPasteById", "time", time.Since(start))
	fmt.Println(p)
	return p, nil
}

func CreatePaste(p *pb.PasteRequest) (*pb.Paste, error) {
	start := time.Now()
	paste := &pb.Paste{
		Content:   p.Content,
		Encrypted: p.Encrypted,
	}

	paste.Id = generateId(8)

	now := time.Now()
	switch p.Expiration {
	case "5m":
		paste.Expiration = now.Add(time.Minute * 5).Format(time.RFC3339)
	case "1h":
		paste.Expiration = now.Add(time.Hour).Format(time.RFC3339)
	case "1d":
		paste.Expiration = now.Add(time.Hour * 24).Format(time.RFC3339)
	case "1w":
		paste.Expiration = now.Add(time.Hour * 24 * 7).Format(time.RFC3339)
	case "1m":
		paste.Expiration = now.Add(time.Hour * 24 * 30).Format(time.RFC3339)
	case "1y":
		paste.Expiration = now.Add(time.Hour * 24 * 365).Format(time.RFC3339)
	default:
		paste.Expiration = now.Add(time.Hour * 24).Format(time.RFC3339)
	}

	if err := db.DB.Create(&paste).Error; err != nil {
		Logger.Error("Error creating paste: ", "paste.CreatePaste", err)
		return nil, err
	}

	Logger.Info("CreatePaste", "time", time.Since(start))
	return paste, nil
}

func GetPasteHandler(c *gin.Context) {
	id := c.Param("uuid")
	var p Paste
	if err := db.DB.Where("uuid = ?", id).First(&p).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Paste not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"uuid": p.Id, "expiration": p.Expiration, "content": p.Content, "encrypted": p.Encrypted})
}

func GetRawPasteHandler(c *gin.Context) {
	id := c.Param("uuid")
	var p Paste
	if err := db.DB.Where("uuid = ?", id).First(&p).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Paste not found"})
		return
	}

	c.String(http.StatusOK, p.Content)

}

func NewPasteHandler(c *gin.Context) {
	var r pasteRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p := Paste{
		Content:   r.Content,
		Encrypted: r.Encrypted,
	}

	//	Generate a UUID for the paste
	p.Id = uuid.New().String()

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
	c.JSON(201, gin.H{"id": p.Id, "expiration": p.Expiration})
}
