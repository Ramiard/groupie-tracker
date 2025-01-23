package Api

type GroupInfos struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	QtyOfMembers int
	CreationDate int    `json:"creationDate"`
	FirstAlbum   string `json:"firstAlbum"`
	RelationsUrl string `json:"relations"`
	Relations    Relation
}

type Relation struct {
	Id             int                  `json:"id"`
	DatesLocations map[string][]string  `json:"datesLocations"`
	Coordinates    map[string][]float64 `json:"coordinates"`
}

//var Groups GroupList
//var Group GroupInfos
