import { Provider } from '@nestjs/common';
import { AssetRepository } from './asset.repository';
import { IAssetRepository } from './asset.repository.interface';

export const assetRepositoryProvider: Provider = {
  provide: IAssetRepository,
  useExisting: AssetRepository,
};
