package lib

import (
	"context"
	"google.golang.org/api/option"
	"time"

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
	ProjectID string
}

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
