package persistence

import (
	"context"
	"strings"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/specification"
)

func (r *SqlcVetRepository) getTotalCountWithFilters(ctx context.Context, spec *specification.VetSearchSpecification) (int, error) {
	query, params := r.buildCountQuery(spec)

	var totalCount int
	err := r.pool.QueryRow(ctx, query, params...).Scan(&totalCount)
	if err != nil {
		return 0, r.dbError(OpCount, "failed to count veterinarians", err)
	}

	return totalCount, nil
}

func (r *SqlcVetRepository) buildCountQuery(spec *specification.VetSearchSpecification) (string, []any) {
	query, params := spec.ToSQL()

	// Extraer solo la parte WHERE de la consulta original
	whereIndex := strings.Index(query, "WHERE")
	orderByIndex := strings.Index(query, "ORDER BY")
	limitIndex := strings.Index(query, "LIMIT")

	var whereClause string
	if whereIndex != -1 {
		if orderByIndex != -1 {
			whereClause = query[whereIndex:orderByIndex]
		} else if limitIndex != -1 {
			whereClause = query[whereIndex:limitIndex]
		} else {
			whereClause = query[whereIndex:]
		}
	}

	// Construir consulta COUNT
	countQuery := "SELECT COUNT(*) FROM veterinarians"
	if whereClause != "" {
		countQuery += " " + whereClause
	}

	// Remover los parámetros de paginación (los últimos 2 parámetros)
	if len(params) > 2 {
		params = params[:len(params)-2]
	}

	return countQuery, params
}
