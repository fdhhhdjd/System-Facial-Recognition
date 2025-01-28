package models

type Face struct {
	ID     string `json:"id"`     // Unique ID for the face
	X      int    `json:"x"`      // X coordinate of the face
	Y      int    `json:"y"`      // Y coordinate of the face
	Width  int    `json:"width"`  // Width of the detected face
	Height int    `json:"height"` // Height of the detected face
}

type FaceRegistrationResponse struct {
	Message string `json:"message"`
	Face    Face   `json:"face"`
}
