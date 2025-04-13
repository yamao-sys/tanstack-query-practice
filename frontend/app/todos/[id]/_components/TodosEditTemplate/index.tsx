import { getTodo } from "@/apis/todos.api";
import { FC } from "react";
import { TodosEditForm } from "../TodosEditForm";

type Props = {
  todoId: string;
};

export const TodosEditTemplate: FC<Props> = async ({ todoId }: Props) => {
  const todo = await getTodo(todoId);

  return (
    <>
      <TodosEditForm todo={todo} />
    </>
  );
};
