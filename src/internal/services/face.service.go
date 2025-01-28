package services

import (
	"bytes"
	"facial-recognition/src/internal/models"
	"image"
	_ "image/jpeg" // Import JPEG decoder
	_ "image/png"  // Import PNG decoder
	"io"
	"log"
	"net/http"
	"os"

	pigo "github.com/esimov/pigo/core"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func DetectFaces(c *gin.Context) *models.FaceRegistrationResponse {
	var requestBody struct {
		Image string `json:"image"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return nil
	}

	// Fetch the image from the URL
	resp, err := http.Get(requestBody.Image)
	if err != nil || resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to fetch image"})
		return nil
	}
	defer resp.Body.Close()

	// Read the image from the response body
	imgData, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read image data"})
		return nil
	}

	// Decode to image.Image
	img, _, err := image.Decode(bytes.NewReader(imgData))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to decode image"})
		return nil
	}

	// Load facefinder binary file
	facefinderData, err := os.ReadFile("cascade/facefinder")
	if err != nil {
		log.Printf("Error reading facefinder file: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load facefinder"})
		return nil
	}

	// Initialize Pigo
	classifier, err := pigo.NewPigo().Unpack(facefinderData)
	if err != nil {
		log.Printf("Error unpacking Pigo classifier: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Face detection failed"})
		return nil
	}

	// Convert image to grayscale
	grayImg := pigo.RgbToGrayscale(img)

	cols, rows := img.Bounds().Dx(), img.Bounds().Dy()
	cascadeParams := pigo.CascadeParams{
		MinSize:     50,
		MaxSize:     500,
		ShiftFactor: 0.1,
		ScaleFactor: 1.1,
		ImageParams: pigo.ImageParams{
			Pixels: grayImg,
			Rows:   rows,
			Cols:   cols,
			Dim:    cols,
		},
	}

	// Detect faces
	detections := classifier.RunCascade(cascadeParams, 0)
	detections = classifier.ClusterDetections(detections, 0.2)

	if len(detections) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No face detected"})
		return nil
	}

	// Register first detected face
	bestFace := detections[0]
	faceID := uuid.New().String()
	faceData := models.Face{
		ID:     faceID,
		X:      bestFace.Col,
		Y:      bestFace.Row,
		Width:  bestFace.Scale,
		Height: bestFace.Scale,
	}

	response := &models.FaceRegistrationResponse{
		Message: "Face registered successfully",
		Face:    faceData,
	}

	return response
}
