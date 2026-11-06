import { Module, ValidationPipe } from '@nestjs/common';
import { ConfigModule } from '@nestjs/config';
import { EnvConfig } from '@src/modules/infrastructure/config/config';
import { APP_PIPE } from '@nestjs/core';

@Module({
  imports: [ConfigModule.forRoot(EnvConfig)],
  providers: [
    {
      provide: APP_PIPE,
      useClass: ValidationPipe,
    },
  ]
})
export class AppModule {}
