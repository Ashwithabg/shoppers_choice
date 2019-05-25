package viewmodel

type Base struct {
	Title string
	Active string
}

func NewBase() Base {
	return Base{
		Title: "Shopper's Choice",
		Active: "home",
	}
}
