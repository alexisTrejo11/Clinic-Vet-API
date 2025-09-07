// Package repositoryimpl implements the NotificationRepository interface using MongoDB as the data store.
package repositoryimpl

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/notification"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoNotificationRepository struct {
	client                 *mongo.Client
	notificationCollection *mongo.Collection
}

func NewMongoNotificationRepository(client *mongo.Client) repository.NotificationRepository {
	return &MongoNotificationRepository{
		client:                 client,
		notificationCollection: client.Database("notifications").Collection("notifications"),
	}
}

func (r *MongoNotificationRepository) Create(ctx context.Context, notification *notification.Notification) error {
	result, err := r.notificationCollection.InsertOne(ctx, notification)
	if err != nil {
		return err
	}

	notification.ID = result.InsertedID.(bson.ObjectID).Hex()

	return nil
}

func (r *MongoNotificationRepository) GetByID(ctx context.Context, id string) (notification.Notification, error) {
	var notification notification.Notification
	if err := r.notificationCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&notification); err != nil {
		return notification, err
	}
	return notification, nil
}

func (r *MongoNotificationRepository) ListByUser(ctx context.Context, userID valueobject.UserID, pagination page.PageInput) (page.Page[[]notification.Notification], error) {
	findOptions := r.paginateSearch(pagination)
	queryFilter := bson.M{"user_id": userID.String()}

	notifications, err := r.findByFilter(ctx, queryFilter, findOptions)
	if err != nil {
		return page.EmptyPage[[]notification.Notification](), err
	}

	pageMetadata, err := r.getPaginationMetadata(ctx, queryFilter, pagination)
	if err != nil {
		return page.EmptyPage[[]notification.Notification](), err
	}

	return page.NewPage(notifications, pageMetadata), nil
}

func (r *MongoNotificationRepository) ListByType(ctx context.Context, notificationType string, pagination page.PageInput) (page.Page[[]notification.Notification], error) {
	findOptions := r.paginateSearch(pagination)
	queryFilter := bson.M{"type": notificationType}

	notifications, err := r.findByFilter(ctx, queryFilter, findOptions)
	if err != nil {
		return page.EmptyPage[[]notification.Notification](), err
	}

	pageMetadata, err := r.getPaginationMetadata(ctx, queryFilter, pagination)
	if err != nil {
		return page.EmptyPage[[]notification.Notification](), err
	}

	return page.NewPage(notifications, pageMetadata), nil
}

func (r *MongoNotificationRepository) ListByChannel(ctx context.Context, channel string, pagination page.PageInput) (page.Page[[]notification.Notification], error) {
	findOptions := r.paginateSearch(pagination)
	queryFilter := bson.M{"channel": channel}

	notifications, err := r.findByFilter(ctx, queryFilter, findOptions)
	if err != nil {
		return page.EmptyPage[[]notification.Notification](), err
	}

	pageMetadata, err := r.getPaginationMetadata(ctx, queryFilter, pagination)
	if err != nil {
		return page.EmptyPage[[]notification.Notification](), err
	}

	return page.NewPage(notifications, pageMetadata), nil
}

func (r *MongoNotificationRepository) paginateSearch(pagination page.PageInput) *options.FindOptionsBuilder {
	pageSize := int64(pagination.PageSize)
	pageNumber := int64(pagination.PageNumber)

	offset := (pageNumber - 1) * pageSize

	if offset < 0 {
		offset = 0
	}

	findOptions := options.Find().SetLimit(pageSize).SetSkip(offset)

	return findOptions
}

func (r *MongoNotificationRepository) countQueryTotalItems(ctx context.Context, queryFilter bson.M) (int64, error) {
	totalItems, err := r.notificationCollection.CountDocuments(ctx, queryFilter)
	if err != nil {
		return 0, err
	}

	return totalItems, nil
}

func (r *MongoNotificationRepository) getPaginationMetadata(ctx context.Context, queryFilter bson.M, pagination page.PageInput) (page.PageMetadata, error) {
	totalItems, err := r.countQueryTotalItems(ctx, queryFilter)
	if err != nil {
		return page.PageMetadata{}, err
	}

	pageMetadata := *page.GetPageMetadata(int(totalItems), pagination)
	return pageMetadata, nil
}

func (r *MongoNotificationRepository) findByFilter(ctx context.Context,
	queryFilter bson.M,
	findBuilderOptions *options.FindOptionsBuilder,
) ([]notification.Notification, error) {
	var notifications []notification.Notification
	cursor, err := r.notificationCollection.Find(ctx, queryFilter, findBuilderOptions)
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var noti notification.Notification
		if err := cursor.Decode(&noti); err != nil {
			return []notification.Notification{}, err
		}
		notifications = append(notifications, noti)
	}

	return notifications, nil
}
