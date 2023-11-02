package domain

import (
	"fmt"
)

func SearchTopFilm(position int) (string, error) {
	if position < 0 || position > len(ListFilm) {
		return "", fmt.Errorf("Film not found")
	} else {
		return ListFilm[position], nil
	}
}
