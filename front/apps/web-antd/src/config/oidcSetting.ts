import { WebStorageStateStore } from 'oidc-client-ts';

export const oidcSetting: any = {
  authority: 'http://localhost:9097/idx',
  client_id: 'code_client',
  client_secret: 'secret',
  redirect_uri: 'http://localhost:5666/auth/callback',
  response_type: 'code',
  scope: 'openid profile email',
  stateStore: new WebStorageStateStore({ store: window.localStorage }),
};
