import { StoreTodoInput, StoreTodoValidationError } from "@/apis/types/todos";
import { useState } from "react";
import { useForm } from "react-hook-form";

export const INITIAL_VALIDATION_ERRORS = {
  title: [],
  content: [],
};

export const useStoreTodo = (doStoreTodoInput: StoreTodoInput) => {
  const { register, handleSubmit } = useForm<StoreTodoInput>({
    defaultValues: doStoreTodoInput,
  });

  const [validationErrors, setValidationErrors] =
    useState<StoreTodoValidationError>(INITIAL_VALIDATION_ERRORS);

  return {
    register,
    handleSubmit,
    validationErrors,
    setValidationErrors,
  };
};
