"use server";

import { getRequestHeaders } from "@/apis/clients/base";
import { paths } from "@/apis/generated/auth/apiSchema";
import { AuthSignUpInput } from "@/apis/types/auth";
import createClient from "openapi-fetch";

const client = createClient<paths>({
  baseUrl: `${process.env.API_ENDPOINT_URI ?? "http://api_server:8080"}/`,
  credentials: "include",
});

export async function postSignUp(input: AuthSignUpInput) {
  const { data, error } = await client.POST("/auth/signUp", {
    ...(await getRequestHeaders()),
    body: input,
    bodySerializer(body) {
      const formData = new FormData();

      if (body) {
        for (const [key, value] of Object.entries(input)) {
          if (value instanceof File) {
            formData.append(key, value, encodeURI(value.name));
          } else {
            formData.append(key, value);
          }
        }
      }
      return formData;
    },
  });
  if (error?.code === 500 || data === undefined) {
    throw Error("Internal Server Error");
  }

  return data;
}
