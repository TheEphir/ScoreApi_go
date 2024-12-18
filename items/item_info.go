package items

import "errors"

type Item struct {
	Name        string `json:"name"`
	Score       int    `json:"score"`
	Image       string `json:"image"`
	Description string `json:"description"`
}

func NewItem(name string, score int, image string, description string) (*Item, error) {
	if name == "" {
		return nil, errors.New("ITEM_SHOULD_BE_NAMED")
	}
	if score < 0 {
		return nil, errors.New("SCORE_LESS_0")
	}
	if image == "" {
		image = "There is no image"
	}
	if description == "" {
		description = "There is no description"
	}

	return &Item{
		Name:        name,
		Score:       score,
		Image:       image,
		Description: description,
	}, nil
}
