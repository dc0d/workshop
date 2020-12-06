package workshop

type Item struct {
	Name            string
	SellIn, Quality int
}

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {
		updateQuality(items[i])
	}
}

func updateQuality(item *Item) {
	if item.Name == "Aged Brie" {
		if item.Quality < 50 {
			item.Quality = item.Quality + 1
		}

		item.SellIn = item.SellIn - 1

		if item.SellIn < 0 {
			if item.Quality < 50 {
				item.Quality = item.Quality + 1
			}
		}

		return
	}

	if item.Name != "Aged Brie" && item.Name != "Backstage passes to a TAFKAL80ETC concert" {
		if item.Quality > 0 {
			if item.Name != "Sulfuras, Hand of Ragnaros" {
				item.Quality = item.Quality - 1
			}
		}
	} else {
		if item.Quality < 50 {
			item.Quality = item.Quality + 1
			if item.Name == "Backstage passes to a TAFKAL80ETC concert" {
				if item.SellIn < 11 {
					if item.Quality < 50 {
						item.Quality = item.Quality + 1
					}
				}
				if item.SellIn < 6 {
					if item.Quality < 50 {
						item.Quality = item.Quality + 1
					}
				}
			}
		}
	}

	if item.Name != "Sulfuras, Hand of Ragnaros" {
		item.SellIn = item.SellIn - 1
	}

	if item.SellIn < 0 {
		if item.Name != "Aged Brie" {
			if item.Name != "Backstage passes to a TAFKAL80ETC concert" {
				if item.Quality > 0 {
					if item.Name != "Sulfuras, Hand of Ragnaros" {
						item.Quality = item.Quality - 1
					}
				}
			} else {
				item.Quality = item.Quality - item.Quality
			}
		}
	}
}
