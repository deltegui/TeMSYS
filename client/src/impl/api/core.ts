/* eslint-disable no-throw-literal */

const apiURL = (process.env.NODE_ENV === 'production') ? 'https://192.168.1.31/api' : 'http://localhost:3000/api';

const tokenExpiredEventHandlers: TokenExpiredHandler[] = [];

type RawError = {
  Code: number;
  Reason: string;
}

export type TokenExpiredHandler = (err: RawError) => void;

export function onTokenExpired(handler: TokenExpiredHandler) {
  tokenExpiredEventHandlers.push(handler);
}

function dispatchTokenExpiredEvent(err: RawError) {
  tokenExpiredEventHandlers.forEach((handler) => handler(err));
}

function isApiError(err: RawError): boolean {
  return (!!err.Code || err.Code === 0) && !!err.Reason;
}

function isErrorTokenExpired(err: RawError): boolean {
  return err.Code === 302;
}

async function handleResponse(res: any) {
  if (res.ok) {
    return res.json();
  }
  throw await res.json();
}

async function handleError(err: RawError) {
  if (!isApiError(err)) {
    throw {
      code: -1,
      reason: err.toString(),
    };
  }
  if (isErrorTokenExpired(err)) {
    dispatchTokenExpiredEvent(err);
  }
  throw {
    code: err.Code,
    reason: err.Reason,
  };
}

type RequestOptions = {
  body?: any;
  method: string;
  token?: string;
}

export default async function makeRequest(
  endpoint: string,
  options: RequestOptions = {
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
    .then(handleResponse)
    .catch(handleError);
}
