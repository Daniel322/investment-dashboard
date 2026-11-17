import {
  Injectable,
  Logger,
  OnApplicationShutdown,
  OnModuleInit,
} from '@nestjs/common';
import { ConfigService } from '@nestjs/config';
import { Pool, PoolClient } from 'pg';
import { IContext } from '@common/context/context.interface';
import { Result } from '@common/utils/result';
import { IDatabaseConnection } from './database.interface';

@Injectable()
export class DatabaseConnection
  implements OnModuleInit, OnApplicationShutdown, IDatabaseConnection
{
  private readonly logger = new Logger(DatabaseConnection.name);
  private readonly pool: Pool;

  constructor(configService: ConfigService) {
    this.pool = new Pool({
      connectionString: configService.getOrThrow<string>('db.url'),
      connectionTimeoutMillis: 5000,
    });
    this.pool.on('error', (error) => {
      this.logger.error(
        'Unexpected error on idle database client',
        error.stack,
      );
    });
  }

  public async onModuleInit(): Promise<void> {
    try {
      const client = await this.pool.connect();
      try {
        await client.query('SELECT 1');
      } finally {
        client.release();
      }
      this.logger.log('Successfully connected to PostgreSQL database');
    } catch (error) {
      try {
        await this.closePool();
      } catch (closeError) {
        this.logger.error(
          'Failed to close database pool',
          closeError instanceof Error ? closeError.stack : String(closeError),
        );
      }
      throw error;
    }
  }

  public async onApplicationShutdown(): Promise<void> {
    this.logger.log('Shutting down database connection');
    await this.closePool();
    this.logger.log('Database connection shut down');
  }

  public async query<T>(
    context: IContext,
    query: string,
    params?: unknown[],
  ): Promise<Result<T[], Error>> {
    try {
      const client = context.getClient();
      const result = client
        ? await client.query(query, params)
        : await this.pool.query(query, params);

      return Result.success(result.rows as T[]);
    } catch (error) {
      return Result.fail(this.toError(error));
    }
  }

  public async startTransaction(
    context: IContext,
  ): Promise<Result<void, Error>> {
    if (context.getClient()) {
      return Result.fail(
        new Error(`Transaction already started (${context.requestId})`),
      );
    }

    let client: PoolClient;
    try {
      client = await this.pool.connect();
    } catch (error) {
      return Result.fail(this.toError(error));
    }

    try {
      await client.query('BEGIN');
      const setClientResult = context.setClient(client);
      if (setClientResult.isExceptional()) {
        client.release(setClientResult.exception);
        return Result.fail(setClientResult.exception);
      }

      return Result.success();
    } catch (error) {
      const exception = this.toError(error);
      client.release(exception);
      return Result.fail(exception);
    }
  }

  public async commitTransaction(
    context: IContext,
  ): Promise<Result<void, Error>> {
    return this.finishTransaction(context, 'COMMIT');
  }

  public async rollbackTransaction(
    context: IContext,
  ): Promise<Result<void, Error>> {
    return this.finishTransaction(context, 'ROLLBACK');
  }

  private async finishTransaction(
    context: IContext,
    statement: 'COMMIT' | 'ROLLBACK',
  ): Promise<Result<void, Error>> {
    const client = context.getClient();
    if (!client) {
      return Result.fail(
        new Error(`No active transaction (${context.requestId})`),
      );
    }

    context.clearClient();
    try {
      await client.query(statement);
      client.release();
      return Result.success();
    } catch (error) {
      const exception = this.toError(error);
      client.release(exception);
      return Result.fail(exception);
    }
  }

  private async closePool(): Promise<void> {
    if (this.pool.ending) {
      return;
    }

    await this.pool.end();
  }

  private toError(error: unknown): Error {
    return error instanceof Error ? error : new Error(String(error));
  }
}
