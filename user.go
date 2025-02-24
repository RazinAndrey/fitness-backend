package fitnessbackend

type User struct {
	Id       int    `json: "-"`
	Name     string `json: "name"`
	Username string `json: "username"`
	Password string `json: "password"`
}

// json тэги нужны чтоб коректно принимать и выводить данные на наших http запросах
