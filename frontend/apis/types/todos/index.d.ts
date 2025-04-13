import { components } from "@/apis/generated/todos/apiSchema";

export type StoreTodoInput =
  components["requestBodies"]["StoreTodoInput"]["content"]["application/json"];

export type StoreTodoValidationError = components["schemas"]["StoreTodoValidationError"];

export type Todo =
  components["responses"]["ShowTodoResponse"]["content"]["application/json"]["todo"];
