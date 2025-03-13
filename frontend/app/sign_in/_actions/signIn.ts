"use server";

import { getRequestHeaders } from "@/apis/clients/base";
import { paths } from "@/apis/generated/auth/apiSchema";
import { AuthSignInInput } from "@/apis/types/auth";
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
  const statusCode = response.status;
  if (statusCode === 500) {
    throw Error("Internal Server Error");
  }

  return statusCode === 400 ? "メールアドレスまたはパスワードに該当するユーザが存在しません。" : "";
}
