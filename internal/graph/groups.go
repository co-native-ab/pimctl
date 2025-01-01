package graph

type Group struct {
	ID              string   `json:"id"`
	DisplayName     string   `json:"displayName"`
	Description     string   `json:"description"`
	SecurityEnabled bool     `json:"securityEnabled"`
	GroupTypes      []string `json:"groupTypes"`
}

func (g Group) GroupType() string {
	if len(g.GroupTypes) == 0 {
		return "Security"
	}
	return "Microsoft 365"
}
