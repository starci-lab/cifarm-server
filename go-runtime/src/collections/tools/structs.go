package collections_tools

type Tool struct {
	Key         string `json:"key,omitempty"`
	AvailableIn int    `json:"availableIn,omitempty"`
	Index       int    `json:"index,omitempty"`
}
