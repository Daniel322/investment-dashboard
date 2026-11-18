import { EAssetType } from './e-asset-type';

export type AssetRecord = {
  id: string;
  name: string;
  slug: string;
  type: EAssetType;
  price: number; // decimal in pg
  rate: number;
  created_at: Date;
  updated_at: Date;
};
