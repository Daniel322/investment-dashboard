import { Provider } from '@nestjs/common';
import { CreateAssetCommandHandler } from './create-asset.handler';

export const ICreateAssetCommandHandler = Symbol('ICreateAssetCommandHandler');

export const createAssetCommandHandlerProvider: Provider = {
  provide: ICreateAssetCommandHandler,
  useExisting: CreateAssetCommandHandler,
};
