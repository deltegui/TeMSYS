/* eslint-disable no-throw-literal */

const apiURL = 'http://localhost:3000/api';

function isApiError(err: any): boolean {
  return (!!err.Code || err.Code === 0) && !!err.Reason;
}

export default async function makeRequest(
  endpoint: string,
  body: any = undefined,
  method = 'GET',
): Promise<any> {
  let reqConfig;
  if (body) {
    reqConfig = {
      method,
      body: JSON.stringify(body),
      headers: {
        'Content-Type': 'application/json',
      },
    };
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
