import { IResult } from './result.interface';

export class Result<TValue, TError = Error> implements IResult<TValue, TError> {
  private readonly _value?: TValue;
  private readonly error?: TError;
  private readonly successful: boolean;

  private constructor(successful: boolean, value?: TValue, error?: TError) {
    this.successful = successful;
    this._value = value;
    this.error = error;
  }

  public static success(): Result<void, never>;
  public static success<TValue>(value: TValue): Result<TValue, never>;
  public static success<TValue>(value?: TValue): Result<TValue | void, never> {
    return new Result<TValue | void, never>(true, value);
  }

  public static fail<TValue = never, TError = Error>(
    error: TError,
  ): Result<TValue, TError> {
    return new Result<TValue, TError>(false, undefined, error);
  }

  public get value(): TValue {
    if (!this.successful) {
      throw new Error('Cannot get value of an exceptional result');
    }

    return this._value as TValue;
  }

  public get exception(): TError {
    if (this.successful) {
      throw new Error('Cannot get error of a successful result');
    }

    return this.error as TError;
  }

  public isSuccess(): boolean {
    return this.successful;
  }

  public isExceptional(): boolean {
    return !this.successful;
  }
}
