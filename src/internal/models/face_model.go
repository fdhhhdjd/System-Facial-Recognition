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
