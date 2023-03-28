package purchase

import (
	"context"
	coffeeco "example/coffeeco/internal"
	"example/coffeeco/internal/payment"
	"example/coffeeco/internal/store"
	"fmt"
	"time"

	"github.com/Rhymond/go-money"
	"go.mongodb.org/mongo-driver/internal/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	Store(ctx context.Context, purchase Purchase) error
}

type MongoRepository struct {
	purchases *mongo.Collection
}

type mongoPurchase struct {
	id                 uuid.UUID
	store              store.Store
	productsToPurchase []coffeeco.Product
	total              money.Money
	paymentMeans       payment.Means
	timeOfPurchase     time.Time
	cardToken          *string
}

func NewMongoRepo(ctx context.Context, connectionString string) (*MongoRepository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))

	if err != nil {
		return nil, fmt.Errorf("failed to create mongo client: %v", err)
	}

	purchases := client.Database("coffeeco").Collection("purchases")

	return &MongoRepository{purchases: purchases}, nil
}

func (mr *MongoRepository) Store(ctx context.Context, purchase Purchase) error {
	mongopP := new(Purchase)

	_, err := mr.purchases.InsertOne(ctx, mongopP)

	if err != nil {
		return fmt.Errorf("failed to persist purchase: %v", err)
	}

	return nil
}

func toMongoPurchase(p Purchase) mongoPurchase {
	return mongoPurchase{
		id:                 p.ID,
		store:              p.Store,
		productsToPurchase: p.ProductsToPurchase,
		total:              p.total,
		paymentMeans:       p.PaymentMeans,
		timeOfPurchase:     p.timeOfPurchase,
		cardToken:          p.CardToken,
	}
}

func (mr *MongoRepository) Ping(ctx context.Context) error {
	if _, err := mr.purchases.EstimatedDocumentCount(ctx); err != nil {
		return fmt.Errorf("failed to ping DB: %w", err)
	}
	return nil
}
