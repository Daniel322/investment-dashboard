export interface IResult<TValue, TError = Error> {
  isSuccess(): boolean;
  isExceptional(): boolean;
  value: TValue;
  exception: TError;
}
