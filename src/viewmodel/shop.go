package viewmodel

type Category struct {
	Title       string
	Description string
	URL         string
	ImageURL    string
	IsOrientLeft bool
}

type Shop struct {
	Title      string
	Active     string
	Categories []Category
}

func NewShop() Shop {
	shop := Shop{
		Title:  "Shopper's choice shop",
		Active: "shop",
	}

	uncookedCategory := Category{
		URL:      "/shop_details",
		ImageURL: "uncooked_fish.png",
		Title:    "Uncooked fishes",
		IsOrientLeft: true,
		Description: `We have wide varities of fresh fishes available.
		The most important characteristic of fish in transport is its freshness –
		and the freshest fish is the one just caught.
		As opposed to other kinds of meat that require maturing,
		fish is best sold and eaten fresh, preferably killed right before being cooked.
		Transport of fresh fish is considerably expensive,
		which is probably the reason why in total fish and fish products consumption only 3% of
		fresh fish is used.`,
	}

	cookedCategory := Category{
		URL:      "/shop_details",
		ImageURL: "cooked_fish.png",
		Title:    "Cooked fishes",
		IsOrientLeft: false,
		Description: `Save your butter for baking. We've got five healthier ideas for cooking
                        fresh fish so it turns out moist and delicious. Some of them are Grilled fish,
                        Pooched fish, steamed fish, Baked fish, fried fish etc.Choosing fish for your entree is smart --
                        until it's salted, breaded and fried.
                        For the healthiest fish preparation, start by selecting a sustainable fillet.
                        The Monterey Bay Aquarium's Seafood Watch notes U.S. catfish, Alaskan salmon, U.S.-farmed
                        rainbow trout and Pacific halibut as some of the top options.
                        Then, cook it in a manner that limits sodium, fat and calories while still offering the taste
                        you desire.`,
	}
	shop.Categories = []Category{uncookedCategory, cookedCategory}
	return shop
}
