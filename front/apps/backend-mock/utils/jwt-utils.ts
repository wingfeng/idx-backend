import type { EventHandlerRequest, H3Event } from 'h3';

import jwt from 'jsonwebtoken';
import jwksClient from 'jwks-rsa';

import { UserInfo } from './mock-data';

// TODO: Replace with your own secret key
const ACCESS_TOKEN_SECRET = 'access_token_secret';
const REFRESH_TOKEN_SECRET = 'refresh_token_secret';

const client = jwksClient({
  jwksUri: 'http://localhost:9097/idx/.well-known/jwks',
});
function getKey() {
  client.getSigningKey('d2a820a8916647f7ac72627ec0ae4f94', (err, key) => {
    if (err) {
      console.log(err);
    }
    const signingKey = key.getPublicKey;
    return signingKey;
  });
}

export interface UserPayload extends UserInfo {
  iat: number;
  exp: number;
}

export function generateAccessToken(user: UserInfo) {
  return jwt.sign(user, ACCESS_TOKEN_SECRET, { expiresIn: '7d' });
}

export function generateRefreshToken(user: UserInfo) {
  return jwt.sign(user, REFRESH_TOKEN_SECRET, {
    expiresIn: '30d',
  });
}

export function verifyAccessToken(
  event: H3Event<EventHandlerRequest>,
): null | Omit<UserInfo, 'password'> {
  const authHeader = getHeader(event, 'Authorization');
  if (!authHeader?.startsWith('Bearer')) {
    return null;
  }

  const token = authHeader.split(' ')[1];
  console.log('verifyAccessToken begin');
  try {
    jwt.verify(token, getKey, null, (err, decoded) => {
      console.log(err, decoded);
      const username = decoded.payload;
      const user = MOCK_USERS.find((item) => item.username === username);
      const { password: _pwd, ...userinfo } = user;
      return userinfo;
    });
  } catch (error) {
    console.log('verifyAccessToken error', error);
    return null;
  }
}

export function verifyRefreshToken(
  token: string,
): null | Omit<UserInfo, 'password'> {
  try {
    jwt.verify(token, getKey, null, (err, decoded) => {
      console.log(err, decoded);
      const username = decoded.payload;
      const user = MOCK_USERS.find((item) => item.username === username);
      const { password: _pwd, ...userinfo } = user;
      return userinfo;
    });
  } catch (error) {
    console.log(error);
    return null;
  }
}
