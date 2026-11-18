import { IContext } from '@common/context/context.interface';
import { Result } from '@common/utils/result';
import { IDatabaseConnection } from '@infrastructure/database/database.interface';
import { AssetRecord } from '../../domain/asset.record';
import { EAssetType } from '../../domain/e-asset-type';
import { AssetRepository } from './asset.repository';

describe('AssetRepository', () => {
  const context = createContext();
  let query: jest.MockedFunction<IDatabaseConnection['query']>;
  let startTransaction: jest.MockedFunction<
    IDatabaseConnection['startTransaction']
  >;
  let commitTransaction: jest.MockedFunction<
    IDatabaseConnection['commitTransaction']
  >;
  let rollbackTransaction: jest.MockedFunction<
    IDatabaseConnection['rollbackTransaction']
  >;
  let repository: AssetRepository;

  beforeEach(() => {
    query = jest.fn();
    startTransaction = jest.fn().mockResolvedValue(Result.success());
    commitTransaction = jest.fn().mockResolvedValue(Result.success());
    rollbackTransaction = jest.fn().mockResolvedValue(Result.success());
    repository = new AssetRepository({
      query,
      startTransaction,
      commitTransaction,
      rollbackTransaction,
    });
  });

  describe('save', () => {
    it('inserts an asset when the locked row does not exist', async () => {
      const record = createRecord();
      const stored = createRecord({
        updated_at: new Date('2024-07-01T00:00:00.000Z'),
      });
      query
        .mockResolvedValueOnce(Result.success([]))
        .mockResolvedValueOnce(Result.success([stored]));

      const result = await repository.save(context, record);

      expect(startTransaction).not.toHaveBeenCalled();
      expect(commitTransaction).not.toHaveBeenCalled();
      expect(rollbackTransaction).not.toHaveBeenCalled();
      expect(calledSql(query)).toContain('WHERE id = $1');
      expect(calledSql(query).indexOf('FOR UPDATE')).toBeGreaterThan(
        calledSql(query).indexOf('ORDER BY'),
      );
      expect(calledParams(query)).toEqual([record.id]);
      expect(calledSql(query, 1)).toContain('INSERT INTO assets');
      expect(returningColumns(calledSql(query, 1))).toEqual(assetColumns);
      expect(calledParams(query, 1)).toEqual([
        record.id,
        record.name,
        record.slug,
        record.type,
        record.price,
        record.rate,
        record.created_at,
        record.updated_at,
      ]);
      expect(result.isSuccess()).toBe(true);
      expect(result.value).toEqual(stored);
    });

    it('locks the existing row and updates it', async () => {
      const record = createRecord({
        name: 'Bitcoin Cash',
        updated_at: new Date('2024-07-01T00:00:00.000Z'),
      });
      const stored = createRecord({
        name: 'Bitcoin Cash',
        created_at: new Date('2024-01-01T00:00:00.000Z'),
        updated_at: new Date('2024-07-01T00:00:00.000Z'),
      });
      query
        .mockResolvedValueOnce(Result.success([createRecord()]))
        .mockResolvedValueOnce(Result.success([stored]));

      const result = await repository.save(context, record);

      expect(calledSql(query)).toContain('FOR UPDATE');
      expect(calledParams(query)).toEqual([record.id]);
      expect(calledSql(query, 1)).toContain('UPDATE assets');
      expect(calledSql(query, 1)).not.toContain('created_at =');
      expect(returningColumns(calledSql(query, 1))).toEqual(assetColumns);
      expect(calledParams(query, 1)).toEqual([
        record.id,
        record.name,
        record.slug,
        record.type,
        record.price,
        record.rate,
        record.updated_at,
      ]);
      expect(result.isSuccess()).toBe(true);
      expect(result.value).toEqual(stored);
    });

    it('returns the database error when the row cannot be locked', async () => {
      const error = new Error('select failed');
      query.mockResolvedValue(Result.fail(error));

      const result = await repository.save(context, createRecord());

      expect(query).toHaveBeenCalledTimes(1);
      expect(startTransaction).not.toHaveBeenCalled();
      expect(rollbackTransaction).not.toHaveBeenCalled();
      expect(result.isExceptional()).toBe(true);
      expect(result.exception).toBe(error);
    });

    it('returns the database error when the row cannot be written', async () => {
      const error = new Error('insert failed');
      query
        .mockResolvedValueOnce(Result.success([]))
        .mockResolvedValueOnce(Result.fail(error));

      const result = await repository.save(context, createRecord());

      expect(result.isExceptional()).toBe(true);
      expect(result.exception).toBe(error);
    });

    it('returns an error when the database stores nothing', async () => {
      query
        .mockResolvedValueOnce(Result.success([]))
        .mockResolvedValueOnce(Result.success([]));

      const result = await repository.save(context, createRecord());

      expect(result.isExceptional()).toBe(true);
      expect(result.exception).toEqual(
        new Error('Asset asset-1 was not saved'),
      );
    });
  });

  describe('findById', () => {
    it('returns the record', async () => {
      const record = createRecord();
      query.mockResolvedValue(Result.success([record]));

      const result = await repository.findById(context, 'asset-1');

      expect(calledSql(query)).toContain('WHERE id = $1');
      expect(calledSql(query)).not.toContain('FOR UPDATE');
      expect(calledParams(query)).toEqual(['asset-1']);
      expect(result.isSuccess()).toBe(true);
      expect(result.value).toEqual(record);
    });

    it('returns an empty result when the asset does not exist', async () => {
      query.mockResolvedValue(Result.success([]));

      const result = await repository.findById(context, 'missing');

      expect(result.isSuccess()).toBe(true);
      expect(result.value).toBeUndefined();
    });

    it('returns the database error', async () => {
      const error = new Error('select failed');
      query.mockResolvedValue(Result.fail(error));

      const result = await repository.findById(context, 'asset-1');

      expect(result.isExceptional()).toBe(true);
      expect(result.exception).toBe(error);
    });
  });

  describe('findBySlug', () => {
    it('returns the records with that slug', async () => {
      const records = [createRecord()];
      query.mockResolvedValue(Result.success(records));

      const result = await repository.findBySlug(context, 'bitcoin');

      expect(calledSql(query)).toContain('WHERE slug = $1');
      expect(calledSql(query)).not.toContain('FOR UPDATE');
      expect(calledParams(query)).toEqual(['bitcoin']);
      expect(result.isSuccess()).toBe(true);
      expect(result.value).toEqual(records);
    });

    it('returns an empty list when the slug does not exist', async () => {
      query.mockResolvedValue(Result.success([]));

      const result = await repository.findBySlug(context, 'missing');

      expect(result.isSuccess()).toBe(true);
      expect(result.value).toEqual([]);
    });

    it('returns the database error', async () => {
      const error = new Error('select failed');
      query.mockResolvedValue(Result.fail(error));

      const result = await repository.findBySlug(context, 'bitcoin');

      expect(result.isExceptional()).toBe(true);
      expect(result.exception).toBe(error);
    });
  });

  describe('findByName', () => {
    it('returns every record with that name', async () => {
      const records = [
        createRecord(),
        createRecord({ id: 'asset-2', slug: 'bitcoin-cash' }),
      ];
      query.mockResolvedValue(Result.success(records));

      const result = await repository.findByName(context, 'Bitcoin');

      expect(calledSql(query)).toContain('WHERE name = $1');
      expect(calledSql(query)).toContain('ORDER BY name ASC, id ASC');
      expect(calledParams(query)).toEqual(['Bitcoin']);
      expect(result.isSuccess()).toBe(true);
      expect(result.value).toEqual(records);
    });

    it('returns the database error', async () => {
      const error = new Error('select failed');
      query.mockResolvedValue(Result.fail(error));

      const result = await repository.findByName(context, 'Bitcoin');

      expect(result.isExceptional()).toBe(true);
      expect(result.exception).toBe(error);
    });
  });

  describe('findByType', () => {
    it('returns every record of that type', async () => {
      const records = [
        createRecord(),
        createRecord({ id: 'asset-2', name: 'Ethereum', slug: 'ethereum' }),
      ];
      query.mockResolvedValue(Result.success(records));

      const result = await repository.findByType(context, EAssetType.Crypto);

      expect(calledSql(query)).toContain('WHERE type = $1');
      expect(calledSql(query)).toContain('ORDER BY name ASC, id ASC');
      expect(calledParams(query)).toEqual([EAssetType.Crypto]);
      expect(result.isSuccess()).toBe(true);
      expect(result.value).toEqual(records);
    });

    it('returns an empty list when nothing matches', async () => {
      query.mockResolvedValue(Result.success([]));

      const result = await repository.findByType(context, EAssetType.Gold);

      expect(result.isSuccess()).toBe(true);
      expect(result.value).toEqual([]);
    });

    it('returns the database error', async () => {
      const error = new Error('select failed');
      query.mockResolvedValue(Result.fail(error));

      const result = await repository.findByType(context, EAssetType.Crypto);

      expect(result.isExceptional()).toBe(true);
      expect(result.exception).toBe(error);
    });
  });
});

const assetColumns = [
  'id',
  'name',
  'slug',
  'type',
  'price',
  'rate',
  'created_at',
  'updated_at',
];

function returningColumns(sql: string): string[] {
  const [, returning = ''] = sql.split('RETURNING');
  return returning.split(',').map((column) => column.trim());
}

function createContext(): IContext {
  return {
    requestId: 'request-1',
    getClient: () => undefined,
    setClient: () => Result.success(),
    clearClient: () => Result.success(),
  };
}

function calledSql(
  mock: jest.MockedFunction<IDatabaseConnection['query']>,
  index = 0,
): string {
  const sql = mock.mock.calls[index]?.[1];
  if (typeof sql !== 'string') {
    throw new Error(`Expected SQL at call ${index}`);
  }

  return sql;
}

function calledParams(
  mock: jest.MockedFunction<IDatabaseConnection['query']>,
  index = 0,
): unknown[] | undefined {
  return mock.mock.calls[index]?.[2];
}

function createRecord(overrides: Partial<AssetRecord> = {}): AssetRecord {
  return {
    id: 'asset-1',
    name: 'Bitcoin',
    slug: 'bitcoin',
    type: EAssetType.Crypto,
    price: 100,
    rate: 1,
    created_at: new Date('2024-01-01T00:00:00.000Z'),
    updated_at: new Date('2024-06-01T00:00:00.000Z'),
    ...overrides,
  };
}
