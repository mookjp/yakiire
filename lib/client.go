package lib

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// Client is used to access to Firestore documents
type Client struct {
	config    *ClientConfig
	firestore *firestore.Client
	ctx       context.Context
}

// ClientConfig is a configuration to use Firestore Client
type ClientConfig struct {
	Credentials string
}

// NewClient returns a Client to operate data on Firestore
func NewClient(ctx context.Context, config *ClientConfig) (*Client, error) {
	opt := option.WithCredentialsFile(config.Credentials)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize Firestore app")
	}
	f, err := app.Firestore(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize Firestore client")
	}
	return &Client{
		config:    config,
		firestore: f,
		ctx:       ctx,
	}, nil
}

// Get returns a document by docID
func (c *Client) Get(collection string, docID string) (*Doc, error) {
	ctx, cancel := context.WithTimeout(c.ctx, 5*time.Second)
	defer cancel()

	res, err := c.firestore.Collection(collection).Doc(docID).Get(ctx)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("counldn't doc(id: %s) in collection(%s)", docID, collection))
	}
	return &Doc{data: res.Data()}, nil
}
