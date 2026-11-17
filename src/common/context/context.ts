import { PoolClient } from 'pg';
import { IContext } from './context.interface';
import { Result } from '../utils/result';

export class Context implements IContext {
  public readonly requestId: string;
  private client?: PoolClient;

  constructor(requestId: string) {
    this.requestId = requestId;
  }

  public getClient(): PoolClient | undefined {
    return this.client;
  }

  public setClient(client: PoolClient): Result<void, Error> {
    if (this.client) {
      return Result.fail(
        new Error(`Transaction already started (${this.requestId})`),
      );
    }

    this.client = client;
    return Result.success();
  }

  public clearClient(): Result<void, Error> {
    this.client = undefined;
    return Result.success();
  }
}
