package main

type Metadata struct {
	Annotation       string   `json:"annotation"`
	Ext              string   `json:"ext"`
	Folders          []string `json:"folders"`
	Height           int      `json:"height"`
	ID               string   `json:"id"`
	IsDeleted        bool     `json:"isDeleted"`
	LastModified     int64    `json:"lastModified"`
	ModificationTime int64    `json:"modificationTime"`
	Name             string   `json:"name"`
	NoThumbnail      bool     `json:"noThumbnail"`
	Palettes         []struct {
		Color []int   `json:"color"`
		Ratio float64 `json:"ratio"`
	} `json:"palettes"`
	Size  int           `json:"size"`
	Tags  []interface{} `json:"tags"`
	URL   string        `json:"url"`
	Width int           `json:"width"`
}
