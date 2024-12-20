package types

type SiteMapRequest struct {
	ID string `json:"crawl_id"`
}

type SiteMapResponse struct {
	Message string `json:"message"`
	Data    string `json:"data"`
}
