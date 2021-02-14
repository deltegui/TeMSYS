import {
  inject,
  reactive,
  readonly,
} from 'vue';
import { Token } from '@/services/models';
import TokenRepo from '@/impl/tokenstorage';

export type State = {
  token?: Token;
};

const tokenRepo = new TokenRepo();
const prevToken = tokenRepo.load();

export const stateSymbol = Symbol('state');
export const useState = () => inject<Readonly<State>>(stateSymbol);

const state = reactive({
  token: prevToken,
});

export const store = readonly(state);

export const actions = {
  tokens: {
    set: (other: Token) => {
      state.token = other;
    },

    delete: () => {
      state.token = undefined;
    },
  },
};
