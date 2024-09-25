import type { PageParams } from '#/types/page';

import { delClient, getClientList, saveClient } from '#/api/oauth2/client';

export async function GetClientsPage(params: PageParams) {
  return getClientList(params);
  //  return result.data;
}

export async function SaveClient(data: any) {
  return saveClient(data);
}

export async function DeleteClient(id: string) {
  return delClient(id);
}
