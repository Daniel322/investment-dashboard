import { randomUUID } from 'node:crypto';
import { Inject, Injectable } from '@nestjs/common';
import { CommandHandler } from '@common/cmd-handler/command.handler';
import { IContext } from '@common/context/context.interface';
import { Result } from '@common/utils/result';
import { Asset } from '@domain/asset/domain/asset.model';
import { AssetRecord } from '@domain/asset/domain/asset.record';
import { IAssetRepository } from '@domain/asset/infrastructure/repositories/asset.repository.interface';
import { CreateAssetCommand } from '../commands/create-asset.cmd';

@Injectable()
export class CreateAssetCommandHandler implements CommandHandler<
  CreateAssetCommand,
  AssetRecord
> {
  constructor(
    @Inject(IAssetRepository)
    private readonly assetRepository: IAssetRepository,
  ) {}

  public async handle(
    context: IContext,
    command: CreateAssetCommand,
  ): Promise<Result<AssetRecord, Error>> {
    const asset = Asset.create({ ...command });

    return this.assetRepository.save(context, asset.toRecord());
  }
}
