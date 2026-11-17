import { IContext } from '@common/context/context.interface';
import { Result } from '@common/utils/result';

export const IDatabaseConnection = Symbol('IDatabaseConnection');

export type IDatabaseConnection = {
  query<T>(
    context: IContext,
    query: string,
    params?: unknown[],
  ): Promise<Result<T[], Error>>;
  startTransaction(context: IContext): Promise<Result<void, Error>>;
  commitTransaction(context: IContext): Promise<Result<void, Error>>;
  rollbackTransaction(context: IContext): Promise<Result<void, Error>>;
};
