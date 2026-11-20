import { randomUUID } from 'node:crypto';
import { AssetRecord } from './asset.record';
import { EAssetType } from './e-asset-type';

type CreateAssetData = Pick<AssetRecord, 'name' | 'slug' | 'type' | 'rate' | 'price'>;

export class Asset {
  public readonly id: string;
  public readonly name: string;
  public readonly slug: string;
  public readonly type: EAssetType;
  public readonly price: number;
  public readonly rate: number;
  public readonly createdAt: Date;
  public readonly updatedAt: Date;

  constructor({
    id = randomUUID(),
    name,
    slug,
    type,
    rate,
    price,
    created_at = new Date(),
    updated_at = new Date(),
  }: AssetRecord) {
    this.id = id;
    this.name = name;
    this.slug = slug;
    this.type = type;
    this.price = price;
    this.rate = rate;
    this.createdAt = created_at;
    this.updatedAt = updated_at;
  }

  public static create(record: CreateAssetData): Asset {
    return new Asset(record as AssetRecord);
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
