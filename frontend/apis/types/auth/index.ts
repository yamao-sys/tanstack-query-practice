import { components } from "@/apis/generated/auth/apiSchema";

export type AuthSignUpInput =
  components["requestBodies"]["SignUpInput"]["content"]["multipart/form-data"];

export type AuthSignUpValidationError = components["schemas"]["SignUpValidationError"];
