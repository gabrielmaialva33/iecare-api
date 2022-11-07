package paginate

import (
	"iecare-api/src/app/utils"
	"strings"
)

type Pagination struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

func (m *Meta) GetOffset() int {
	return (m.GetCurrentPage() - 1) * m.GetPerPage()
}

func (m *Meta) GetPerPage() int {
	if m.PerPage == 0 || m.PerPage < 0 {
		return 10
	} else if m.PerPage > 100 {
		return 100
	}

	return m.PerPage
}

func (m *Meta) GetCurrentPage() int {
	if m.CurrentPage == 0 || m.CurrentPage < 0 {
		return 1
	}
	return m.CurrentPage
}

func (m *Meta) GetSort(fields []string) string {
	if m.Sort == "" {
		return "id" + " " + m.GetOrder()
	}

	if utils.Contains(fields, m.Sort) != true {
		return "id" + " " + m.GetOrder()
	}

	return strings.ToLower(m.Sort) + " " + m.GetOrder()
}

func (m *Meta) GetOrder() string {
	orders := []string{"asc", "desc"}

	if m.Order == "" {
		return "ASC"
	}

	if utils.Contains(orders, strings.ToLower(m.Order)) != true {
		return "ASC"
	}

	return strings.ToUpper(m.Order)
}

func (m *Meta) GetSearch() string {
	return m.Search
}

func (p *Pagination) SetData(data interface{}) {
	p.Data = data
}
