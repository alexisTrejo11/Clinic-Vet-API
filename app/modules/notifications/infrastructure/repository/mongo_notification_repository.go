// Package repositoryimpl implements the NotificationRepository interface using MongoDB as the data store.
package repositoryimpl

import (
	"context"
	"errors"
	"fmt"

	"clinic-vet-api/app/core/domain/entity/notification"
	"clinic-vet-api/app/core/domain/valueobject"
	"clinic-vet-api/app/core/repository"
	"clinic-vet-api/app/shared/page"

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
		return r.mongoError(OpInsertMongo, ErrMsgCreateNotification, err)
	}

	// Asumiendo que el ID se genera como ObjectID en MongoDB
	if oid, ok := result.InsertedID.(bson.ObjectID); ok {
		notification.ID = oid.Hex()
	} else {
		return r.mongoError(OpInsertMongo, "failed to get inserted notification ID", errors.New("invalid inserted ID type"))
	}

	return nil
}

func (r *MongoNotificationRepository) GetByID(ctx context.Context, id string) (notification.Notification, error) {
	var notification notification.Notification

	// Convertir el string ID a ObjectID si es necesario
	objectID, err := convertStringToObjectID(id)
	if err != nil {
		return notification, r.mongoError(OpFindMongo, "invalid notification ID format", err)
	}

	err = r.notificationCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&notification)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return notification, r.notFoundError("id", id)
		}
		return notification, r.mongoError(OpFindMongo, fmt.Sprintf("%s with ID %s", ErrMsgGetNotification, id), err)
	}

	return notification, nil
}

func (r *MongoNotificationRepository) ListByUser(ctx context.Context, userID valueobject.UserID, pagination page.PageInput) (page.Page[notification.Notification], error) {
	findOptions := r.paginateSearch(pagination)
	queryFilter := bson.M{"user_id": userID.String()}

	notifications, err := r.findByFilter(ctx, queryFilter, findOptions)
	if err != nil {
		return page.EmptyPage[notification.Notification](), err
	}

	pageMetadata, err := r.getPaginationMetadata(ctx, queryFilter, pagination)
	if err != nil {
		return page.EmptyPage[notification.Notification](), err
	}

	return page.NewPage(notifications, pageMetadata), nil
}

func (r *MongoNotificationRepository) ListByType(ctx context.Context, notificationType string, pagination page.PageInput) (page.Page[notification.Notification], error) {
	findOptions := r.paginateSearch(pagination)
	queryFilter := bson.M{"type": notificationType}

	notifications, err := r.findByFilter(ctx, queryFilter, findOptions)
	if err != nil {
		return page.EmptyPage[notification.Notification](), err
	}

	pageMetadata, err := r.getPaginationMetadata(ctx, queryFilter, pagination)
	if err != nil {
		return page.EmptyPage[notification.Notification](), err
	}

	return page.NewPage(notifications, pageMetadata), nil
}

func (r *MongoNotificationRepository) ListByChannel(ctx context.Context, channel string, pagination page.PageInput) (page.Page[notification.Notification], error) {
	findOptions := r.paginateSearch(pagination)
	queryFilter := bson.M{"channel": channel}

	notifications, err := r.findByFilter(ctx, queryFilter, findOptions)
	if err != nil {
		return page.EmptyPage[notification.Notification](), err
	}

	pageMetadata, err := r.getPaginationMetadata(ctx, queryFilter, pagination)
	if err != nil {
		return page.EmptyPage[notification.Notification](), err
	}

	return page.NewPage(notifications, pageMetadata), nil
}

func (r *MongoNotificationRepository) paginateSearch(pagination page.PageInput) *options.FindOptionsBuilder {
	pageSize := int64(pagination.PageSize)
	pageNumber := int64(pagination.Page)

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
		return 0, r.mongoError(OpCountMongo, ErrMsgCountNotifications, err)
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
		return nil, r.mongoError(OpFindMongo, ErrMsgListNotifications, err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var noti notification.Notification
		if err := cursor.Decode(&noti); err != nil {
			return nil, r.mongoError(OpDecodeMongo, ErrMsgDecodeNotification, err)
		}
		notifications = append(notifications, noti)
	}

	if err := cursor.Err(); err != nil {
		return nil, r.mongoError(OpFindMongo, "cursor error while listing notifications", err)
	}

	return notifications, nil
}
