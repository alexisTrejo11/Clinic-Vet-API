package repository

import (
	"clinic-vet-api/app/modules/core/domain/specification"
	"context"
	"strings"
)

func (r *SqlcEmployeeRepository) getTotalCountWithFilters(ctx context.Context, spec *specification.EmployeeSearchSpecification) (int, error) {
	query, params := r.buildCountQuery(spec)

	var totalCount int
	err := r.pool.QueryRow(ctx, query, params...).Scan(&totalCount)
	if err != nil {
		return 0, r.dbError(OpCount, "failed to count veterinarians", err)
	}

	return totalCount, nil
}

func (r *SqlcEmployeeRepository) buildCountQuery(spec *specification.EmployeeSearchSpecification) (string, []any) {
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

/*
func (r *SqlcEmployeeRepository) Search(ctx context.Context, spec specification) (page.Page[[]customer.Customer], error) {
	query, params := spec.ToSQL()

	// Execute the query
	rows, err := r.pool.Query(ctx, query, params...)
	if err != nil {
		return page.Page[[]customer.Customer]{}, r.dbError(OpSelect, "failed to search veterinarians", err)
	}
	defer rows.Close()

	// Iterate through the rows and scan into Customer structs
	var vets []customer.Customer
	for rows.Next() {
		var veterinarian customer.Customer
		err := r.scanEmployeeFromRow(rows, &veterinarian)
		if err != nil {
			return page.Page[[]customer.Customer]{}, r.wrapConversionError(err)
		}
		vets = append(vets, veterinarian)
	}

	if err := rows.Err(); err != nil {
		return page.Page[[]customer.Customer]{}, r.dbError(OpSelect, "error iterating search results", err)
	}

	// Get total count for pagination
	totalCount, err := r.getTotalCountWithFilters(ctx, &spec)
	if err != nil {
		return page.Page[[]customer.Customer]{}, err
	}

	// Handle pagination
	pageMetadata := page.GetPageMetadata(totalCount, page.PageInput{
		PageNumber: spec.GetPagination().Page,
		Limit:   spec.GetPagination().Limit,
	})

	return page.NewPage(vets, *pageMetadata), nil
}
*/
