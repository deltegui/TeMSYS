import { Inject, Injectable } from "@angular/core";
import { ApiUserRepository } from "./api";
import { LocalStorageTokenRepository } from './localstoragetokenrepo';
import { User } from './models';

@Injectable()
export class UserService {
  constructor(
    private userRepo: ApiUserRepository,
    private tokenRepo: LocalStorageTokenRepository,
  ) {}

  async login({ name, password }: { name: string, password: string }): Promise<User> {
    return this.userRepo.login({
      name,
      password,
    }).then(user => {
      this.tokenRepo.save(user.token);
      return user;
    });
  }
}