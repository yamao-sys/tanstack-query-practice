import { cookies } from "next/headers";
import createClient from "openapi-fetch";
import { paths } from "../apiSchema";

export const getRequestHeaders = async () => {
  const csrfToken = (await cookies()).get("_csrf")?.value ?? "";
  const authenticateToken = (await cookies()).get("token")?.value ?? "";
  return {
    headers: {
      "X-CSRF-Token": csrfToken,
      Cookie: `_csrf=${csrfToken}; token=${authenticateToken}`,
    },
  };
};

export async function setCsrfToken() {
  const client = createClient<paths>({
    baseUrl: `${process.env.API_ENDPOINT_URI ?? "http://api_server:8080"}/`,
    credentials: "include",
  });

  const { data, error, response } = await client.GET("/auth/csrf", {});
  if (error?.code === 500 || data === undefined) {
    throw Error();
  }

  // NOTE: クライアントにCookieをセット
  const setCookie = response.headers.get("set-cookie");
  if (!setCookie) {
    throw Error();
  }
  const csrfToken = setCookie?.split(";")[0]?.split("=")[1];
  if (!csrfToken) {
    throw Error();
  }

  // TODO: cookieの属性は環境変数に切り出す
  (await cookies()).set({
    name: "_csrf",
    value: csrfToken,
    secure: true,
    sameSite: "none",
    httpOnly: true,
  });
}
