package domain

import (
	"awesomeProject/src/data"
	"fmt"
)

func SearchTopFilm(position int) (string, error) {
	if position < 0 || position > len(data.ListFilm) {
		return "", fmt.Errorf("film not found")
	} else {
		return data.ListFilm[position], nil
	}
}
