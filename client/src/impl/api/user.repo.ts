/* eslint-disable class-methods-use-this */

import {
  Token,
} from '@/services/models';

import {
  UserRepository,
} from '@/services/gateways';

import makeRequest from './core';

export default class ApiUserRepository implements UserRepository {
  async login(body: { name: string; password: string }): Promise<Token> {
    return makeRequest('/user/login', body, 'POST')
      .then(({ token }) => ({
        value: token.value,
        expires: token.expires,
        owner: token.owner,
        role: token.role,
      }));
  }
}
