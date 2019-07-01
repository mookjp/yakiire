package test

import (
	"context"
	"fmt"

	"google.golang.org/api/iterator"

	"cloud.google.com/go/firestore"
)

// Helper is used for testing as a helper
type Helper struct {
	client *firestore.Client
}

// Product is type for test document
type Product struct {
	ID          string
	Name        string
	CategoryIDs []string
	Attributes  map[string]interface{}
}

const products = "products"

// NewHelper is a Helper for testing
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

// CreateData inserts test documents
func (h *Helper) CreateData() {
	ctx := context.Background()
	products := h.client.Collection(products)
	createProduct1(ctx, products)
	createProduct2(ctx, products)
	createProduct3(ctx, products)
}

// DeleteAll deletes all documents for test
func (h *Helper) DeleteAll() error {
	ctx := context.Background()
	iter := h.client.Collection(products).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err != nil {
			if err == iterator.Done {
				break
			} else {
				return err
			}
		}
		if _, err := doc.Ref.Delete(ctx); err != nil {
			return err
		}
	}
	return nil
}

// Close closes the connection
func (h *Helper) Close() error {
	return h.client.Close()
}

func createProduct1(ctx context.Context, products *firestore.CollectionRef) {
	res, err := products.Doc("1").Set(ctx, Product{
		ID:          "1",
		Name:        "Test Product",
		CategoryIDs: []string{"1", "2", "3"},
		Attributes: map[string]interface{}{
			"color": "red",
			"size":  100,
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("docs were inserted. res: %s\n", res)
}

func createProduct2(ctx context.Context, products *firestore.CollectionRef) {
	res, err := products.Doc("2").Set(ctx, Product{
		ID:          "2",
		Name:        "Another Test Product",
		CategoryIDs: []string{"3", "4", "5"},
		Attributes: map[string]interface{}{
			"color": "blue",
			"size":  200,
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("docs were inserted. res: %s\n", res)
}

func createProduct3(ctx context.Context, products *firestore.CollectionRef) {
	res, err := products.Doc("3").Set(ctx, Product{
		ID:          "3",
		Name:        "Another Great Test Product",
		CategoryIDs: []string{"5", "6", "7"},
		Attributes: map[string]interface{}{
			"color": "yellow",
			"size":  300,
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("docs were inserted. res: %s\n", res)
}
