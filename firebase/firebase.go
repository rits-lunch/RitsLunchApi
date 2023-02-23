package firebase

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

type shop struct {
	ShopName    string   `firestore:"shopName"`
	IsActive    bool     `firestore:"isActive"`
	Description string   `firestore:"description"`
	Items       []string `firestore:"items"`
}
type shopitem struct {
	Image string `firestore:"image"`
	Price int    `firestore:"price"`
	Name  string `firestore:"name"`
}

func Test() {
	// クライアント接続
	ctx := context.Background()
	opt := option.WithCredentialsFile("key.json")
	client, err := firestore.NewClient(ctx, "ritslunch-f747e", opt)
	if err != nil {
		fmt.Printf("error get data: %v", err)
	}
	collection := client.Collection("shop")
	datasnap, err := collection.Doc("0001").Get(ctx)
	// fmt.Printf("%#v",datasnap.Data())
	if err != nil {
		fmt.Printf("error get data: %v", err)
	}
	var shop = shop{}
	if err := datasnap.DataTo(&shop); err != nil {
		fmt.Printf("error get data: %v", err)
	}
	fmt.Printf("%#v \n", shop)

	for _, itemid := range shop.Items {
		var item = shopitem{}
		collection := client.Collection("items")
		datasnap, err := collection.Doc(itemid).Get(ctx)
		if err != nil {
			fmt.Printf("error get data: %v", err)
		}
		if err := datasnap.DataTo(&item); err != nil {

		}
		fmt.Printf("%#v \n", item)
	}
}
