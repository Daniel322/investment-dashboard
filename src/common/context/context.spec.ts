import { PoolClient } from 'pg';
import { Context } from './context';

describe('Context', () => {
  const requestId = 'request-1';

  it('stores the request id', () => {
    const context = new Context(requestId);

    expect(context.requestId).toBe(requestId);
  });

  it('has no client before one is set', () => {
    const context = new Context(requestId);

    expect(context.getClient()).toBeUndefined();
  });

  it('stores a client', () => {
    const context = new Context(requestId);
    const client = createClient();

    const result = context.setClient(client);

    expect(result.isSuccess()).toBe(true);
    expect(result.isExceptional()).toBe(false);
    expect(result.value).toBeUndefined();
    expect(context.getClient()).toBe(client);
  });

  it('returns an error when a transaction is already started', () => {
    const context = new Context(requestId);
    const client = createClient();
    const nextClient = createClient();
    context.setClient(client);

    const result = context.setClient(nextClient);

    expect(result.isSuccess()).toBe(false);
    expect(result.isExceptional()).toBe(true);
    expect(result.exception).toEqual(
      new Error(`Transaction already started (${requestId})`),
    );
    expect(context.getClient()).toBe(client);
  });

  it('clears the client', () => {
    const context = new Context(requestId);
    const client = createClient();
    context.setClient(client);

    const result = context.clearClient();

    expect(result.isSuccess()).toBe(true);
    expect(result.isExceptional()).toBe(false);
    expect(result.value).toBeUndefined();
    expect(context.getClient()).toBeUndefined();
  });

  it('stores a new client after the previous one was cleared', () => {
    const context = new Context(requestId);
    const client = createClient();
    const nextClient = createClient();
    context.setClient(client);
    context.clearClient();

    const result = context.setClient(nextClient);

    expect(result.isSuccess()).toBe(true);
    expect(context.getClient()).toBe(nextClient);
  });

  it('clears a context that has no client', () => {
    const context = new Context(requestId);

    const result = context.clearClient();

    expect(result.isSuccess()).toBe(true);
    expect(context.getClient()).toBeUndefined();
  });
});

function createClient(): PoolClient {
  return {} as PoolClient;
}
