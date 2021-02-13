import {
  inject,
  reactive,
  readonly,
} from 'vue';
import { Token } from '@/services/models';
import TokenRepo from '@/impl/tokenstorage';

export type State = {
  state: Readonly<{ token?: Token }>;
  deleteToken(): void;
  setToken(other: Token): void;
};

const tokenRepo = new TokenRepo();
const prevToken = tokenRepo.load();

export const stateSymbol = Symbol('state');

export const createState = (): State => {
  const state = reactive({
    token: prevToken,
  });
  const setToken = (other: Token) => {
    state.token = other;
  };
  const deleteToken = () => {
    state.token = undefined;
  };

  return {
    setToken,
    deleteToken,
    state: readonly(state),
  };
};

export const useState = () => inject<State>(stateSymbol);
