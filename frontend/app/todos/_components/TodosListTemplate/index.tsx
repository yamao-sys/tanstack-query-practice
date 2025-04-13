import { getTodos } from "@/apis/todos.api";
import { FC } from "react";
import { TodosList } from "../TodosList";

export const TodosListTemplate: FC = async () => {
  const todos = await getTodos();

  return (
    <>
      <TodosList todos={todos} />
    </>
  );
};
