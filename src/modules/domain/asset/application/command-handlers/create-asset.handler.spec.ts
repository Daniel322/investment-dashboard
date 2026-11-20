import { randomUUID } from 'node:crypto';
import { IContext } from '@common/context/context.interface';
import { Result } from '@common/utils/result';
import { AssetRecord } from '@domain/asset/domain/asset.record';
import { EAssetType } from '@domain/asset/domain/e-asset-type';
import { IAssetRepository } from '@domain/asset/infrastructure/repositories/asset.repository.interface';
import { CreateAssetCommand } from '../commands/create-asset.cmd';
import { CreateAssetCommandHandler } from './create-asset.handler';

jest.mock('node:crypto', () => ({
  randomUUID: jest.fn(),
}));

describe('CreateAssetCommandHandler', () => {
  const context = createContext();
  const now = new Date('2024-06-01T00:00:00.000Z');
  const assetId = 'asset-1';

  let save: jest.MockedFunction<IAssetRepository['save']>;
  let handler: CreateAssetCommandHandler;

  beforeEach(() => {
    jest.useFakeTimers();
    jest.setSystemTime(now);
    (randomUUID as jest.Mock).mockReturnValue(assetId);

    save = jest.fn();
    handler = new CreateAssetCommandHandler({
      save,
      findById: jest.fn(),
      findBySlug: jest.fn(),
      findByName: jest.fn(),
      findByType: jest.fn(),
    });
  });

  afterEach(() => {
    jest.useRealTimers();
    jest.clearAllMocks();
  });

  it('creates an asset and saves it through the repository', async () => {
    const command = createCommand();
    const stored = createRecord();
    save.mockResolvedValue(Result.success(stored));

    const result = await handler.handle(context, command);

    expect(save).toHaveBeenCalledWith(context, {
      id: assetId,
      name: command.name,
      slug: command.slug,
      type: command.type,
      price: command.price,
      rate: command.rate,
      created_at: now,
      updated_at: now,
    });
    expect(result.isSuccess()).toBe(true);
    expect(result.value).toEqual(stored);
  });

  it('returns the repository error', async () => {
    const error = new Error('save failed');
    save.mockResolvedValue(Result.fail(error));

    const result = await handler.handle(context, createCommand());

    expect(result.isExceptional()).toBe(true);
    expect(result.exception).toBe(error);
  });
});

function createCommand(): CreateAssetCommand {
  return new CreateAssetCommand(
    'Bitcoin',
    'bitcoin',
    EAssetType.Crypto,
    100,
    1,
  );
}

function createRecord(overrides: Partial<AssetRecord> = {}): AssetRecord {
  return {
    id: 'asset-1',
    name: 'Bitcoin',
    slug: 'bitcoin',
    type: EAssetType.Crypto,
    price: 100,
    rate: 1,
    created_at: new Date('2024-06-01T00:00:00.000Z'),
    updated_at: new Date('2024-06-01T00:00:00.000Z'),
    ...overrides,
  };
}

function createContext(): IContext {
  return {
    requestId: 'request-1',
    getClient: () => undefined,
    setClient: () => Result.success(),
    clearClient: () => Result.success(),
  };
}
