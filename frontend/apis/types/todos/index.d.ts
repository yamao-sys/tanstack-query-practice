import { components } from "@/apis/generated/todos/apiSchema";

export type StoreTodoInput =
  components["requestBodies"]["StoreTodoInput"]["content"]["application/json"];

export type StoreTodoValidationError = components["schemas"]["StoreTodoValidationError"];
