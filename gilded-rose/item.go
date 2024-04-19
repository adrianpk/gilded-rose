package gildedrose

import "fmt"

// Item describes an item sold by the Gilded Rose Inn
type Item struct {
	name    string
	days    int
	quality int
}

var things []Updatable

type (
	Default struct {
		*Item
	}

	AgedBrie struct {
		*Item
	}

	Sulfuras struct {
		*Item
	}

	BackstagePasses struct {
		*Item
	}

	Conjured struct {
		*Item
	}

	Updatable interface {
		Update()
	}
)

func (i *Item) DecreaseQuality(amount int) {
	result := i.quality - amount
	if result > 0 {
		i.quality = result
	}
}

func (i *Item) IncreaseQuality(amount int) {
	result := i.quality + amount
	if result <= 50 {
		i.quality = result
	}
}

func (i *Default) Update() {
	i.days--
	if i.days > 0 {
		i.DecreaseQuality(1)
		return
	}

	i.DecreaseQuality(2)
}

func (i *AgedBrie) Update() {
	i.days--
	i.IncreaseQuality(1)
}

func (i *Sulfuras) Update() {
	// Do nothing
}

func (i *BackstagePasses) Update() {
	i--
	switch {

	}

}

func (i *Conjured) Update() {

}

// String outputs a string representation of an Item
func (i *Item) String() string {
	return fmt.Sprintf("%s: %d days left, quality is %d", i.name, i.days, i.quality)
}

// New creates a new Item
func New(name string, days, quality int) *Item {
	return &Item{
		name:    name,
		days:    days,
		quality: quality,
	}
}

func Update(items ...Updatable) {
	for _, item := range items {
		item.Update()
	}
}

// UpdateQuality2 ages the item by a day, and updates the quality of the item
func UpdateQuality2(items ...*Item) {
	for _, item := range items {
		if item.name != "Aged Brie" && item.name != "Backstage passes to a TAFKAL80ETC concert" {
			if item.quality > 0 {
				if item.name != "Sulfuras, Hand of Ragnaros" {
					item.quality = item.quality - 1
				}
			}
		} else {
			if item.quality < 50 {
				item.quality = item.quality + 1
				if item.name == "Backstage passes to a TAFKAL80ETC concert" {
					if item.days < 11 {
						if item.quality < 50 {
							item.quality = item.quality + 1
						}
					}
					if item.days < 6 {
						if item.quality < 50 {
							item.quality = item.quality + 1
						}
					}
				}
			}
		}

		if item.name != "Sulfuras, Hand of Ragnaros" {
			item.days = item.days - 1
		}

		if item.days < 0 {
			if item.name != "Aged Brie" {
				if item.name != "Backstage passes to a TAFKAL80ETC concert" {
					if item.quality > 0 {
						if item.name != "Sulfuras, Hand of Ragnaros" {
							item.quality = item.quality - 1
						}
					}
				} else {
					item.quality = item.quality - item.quality
				}
			} else {
				if item.quality < 50 {
					item.quality = item.quality + 1
				}
			}
		}
	}
}
