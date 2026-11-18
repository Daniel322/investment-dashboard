import { Module, DynamicModule, Global } from '@nestjs/common';
import { DatabaseConnection } from './database.connection';
import { IDatabaseConnection } from './database.interface';

@Global()
@Module({})
export class DatabaseModule {
  public static forRoot(): DynamicModule {
    const databaseConnectionProvider = {
      provide: IDatabaseConnection,
      useExisting: DatabaseConnection,
    };

    return {
      module: DatabaseModule,
      providers: [DatabaseConnection, databaseConnectionProvider],
      exports: [IDatabaseConnection],
    };
  }
}
