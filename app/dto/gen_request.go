package dto

type SimpleRouteRequest struct {
	DepartureTime string  `json:"departureTime"`
	SelectedCity  string  `json:"selectedCity"`
	SourceLat     float64 `json:"sourceLat"`
	SourceLong    float64 `json:"sourceLong"`
	Direction     string  `json:"direction"`
	StartStopId   int64   `json:"startStopId"`
	EndStopId     int64   `json:"endStopId"`
	RouteId       string  `json:"routeId"`
	RouteOrder    int     `json:"routeOrder"`
	Service       string  `json:"service"`
	Search        string  `json:"search"`
	Limit         int     `json:"limit"`
	Page          int     `json:"page"`
	Code          string  `json:"code"`
	StopName      string  `json:"stopName"`
	Address       string  `json:"address"`
	Action        string  `json:"action"`
	ActionId      string  `json:"actionId"`
	ActionOrder   string  `json:"actionOrder"`
}

type Request struct {
	RepoURL      string   `json:"repo_url"`
	Dir          string   `json:"dir"`
	Username     string   `json:"username"`
	Password     string   `json:"password"`
	Branch       string   `json:"branch"`
	RemoteBranch string   `json:"remoteBranch"`
	Flag         string   `json:flag`
	Files        []string `json:files`
	Message      string   `json:message`
}
