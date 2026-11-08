import { Context } from "@common/context/context";

export const IDatabaseConnection = Symbol('IDatabaseConnection');

export type IDatabaseConnection = {
  query<T>(context: Context, query: string, params?: unknown[]): Promise<T[]>;
  startTransaction(context: Context): Promise<void>;
  commitTransaction(context: Context): Promise<void>;
  rollbackTransaction(context: Context): Promise<void>;
};