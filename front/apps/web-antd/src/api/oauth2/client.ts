import type { ClientInfo } from '#/types/client';
import type { PageParams, PageResult } from '#/types/page';
import type { APIResult } from '#/types/result';

import { requestClient } from '#/api/request';

const apiURL: any = {
  Page: '/api/v1/oauth2/client/page',
  Get: '/api/v1/oauth2/client/get',
  PUT: '/api/v1/oauth2/client',
  Delete: '/api/v1/oauth2/client/del',
};

export async function getClientList(params: PageParams) {
  return requestClient.post<APIResult<PageResult<ClientInfo>>>(
    apiURL.Page,
    params,
  );
}

export async function getClient(id: string) {
  return requestClient.get<APIResult<ClientInfo>>(apiURL.Get, {
    params: { id },
  });
}
export async function delClient(id: string) {
  return requestClient.delete<APIResult<ClientInfo>>(apiURL.Delete, {
    params: {
      id,
    },
  });
}
export async function saveClient(info: ClientInfo) {
  return requestClient.put<APIResult<ClientInfo>>(apiURL.PUT, info);
}
