package foo

import "github.com/sirupsen/logrus"

// Hello is the v2 implementation of Hello.
type Hello struct {
	Name string
}

// String is the v2 implementation to return Hello's name.
func (h *Hello) String() string {
	logrus.Infof("printing v2: %s", h.Name)
	return h.Name
}
