package services

import (
	"bytes"
	"encoding/base64"
	"facial-recognition/src/internal/handlers"
	"facial-recognition/src/internal/models"
	pkg "facial-recognition/src/pkg/constants"
	"image"
	"image/jpeg"
	_ "image/jpeg" // Import JPEG decoder
	_ "image/png"  // Import PNG decoder
	"io"
	"log"
	"net/http"
	"os"

	"github.com/Kagami/go-face"
	pigo "github.com/esimov/pigo/core"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func DetectFaces(c *gin.Context) *models.FaceDetectionResponse {
	var requestBody struct {
		Image string `json:"image"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		handlers.BadRequestError(c, pkg.StatusBadRequest)
		return nil
	}

	// Fetch the image from the URL
	resp, err := http.Get(requestBody.Image)
	if err != nil || resp.StatusCode != http.StatusOK {
		handlers.BadRequestError(c, pkg.StatusBadRequest, "Failed to fetch image")
		return nil
	}
	defer resp.Body.Close()

	// Read the image from the response body
	imgData, err := io.ReadAll(resp.Body)
	if err != nil {
		handlers.InternalServerError(c, pkg.StatusInternalServerError, "Failed to read image")
		return nil
	}

	// Decode to image.Image
	img, _, err := image.Decode(bytes.NewReader(imgData))
	if err != nil {
		handlers.InternalServerError(c, pkg.StatusInternalServerError, "Failed to decode image")
		return nil
	}

	// Load facefinder binary file
	facefinderData, err := os.ReadFile("cascade/facefinder")
	if err != nil {
		log.Printf("Error reading facefinder file: %v", err)
		handlers.InternalServerError(c, pkg.StatusInternalServerError, "Failed to read facefinder file")
		return nil
	}

	// Initialize Pigo
	classifier, err := pigo.NewPigo().Unpack(facefinderData)
	if err != nil {
		log.Printf("Error unpacking Pigo classifier: %v", err)
		handlers.InternalServerError(c, pkg.StatusInternalServerError, "Failed to unpack Pigo classifier")
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
		handlers.NotFoundError(c, pkg.StatusNotFound, "No faces detected")
		return nil
	}

	// Register detected faces
	var faces []models.Face
	for _, det := range detections {
		faceID := uuid.New().String()
		faceData := models.Face{
			ID:     faceID,
			X:      det.Col,
			Y:      det.Row,
			Width:  det.Scale,
			Height: det.Scale,
		}
		faces = append(faces, faceData)
	}

	response := &models.FaceDetectionResponse{
		Faces: faces,
	}

	return response
}

func RegisterFace(c *gin.Context) *models.FaceRegistrationResponse {
	var (
		recognizer *face.Recognizer // Khởi tạo recognizer toàn cục
	)

	var requestBody struct {
		Image string `json:"image"` // Base64 encoded image
	}

	// Parse request body
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		handlers.BadRequestError(c, pkg.StatusBadRequest, "Invalid request body")
		return nil
	}

	// Decode base64 image
	decodedImage, err := base64.StdEncoding.DecodeString(requestBody.Image)
	if err != nil {
		handlers.BadRequestError(c, pkg.StatusBadRequest, "Invalid base64 image data")
		return nil
	}

	// Decode image to image.Image
	img, _, err := image.Decode(bytes.NewReader(decodedImage))
	if err != nil {
		handlers.BadRequestError(c, pkg.StatusBadRequest, "Failed to decode image")
		return nil
	}

	// Save image to a temporary file (GoFace cần đọc từ file)
	tempFile, err := os.CreateTemp("", "face-*.jpg")
	if err != nil {
		handlers.InternalServerError(c, pkg.StatusInternalServerError, "Failed to create temp file")
		return nil
	}
	defer os.Remove(tempFile.Name()) // Xóa file tạm sau khi xử lý xong

	// Encode image to JPEG and write to temp file
	if err := jpeg.Encode(tempFile, img, nil); err != nil {
		handlers.InternalServerError(c, pkg.StatusInternalServerError, "Failed to encode image to JPEG")
		return nil
	}

	// Trích xuất embedding từ ảnh
	faces, err := recognizer.RecognizeFile(tempFile.Name())
	if err != nil {
		handlers.InternalServerError(c, pkg.StatusInternalServerError, "Failed to recognize face")
		return nil
	}

	if len(faces) == 0 {
		handlers.NotFoundError(c, pkg.StatusNotFound, "No face detected in the image")
		return nil
	}

	// Lấy embedding của khuôn mặt đầu tiên
	faceEmbedding := faces[0].Descriptor

	// Tạo ID cho khuôn mặt
	faceID := uuid.New().String()

	// Trả về response
	response := &models.FaceRegistrationResponse{
		FaceID:    faceID,
		Embedding: faceEmbedding[:], // Convert face.Descriptor to []float32
	}

	return response
}
