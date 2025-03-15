package mongo

import (
	"context"
	"fmt"
	"os"

	"first_fiber"
	"github.com/charmbracelet/log"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// MC is equal to MongoClient
// simply return the mongo DB from the path
func MC() *mongo.Client {
	if DB != nil {
		return DB
	}
	dsn := fmt.Sprintf(
		"mongodb://%s:%s@%s:%s/%s",
		first_fiber.MongoUser,
		first_fiber.MongoPassword,
		first_fiber.MongoHost,
		first_fiber.MongoPort,
		first_fiber.MongoName,
	)
	opts.ApplyURI(dsn)

	var err error
	DB, err = mongo.Connect(opts)
	if err != nil {
		log.Fatalf("can not connect to mongo. err: %s", err.Error())
	}
	return DB
}

// MCD is equal  to MongoClientDatabase
// simply use the MC and the default database in .env file
func MCD() *mongo.Database {
	return MC().Database(os.Getenv("MONGO_DB_NAME"))
}

// Checks if given collection name exists or not
// func CollectionExists(collection string, filter interface{}) bool {
// 	collectionNames, err := MCD().ListCollectionNames(context.TODO(), filter)
// 	if err != nil {
// 		fmt.Printf(
// 			"%v%v%v\n",
// 			internal.GetColor("red"),
// 			fmt.Sprintf("err when get collection names, err: %s", err.Error()),
// 			internal.GetColor("reset"),
// 		)
// 	}
//
// 	for _, name := range collectionNames {
// 		if collection == name {
// 			return true
// 		}
// 	}
// 	return false
// }
//
// func CreateCollection(name string) {
// 	err := MCD().CreateCollection(context.TODO(), name)
// 	if err != nil {
// 		fmt.Printf(
// 			"%v%v%v\n",
// 			internal.GetColor("red"),
// 			fmt.Sprintf("failed to create new MONGODB COLLECTION, err: %s", err.Error()),
// 			internal.GetColor("reset"),
// 		)
// 	}
// }

// TODO: complete this function plz!
// func CreateIndexes(coll string, indexModels []internal.Item) {
//
// 	var models []mongo.IndexModel
//
// 	for _, model := range indexModels {
//
// 		for _, v := range model {
// 			models = append(models, mongo.IndexModel{
// 				Keys:    v,
// 				Options: nil,
// 			})
// 		}
// 	}
//
// 	fmt.Println(models)
// 	return
//
// 	indexName, err := MCDColl(coll).Indexes().CreateOne(context.TODO(), mongo.IndexModel{
// 		Keys: internal.Item{
// 			"carID":    1,5
// 			"postTime": -1,
// 		},
// 		Options: nil,
// 	})
// 	indexName, err := MCDColl(coll).Indexes().CreateMany(context.TODO(), models)
// 	if err != nil {
// 		fmt.Printf(
// 			"%v%v%v\n",
// 			internal.GetColor("red"),
// 			fmt.Sprintf("failed to create INDEX for collection %s, err: %s", coll, err.Error()),
// 			internal.GetColor("reset"),
// 		)
// 	}
//
// 	fmt.Println(indexName)
// }

// MCDColl is equal to MongoClientDatabaseCollection
// simply get the given collection from the MCD
func MCDColl(coll string) *mongo.Collection {
	return MCD().Collection(coll)
}

func CountDocuments(coll *mongo.Collection, filter any) int64 {
	count, err := coll.CountDocuments(context.TODO(), filter)
	if err != nil {
		log.Errorf("could not calc count. %s", err)
	}
	return count
}

func FindOne(coll *mongo.Collection, filter any) *mongo.SingleResult {
	return coll.FindOne(context.TODO(), filter)
}

// Aggregate simply do aggregation in a mongo collection
func Aggregate(coll *mongo.Collection, pipelines any) (results []bson.M, err error) {
	cursor, err := coll.Aggregate(context.TODO(), pipelines)
	if err != nil {
		return nil, err
	}

	err = cursor.All(context.TODO(), &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func InsertOne(coll *mongo.Collection, document map[string]any) {
	res, err := coll.InsertOne(context.TODO(), document)
	if err != nil || res == nil {
		log.Warnf("insert one did not saved. %s", err)
		return
	}
	document["_id"] = res.InsertedID
}

func InsertMany(coll *mongo.Collection, documents []any) {
	_, err := coll.InsertMany(context.TODO(), documents)
	if err != nil {
		fmt.Printf(
			"%v\n",
			// internal.GetColor("red"),
			fmt.Sprintf("InsertMany: data did not saved in mongo!. err: %v", err.Error()),
			// internal.GetColor("reset"),
		)
		// custom_log.Fatalf("InsertMany. err: %v\n", err)
	}
}

func UpdateOne(coll *mongo.Collection, filter any, update any) {
	res, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil || res == nil || res.MatchedCount == 0 {
		log.Warnf("update one did not saved. %s", err)
	}
}
