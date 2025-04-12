"use client";

import { StoreTodoInput, StoreTodoValidationError } from "@/apis/types/todos";
import { FC } from "react";
import { postTodos } from "../../_actions/todos";
import { useRouter } from "next/navigation";
import { useMutation } from "@tanstack/react-query";
import { TodoStoreForm } from "@/app/todos/_components/TodoStoreForm";
import { useStoreTodo } from "@/app/todos/_hooks/useStoreTodo";

const INITIAL_VALIDATION_ERRORS = {
  title: [],
  content: [],
};

export const TodosNewForm: FC = () => {
  const doStoreTodoInput: StoreTodoInput = { title: "", content: "" };
  const { register, handleSubmit, validationErrors, setValidationErrors } =
    useStoreTodo(doStoreTodoInput);

  const router = useRouter();

  const mutation = useMutation({
    onMutate: () => setValidationErrors(INITIAL_VALIDATION_ERRORS),
    mutationFn: async (data: StoreTodoInput): Promise<StoreTodoValidationError> =>
      await postTodos(data),
    onSuccess: (data) => {
      if (data === undefined) return;

      if (Object.keys(data).length > 0) {
        setValidationErrors(data);
        return;
      }

      window.alert("TODOの作成に成功しました!");
      router.push("/");
    },
    onError: () => window.alert("予期しないエラーが発生しました."),
  });

  const onSubmit = handleSubmit((data) => mutation.mutate(data));

  return (
    <>
      {mutation.isPending && <p>Processing...</p>}

      <TodoStoreForm
        header='TODO作成'
        register={register}
        onSubmit={onSubmit}
        validationErrors={validationErrors}
      />
    </>
  );
};
