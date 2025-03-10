export interface paths {
  "/auth/validateSignUp": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    get?: never;
    put?: never;
    /**
     * Validate SignUp
     * @description validate sign up
     */
    post: operations["post-auth-validate_sign_up"];
    delete?: never;
    options?: never;
    head?: never;
    patch?: never;
    trace?: never;
  };
  "/auth/signUp": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    get?: never;
    put?: never;
    /** SignUp */
    post: operations["post-auth-sign_up"];
    delete?: never;
    options?: never;
    head?: never;
    patch?: never;
    trace?: never;
  };
  "/auth/signIn": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    get?: never;
    put?: never;
    /** Sign In */
    post: operations["post-auth-sign_in"];
    delete?: never;
    options?: never;
    head?: never;
    patch?: never;
    trace?: never;
  };
  "/auth/csrf": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    /** Get Csrf */
    get: operations["get-auth-csrf"];
    put?: never;
    post?: never;
    delete?: never;
    options?: never;
    head?: never;
    patch?: never;
    trace?: never;
  };
}
export type webhooks = Record<string, never>;
export interface components {
  schemas: {
    /** SignUpValidationError */
    SignUpValidationError: {
      firstName?: string[];
      lastName?: string[];
      email?: string[];
      password?: string[];
      birthday?: string[];
      frontIdentification?: string[];
      backIdentification?: string[];
    };
  };
  responses: {
    SignUpResponse: {
      headers: {
        [name: string]: unknown;
      };
      content: {
        "application/json": {
          /** Format: int64 */
          code: number;
          errors: components["schemas"]["SignUpValidationError"];
        };
      };
    };
    /** @description SignIn Response */
    SignInOkResponse: {
      headers: {
        "Set-Cookie"?: string;
        [name: string]: unknown;
      };
      content: {
        "application/json": Record<string, never>;
      };
    };
    /** @description SignIn BadRequest Response */
    SignInBadRequestResponse: {
      headers: {
        [name: string]: unknown;
      };
      content: {
        "application/json": {
          errors: string[];
        };
      };
    };
    /** @description Csrf response */
    CsrfResponse: {
      headers: {
        [name: string]: unknown;
      };
      content: {
        "application/json": {
          csrf_token: string;
        };
      };
    };
    /** @description Internal Server Error Response */
    InternalServerErrorResponse: {
      headers: {
        [name: string]: unknown;
      };
      content: {
        "application/json": {
          /** Format: int64 */
          code: number;
          message: string;
        };
      };
    };
  };
  parameters: never;
  requestBodies: {
    /** @description SignUp  Iuput */
    SignUpInput: {
      content: {
        "multipart/form-data": {
          firstName: string;
          lastName: string;
          email: string;
          password: string;
          /** Format: date */
          birthday?: string;
          /** Format: binary */
          frontIdentification?: Blob;
          /** Format: binary */
          backIdentification?: Blob;
        };
      };
    };
    /** @description SignIn  Input */
    SignInInput: {
      content: {
        "application/json": {
          email: string;
          password: string;
        };
      };
    };
  };
  headers: never;
  pathItems: never;
}
export type $defs = Record<string, never>;
export interface operations {
  "post-auth-validate_sign_up": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    requestBody?: components["requestBodies"]["SignUpInput"];
    responses: {
      200: components["responses"]["SignUpResponse"];
      400: components["responses"]["SignUpResponse"];
      500: components["responses"]["InternalServerErrorResponse"];
    };
  };
  "post-auth-sign_up": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    requestBody?: components["requestBodies"]["SignUpInput"];
    responses: {
      200: components["responses"]["SignUpResponse"];
      400: components["responses"]["SignUpResponse"];
      500: components["responses"]["InternalServerErrorResponse"];
    };
  };
  "post-auth-sign_in": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    requestBody?: components["requestBodies"]["SignInInput"];
    responses: {
      200: components["responses"]["SignInOkResponse"];
      400: components["responses"]["SignInBadRequestResponse"];
      500: components["responses"]["InternalServerErrorResponse"];
    };
  };
  "get-auth-csrf": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    requestBody?: never;
    responses: {
      200: components["responses"]["CsrfResponse"];
      500: components["responses"]["InternalServerErrorResponse"];
    };
  };
}
