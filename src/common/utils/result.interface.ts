export interface IResult<TValue, TError = Error> {
  isSuccess(): boolean;
  isExceptional(): boolean;
}
