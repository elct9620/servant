package schema

type LivenessResponse struct {
	Ok bool `json:"ok"`
}

type ReadinessResponse struct {
	Ok bool `json:"ok"`
}
