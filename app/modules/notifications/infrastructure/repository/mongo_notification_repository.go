// Package repositoryimpl implements the NotificationRepository interface using MongoDB as the data store.
package repositoryimpl

import (
	"context"
	"errors"
	"fmt"

	noti "clinic-vet-api/app/modules/core/domain/entity/notification"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/modules/core/repository"
	p "clinic-vet-api/app/shared/page"

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

func (r *MongoNotificationRepository) Create(ctx context.Context, notification *noti.Notification) error {
	result, err := r.notificationCollection.InsertOne(ctx, notification)
	if err != nil {
		return r.mongoError(OpInsertMongo, ErrMsgCreateNotification, err)
	}

	if oid, ok := result.InsertedID.(bson.ObjectID); ok {
		notification.ID = oid.Hex()
	} else {
		return r.mongoError(OpInsertMongo, "failed to get inserted notification ID", errors.New("invalid inserted ID type"))
	}

	return nil
}

func (r *MongoNotificationRepository) GetByID(ctx context.Context, id string) (noti.Notification, error) {
	var notification noti.Notification

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

func (r *MongoNotificationRepository) ListByUser(ctx context.Context, userID valueobject.UserID, pagination p.PaginationRequest) (p.Page[noti.Notification], error) {
	findOptions := r.paginateSearch(pagination)
	queryFilter := bson.M{"user_id": userID.String()}

	notifications, err := r.findByFilter(ctx, queryFilter, findOptions)
	if err != nil {
		return p.Page[noti.Notification]{}, err
	}

	totalItems, err := r.countQueryTotalItems(ctx, queryFilter)
	if err != nil {
		return p.Page[noti.Notification]{}, err
	}

	return p.NewPage(notifications, int(totalItems), pagination), nil
}

func (r *MongoNotificationRepository) ListBySpecies(ctx context.Context, notificationType string, pagination p.PaginationRequest) (p.Page[noti.Notification], error) {
	findOptions := r.paginateSearch(pagination)
	queryFilter := bson.M{"type": notificationType}

	notifications, err := r.findByFilter(ctx, queryFilter, findOptions)
	if err != nil {
		return p.Page[noti.Notification]{}, err
	}

	totalItems, err := r.countQueryTotalItems(ctx, queryFilter)
	if err != nil {
		return p.Page[noti.Notification]{}, err
	}

	return p.NewPage(notifications, int(totalItems), pagination), nil
}

func (r *MongoNotificationRepository) ListByChannel(ctx context.Context, channel string, pagination p.PaginationRequest) (p.Page[noti.Notification], error) {
	findOptions := r.paginateSearch(pagination)
	queryFilter := bson.M{"channel": channel}

	notifications, err := r.findByFilter(ctx, queryFilter, findOptions)
	if err != nil {
		return p.Page[noti.Notification]{}, err
	}

	totalItems, err := r.countQueryTotalItems(ctx, queryFilter)
	if err != nil {
		return p.Page[noti.Notification]{}, err
	}

	return p.NewPage(notifications, int(totalItems), pagination), nil
}

func (r *MongoNotificationRepository) paginateSearch(pagination p.PaginationRequest) *options.FindOptionsBuilder {
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

func (r *MongoNotificationRepository) findByFilter(ctx context.Context,
	queryFilter bson.M,
	findBuilderOptions *options.FindOptionsBuilder,
) ([]noti.Notification, error) {
	var notifications []noti.Notification

	cursor, err := r.notificationCollection.Find(ctx, queryFilter, findBuilderOptions)
	if err != nil {
		return nil, r.mongoError(OpFindMongo, ErrMsgListNotifications, err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var noti noti.Notification
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
