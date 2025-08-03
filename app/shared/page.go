package shared

type SortDirection string

const ASC SortDirection = "ASC"
const DESC SortDirection = "DESC"

type PageInput struct {
	Limit         int           `json:"limit" validate:"omitempty,min=1,max=100"`
	Offset        int           `json:"offset" validate:"omitempty,min=0"`
	SortDirection SortDirection `json:"sort_direction" validate:"omitempty,min=0"`
}
