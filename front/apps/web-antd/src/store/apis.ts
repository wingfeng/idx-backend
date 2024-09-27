import type { PageParams } from '#/types/page';

import { delAPI, getAPIList, saveAPI } from '#/api/oauth2/apis';

export async function GetAPIPage(params: PageParams) {
  return getAPIList(params);
  //  return result.data;
}

export async function SaveAPI(data: any) {
  return saveAPI(data);
}

export async function DeleteAPI(id: string) {
  return delAPI(id);
}
