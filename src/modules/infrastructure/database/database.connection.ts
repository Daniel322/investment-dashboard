import {
  Injectable,
  Logger,
  OnApplicationShutdown,
  OnModuleInit,
} from '@nestjs/common';
import { ConfigService } from '@nestjs/config';
import { Pool, type PoolClient } from 'pg';
import { IDatabaseConnection } from './database.interface';
import { Context } from '@common/context/context';

@Injectable()
export class DatabaseConnection implements OnModuleInit, OnApplicationShutdown, IDatabaseConnection {
  private readonly logger = new Logger(DatabaseConnection.name);
  private readonly pool: Pool;
  constructor(configService: ConfigService) {
    this.pool = new Pool({
      connectionString: configService.getOrThrow<string>('db.url'),
    });
  }

  public async onModuleInit(): Promise<void> {
    try {
      const client = await this.pool.connect();
      await client.query('SELECT NOW()');
      client.release();
      this.logger.log('Successfully connected to PostgreSQL database');
    } catch (error) {
      await this.pool.end();
      this.logger.error(
        'Failed to connect to PostgreSQL database:',
        error instanceof Error ? error.message : String(error),
      );
      throw error;
    }
  }

  public async onApplicationShutdown(): Promise<void> {
    this.logger.log('Shutting down database connection');
    if (!this.pool.ended) {
      await this.pool.end();
    }
    this.logger.log('Database connection shut down');
  }

  public async startTransaction(context: Context): Promise<void> {
    const poolClient = context.pool ?? await this.pool.connect();
    await poolClient.query('BEGIN');
    context.setPool(poolClient);
    poolClient.on('error', (error) => {
      this.logger.error(`Error on pool client (${context.requestId}):`, error);
      poolClient.release();
    });
  }

  public async commitTransaction(context: Context): Promise<void> {
    const poolClient = context.pool!;
    try {
      await poolClient.query('COMMIT');
    } catch (error) {
      this.logger.error(`Error on commit transaction (${context.requestId}):`, error);
    } finally {
      poolClient.release();
    }
  }

  public async rollbackTransaction(context: Context): Promise<void> {
    const poolClient = context.pool!;
    try {
      await poolClient.query('ROLLBACK');
    } catch (error) {
      this.logger.error(`Error on rollback transaction (${context.requestId}):`, error);
    } finally {
      poolClient.release();
    }
  }

  public async query<T>(context: Context, query: string, params?: unknown[]): Promise<T[]> {
    let isQueryInTransaction = context.pool ? true : false;
    const poolClient = context.pool ?? await this.pool.connect();
    try {
      const result = await poolClient.query(query, params);

      return result.rows as T[];
    } catch (error) {
      this.logger.error(error);
      throw error;
    } finally {
      if (!isQueryInTransaction) {
        poolClient.release();
      }
    }
  }
}
