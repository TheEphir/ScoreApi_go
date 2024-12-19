package items

import (
	"Scoreapi_go/files"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type TypeItems struct {
	ItemType string `json:"type"`
	Items    []Item `json:"items"`
}

func NewTypeItems(itemType string) (*TypeItems, error) {
	if itemType == "" {
		return nil, errors.New("GOT_EMPTY_TYPE")
	}

	file, err := files.ReadFile(itemType)
	if err != nil {
		return &TypeItems{
			ItemType: itemType,
			Items:    []Item{},
		}, nil
	}

	var itemsList TypeItems

	err = json.Unmarshal(file, &itemsList)
	if err != nil {
		return nil, err
	}
	return &itemsList, nil
}

func (itemsList *TypeItems) AddItem(item Item) {
	itemsList.Items = append(itemsList.Items, item)

	data, err := itemsList.ToBytes()
	if err != nil {
		fmt.Println("CANNNOT SLICE TO BYTES")
	}
	files.WriteFile(itemsList.ItemType, data)
}

func (itemsList *TypeItems) ToBytes() ([]byte, error) {
	data, err := json.Marshal(itemsList)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Find any item that contain toSearch.
func (itemsList *TypeItems) SearchItemByName(toSearch string) []Item {
	var res []Item
	for _, item := range itemsList.Items {
		isMatched := strings.Contains(item.Name, toSearch)
		if isMatched {
			res = append(res, item)
		}
	}
	return res
}

// Find item that compare with itemToRemove, and don't append it into "new items".
// Then rewrite file with new data.
func (itemsList *TypeItems) RemoveItemByName(itemToRemove string) error {
	var res []Item
	for _, item := range itemsList.Items {
		isMatched := strings.Compare(item.Name, itemToRemove) // Compare method cause if you use Contains it will remove all items that contain even part of name
		if isMatched != 0 {
			res = append(res, item)
		}
	}

	itemsList.Items = res

	newItems, err := itemsList.ToBytes()
	if err != nil {
		return err
	}

	err = files.WriteFile(itemsList.ItemType, newItems)
	if err != nil {
		return err
	}
	return nil
}
