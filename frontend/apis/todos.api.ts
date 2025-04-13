"use server";

import { getRequestHeaders } from "@/apis/clients/base";
import { paths } from "@/apis/generated/todos/apiSchema";
import { StoreTodoInput, StoreTodoValidationError, Todo } from "@/apis/types/todos";
import createClient from "openapi-fetch";

const client = createClient<paths>({
  baseUrl: `${process.env.API_ENDPOINT_URI ?? "http://api_server:8080"}/`,
  credentials: "include",
});

export async function getTodos(): Promise<Todo[]> {
  const { data, response } = await client.GET("/todos", {
    ...(await getRequestHeaders()),
  });
  if (data === undefined || response.status === 404) {
    throw Error("Not Found Error");
  }

  return data.todos;
}

export async function postTodos(input: StoreTodoInput): Promise<StoreTodoValidationError> {
  const { data, response } = await client.POST("/todos", {
    ...(await getRequestHeaders()),
    body: input,
  });
  if (response.status === 500 || !data?.errors) {
    throw Error("Internal Server Error");
  }

  return data.errors;
}

export async function getTodo(id: string): Promise<Todo> {
  const { data, response } = await client.GET("/todos/{id}", {
    ...(await getRequestHeaders()),
    params: {
      path: {
        id,
      },
    },
  });
  if (data === undefined || response.status === 404) {
    throw Error("Not Found Error");
  }

  return data.todo;
}

export async function patchTodo(id: string, input: StoreTodoInput) {
  console.log(input);
  const { data, response } = await client.PATCH("/todos/{id}", {
    ...(await getRequestHeaders()),
    params: {
      path: {
        id,
      },
    },
    body: input,
  });
  if (data === undefined || response.status === 500) {
    throw Error("Internal Server Error");
  }
  if (response.status === 404) {
    throw Error("Not Found Error");
  }

  return data.errors;
}
