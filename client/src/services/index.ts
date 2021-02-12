import { UserService } from './user.service';

import {
  ApiUserRepository,
} from '@/impl/api';

import {
  LocalStorageTokenRepository,
} from '@/impl/tokenstorage';

const apiUserRepo = new ApiUserRepository();
const storageTokenRepo = new LocalStorageTokenRepository();

export const userService = new UserService(apiUserRepo, storageTokenRepo);