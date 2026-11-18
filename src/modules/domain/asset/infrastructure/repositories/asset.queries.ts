export const ASSET_INSERT_QUERY = `
            INSERT INTO assets (
              id,
              name,
              slug,
              type,
              price,
              rate,
              created_at,
              updated_at
            )
            VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
            RETURNING
              id,
              name,
              slug,
              type,
              price,
              rate,
              created_at,
              updated_at
          `;

export const ASSET_UPDATE_QUERY = `
            UPDATE assets
            SET
              name = $2,
              slug = $3,
              type = $4,
              price = $5,
              rate = $6,
              updated_at = $7
            WHERE id = $1
            RETURNING
              id,
              name,
              slug,
              type,
              price,
              rate,
              created_at,
              updated_at
          `;
