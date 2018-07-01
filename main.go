package tps

import (
	"github.com/gobuffalo/packr"
)

// GetPageStrategyElement retrieves the element for a specific page strategy
func GetPageStrategyElement(pageStrategy string, element string) (string, error) {
	box := packr.NewBox("./strategies")
	s, err := box.MustString(pageStrategy + "/" + element)
	if err != nil {
		return "", err
	}
	return s, nil
}
