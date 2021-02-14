/* eslint-disable no-useless-constructor */
import {
  store,
  actions,
} from '@/store';
import {
  UserRepository,
  TokenRepository,
} from './gateways';
import {
  Token, UserResponse,
} from './models';

export default class UserService {
  constructor(
    private userRepo: UserRepository,
    private tokenRepo: TokenRepository,
  ) {}

  async login({ name, password }: { name: string; password: string }): Promise<Token> {
    return this.userRepo.login({
      name,
      password,
    }).then((token) => {
      actions.tokens.set(token);
      this.tokenRepo.save(token);
      return token;
    });
  }

  logout(): void {
    actions.tokens.delete();
    this.tokenRepo.clear();
  }

  async create(user: { name: string; password: string }) {
    return this.userRepo.createUser(user, store.token?.value ?? '');
  }

  async delete(user: string): Promise<string> {
    return this.userRepo.deleteUser(user, store.token?.value ?? '');
  }

  async getAll(): Promise<UserResponse[]> {
    return this.userRepo.getAll(store.token?.value ?? '');
  }
}
