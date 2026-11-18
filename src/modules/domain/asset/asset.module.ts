import { Module } from '@nestjs/common';
import { AssetRepository } from './infrastructure/repositories/asset.repository';
import { IAssetRepository } from './infrastructure/repositories/asset.repository.interface';
import { assetRepositoryProvider } from './infrastructure/repositories/asset.repository.provider';

@Module({
  providers: [AssetRepository, assetRepositoryProvider],
  exports: [IAssetRepository],
})
export class AssetModule {}
