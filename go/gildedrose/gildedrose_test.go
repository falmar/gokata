package gildedrose_test

import (
	"fmt"
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func Test_Legendary(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Sulfuras, Hand of Ragnaros", 3, 1},
		{"Sulfuras, Hand of Ragnaros", 0, 1},
		{"Sulfuras, Hand of Ragnaros", -1, 30},
		{"Sulfuras, Hand of Ragnaros", 10, 0},
		{"Sulfuras, Hand of Ragnaros", 20, 10000},
	}

	expectItems := make([][]int, len(items))

	// no change
	for j := 0; j < len(items); j++ {
		expectItems[j] = []int{items[j].SellIn, items[j].Quality}
	}

	for i := 0; i < 30; i++ {
		gildedrose.UpdateQuality(items)

		for j := 0; j < len(items); j++ {
			// nameshould not change
			if items[j].Name != "Sulfuras, Hand of Ragnaros" {
				t.Errorf("Name: Expected \"Sulfuras, Hand of Ragnaros\" but got %s ", items[j].Name)
			}

			// sellIn should decrease
			if items[j].SellIn != expectItems[j][0] {
				t.Errorf("SellIn: Expected %d but got %d ", expectItems[j][0]-1, items[j].SellIn)
			}

			// quality should decrease by 2 and not go below 0 or above 50
			if items[j].Quality != expectItems[j][1] {
				t.Errorf("Quality: Expected %d but got %d ", expectItems[j][1]-2, items[j].Quality)
			}
		}
	}
}

func Test_Cheese(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Aged Brie", 3, 1},
		{"Aged Brie", 0, 1},
		{"Aged Brie", -1, 30},
		{"Aged Brie", 10, 0},
		{"Aged Brie", 20, 10000},
	}

	expectItems := make([][]int, len(items))

	for i := 0; i < 30; i++ {
		for j := 0; j < len(items); j++ {
			expectItems[j] = []int{items[j].SellIn, items[j].Quality}

			var by int = 1

			if items[j].SellIn <= 0 {
				by = 2
			}

			// quality should decrease by 2 and not go below 0 or above 50
			if items[j].Quality+by < 0 {
				expectItems[j][1] = 0
			} else if items[j].Quality+by > 50 {
				expectItems[j][1] = 50
			} else {
				expectItems[j][1] = items[j].Quality + by
			}
		}

		gildedrose.UpdateQuality(items)

		for j := 0; j < len(items); j++ {
			// name should not change
			if items[j].Name != "Aged Brie" {
				t.Errorf("Name: Expected \"Aged Brie\" but got %s ", items[j].Name)
			}

			// sellIn should decrease
			if items[j].SellIn != expectItems[j][0]-1 {
				t.Errorf("SellIn: Expected %d but got %d ", expectItems[j][0]-1, items[j].SellIn)
			}

			// quality should increase by 1 and not go below 0 or above 50
			if items[j].Quality < 0 {
				t.Errorf("Quality: Expected %d but got %d ", 0, items[j].Quality)
			} else if items[j].Quality > 50 {
				t.Errorf("Quality: Expected %d but got %d ", 50, items[j].Quality)
			} else if items[j].Quality != expectItems[j][1] {
				t.Errorf("Quality: Expected %d but got %d ", expectItems[j][1]-2, items[j].Quality)
			}
		}
	}
}

func Test_Conjured(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Conjured Mana Cake", 3, 6},
		{"Conjured Shield", 10, 47},
	}

	expectItems := make([][]int, len(items))
	var expectNames = []string{"Conjured Mana Cake", "Conjured Shield"}

	for i := 0; i < 30; i++ {
		for j := 0; j < len(items); j++ {
			expectItems[j] = []int{items[j].SellIn, items[j].Quality}

			var by int = 1

			if items[j].SellIn <= 0 {
				by = 2
			}

			if items[j].Quality-by < 0 {
				expectItems[j][1] = 0
			} else if items[j].Quality-by > 50 {
				expectItems[j][1] = 50
			} else {
				expectItems[j][1] = items[j].Quality - by
			}
		}

		gildedrose.UpdateQuality(items)

		for j := 0; j < len(items); j++ {
			// nameshould not change
			if items[j].Name != expectNames[j] {
				t.Errorf("Name: Expected %s but got %s ", expectNames[j], items[j].Name)
			}

			// sellIn should decrease
			if items[j].SellIn != expectItems[j][0]-1 {
				t.Errorf("SellIn: Expected %d but got %d ", expectItems[j][0]-1, items[j].SellIn)
			}

			// quality should decrease by 2 and not go below 0 or above 50
			if items[j].Quality < 0 {
				t.Errorf("Quality: Expected %d but got %d ", 0, items[j].Quality)
			} else if items[j].Quality > 50 {
				t.Errorf("Quality: Expected %d but got %d ", 50, items[j].Quality)
			} else if items[j].Quality != expectItems[j][1] {
				t.Errorf("Quality: Expected %d but got %d ", expectItems[j][1], items[j].Quality)
			}
		}
	}
}

func Test_Backstage(t *testing.T) {
	items := []*gildedrose.Item{
		{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
		{"Backstage passes to a Wind Rose concert", 10, 49},
		{"Backstage passes to a Siames concert", 5, 49},
		{"Backstage passes to a In Flames concert", 3, 6},
		{"Backstage passes to a Thrice concert", 0, 6},
	}

	expectItems := make([][]int, len(items))
	var expectNames = []string{
		"Backstage passes to a TAFKAL80ETC concert",
		"Backstage passes to a Wind Rose concert",
		"Backstage passes to a Siames concert",
		"Backstage passes to a In Flames concert",
		"Backstage passes to a Thrice concert",
	}

	for i := 0; i < 30; i++ {
		for j := 0; j < len(items); j++ {
			expectItems[j] = []int{items[j].SellIn, items[j].Quality}

			var by int = 1

			if items[j].SellIn <= 0 {
				expectItems[j][1] = 0
				continue
			} else if items[j].SellIn <= 5 {
				by = 3
			} else if items[j].SellIn <= 10 {
				by = 2
			}

			// quality should decrease by 1 and not go below 0 or above 50
			if items[j].Quality+by > 50 {
				expectItems[j][1] = 50
			} else if items[j].Quality < 0 {
				expectItems[j][1] = 0
			} else {
				expectItems[j][1] = items[j].Quality + by
			}
		}

		gildedrose.UpdateQuality(items)

		for j := 0; j < len(items); j++ {
			// nameshould not change
			if items[j].Name != expectNames[j] {
				t.Errorf("Name: Expected %s but got %s ", expectNames[j], items[j].Name)
			}

			// sellIn should decrease
			if items[j].SellIn != expectItems[j][0]-1 {
				t.Errorf("SellIn: Expected %d but got %d ", expectItems[j][0]-1, items[j].SellIn)
			}

			// quality should decrease by 1 and not go below 0 or above 50
			if items[j].Quality < 0 {
				t.Errorf("Quality: Expected %d but got %d ", 0, items[j].Quality)
			} else if items[j].Quality > 50 {
				t.Errorf("Quality: Expected %d but got %d ", 50, items[j].Quality)
			} else if items[j].Quality != expectItems[j][1] {
				fmt.Println(items[j])
				t.Errorf("Quality: Expected %d but got %d ", expectItems[j][1], items[j].Quality)
			}
		}
	}
}

func Test_Regular(t *testing.T) {
	var items = []*gildedrose.Item{
		{"+5 Dexterity Vest", 10, 20},
		{"Elixir of the Mongoose", 5, 7},
		{"Mystical Ruby Amulet", 15, 30},
		{"Potion of Infinite Wisdom", 7, 10},
		{"Elven Leather Boots", 12, 25},
	}

	expectItems := make([][]int, len(items))
	var expectNames = []string{
		"+5 Dexterity Vest",
		"Elixir of the Mongoose",
		"Mystical Ruby Amulet",
		"Potion of Infinite Wisdom",
		"Elven Leather Boots",
	}

	for i := 0; i < 30; i++ {
		// make sure we have the right expectations from their current state
		for j := 0; j < len(items); j++ {
			expectItems[j] = []int{items[j].SellIn, items[j].Quality}

			var by int = 1

			if items[j].SellIn <= 0 {
				by = 2
			}

			// quality should decrease by 1 and not go below 0 or above 50
			if items[j].Quality-by < 0 {
				expectItems[j][1] = 0
			} else if items[j].Quality-by > 50 {
				expectItems[j][1] = 50
			} else {
				expectItems[j][1] = items[j].Quality - by
			}
		}

		gildedrose.UpdateQuality(items)

		for j := 0; j < len(items); j++ {
			// nameshould not change
			if items[j].Name != expectNames[j] {
				t.Errorf("Name: Expected %s but got %s ", expectNames[j], items[j].Name)
			}

			// sellIn should decrease
			if items[j].SellIn != expectItems[j][0]-1 {
				t.Errorf("SellIn: Expected %d but got %d ", expectItems[j][0]-1, items[j].SellIn)
			}

			// quality should decrease by 1 and not go below 0 or above 50
			if items[j].Quality < 0 {
				t.Errorf("Quality: Expected %d but got %d ", 0, items[j].Quality)
			} else if items[j].Quality > 50 {
				t.Errorf("Quality: Expected %d but got %d ", 50, items[j].Quality)
			} else if items[j].Quality != expectItems[j][1] {
				t.Errorf("Quality: Expected %d but got %d ", expectItems[j][1], items[j].Quality)
			}
		}
	}
}
