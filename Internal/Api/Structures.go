package Api

type Data struct {
	Groups          []GroupInfos `json:"groups"`
	AllGroups       []GroupInfos `json:"allGroups"`
	Countries       []string     `json:"countries"`
	MinCreationDate int
	MaxCreationDate int
	QtyOfMemberList []int
	MinFirstAlbum   int
	MaxFirstAlbum   int
}

type GroupInfos struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	QtyOfMembers int      `json:"qtyOfMembers"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	RelationsUrl string   `json:"relations"`
	Relations    Relation `json:"allRelations"`
}

type Relation struct {
	Id             int                  `json:"id"`
	DatesLocations map[string][]string  `json:"datesLocations"`
	Coordinates    map[string][]float64 `json:"coordinates"`
	CountriesList  []string             `json:"countriesList"`
}

type Filters struct {
	// 'CreationDate' filter
	IsCreationDateFilter bool
	MinCreationDate      string
	MaxCreationDate      string

	// 'QtyOfMembers' filter
	IsQtyOfMembersFilter bool
	QtyOfMembersList     []string

	// 'FirstAlbumDate' filter
	IsFirstAlbumDateFilter bool
	MinFirstAlbumDate      string
	MaxFirstAlbumDate      string

	// 'Country' filter
	IsCountryFilter bool
	CountryToFilter string
}
