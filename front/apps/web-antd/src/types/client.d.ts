export interface ClientInfo {
  Id: number;
  Enabled: boolean;
  ClientId: string;
  ClientName: string;
  ClientURI?: string;
  Description?: string;
  Scopes: string;
  GrantTypes: string;
  RedirectUris?: string;
  RequireConsent: boolean;
  RequirePkce: boolean;
  RequireSecret: boolean;
  LogoURI?: string;
}
