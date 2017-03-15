package dbr

type union struct {
	builder []SelectBuilder
	all     bool
}

func Union(builder ...SelectBuilder) interface {
	Builder
	As(string) Builder
} {
	return &union{
		builder: builder,
	}
}

func UnionAll(builder ...SelectBuilder) interface {
	Builder
	As(string) Builder
} {
	return &union{
		builder: builder,
		all:     true,
	}
}

func (u *union) Build(d Dialect, buf Buffer) error {
	for i, b := range u.builder {
		if i > 0 {
			buf.WriteString(" UNION ")
			if u.all {
				buf.WriteString("ALL ")
			}
		}
		buf.WriteString(placeholder)
		buf.WriteValue(b)
	}
	return nil
}

func (u *union) As(alias string) Builder {
	return as(u, alias)
}
