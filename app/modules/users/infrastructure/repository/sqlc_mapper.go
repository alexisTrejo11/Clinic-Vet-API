package repositoryimpl

import (
	"errors"

	"clinic-vet-api/app/core/domain/entity/user"
	"clinic-vet-api/sqlc"
)

func sqlcRowToEntity(sqlRow sqlc.User) (*user.User, error) {
	return nil, errors.New("not implemented")
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

func sqlcRowsToEntities(sqlRows []sqlc.User) ([]user.User, error) {
	users := make([]user.User, 0, len(sqlRows))
	for i, sqlRow := range sqlRows {
		user, err := sqlcRowToEntity(sqlRow)
		if err != nil {
			return nil, err
		}
		users[i] = *user
	}

	return users, nil
}
