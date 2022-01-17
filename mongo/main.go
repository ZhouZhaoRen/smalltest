package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"smalltest/model"
	"time"
)

var mongoCli *mongo.Collection

func main() {
	var (
		client     *mongo.Client
		err        error
		db         *mongo.Database
		collection *mongo.Collection
	)
	//1.建立连接
	//if client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017").SetConnectTimeout(5*time.Second)); err != nil {
	//	fmt.Print(err)
	//	return
	//}
	if client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://1.117.76.139:27017").SetConnectTimeout(5*time.Second)); err != nil {
		fmt.Print(err)
		return
	}
	//2.选择数据库 my_db
	db = client.Database("small")

	//3.选择表 my_collection
	collection = db.Collection("stu")
	mongoCli = collection
	insertOne()
	//insertMany()
	//getOne()
	//insertManyStudent()
	//bsonFind()
	//aggregate()
	//findRegex()
}

// findRegex 通过正则表达式查找数据
func findRegex() {
	filter := bson.M{
		"name": bson.M{"$regex":"small1.*"},
	}
	res,err:=mongoCli.Find(context.TODO(),filter)
	if err != nil {
		panic(err)
	}
	var persons []model.Person
	err=res.All(context.TODO(),&persons)
	if err!=nil {
		panic(err)
	}
	fmt.Printf("res==%+v\n",persons)
}

func insertManyStudent() {
	stu := []interface{}{
		model.StudentMongo{
			Id:      1,
			Name:    "small1",
			Gender:  0,
			Country: "china",
		},
		model.StudentMongo{
			Id:      2,
			Name:    "small2",
			Gender:  1,
			Country: "use",
		},
		model.StudentMongo{
			Id:      3,
			Name:    "small3",
			Gender:  0,
			Country: "china",
		},
		model.StudentMongo{
			Id:      4,
			Name:    "small4",
			Gender:  1,
			Country: "canada",
		},
	}
	//
	_, err := mongoCli.InsertMany(context.TODO(), stu)
	if err != nil {
		panic(err)
	}
}

func insertMany() {
	persons := []interface{}{
		model.Person{
			Id:      6,
			Name:    "zzr4",
			Address: "shenzhen4",
			Student: model.Student{
				Id:    44,
				Class: "三年4班",
				Score: 4,
			},
		},
		model.Person{
			Id:      8,
			Name:    "zzr5",
			Address: "shenzhen5",
			Student: model.Student{
				Id:    55,
				Class: "三年5班",
				Score: 5,
			},
		},
	}
	//
	res, err := mongoCli.InsertMany(context.TODO(), persons)
	if err != nil {
		panic(err)
	}
	for _, value := range res.InsertedIDs {
		fmt.Println("id==", value)
	}
}

func insertOne() {
	person := model.Person{
		Id:      4,
		Name:    "zzr",
		Address: "shenzhen",
		Student: model.Student{
			Id:    22,
			Class: "三年二班",
			Score: 100,
		},
	}
	res, err := mongoCli.InsertOne(context.Background(), person)
	if err != nil {
		panic(err)
	}
	fmt.Println("id==", res.InsertedID)
}

func getOne() {
	filter := bson.M{
		"name": "zzr",
	}
	cursor, err := mongoCli.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var result []model.Person
	if err := cursor.All(context.TODO(), &result); err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func bsonFind() {
	option := bson.D{{
		"name",
		bson.D{{
			"$in",
			bson.A{
				"small1", "small2",
			},
		}},
	}}
	//
	var result []model.StudentMongo
	res, err := mongoCli.Find(context.TODO(), option)
	if err != nil {
		panic(err)
	}
	err = res.All(context.TODO(), &result)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func aggregate() {
	groupAggregate := mongo.Pipeline{
		bson.D{
			{
				"$group",
				bson.D{
					{
						"_id",
						"$country",
					},
					{
						"countGender",
						bson.D{
							{"$sum", 1},
						},
					},
				},
			},
			{
				"$sort",
				bson.D{
					{"countGender", 1},
				},
			},
		},
	}

	res, err := mongoCli.Aggregate(context.TODO(), groupAggregate)
	if err != nil {
		panic(err)
	}
	var result []bson.M
	err = res.All(context.TODO(), &result)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
