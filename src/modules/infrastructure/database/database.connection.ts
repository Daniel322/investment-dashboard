import {
  Injectable,
  Logger,
  OnApplicationShutdown,
  OnModuleInit,
} from '@nestjs/common';
import { ConfigService } from '@nestjs/config';
import { Pool, type PoolClient } from 'pg';
import { IConnection } from './database.interface';

@Injectable()
export class DatabaseConnection implements OnModuleInit, OnApplicationShutdown, IConnection {
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

  public getPool(): Pool {
    return this.pool;
  }

  public async query<T>(query: string, params: any[]): Promise<T> {
    const client = await this.pool.connect();
    const result = await client.query(query, params);
    client.release();
    return result.rows as T;
  }
  
   /** Runs work inside BEGIN/COMMIT (ROLLBACK on error). */
   public async withTransaction<T>(
    fn: (client: PoolClient) => Promise<T>,
  ): Promise<T> {
    const client = await this.pool.connect();
    try {
      await client.query('BEGIN');
      const result = await fn(client);
      await client.query('COMMIT');
      return result;
    } catch (error) {
      await client.query('ROLLBACK');
      this.logger.error(error);
      throw error;
    } finally {
      client.release();
    }
  }
}
