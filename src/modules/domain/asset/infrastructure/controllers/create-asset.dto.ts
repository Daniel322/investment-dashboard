import { IsEnum, IsNotEmpty, IsNumber, IsString } from 'class-validator';
import { EAssetType } from '../../domain/e-asset-type';

export class CreateAssetDto {
  @IsString()
  @IsNotEmpty()
  public readonly name!: string;

  @IsString()
  @IsNotEmpty()
  public readonly slug!: string;

  @IsEnum(EAssetType)
  public readonly type!: EAssetType;

  @IsNumber()
  public readonly price!: number;

  @IsNumber()
  public readonly rate!: number;
}
