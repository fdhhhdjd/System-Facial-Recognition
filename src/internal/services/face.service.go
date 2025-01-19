package services

import (
	"facial-recognition/src/internal/models"

	"github.com/gin-gonic/gin"
	// "gocv.io/x/gocv"
)

func DetectFaces(c *gin.Context) *models.FaceDetectionResult {
	// Upload image from URL
	// imageURL := c.Query("image_url")
	// if imageURL == "" {
	// 	handlers.BadRequestError(c, pkg.StatusBadRequest)
	// 	return nil
	// }

	// resp, err := http.Get(imageURL)
	// if err != nil {
	// 	handlers.BadRequestError(c, pkg.StatusBadRequest, fmt.Sprintf("failed to create temp file: %v", err))
	// 	return nil
	// }

	// // Close the response body
	// defer resp.Body.Close()

	// // Decode the image
	// file, err := os.CreateTemp("", "face-detection-*.jpg")
	// if err != nil {
	// 	handlers.BadRequestError(c, pkg.StatusBadRequest, fmt.Sprintf("failed to create temp file: %v", err))
	// 	return nil
	// }

	// defer os.Remove(file.Name())

	// _, err = io.Copy(file, resp.Body)
	// if err != nil {
	// 	handlers.BadRequestError(c, pkg.StatusBadRequest, fmt.Sprintf("failed to save image: %v", err))
	// 	return nil
	// }

	// img := gocv.IMRead(file.Name(), gocv.IMReadColor)
	// defer img.Close()

	// if img.Empty() {
	// 	handlers.BadRequestError(c, pkg.StatusBadRequest, "failed to read image")
	// 	return nil
	// }

	// // Detect faces
	// faceCascade := gocv.NewCascadeClassifier()
	// defer faceCascade.Close()

	// if !faceCascade.Load("haarcascade_frontalface_default.xml") {
	// 	handlers.BadRequestError(c, pkg.StatusBadRequest, "failed to load classifier")
	// 	return nil
	// }

	// rects := faceCascade.DetectMultiScale(img)

	return &models.FaceDetectionResult{}

}
