package repositoryimpl

import (
	"errors"
	"fmt"

	"clinic-vet-api/app/core/domain/entity/user"
	"clinic-vet-api/app/core/domain/enum"
	"clinic-vet-api/app/core/domain/valueobject"
	"clinic-vet-api/sqlc"
)

func MapUserFromSQLC(sqlRow sqlc.User) (*user.User, error) {
	userID, err := valueobject.NewUserID(int(sqlRow.ID))
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %w", err)
	}

	var email valueobject.Email
	if sqlRow.Email.Valid && sqlRow.Email.String != "" {
		email, err = valueobject.NewEmail(sqlRow.Email.String)
		if err != nil {
			return nil, fmt.Errorf("invalid email: %w", err)
		}
	}

	var phone valueobject.PhoneNumber
	if sqlRow.PhoneNumber.Valid && sqlRow.PhoneNumber.String != "" {
		phone, err = valueobject.NewPhoneNumber(sqlRow.PhoneNumber.String)
		if err != nil {
			return nil, fmt.Errorf("invalid phone number: %w", err)
		}
	}

	var role enum.UserRole
	roleVal, err := sqlRow.Role.Value()
	if err != nil {
		return nil, fmt.Errorf("failed to get role value: %w", err)
	}

	roleStr, ok := roleVal.(string)
	if !ok {
		return nil, errors.New("user role value is not a string")
	}

	role, err = enum.ParseUserRole(roleStr)
	if err != nil {
		return nil, fmt.Errorf("invalid user role: %w", err)
	}

	var status enum.UserStatus
	statusVal, err := sqlRow.Status.Value()
	if err != nil {
		return nil, fmt.Errorf("failed to get status value: %w", err)
	}

	statusStr, ok := statusVal.(string)
	if !ok {
		return nil, errors.New("user status value is not a string")
	}

	status, err = enum.ParseUserStatus(statusStr)
	if err != nil {
		return nil, fmt.Errorf("invalid user status: %w", err)
	}

	// Validar que la contraseña no esté vacía
	if !sqlRow.Password.Valid || sqlRow.Password.String == "" {
		return nil, errors.New("password is required")
	}

	// Crear options
	opts := []user.UserOption{
		user.WithEmail(email),
		user.WithPhoneNumber(phone),
		user.WithPassword(sqlRow.Password.String),
	}
	/*
		// Campos opcionales
		if sqlRow.LastLoginAt.Valid {
			opts = append(opts, user.WithLastLoginAt(sqlRow.LastLoginAt.Time))
		}

		if sqlRow.JoinedAt.Valid {
			opts = append(opts, user.WithJoinedAt(sqlRow.JoinedAt.Time))
		} else {
			// Si joinedAt es NULL, usar created_at como fallback
			if sqlRow.CreatedAt.Valid {
				opts = append(opts, user.WithJoinedAt(sqlRow.CreatedAt.Time))
			}
		}

		// Campos de auditoría
		if sqlRow.CreatedAt.Valid {
			opts = append(opts, user.WithCreatedAt(sqlRow.CreatedAt.Time))
		}

		if sqlRow.UpdatedAt.Valid {
			opts = append(opts, user.WithUpdatedAt(sqlRow.UpdatedAt.Time))
		}

		// 2FA settings (si están disponibles en la base de datos)
		if sqlRow.TwoFactorEnabled.Valid {
			twoFAOpts, err := getTwoFactorAuthOptions(sqlRow)
			if err != nil {
				return nil, fmt.Errorf("failed to get 2FA options: %w", err)
			}
			opts = append(opts, twoFAOpts...)
		}
	*/

	userEntity, err := user.NewUser(
		userID,
		role,
		status,
		opts...,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return userEntity, nil
}

/*
// Helper para opciones de 2FA
func getTwoFactorAuthOptions(sqlRow sqlc.User) ([]user.UserOption, error) {
	var opts []user.UserOption

	if sqlRow.TwoFactorEnabled.Valid && sqlRow.TwoFactorEnabled.Bool {
		// Si 2FA está habilitado, necesitamos el método y secreto
		if !sqlRow.TwoFactorMethod.Valid || sqlRow.TwoFactorMethod.String == "" {
			return nil, errors.New("2FA method is required when 2FA is enabled")
		}

		if !sqlRow.TwoFactorSecret.Valid || sqlRow.TwoFactorSecret.String == "" {
			return nil, errors.New("2FA secret is required when 2FA is enabled")
		}

		method := sqlRow.TwoFactorMethod.String
		secret := sqlRow.TwoFactorSecret.String

		// Parsear backup codes si están disponibles
		var backupCodes []string
		if sqlRow.TwoFactorBackupCodes.Valid && sqlRow.TwoFactorBackupCodes.String != "" {
			if err := json.Unmarshal([]byte(sqlRow.TwoFactorBackupCodes.String), &backupCodes); err != nil {
				return nil, fmt.Errorf("failed to parse 2FA backup codes: %w", err)
			}
		}

		twoFA := auth.NewTwoFactorAuth(true, method, secret, backupCodes)
		opts = append(opts, user.WithTwoFactorAuth(twoFA))
	}

	return opts, nil
}
*/

func MapUsersFromSQLC(sqlRows []sqlc.User) ([]user.User, error) {
	users := make([]user.User, 0, len(sqlRows))
	for i, sqlRow := range sqlRows {
		user, err := MapUserFromSQLC(sqlRow)
		if err != nil {
			return nil, err
		}
		users[i] = *user
	}

	return users, nil
}
