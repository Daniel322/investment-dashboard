import { PoolClient } from "pg";

export class Context {
  public readonly requestId: string;
  private client?: PoolClient;

  constructor(requestId: string) {
    this.requestId = requestId;
  }

  public getClient(): PoolClient | undefined {
    return this.client;
  }

  public setClient(client: PoolClient): void {
    if (this.client) {
      throw new Error(`Transaction already started (${this.requestId})`);
    }

    this.client = client;
  }

  public clearClient(): void {
    this.client = undefined;
  }
}
