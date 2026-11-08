import { Module, ValidationPipe } from '@nestjs/common';
import { ConfigModule } from '@nestjs/config';
import { EnvConfig } from '@infrastructure/config/config';
import { APP_PIPE } from '@nestjs/core';
import { DatabaseModule } from '@infrastructure/database/database.module';

@Module({
  imports: [
    ConfigModule.forRoot(EnvConfig),
    DatabaseModule.forRoot(),
  ],
  providers: [
    {
      provide: APP_PIPE,
      useClass: ValidationPipe,
    },
  ]
})
export class AppModule {}
