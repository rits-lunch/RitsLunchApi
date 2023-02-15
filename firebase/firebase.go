package firebase

import (
	"fmt"

	firebase "firebase.google.com/go"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

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
	//collection2 := client.Collection("foods")
	//doc := collection2.Doc("0001")
	field, err := doc.Get(ctx)
	if err != nil {
		// fmt.Errorf("error get data: %v", err)
		fmt.Print("Hello world!")
	}
	data := field.Data()
	for key, value := range data {
		fmt.Printf("key: %v, value: %v\n", key, value)
	}
}
