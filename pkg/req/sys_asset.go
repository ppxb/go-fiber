package req

import "time"

type CreateAssetDto struct {
	AssetId     string    `json:"assetId"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	ChildType   string    `json:"childType"`
	ProjectName string    `json:"projectName"`
	Origin      string    `json:"origin"`
	Model       string    `json:"model"`
	Value       float64   `json:"value"`
	Unit        string    `json:"unit"`
	InDate      time.Time `json:"inDate"`
	OpDate      time.Time `json:"opDate"`
	DepYear     int       `json:"depYear"`
	Location    string    `json:"location"`
}
