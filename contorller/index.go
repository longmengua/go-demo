package contorller

func (s *Struct) Init(routes []Route) {
	if s.Routes != nil && cap(routes) > 0 {
		s.Routes = append(s.Routes, routes...)
	}
}
