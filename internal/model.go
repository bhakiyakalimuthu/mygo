package internal

type MusicBraizResponse struct {
	Relations []struct {
		Type string `json:"type"`
		URL  struct {
			ID       string `json:"id"`
			Resource string `json:"resource"`
		} `json:"url"`
		AttributeValues struct {
		} `json:"attribute-values"`
		Begin        interface{}   `json:"begin"`
		SourceCredit string        `json:"source-credit"`
		Ended        bool          `json:"ended"`
		Attributes   []interface{} `json:"attributes"`
		AttributeIds struct {
		} `json:"attribute-ids"`
		TypeID       string      `json:"type-id"`
		TargetCredit string      `json:"target-credit"`
		Direction    string      `json:"direction"`
		End          interface{} `json:"end"`
		TargetType   string      `json:"target-type"`
	} `json:"relations"`
	Gender  string `json:"gender"`
	Country string `json:"country"`
	TypeID  string `json:"type-id"`
	Area    struct {
		Name           string   `json:"name"`
		ID             string   `json:"id"`
		Iso31661Codes  []string `json:"iso-3166-1-codes"`
		SortName       string   `json:"sort-name"`
		Disambiguation string   `json:"disambiguation"`
	} `json:"area"`
	Beginarea struct {
		ID             string `json:"id"`
		Name           string `json:"name"`
		Disambiguation string `json:"disambiguation"`
		SortName       string `json:"sort-name"`
	} `json:"begin-area"`
	Disambiguation string `json:"disambiguation"`
	ID             string `json:"id"`
	Endarea        struct {
		Name           string `json:"name"`
		ID             string `json:"id"`
		SortName       string `json:"sort-name"`
		Disambiguation string `json:"disambiguation"`
	} `json:"end_area"`
	Ipis    []string `json:"ipis"`
	Isnis   []string `json:"isnis"`
	EndArea struct {
		Name           string `json:"name"`
		ID             string `json:"id"`
		SortName       string `json:"sort-name"`
		Disambiguation string `json:"disambiguation"`
	} `json:"end-area"`
	ReleaseGroups []struct {
		SecondaryTypes   []interface{} `json:"secondary-types"`
		ID               string        `json:"id"`
		Title            string        `json:"title"`
		PrimaryType      string        `json:"primary-type"`
		PrimaryTypeID    string        `json:"primary-type-id"`
		FirstReleaseDate string        `json:"first-release-date"`
		SecondaryTypeIds []interface{} `json:"secondary-type-ids"`
		Disambiguation   string        `json:"disambiguation"`
	} `json:"release-groups"`
	Type      string `json:"type"`
	BeginArea struct {
		ID             string `json:"id"`
		Name           string `json:"name"`
		Disambiguation string `json:"disambiguation"`
		SortName       string `json:"sort-name"`
	} `json:"begin_area"`
	Name     string `json:"name"`
	GenderID string `json:"gender-id"`
	LifeSpan struct {
		Ended bool   `json:"ended"`
		End   string `json:"end"`
		Begin string `json:"begin"`
	} `json:"life-span"`
	SortName string `json:"sort-name"`
}
