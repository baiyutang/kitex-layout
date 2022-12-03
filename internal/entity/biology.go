package entity

type Biology struct {
	ID       uint64
	Kingdom  string
	Division string
	Class    string
	Order    string
	Family   string
	Genus    string
	Species  string
	Name     string
}

func NewBiology() *Biology {
	return &Biology{
		Kingdom:  "1",
		Division: "2",
		Class:    "3",
		Order:    "4",
		Family:   "5",
		Genus:    "6",
		Species:  "7",
		Name:     "8",
	}
}
