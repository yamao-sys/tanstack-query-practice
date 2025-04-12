export interface paths {
  "/todos": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    /**
     * Fetch Todos
     * @description Fetch Todos Schema
     */
    get: operations["get-todos"];
    put?: never;
    /**
     * Create Todo
     * @description Create Todo Schema
     */
    post: operations["post-todos"];
    delete?: never;
    options?: never;
    head?: never;
    patch?: never;
    trace?: never;
  };
  "/todos/{id}": {
    parameters: {
      query?: never;
      header?: never;
      path: {
        id: string;
      };
      cookie?: never;
    };
    /**
     * Show Todo
     * @description Show Todo Schema
     */
    get: operations["get-todo"];
    put?: never;
    post?: never;
    /**
     * Delete Todo
     * @description Delete Todo Schema
     */
    delete: operations["delete-todo"];
    options?: never;
    head?: never;
    /**
     * Update Todo
     * @description Update Todo Schema
     */
    patch: operations["patch-todo"];
    trace?: never;
  };
}
export type webhooks = Record<string, never>;
export interface components {
  schemas: {
    /** Todo Object */
    Todo: {
      id: number;
      title: string;
      content: string;
    };
    /** StoreTodoValidationError */
    StoreTodoValidationError: {
      title?: string[];
      content?: string[];
    };
  };
  responses: {
    /** @description Fetch Todos Response */
    FetchTodosResponse: {
      headers: {
        [name: string]: unknown;
      };
      content: {
        "application/json": {
          todos: components["schemas"]["Todo"][];
        };
      };
    };
    /** @description Show Todo Response */
    ShowTodoResponse: {
      headers: {
        [name: string]: unknown;
      };
      content: {
        "application/json": {
          todo: components["schemas"]["Todo"];
        };
      };
    };
    StoreTodoResponse: {
      headers: {
        [name: string]: unknown;
      };
      content: {
        "application/json": {
          /** Format: int64 */
          code: number;
          errors: components["schemas"]["StoreTodoValidationError"];
        };
      };
    };
    DeleteTodoResponse: {
      headers: {
        [name: string]: unknown;
      };
      content: {
        "application/json": {
          /** Format: int64 */
          code: number;
          result: boolean;
        };
      };
    };
    /** @description Unauthorized Error Response */
    UnauthorizedErrorResponse: {
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
    /** @description Not Found Error Response */
    NotFoundErrorResponse: {
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
    /** @description Todo Iuput */
    StoreTodoInput: {
      content: {
        "application/json": {
          title: string;
          content: string;
        };
      };
    };
  };
  headers: never;
  pathItems: never;
}
export type $defs = Record<string, never>;
export interface operations {
  "get-todos": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    requestBody?: never;
    responses: {
      200: components["responses"]["FetchTodosResponse"];
      401: components["responses"]["UnauthorizedErrorResponse"];
      500: components["responses"]["InternalServerErrorResponse"];
    };
  };
  "post-todos": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    requestBody?: components["requestBodies"]["StoreTodoInput"];
    responses: {
      200: components["responses"]["StoreTodoResponse"];
      400: components["responses"]["StoreTodoResponse"];
      401: components["responses"]["UnauthorizedErrorResponse"];
      500: components["responses"]["InternalServerErrorResponse"];
    };
  };
  "get-todo": {
    parameters: {
      query?: never;
      header?: never;
      path: {
        id: string;
      };
      cookie?: never;
    };
    requestBody?: never;
    responses: {
      200: components["responses"]["ShowTodoResponse"];
      401: components["responses"]["UnauthorizedErrorResponse"];
      404: components["responses"]["NotFoundErrorResponse"];
      500: components["responses"]["InternalServerErrorResponse"];
    };
  };
  "delete-todo": {
    parameters: {
      query?: never;
      header?: never;
      path: {
        id: string;
      };
      cookie?: never;
    };
    requestBody?: never;
    responses: {
      200: components["responses"]["DeleteTodoResponse"];
      401: components["responses"]["UnauthorizedErrorResponse"];
      404: components["responses"]["NotFoundErrorResponse"];
      500: components["responses"]["InternalServerErrorResponse"];
    };
  };
  "patch-todo": {
    parameters: {
      query?: never;
      header?: never;
      path: {
        id: string;
      };
      cookie?: never;
    };
    requestBody?: components["requestBodies"]["StoreTodoInput"];
    responses: {
      200: components["responses"]["StoreTodoResponse"];
      400: components["responses"]["StoreTodoResponse"];
      401: components["responses"]["UnauthorizedErrorResponse"];
      404: components["responses"]["NotFoundErrorResponse"];
      500: components["responses"]["InternalServerErrorResponse"];
    };
  };
}
