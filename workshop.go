package workshop

func UpdateQuality(items []*Item) {
	var (
		qualifiers = make(map[string]qualifier)
	)

	qualifiers["Aged Brie"] = qualifyAgedBrie()
	qualifiers["Backstage passes to a TAFKAL80ETC concert"] = qualifyBackstagePassesToATafkal80EtcConcert()
	qualifiers["Sulfuras, Hand of Ragnaros"] = qualifySulfurasHandOfRagnaros()
	qualifiers[""] = defaultQualifier()

	for i := 0; i < len(items); i++ {
		if qualifier, ok := qualifiers[items[i].Name]; ok {
			qualifier.updateQuality(items[i])
		} else {
			qualifiers[""].updateQuality(items[i])
		}
	}
}

type Item struct {
	Name            string
	SellIn, Quality int
}

type qualifier interface{ updateQuality(*Item) }

type qualifierFunc func(*Item)

func (qf qualifierFunc) updateQuality(item *Item) { qf(item) }

func qualifyAgedBrie() qualifierFunc {
	return func(item *Item) {
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
}

func qualifyBackstagePassesToATafkal80EtcConcert() qualifierFunc {
	return func(item *Item) {
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
}

func qualifySulfurasHandOfRagnaros() qualifierFunc { return func(item *Item) {} }

func defaultQualifier() qualifierFunc {
	return func(item *Item) {
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
}
