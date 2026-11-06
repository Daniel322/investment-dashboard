import { Module, DynamicModule, Global } from '@nestjs/common';
import { DatabaseConnection } from './database.connection';

@Global()
@Module({})
export class DatabaseModule {
  public static forRootAsync(): DynamicModule {
    const databaseConnectionProvider = {
      provide: 'DATABASE_CONNECTION',
      useFactory: (connection: DatabaseConnection) => connection.getPool(),
      inject: [DatabaseConnection],
    };

    return {
      module: DatabaseModule,
      providers: [
        DatabaseConnection,
        databaseConnectionProvider,
      ],
      exports: ['DATABASE_CONNECTION'],
    };
  }
}