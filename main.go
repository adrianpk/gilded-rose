package main

import (
	"fmt"

	gildedrose "github.com/robphoenix/gilded-rose/gilded-rose"
)

func init() {
	things = []Updatable{
		&Default{
			Item: &Item{"+5 Dexterity Vest", 10, 20},
		},
		&AgedBrie{
			Item: &Item{"Aged Brie", 2, 0},
		},
		&Default{
			Item: &Item{"Elixir of the Mongoose", 5, 7},
		},
		&Sulfuras{
			Item: &Item{"Sulfuras, Hand of Ragnaros", 0, 80},
		},
		&BackstagePasses{
			Item: &Item{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
		},
		&Conjured{
			Item: &Item{"Conjured Mana Cake", 3, 6},
		},
	}
}

func main() {
	gildedrose.UpdateQuality(items...)

	for _, item := range items {
		fmt.Println(item)
	}
}
