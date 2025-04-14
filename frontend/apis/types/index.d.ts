import { components } from "../apiSchema";

export type AuthSignUpInput =
  components["requestBodies"]["SignUpInput"]["content"]["multipart/form-data"];

export type AuthSignUpValidationError = components["schemas"]["SignUpValidationError"];

export type AuthSignInInput =
  components["requestBodies"]["SignInInput"]["content"]["application/json"];

export type StoreTodoInput =
  components["requestBodies"]["StoreTodoInput"]["content"]["application/json"];

export type StoreTodoValidationError = components["schemas"]["StoreTodoValidationError"];

export type Todo =
  components["responses"]["ShowTodoResponse"]["content"]["application/json"]["todo"];
