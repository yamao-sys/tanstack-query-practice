import { Suspense } from "react";
import { TodosListTemplate } from "./_components/TodosListTemplate";

export default function TodosPage() {
  return (
    <>
      <Suspense fallback={<>loading...</>}>
        <TodosListTemplate />
      </Suspense>
    </>
  );
}
