import { EAssetType } from '../../domain/e-asset-type';

export class CreateAssetCommand {
  constructor(
    public readonly name: string,
    public readonly slug: string,
    public readonly type: EAssetType,
    public readonly price: number,
    public readonly rate: number,
  ) {}
}
