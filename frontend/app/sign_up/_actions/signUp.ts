"use server";

import { paths } from "@/apis/apiSchema";
import { getRequestHeaders } from "@/apis/clients/base";
import { AuthSignUpInput } from "@/apis/types";
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
