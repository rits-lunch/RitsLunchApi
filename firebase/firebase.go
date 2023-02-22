package firebase

import (
	"fmt"

	firebase "firebase.google.com/go"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

type shop struct {
	shopName    string
	isActive    bool
	description string
	items       []string
}

func Test() {
	// クライアント接続
	opt := option.WithCredentialsFile("key.json")
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil, opt)

	if err != nil {
		// fmt.Errorf("error initializing app: %v", err)
		fmt.Print("Hello world!")
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		// fmt.Errorf("error initializing client: %v", err)
		fmt.Print("Hello world!")
	}
	defer client.Close()
	fmt.Println("Connection done")
	// 値の取得
	collection := client.Collection("shop")
	doc := collection.Doc("0001") //.Collection("foods").Doc("0001")
	//collection = client.Collection("item")
	//doc = collection.Doc("EIsr8JFtrvSMk5DVydn6")
	field, err := doc.Get(ctx)
	if err != nil {
		// fmt.Errorf("error get data: %v", err)
		fmt.Print("Hello world!")
	}

	//var dataShop shop
	dataShop := field.Data()
	for key, value := range dataShop {
		fmt.Printf("key: %v, value: %v\n", key, value)
	}

	collectionItem := client.Collection("items")
	docItem := collectionItem.Doc((dataShop.(shop)).items[0]) //.Collection("foods").Doc("0001")
	//collection = client.Collection("item")
	//doc = collection.Doc("EIsr8JFtrvSMk5DVydn6")
	fieldItem, errItem := docItem.Get(ctx)
	if errItem != nil {
		// fmt.Errorf("error get data: %v", err)
		fmt.Print("Hello world!")
	}
	dataItem := fieldItem.Data()
	for key, value := range dataItem {
		fmt.Printf("key: %v, value: %v\n", key, value)
	}
}
