import { randomUUID } from 'node:crypto';
import {
  Body,
  Controller,
  Headers,
  HttpCode,
  HttpStatus,
  Inject,
  InternalServerErrorException,
  Post,
} from '@nestjs/common';
import { Context } from '@common/context/context';
import type { AssetRecord } from '../../domain/asset.record';
import { CreateAssetCommand } from '../../application/commands/create-asset.cmd';
import { CreateAssetCommandHandler } from '../../application/command-handlers/create-asset.handler';
import { ICreateAssetCommandHandler } from '../../application/command-handlers/create-asset.handler.provider';
import { CreateAssetDto } from './create-asset.dto';

@Controller('api/v1/assets')
export class AssetController {
  constructor(
    @Inject(ICreateAssetCommandHandler)
    private readonly createAssetHandler: CreateAssetCommandHandler,
  ) {}

  @Post()
  @HttpCode(HttpStatus.CREATED)
  public async create(
    @Body() body: CreateAssetDto,
    @Headers('x-request-id') requestId?: string,
  ): Promise<AssetRecord> {
    const context = new Context(requestId ?? randomUUID());
    const command = new CreateAssetCommand(
      body.name,
      body.slug,
      body.type,
      body.price,
      body.rate,
    );

    const result = await this.createAssetHandler.handle(context, command);
    if (result.isExceptional()) {
      throw new InternalServerErrorException(result.exception.message);
    }

    return result.value;
  }
}
