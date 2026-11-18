import { AssetRecord } from './asset.record';
import { EAssetType } from './e-asset-type';

export class Asset {
  public readonly id: string;
  public readonly name: string;
  public readonly slug: string;
  public readonly type: EAssetType;
  public readonly price: number;
  public readonly rate: number;
  public readonly createdAt: Date;
  public readonly updatedAt: Date;

  constructor(record: AssetRecord) {
    this.id = record.id;
    this.name = record.name;
    this.slug = record.slug;
    this.type = record.type;
    this.price = record.price;
    this.rate = record.rate;
    this.createdAt = record.created_at;
    this.updatedAt = record.updated_at;
  }

  public static create(record: AssetRecord): Asset {
    return new Asset(record);
  }

  public static fromRecord(record: AssetRecord): Asset {
    return new Asset(record);
  }

  public toRecord(): AssetRecord {
    return {
      id: this.id,
      name: this.name,
      slug: this.slug,
      type: this.type,
      price: this.price,
      rate: this.rate,
      created_at: this.createdAt,
      updated_at: this.updatedAt,
    };
  }
}
