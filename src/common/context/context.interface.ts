import { PoolClient } from 'pg';
import { Result } from '../utils/result';

export interface IContext {
  readonly requestId: string;
  getClient(): PoolClient | undefined;
  setClient(client: PoolClient): Result<void, Error>;
  clearClient(): Result<void, Error>;
}
