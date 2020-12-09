package workshop

type Item struct {
	name            string
	sellIn, quality int
}

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {
		qualifier := getQualifier(items[i].name)
		qualifier.qualify(items[i])
	}
}

func getQualifier(name string) qualifier {
	switch name {
	case "Aged Brie":
		return qualifierFunc(qualifyAgedBrie)
	case "Backstage passes to a TAFKAL80ETC concert":
		return qualifierFunc(qualifyBackstagePassesToATafkal80EtcConcert)
	case "Sulfuras, Hand of Ragnaros":
		return qualifierFunc(qualifySulfurasHandOfRagnaros)
	}

	return qualifierFunc(qualifyUnknown)
}

func qualifyAgedBrie(item *Item) {
	if item.quality < 50 {
		item.quality = item.quality + 1
	}

	item.sellIn = item.sellIn - 1

	if item.sellIn < 0 {
		if item.quality < 50 {
			item.quality = item.quality + 1
		}
	}
}

func qualifyBackstagePassesToATafkal80EtcConcert(item *Item) {
	if item.quality < 50 {
		item.quality = item.quality + 1

		if item.sellIn < 11 {
			if item.quality < 50 {
				item.quality = item.quality + 1
			}
		}
		if item.sellIn < 6 {
			if item.quality < 50 {
				item.quality = item.quality + 1
			}
		}
	}

	item.sellIn = item.sellIn - 1

	if item.sellIn < 0 {
		item.quality = 0
	}
}

func qualifySulfurasHandOfRagnaros(item *Item) {}

func qualifyUnknown(item *Item) {
	if item.quality > 0 {
		item.quality = item.quality - 1
	}

	item.sellIn = item.sellIn - 1

	if item.sellIn < 0 {
		if item.quality > 0 {
			item.quality = item.quality - 1
		}
	}
}

type qualifierFunc func(item *Item)

func (f qualifierFunc) qualify(item *Item) { f(item) }

type qualifier interface {
	qualify(item *Item)
}
