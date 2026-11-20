import { Module } from '@nestjs/common';
import { CreateAssetCommandHandler } from './application/command-handlers/create-asset.handler';
import {
  createAssetCommandHandlerProvider,
  ICreateAssetCommandHandler,
} from './application/command-handlers/create-asset.handler.provider';
import { AssetController } from './infrastructure/controllers/asset.controller';
import { AssetRepository } from './infrastructure/repositories/asset.repository';
import { IAssetRepository } from './infrastructure/repositories/asset.repository.interface';
import { assetRepositoryProvider } from './infrastructure/repositories/asset.repository.provider';

@Module({
  controllers: [AssetController],
  providers: [
    AssetRepository,
    assetRepositoryProvider,
    CreateAssetCommandHandler,
    createAssetCommandHandlerProvider,
  ],
  exports: [IAssetRepository, ICreateAssetCommandHandler],
})
export class AssetModule {}
