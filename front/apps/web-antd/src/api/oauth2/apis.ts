import type { APIInfo } from '#/types/apis';
import type { PageParams, PageResult } from '#/types/page';
import type { APIResult } from '#/types/result';

import { requestClient } from '#/api/request';

const apiURL: any = {
  Page: '/api/v1/oauth2/api/page',
  Get: '/api/v1/oauth2/api/get',
  PUT: '/api/v1/oauth2/api',
  Delete: '/api/v1/oauth2/api/del',
};

export async function getAPIList(params: PageParams) {
  return requestClient.post<APIResult<PageResult<APIInfo>>>(
    apiURL.Page,
    params,
  );
}

export async function getAPI(id: string) {
  return requestClient.get<APIResult<APIInfo>>(apiURL.Get, {
    params: { id },
  });
}
export async function delAPI(id: string) {
  return requestClient.delete<APIResult<APIInfo>>(apiURL.Delete, {
    params: {
      id,
    },
  });
}
export async function saveAPI(info: APIInfo) {
  return requestClient.put<APIResult<APIInfo>>(apiURL.PUT, info);
}
