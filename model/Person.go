package model

type Person struct {
	Id      int    `bson:"_id"`
	Name    string `bson:"name"`
	Address string `bson:"address"`
	Student Student `json:"student"`
}

type Student struct {
	Id    int `bson:"_id"`
	Class string `bson:"class"`
	Score int `bson:"score"`
}
