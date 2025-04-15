import { getTodos } from "@/apis/todos.api";
import { Todo } from "@/apis/types";
import { useQuery } from "@tanstack/react-query";

const todoKeys = {
  all: ["todos"] as const,
  lists: () => [...todoKeys.all, "list"] as const,
  list: (filters: string) => [...todoKeys.lists(), { filters }] as const,
  details: () => [...todoKeys.all, "detail"] as const,
  detail: (id: number) => [...todoKeys.details(), id] as const,
};

export const useTodoLists = () => {
  return useQuery({
    queryKey: todoKeys.all,
    queryFn: async (): Promise<Todo[]> => await getTodos(),
  });
};
