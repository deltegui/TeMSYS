import { store } from '@/store';

if (process.env.NODE_ENV !== 'production') {
  const temsys = {
    getStore() {
      return store;
    },
  };
  window.temsys = temsys;
}
