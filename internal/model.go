package internal

type MusicBraizResponse struct {
	Disambiguation string      `json:"disambiguation"`
	EndArea        interface{} `json:"end_area"`
	Country        string      `json:"country"`
	SortName       string      `json:"sort-name"`
	ID             string      `json:"id"`
	BeginArea      struct {
	ID             string   `json:"id"`
	Name           string   `json:"name"`
	Disambiguation string   `json:"disambiguation"`
	Iso31662Codes  []string `json:"iso-3166-2-codes"`
	SortName       string   `json:"sort-name"`
	} `json:"begin-area"`
	Area struct {
	Iso31661Codes  []string `json:"iso-3166-1-codes"`
	SortName       string   `json:"sort-name"`
	ID             string   `json:"id"`
	Name           string   `json:"name"`
	Disambiguation string   `json:"disambiguation"`
	} `json:"area"`
	Name     string `json:"name"`
	LifeSpan struct {
	Ended bool        `json:"ended"`
	End   interface{} `json:"end"`
	Begin string      `json:"begin"`
	} `json:"life-span"`
	BeginArea1 struct {
	ID             string   `json:"id"`
	Name           string   `json:"name"`
	Disambiguation string   `json:"disambiguation"`
	Iso31662Codes  []string `json:"iso-3166-2-codes"`
	SortName       string   `json:"sort-name"`
	} `json:"begin_area"`
	EndArea1  interface{} `json:"end-area"`
	Type     string      `json:"type"`
	GenderID string      `json:"gender-id"`
	Isnis    []string    `json:"isnis"`
	Gender   string      `json:"gender"`
	Ipis     []string    `json:"ipis"`
	TypeID   string      `json:"type-id"`
	}

