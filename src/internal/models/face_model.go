package models

type Face struct {
	ID     string `json:"id"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type FaceDetectionResponse struct {
	Faces []Face `json:"faces"`
}

type FaceRegistrationResponse struct {
	FaceID    string    `json:"face_id"`   // ID của khuôn mặt
	Embedding []float32 `json:"embedding"` // Embedding của khuôn mặt
}
