package models

import (
	"context"
	"errors"
	"time"

	"github.com/kotopanjang/url_shortener/helper"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (e *ShortURL) TableName() string {
	return "ShortUrl"
}

func (su *ShortURL) Insert() (*ShortURL, error) {
	helper.Println("Inserting data to database")
	db, err := helper.ConnectDB()
	if err != nil {
		helper.Println(err)
		return nil, err
	}

	collection := db.Collection(su.TableName())

	su.Id = primitive.NewObjectID()
	su.MaskedUrl, err = su.GetRandomStr()
	su.CreateDate = time.Now()
	su.ExpiredDate = su.CreateDate.AddDate(0, 0, 1) //default 1 day

	res, err := collection.InsertOne(context.TODO(), su)
	if err != nil {
		helper.Println(err)
		return nil, err
	}
	helper.Println("Data Inserted. Doc ID ", res)

	defer collection.Database().Client().Disconnect(context.TODO())

	return su, nil
}

func (su *ShortURL) GetRandomStr() (string, error) {
	helper.Println("Generating Random String ...")
	ListMaskedUrl := []string{}
	//prepare 25 set random str in case of the high load. and prevent from multiple checking on database
	for i := 0; i < 25; i++ {
		ListMaskedUrl = append(ListMaskedUrl, helper.RandomStr())
	}

	bsonTT := []bson.D{}
	bsonTT = append(bsonTT, bson.D{{"ExpiredDate", bson.D{{"$gte", time.Now()}}}})
	bsonTT = append(bsonTT, bson.D{{"listMasked", bson.D{{"$in", ListMaskedUrl}}}})
	match := bson.D{
		{
			"$match", bson.D{
				{
					"$or", bsonTT,
				},
			},
		},
	}

	group := bson.D{
		{
			"$group", bson.D{
				{"_id", ""},
				{"listMasked", bson.D{{"$addToSet", "$MaskedUrl"}}},
			},
		},
	}
	db, err := helper.ConnectDB()
	if err != nil {
		helper.Println(err)
		return "", nil
	}
	collection := db.Collection(su.TableName())
	showInfoCursor, err := collection.Aggregate(context.TODO(), mongo.Pipeline{match, group})
	if err != nil {
		helper.Println(err)
		return "", nil
	}
	defer collection.Database().Client().Disconnect(context.TODO())

	var showsWithInfo []bson.M
	if err = showInfoCursor.All(context.TODO(), &showsWithInfo); err != nil {
		helper.Println(err)
		return "", nil
	}

	exist := []string{}
	for _, xx := range showsWithInfo {
		tt := xx["listMasked"].(primitive.A)
		for _, yy := range tt {
			exist = append(exist, yy.(string))
		}
	}

	finalRes := helper.StringDiff(ListMaskedUrl, exist)
	if len(finalRes) < 0 {
		helper.Println("Random String conflict with the others")
		helper.Println("Re-Generating random string ...")
		final, err := su.GetRandomStr()
		if err != nil {
			helper.Println(err)
			return "", nil
		} else {
			return final, nil
		}
	} else {
		firsrt := finalRes[0]
		return firsrt, nil
	}
}

func (su *ShortURL) CheckExistingData() *ShortURL {
	helper.Println("Checking existing URL")
	db, err := helper.ConnectDB()
	if err != nil {
		helper.Println(err)
		return nil
	}

	collection := db.Collection(su.TableName())
	filter := bson.M{"OriginalUrl": su.OriginalUrl, "ExpiredDate": bson.M{"$gte": time.Now()}}
	res, _ := collection.Find(context.TODO(), filter)

	ss := []ShortURL{}
	if err = res.All(context.TODO(), &ss); err != nil {
		helper.Println(err)
		return nil
	}
	defer collection.Database().Client().Disconnect(context.TODO())

	if len(ss) > 0 {
		helper.Println("URL ", su.OriginalUrl, " found on database")
		helper.Println("Retrieving data from database")
		return &ss[0]
	} else {
		helper.Println(su.OriginalUrl, " not found on the database or expired")
		return nil
	}
}

func (su *ShortURL) Retrieve(param string) (string, error) {
	helper.Println("Retrieving Original URL from database...")
	db, err := helper.ConnectDB()
	if err != nil {
		helper.Println(err)
		return "", err
	}
	collection := db.Collection(su.TableName())
	filter := bson.M{"MaskedUrl": param, "ExpiredDate": bson.M{"$gte": time.Now()}}
	res, _ := collection.Find(context.TODO(), filter)

	ss := []ShortURL{}
	if err = res.All(context.TODO(), &ss); err != nil {
		helper.Println(err)
		return "", err
	}

	defer collection.Database().Client().Disconnect(context.TODO())

	if len(ss) > 0 {
		helper.Println("Short URL found on database")
		return ss[0].OriginalUrl, nil
	} else {
		helper.Println("Short URL not found on the database or expired")
		return "", errors.New("Invalid url")
	}
}

type ShortURL struct {
	Id          primitive.ObjectID `bson:"_id" , json:"_id" `
	OriginalUrl string             `bson:"OriginalUrl" , json:"OriginalUrl" `
	MaskedUrl   string             `bson:"MaskedUrl" , json:"MaskedUrl" `
	CreateDate  time.Time          `bson:"CreateDate" , json:"CreateDate" `
	ExpiredDate time.Time          `bson:"ExpiredDate" , json:"ExpiredDate" `
}
