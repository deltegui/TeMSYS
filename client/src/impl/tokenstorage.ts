import { TokenRepository } from '@/services/gateways';
import { Token } from '@/services/models';

export default class LocalStorageTokenRepository implements TokenRepository {
  private storage: Storage;

  constructor() {
    this.storage = window.localStorage;
  }

  save(token: Token): void {
    this.storage.setItem('token', JSON.stringify(token));
  }

  load(): Token | undefined {
    const value = this.storage.getItem('token');
    if (!value) {
      return undefined;
    }
    return JSON.parse(value);
  }

  clear() {
    this.storage.clear();
  }
}
