export type IConnection = {
  query<T>(query: string, params: any[]): Promise<T>;
};