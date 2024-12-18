package main

import "Scoreapi_go/items"

func main() {
	itemsList, err := items.NewTypeItems("beers")
	if err != nil {
		panic(err)
	}
	itemsList.RemoveItemByName("kozel")
}
