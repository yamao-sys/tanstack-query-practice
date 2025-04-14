"use client";

import { patchTodo } from "@/apis/todos.api";
import { StoreTodoInput, StoreTodoValidationError, Todo } from "@/apis/types";
import { TodoStoreForm } from "@/app/todos/_components/TodoStoreForm";
import { INITIAL_VALIDATION_ERRORS, useStoreTodo } from "@/app/todos/_hooks/useStoreTodo";
import { useMutation } from "@tanstack/react-query";
import { useRouter } from "next/navigation";
import { FC } from "react";

type Props = {
  todo: Todo;
};

export const TodosEditForm: FC<Props> = ({ todo }: Props) => {
  const doUpdateTodoInput: StoreTodoInput = {
    title: todo.title,
    content: todo.content,
  };
  const { register, handleSubmit, validationErrors, setValidationErrors } =
    useStoreTodo(doUpdateTodoInput);

  const router = useRouter();

  const mutation = useMutation({
    onMutate: () => setValidationErrors(INITIAL_VALIDATION_ERRORS),
    mutationFn: async (data: StoreTodoInput): Promise<StoreTodoValidationError> =>
      await patchTodo(String(todo.id), data),
    onSuccess: (data) => {
      if (data === undefined) return;

      if (Object.keys(data).length > 0) {
        setValidationErrors(data);
        return;
      }

      window.alert("TODOの更新に成功しました!");
      router.push("/");
    },
    onError: () => window.alert("予期しないエラーが発生しました."),
  });

  const onSubmit = handleSubmit((data) => mutation.mutate(data));

  return (
    <>
      {mutation.isPending && <p>Processing...</p>}

      <TodoStoreForm
        header='TODO編集'
        register={register}
        onSubmit={onSubmit}
        validationErrors={validationErrors}
      />
    </>
  );
};
