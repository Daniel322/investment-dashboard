import { Inject, Injectable } from '@nestjs/common';
import { IContext } from '@common/context/context.interface';
import { Result } from '@common/utils/result';
import { IDatabaseConnection } from '@infrastructure/database/database.interface';
import { AssetRecord } from '@domain/asset/domain/asset.record';
import { IAssetRepository } from './asset.repository.interface';
import { ASSET_INSERT_QUERY, ASSET_UPDATE_QUERY } from './asset.queries';

type AssetColumn = keyof AssetRecord;

type AssetValues<Columns extends readonly AssetColumn[]> = Array<
  AssetRecord[Columns[number]]
>;

@Injectable()
export class AssetRepository implements IAssetRepository {
  constructor(
    @Inject(IDatabaseConnection)
    private readonly database: IDatabaseConnection,
  ) {}

  public async save(
    context: IContext,
    record: AssetRecord,
  ): Promise<Result<AssetRecord, Error>> {
    const existing = await this.select(context, 'id', record.id, true);

    if (existing.isExceptional()) {
      return Result.fail(existing.exception);
    }

    const saved = existing.value[0]
      ? await this.database.query<AssetRecord>(
          context,
          ASSET_UPDATE_QUERY,
          this.values(record, [
            'id',
            'name',
            'slug',
            'type',
            'price',
            'rate',
            'updated_at',
          ]),
        )
      : await this.database.query<AssetRecord>(
          context,
          ASSET_INSERT_QUERY,
          this.values(record, [
            'id',
            'name',
            'slug',
            'type',
            'price',
            'rate',
            'created_at',
            'updated_at',
          ]),
        );

    if (saved.isExceptional()) {
      return Result.fail(saved.exception);
    }

    const row = saved.value[0];
    if (!row) {
      return Result.fail(new Error(`Asset ${record.id} was not saved`));
    }

    return Result.success(row);
  }

  public async findById(
    context: IContext,
    id: AssetRecord['id'],
  ): Promise<Result<AssetRecord | undefined, Error>> {
    const result = await this.select(context, 'id', id);

    if (result.isExceptional()) return Result.fail(result.exception);

    return Result.success(result.value[0]);
  }

  public async findBySlug(
    context: IContext,
    slug: AssetRecord['slug'],
  ): Promise<Result<AssetRecord[], Error>> {
    return this.select(context, 'slug', slug);
  }

  public async findByName(
    context: IContext,
    name: AssetRecord['name'],
  ): Promise<Result<AssetRecord[], Error>> {
    return this.select(context, 'name', name);
  }

  public async findByType(
    context: IContext,
    type: AssetRecord['type'],
  ): Promise<Result<AssetRecord[], Error>> {
    return this.select(context, 'type', type);
  }

  private async select(
    context: IContext,
    filter: keyof AssetRecord,
    value: AssetRecord[keyof AssetRecord],
    isForUpdate: boolean = false,
  ): Promise<Result<AssetRecord[], Error>> {
    const result = await this.database.query<AssetRecord>(
      context,
      `
        SELECT
          id,
          name,
          slug,
          type,
          price,
          rate,
          created_at,
          updated_at
        FROM assets
        WHERE ${filter} = $1
        ORDER BY name ASC, id ASC
        ${isForUpdate ? 'FOR UPDATE' : ''}
      `,
      [value],
    );
    if (result.isExceptional()) {
      return Result.fail(result.exception);
    }

    return Result.success(result.value);
  }

  private values<const Columns extends readonly AssetColumn[]>(
    record: AssetRecord,
    columns: Columns,
  ): AssetValues<Columns> {
    return columns.map((column) => record[column]) as AssetValues<Columns>;
  }
}
