import {
  Injectable,
  Logger,
  OnApplicationShutdown,
  OnModuleInit,
} from '@nestjs/common';
import { ConfigService } from '@nestjs/config';
import { Pool } from 'pg';
import { Context } from '@common/context/context';
import { IDatabaseConnection } from './database.interface';

@Injectable()
export class DatabaseConnection implements OnModuleInit, OnApplicationShutdown, IDatabaseConnection {
  private readonly logger = new Logger(DatabaseConnection.name);
  private readonly pool: Pool;

  constructor(configService: ConfigService) {
    this.pool = new Pool({
      connectionString: configService.getOrThrow<string>('db.url'),
      connectionTimeoutMillis: 5000,
    });
    this.pool.on('error', (error) => {
      this.logger.error('Unexpected error on idle database client', error.stack);
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

  public async query<T>(context: Context, query: string, params?: unknown[]): Promise<T[]> {
    const client = context.getClient();
    const result = client
      ? await client.query(query, params)
      : await this.pool.query(query, params);

    return result.rows as T[];
  }

  public async startTransaction(context: Context): Promise<void> {
    if (context.getClient()) {
      throw new Error(`Transaction already started (${context.requestId})`);
    }

    const client = await this.pool.connect();
    try {
      await client.query('BEGIN');
      context.setClient(client);
    } catch (error) {
      client.release(this.toError(error));
      throw error;
    }
  }

  public async commitTransaction(context: Context): Promise<void> {
    await this.finishTransaction(context, 'COMMIT');
  }

  public async rollbackTransaction(context: Context): Promise<void> {
    await this.finishTransaction(context, 'ROLLBACK');
  }

  private async finishTransaction(context: Context, statement: 'COMMIT' | 'ROLLBACK'): Promise<void> {
    const client = context.getClient();
    if (!client) {
      throw new Error(`No active transaction (${context.requestId})`);
    }

    context.clearClient();
    try {
      await client.query(statement);
      client.release();
    } catch (error) {
      client.release(this.toError(error));
      throw error;
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
