/* eslint-disable class-methods-use-this */
import {
  Token, UserResponse,
} from '@/services/models';
import {
  UserRepository,
} from '@/services/gateways';

import makeRequest from './core';

export default class ApiUserRepository implements UserRepository {
  async login(body: { name: string; password: string }): Promise<Token> {
    return makeRequest('/user/login', { body, method: 'POST' })
      .then(({ token }) => ({
        value: token.value,
        expires: token.expires,
        owner: token.owner,
        role: token.role,
      }));
  }

  async createUser(user: { name: string; password: string }, token: string): Promise<UserResponse> {
    return makeRequest('/user/create', { body: user, method: 'POST', token })
      .then((res) => ({
        name: res.Name,
        role: res.Role,
      }));
  }

  async deleteUser(user: string, token: string): Promise<string> {
    return makeRequest(`/user/${user}`, { method: 'DELETE', token })
      .then((res) => res.Name);
  }

  async getAll(token: string): Promise<UserResponse[]> {
    return makeRequest('/user/all', { method: 'GET', token })
      .then((res) => {
        if (!res) return [];
        return res.map(({ Name, Role }: { Name: string; Role: string }) => ({
          name: Name,
          role: Role,
        }));
      });
  }
}
