import { IContext } from '@common/context/context.interface';
import { Result } from '@common/utils/result';
import { AssetRecord } from '../../domain/asset.record';

export const IAssetRepository = Symbol('IAssetRepository');

export type IAssetRepository = {
  save(
    context: IContext,
    record: AssetRecord,
  ): Promise<Result<AssetRecord, Error>>;
  findById(
    context: IContext,
    id: AssetRecord['id'],
  ): Promise<Result<AssetRecord | undefined, Error>>;
  findBySlug(
    context: IContext,
    slug: AssetRecord['slug'],
  ): Promise<Result<AssetRecord[], Error>>;
  findByName(
    context: IContext,
    name: AssetRecord['name'],
  ): Promise<Result<AssetRecord[], Error>>;
  findByType(
    context: IContext,
    type: AssetRecord['type'],
  ): Promise<Result<AssetRecord[], Error>>;
};
