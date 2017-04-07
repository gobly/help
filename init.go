package help

import "github.com/gobly/core"

func CreateContext(prefix string, r *core.Router) {
	r.AddSubRouter(prefix, initialize)
}

func initialize(s *core.Router) {
	s.AddHandler("/about", AboutHandler)
}
