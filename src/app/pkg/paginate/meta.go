package paginate

type Meta struct {
	Total       int64  `json:"total"`
	TotalPages  int    `json:"total_pages"`
	PerPage     int    `json:"per_page,omitempty;query:per_page"`
	CurrentPage int    `json:"current_page,omitempty;query:page"`
	LastPage    int    `json:"last_page,omitempty"`
	FistPage    int    `json:"first_page,omitempty"`
	Search      string `json:"search,omitempty;query:search"`
	Sort        string `json:"sort,omitempty;query:sort"`
	Order       string `json:"order,omitempty;query:order"`
}

func (p *Pagination) SetMeta(meta Meta) {
	p.Meta = meta
}
