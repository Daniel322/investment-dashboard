import { IContext } from '@common/context/context.interface';
import { IResult } from '@common/utils/result.interface';

export interface CommandHandler<
  TCommand = Record<string, unknown>,
  TResult = void,
> {
  handle(
    context: IContext,
    command: TCommand,
  ): Promise<IResult<TResult, Error>>;
}
