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
		qualifyAgedBrie(item)

		return
	}

	if item.Name == "Backstage passes to a TAFKAL80ETC concert" {
		qualifyBackstagePassesToATafkal80EtcConcert(item)

		return
	}

	if item.Name == "Sulfuras, Hand of Ragnaros" {
		qualifySulfurasHandOfRagnaros(item)
		return
	}

	defaultQualifier(item)
}

func qualifyAgedBrie(item *Item) {
	if item.Quality < 50 {
		item.Quality = item.Quality + 1
	}

	item.SellIn = item.SellIn - 1

	if item.SellIn < 0 {
		if item.Quality < 50 {
			item.Quality = item.Quality + 1
		}
	}
}

func qualifyBackstagePassesToATafkal80EtcConcert(item *Item) {
	if item.Quality < 50 {
		item.Quality = item.Quality + 1
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

	item.SellIn = item.SellIn - 1

	if item.SellIn < 0 {
		item.Quality = item.Quality - item.Quality
	}
}

func qualifySulfurasHandOfRagnaros(item *Item) {}

func defaultQualifier(item *Item) {
	if item.Quality > 0 {
		item.Quality = item.Quality - 1
	}

	item.SellIn = item.SellIn - 1

	if item.SellIn < 0 {
		if item.Quality > 0 {
			item.Quality = item.Quality - 1
		}
	}
}
