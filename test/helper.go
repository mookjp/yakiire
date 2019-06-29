package test

import (
	"context"
	"fmt"

	"google.golang.org/api/iterator"

	"cloud.google.com/go/firestore"
)

type Helper struct {
	client *firestore.Client
}

type Product struct {
	ID          string
	Name        string
	CategoryIDs []string
	Attributes  map[string]string
}

const products = "products"

func NewHelper() *Helper {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "yakiire")
	if err != nil {
		panic(err)
	}
	return &Helper{
		client: client,
	}
}

func (h *Helper) CreateData() {
	ctx := context.Background()

	products := h.client.Collection(products)
	res, err := products.Doc("").Set(ctx, Product{
		ID:          "1",
		Name:        "Test Product",
		CategoryIDs: []string{"1", "2", "3"},
		Attributes: map[string]string{
			"color": "red",
			"size":  "100",
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("docs were inserted. res: %s", res)
}

func (h *Helper) DeleteAll() {
	ctx := context.Background()
	iter := h.client.Collection(products).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if _, err := doc.Ref.Delete(ctx); err != nil {
			panic(err)
		}
	}
}
