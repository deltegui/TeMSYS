/* eslint-disable no-throw-literal */

const apiURL = 'http://localhost:3000/api';

function isApiError(err: any): boolean {
  return (!!err.Code || err.Code === 0) && !!err.Reason;
}

export default async function makeRequest(
  endpoint: string,
  options: {
    body?: any;
    method: string;
    token?: string;
  } = {
    body: undefined,
    method: 'GET',
    token: undefined,
  },
): Promise<any> {
  const reqConfig: any = {
    method: options.method,
    headers: {},
  };
  if (options.body) {
    reqConfig.body = JSON.stringify(options.body);
    reqConfig.headers['Content-Type'] = 'application/json';
  }
  if (options.token) {
    reqConfig.headers.Authorization = `Bearer ${options.token}`;
  }
  return fetch(`${apiURL}${endpoint}`, reqConfig)
    .then(async (res) => {
      if (res.ok) {
        return res.json();
      }
      throw await res.json();
    })
    .catch((err) => {
      if (isApiError(err)) {
        throw {
          code: err.Code,
          reason: err.Reason,
        };
      }
      throw {
        code: -1,
        reason: err.toString(),
      };
    });
}
