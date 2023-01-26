package domain

import "github.com/golangid/candi/candishared"

// FilterBrand model
type FilterBrand struct {
	candishared.Filter
	ID        *string `json:"id"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Preloads  []string `json:"-"`
}
