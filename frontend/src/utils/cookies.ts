export enum CookieName {
  Role = 'role',
  UserId = 'userId',
  Token = 'token',
  Expires = 'expires',
}

export const getCookieValueByName = (name: CookieName): string | undefined =>
  document.cookie
    .split('; ')
    .find((row) => row.startsWith(name))
    ?.split(`${name}=`)[1];

export const setCookie = (name: CookieName, value: string): void => {
  document.cookie = name + '=' + value + ';';
};

export const eraseCookie = (name: CookieName): void => {
  document.cookie = name + '=;';
};
