package lib

import (
	"context"
	"time"

	"google.golang.org/api/option"

	"cloud.google.com/go/firestore"
)

// Client is used to access to Firestore documents
type Client struct {
	config    *ClientConfig
	firestore *firestore.Client
}

// ClientConfig is a configuration to use Firestore Client
type ClientConfig struct {
	Credentials string
	ProjectID   string
}

// Condition is used for Firesstore query
type Condition struct {
	Path  string
	Op    string
	Value interface{}
}

const defaultLimit = 20

// NewClient returns a Client to operate data on Firestore
func NewClient(ctx context.Context, config *ClientConfig) (*Client, error) {
	var client *firestore.Client
	if config.Credentials != "" {
		f, err := firestore.NewClient(ctx, config.ProjectID, option.WithCredentialsFile(config.Credentials))
		if err != nil {
			return nil, err
		}
		client = f
	} else {
		f, err := firestore.NewClient(ctx, config.ProjectID)
		if err != nil {
			return nil, err
		}
		client = f
	}
	return &Client{
		config:    config,
		firestore: client,
	}, nil
}

// Get returns a document by docID
func (c *Client) Get(ctx context.Context, collection string, docID string) (*Doc, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	res, err := c.firestore.Collection(collection).Doc(docID).Get(ctx)
	if err != nil {
		return nil, err
	}
	return &Doc{data: res.Data()}, nil
}

// Query returns documents matched with conditions and limit the results
func (c *Client) Query(ctx context.Context, collection string, conditions []*Condition, limit int) ([]*Doc, error) {
	collectionRef := c.firestore.Collection(collection)
	var query firestore.Query
	for _, condition := range conditions {
		query = collectionRef.Where(condition.Path, condition.Op, condition.Value)
	}
	if limit != 0 {
		query = query.Limit(limit)
	} else {
		query = query.Limit(defaultLimit)
	}
	docs, err := query.Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}
	res := make([]*Doc, 0)
	for _, doc := range docs {
		res = append(res, &Doc{data: doc.Data()})
	}
	return res, nil
}

// Add a document into a collection
func (c *Client) Add(ctx context.Context, collection string, document map[string]interface{}) (*Doc, error) {
	collectionRef := c.firestore.Collection(collection)
	res, _, err := collectionRef.Add(ctx, document)
	if err != nil {
		return nil, err
	}
	doc, err := res.Get(ctx)
	if err != nil {
		return nil, err
	}
	return &Doc{data: doc.Data()}, nil
}

// Delete a document from a collection
func (c *Client) Delete(ctx context.Context, collection string, docID string) error {
	collectionRef := c.firestore.Collection(collection)
	_, err := collectionRef.Doc(docID).Delete(ctx)
	return err
}
