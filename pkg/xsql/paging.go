package xsql

type Paging struct {
	Page    int
	Size    int
	OrderBy string
}

func (p *Paging) Limit() int {
	return p.Size
}

func (p *Paging) Offset() int {
	return p.Page * p.Size
}

func (p *Paging) WithDefaultOrderBy(orderBy string) *Paging {
	if p.OrderBy == "" && orderBy != "" {
		p.OrderBy = orderBy
	}
	return p
}
