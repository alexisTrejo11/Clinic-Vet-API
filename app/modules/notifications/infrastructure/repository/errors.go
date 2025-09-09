package repositoryimpl

import (
	"fmt"

	dberr "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/infrastructure/database"
	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionNotifications = "notifications"
	DriverMongoDB           = "mongodb"

	ErrMsgCreateNotification  = "failed to create notification"
	ErrMsgGetNotification     = "failed to get notification"
	ErrMsgListNotifications   = "failed to list notifications"
	ErrMsgCountNotifications  = "failed to count notifications"
	ErrMsgConvertNotification = "failed to convert notification"
	ErrMsgDecodeNotification  = "failed to decode notification"
)

// Operaciones específicas de MongoDB
const (
	OpInsertMongo = "insert"
	OpFindMongo   = "find"
	OpCountMongo  = "count"
	OpDecodeMongo = "decode"
)

// mongoError crea un error estandarizado para operaciones de MongoDB
func (r *MongoNotificationRepository) mongoError(operation, message string, err error) error {
	return dberr.DatabaseOperationError(operation, CollectionNotifications, DriverMongoDB, fmt.Sprintf("%s: %v", message, err))
}

// notFoundError crea un error estandarizado de entidad no encontrada
func (r *MongoNotificationRepository) notFoundError(parameterName, parameterValue string) error {
	return dberr.EntityNotFoundError(parameterName, parameterValue, OpFindMongo, CollectionNotifications, DriverMongoDB)
}

// convertStringToObjectID convierte un string ID a ObjectID de MongoDB
func convertStringToObjectID(id string) (bson.ObjectID, error) {
	var objectID bson.ObjectID
	err := objectID.UnmarshalText([]byte(id))
	if err != nil {
		return bson.ObjectID{}, fmt.Errorf("invalid ObjectID format: %w", err)
	}
	return objectID, nil
}
