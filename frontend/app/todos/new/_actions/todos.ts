"use server";

import { getRequestHeaders } from "@/apis/clients/base";
import { paths } from "@/apis/generated/todos/apiSchema";
import { StoreTodoInput, StoreTodoValidationError } from "@/apis/types/todos";
import createClient from "openapi-fetch";

const client = createClient<paths>({
  baseUrl: `${process.env.API_ENDPOINT_URI ?? "http://api_server:8080"}/`,
  credentials: "include",
});

export async function postTodos(input: StoreTodoInput): Promise<StoreTodoValidationError> {
  const { data, response } = await client.POST("/todos", {
    ...(await getRequestHeaders()),
    body: input,
  });
  if (response.status === 500) {
    throw Error("Internal Server Error");
  }
  if (!!data?.errors && response.status === 400) {
    return data.errors;
  }

  return {};
}
