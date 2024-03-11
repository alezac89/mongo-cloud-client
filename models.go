package mongoCloud

type Function struct {
	ID string `json:"_id"`
	Name string `json:"name"`
	Source string `json:"source"`
	Private bool `json:"private"`
	LastModified int64 `json:"last_modified"`
	ReadOnly bool `json:"read_only"`
}