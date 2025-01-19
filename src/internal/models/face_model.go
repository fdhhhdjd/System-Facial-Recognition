package models

import "image"

type FaceDetectionResult struct {
	FaceCount int               `json:"face_count"`
	Faces     []image.Rectangle `json:"faces"`
}

type FaceVerificationResult struct {
	Match bool `json:"match"`
}
