import { Suspense } from "react";
import { TodosEditTemplate } from "./_components/TodosEditTemplate";

type TodoEditPageProps = {
  params: {
    id: string;
  };
};

export default async function TodosEditPage({ params }: TodoEditPageProps) {
  const { id } = await params;

  return (
    <>
      <Suspense fallback={<>loading...</>}>
        <TodosEditTemplate todoId={id} />
      </Suspense>
    </>
  );
}
