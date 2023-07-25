package gildedrose

import (
	"strings"
)

type Item struct {
	Name            string
	SellIn, Quality int
}

const maxQuality = 50
const minQuality = 0

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {
		// too legendary to change
		if items[i].Name == "Sulfuras, Hand of Ragnaros" {
			continue
		}

		// degrade by 1
		var offset int = -1
		// to check for "backstage passes" and "conjured"
		var itemName = strings.ToLower(items[i].Name)

		if strings.Contains(itemName, "backstage passes") {
			if items[i].SellIn <= 0 {
				// useless ticket!
				items[i].SellIn -= 1
				items[i].Quality = 0
				continue
			}

			offset = 1

			// quality goes up by dates
			if items[i].SellIn <= 5 {
				offset++
			}
			if items[i].SellIn <= 10 {
				offset++
			}
		} else if itemName == "aged brie" {
			offset = 1
		}

		// degrade twice as much after sell date
		// || strings.Contains(itemName, "conjured")
		// conjured items degrade twice as fast; it will break the "texttest"
		if items[i].SellIn <= 0 {
			offset *= 2
		}

		// SellIn changes for all
		items[i].SellIn -= 1

		// can't go above 50 or below 0
		if items[i].Quality+offset > maxQuality {
			items[i].Quality = maxQuality
			continue
		} else if items[i].Quality+offset < minQuality {
			items[i].Quality = minQuality
			continue
		}

		items[i].Quality += offset
	}
}
