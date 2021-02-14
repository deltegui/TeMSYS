/* eslint-disable no-useless-constructor */
import {
  actions,
} from '@/store';
import {
  UserRepository,
  TokenRepository,
} from './gateways';
import {
  Token,
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
}
