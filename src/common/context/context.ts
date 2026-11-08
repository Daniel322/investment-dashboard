import { PoolClient } from "pg";

export class Context {
  public readonly requestId: string;
  private _pool?: PoolClient;

  constructor(requestId: string, pool?: PoolClient) {
    this.requestId = requestId;
    this._pool = pool;
  }

  public get pool(): PoolClient | undefined {
    return this._pool;
  }

  public setPool(pool: PoolClient) {
    if (this._pool) {
      console.log(`Pool already exist in context (${this.requestId})`);
      return;
    }

    this._pool = pool;
  }
}