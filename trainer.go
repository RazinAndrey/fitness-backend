package fitnessbackend

type Trainer struct {
	Id          int    `json: "-"`
	Name        string `json: "name"`
	Description string `json: "description"`
}
