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

  load(): Token {
    const value = this.storage.getItem('token');
    return JSON.parse(value ?? '');
  }
}
