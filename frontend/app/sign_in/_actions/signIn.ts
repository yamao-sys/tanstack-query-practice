"use server";

import { paths } from "@/apis/apiSchema";
import { getRequestHeaders } from "@/apis/clients/base";
import { AuthSignInInput } from "@/apis/types";
import { cookies } from "next/headers";
import createClient from "openapi-fetch";

const client = createClient<paths>({
  baseUrl: `${process.env.API_ENDPOINT_URI ?? "http://api_server:8080"}/`,
  credentials: "include",
});

export async function postSignIn(input: AuthSignInInput) {
  const { response } = await client.POST("/auth/signIn", {
    ...(await getRequestHeaders()),
    body: input,
  });
  if (response.status === 500) {
    throw Error("Internal Server Error");
  }
  if (response.status === 400) {
    return "メールアドレスまたはパスワードが正しくありません";
  }

  // NOTE: クライアントにCookieをセット
  const setCookie = response.headers.get("set-cookie");
  if (!setCookie) {
    throw Error();
  }
  const token = setCookie?.split(";")[0]?.split("=")[1];
  if (!token) {
    throw Error();
  }

  // TODO: cookieの属性は環境変数に切り出す
  (await cookies()).set({
    name: "token",
    value: token,
    secure: true,
    sameSite: "none",
    httpOnly: true,
  });

  return "";
}
