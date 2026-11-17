import { IResult } from './result.interface';
import { Result } from './result';

type Asset = {
  id: string;
};

type AssetNotFound = {
  code: 'ASSET_NOT_FOUND';
  id: string;
};

describe('Result', () => {
  describe('success', () => {
    it('stores the value and reports a successful result', () => {
      const result = Result.success({ id: 'asset-1' });

      expect(result.isSuccess()).toBe(true);
      expect(result.isExceptional()).toBe(false);
      expect(result.value).toEqual({ id: 'asset-1' });
      expect(() => result.exception).toThrow(
        'Cannot get error of a successful result',
      );
    });

    it('creates a successful result without a value', () => {
      const result = Result.success();

      expect(result.isSuccess()).toBe(true);
      expect(result.isExceptional()).toBe(false);
      expect(result.value).toBeUndefined();
      expect(() => result.exception).toThrow(
        'Cannot get error of a successful result',
      );
    });

    it('keeps falsy values on a successful result', () => {
      const values = [0, false, '', null] as const;

      for (const value of values) {
        const result = Result.success(value);

        expect(result.isSuccess()).toBe(true);
        expect(result.isExceptional()).toBe(false);
        expect(result.value).toBe(value);
        expect(() => result.exception).toThrow(
          'Cannot get error of a successful result',
        );
      }
    });

    it('stays successful when the value is undefined', () => {
      const result = Result.success(undefined);

      expect(result.isSuccess()).toBe(true);
      expect(result.isExceptional()).toBe(false);
      expect(result.value).toBeUndefined();
      expect(() => result.exception).toThrow(
        'Cannot get error of a successful result',
      );
    });
  });

  describe('fail', () => {
    it('stores the error and reports an exceptional result', () => {
      const error = new Error('Asset not found');
      const result = Result.fail(error);

      expect(result.isSuccess()).toBe(false);
      expect(result.isExceptional()).toBe(true);
      expect(result.exception).toBe(error);
      expect(() => result.value).toThrow(
        'Cannot get value of an exceptional result',
      );
    });

    it('stores a non-Error failure', () => {
      const error: AssetNotFound = { code: 'ASSET_NOT_FOUND', id: 'asset-1' };
      const result = Result.fail(error);

      expect(result.isSuccess()).toBe(false);
      expect(result.isExceptional()).toBe(true);
      expect(result.exception).toEqual(error);
      expect(() => result.value).toThrow(
        'Cannot get value of an exceptional result',
      );
    });
  });

  describe('IResult', () => {
    it('accepts Result where another type requires the interface', () => {
      const found: IResult<Asset, Error> = findAsset('asset-1');
      const missing: IResult<Asset, Error> = findAsset('missing');

      expect(found.isSuccess()).toBe(true);
      expect(missing.isExceptional()).toBe(true);
    });

    it('returns the value after a successful check', () => {
      expect(valueOf(Result.success({ id: 'asset-1' }))).toEqual({
        id: 'asset-1',
      });
    });

    it('returns the error after an exceptional check', () => {
      expect(messageOf(Result.fail(new Error('Asset not found')))).toBe(
        'Asset not found',
      );
    });
  });
});

function findAsset(id: string): Result<Asset, Error> {
  if (id !== 'asset-1') {
    return Result.fail(new Error('Asset not found'));
  }

  return Result.success({ id });
}

function valueOf(result: Result<Asset, Error>): Asset {
  if (!result.isSuccess()) {
    throw new Error('Expected a successful result');
  }

  return result.value;
}

function messageOf(result: Result<Asset, Error>): string {
  if (!result.isExceptional()) {
    throw new Error('Expected an exceptional result');
  }

  return result.exception.message;
}
