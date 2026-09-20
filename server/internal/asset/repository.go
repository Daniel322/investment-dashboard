package asset

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"investment-dashboard/pkg/price"
	"investment-dashboard/pkg/rate"
)

type AssetRepository struct {
	db *sql.DB
}

type AssetFilter struct {
	ID        *string
	Name      *string
	Slug      *string
	Price     *float64
	Rate      *float64
	Type      *string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type filterQuery struct {
	conditions []string
	args       []any
}

func (q *filterQuery) add(column string, value any) {
	q.conditions = append(q.conditions, fmt.Sprintf("%s = $%d", column, len(q.args)+1))
	q.args = append(q.args, value)
}

func (q filterQuery) where(query string) string {
	if len(q.conditions) > 0 {
		query += " WHERE " + strings.Join(q.conditions, " AND ")
	}
	return query
}

func buildFilterQuery(filter AssetFilter) filterQuery {
	var q filterQuery

	if filter.ID != nil {
		q.add("id", *filter.ID)
	}
	if filter.Name != nil {
		q.add("name", *filter.Name)
	}
	if filter.Slug != nil {
		q.add("slug", *filter.Slug)
	}
	if filter.Price != nil {
		q.add("price", price.ToMinor(*filter.Price))
	}
	if filter.Rate != nil {
		q.add("rate", rate.ToMinor(*filter.Rate))
	}
	if filter.Type != nil {
		q.add(`"type"`, *filter.Type)
	}
	if filter.CreatedAt != nil {
		q.add("created_at", *filter.CreatedAt)
	}
	if filter.UpdatedAt != nil {
		q.add("updated_at", *filter.UpdatedAt)
	}

	return q
}

const selectAssets = `
SELECT id, name, slug, price, rate, "type", created_at, updated_at
FROM assets
`

const selectAssetByID = selectAssets + `
WHERE id = $1
`

const selectAssetForUpdate = selectAssetByID + `
FOR UPDATE
`

const insertAsset = `
INSERT INTO assets (id, name, slug, price, rate, "type", created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
`

const updateAsset = `
UPDATE assets
SET name = $1, slug = $2, price = $3, rate = $4, "type" = $5, updated_at = $6
WHERE id = $7
`

func (r *AssetRepository) Save(ctx context.Context, entity *AssetRecord) (*AssetRecord, error) {
	if entity == nil {
		return nil, ErrAssetPointerIsNil
	}

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var existing AssetRecord
	err = tx.QueryRowContext(ctx, selectAssetForUpdate, entity.ID).Scan(
		&existing.ID,
		&existing.Name,
		&existing.Slug,
		&existing.Price,
		&existing.Rate,
		&existing.Type,
		&existing.CreatedAt,
		&existing.UpdatedAt,
	)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	if errors.Is(err, sql.ErrNoRows) {
		_, err = tx.ExecContext(ctx, insertAsset,
			entity.ID,
			entity.Name,
			entity.Slug,
			entity.Price,
			entity.Rate,
			entity.Type,
			entity.CreatedAt,
			entity.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
	} else {
		_, err = tx.ExecContext(ctx, updateAsset,
			entity.Name,
			entity.Slug,
			entity.Price,
			entity.Rate,
			entity.Type,
			entity.UpdatedAt,
			entity.ID,
		)
		if err != nil {
			return nil, err
		}
		entity.CreatedAt = existing.CreatedAt
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return entity, nil
}

func (r *AssetRepository) Load(ctx context.Context, id string) (*AssetRecord, error) {
	var entity AssetRecord
	err := r.db.QueryRowContext(ctx, selectAssetByID, id).Scan(
		&entity.ID,
		&entity.Name,
		&entity.Slug,
		&entity.Price,
		&entity.Rate,
		&entity.Type,
		&entity.CreatedAt,
		&entity.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrAssetNotFound
		}
		return nil, err
	}

	return &entity, nil
}

func (r *AssetRepository) Find(ctx context.Context, filter AssetFilter) (*AssetRecord, error) {
	q := buildFilterQuery(filter)
	query := q.where(selectAssets) + " ORDER BY created_at LIMIT 1"

	var entity AssetRecord
	err := r.db.QueryRowContext(ctx, query, q.args...).Scan(
		&entity.ID,
		&entity.Name,
		&entity.Slug,
		&entity.Price,
		&entity.Rate,
		&entity.Type,
		&entity.CreatedAt,
		&entity.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrAssetNotFound
		}
		return nil, err
	}

	return &entity, nil
}

func (r *AssetRepository) List(ctx context.Context, filter AssetFilter) ([]*AssetRecord, error) {
	q := buildFilterQuery(filter)
	query := q.where(selectAssets) + " ORDER BY created_at"

	rows, err := r.db.QueryContext(ctx, query, q.args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	assets := make([]*AssetRecord, 0)
	for rows.Next() {
		var entity AssetRecord
		if err := rows.Scan(
			&entity.ID,
			&entity.Name,
			&entity.Slug,
			&entity.Price,
			&entity.Rate,
			&entity.Type,
			&entity.CreatedAt,
			&entity.UpdatedAt,
		); err != nil {
			return nil, err
		}
		assets = append(assets, &entity)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return assets, nil
}
