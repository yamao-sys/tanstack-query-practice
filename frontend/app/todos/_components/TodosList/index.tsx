"use client";

import { useTodoLists } from "@/feature/todos/hooks/useTodo";
import Link from "next/link";
import { FC } from "react";

export const TodosList: FC = () => {
  const { data: todos, isPending } = useTodoLists();

  if (isPending) return <p>fetching todos...</p>;
  if (!isPending && !todos) return <p>unexpected fetching</p>;

  return (
    <>
      <div className='p-4 md:p-16'>
        <div className='w-4/5 md:w-3/5 mx-auto'>
          <h3 className='text-center text-2xl font-bold'>TODO一覧</h3>

          <div className='w-full mt-4 md:mt-16'>
            {todos.length === 0 ? (
              <p>まだTODOが0件です</p>
            ) : (
              todos.map((todo) => (
                <div key={todo.id} className='max-w-3xl mx-auto space-y-4'>
                  <div className='flex justify-between items-start p-4 bg-white shadow border border-gray-400 hover:bg-gray-50 transition'>
                    <div className='flex gap-3'>
                      <div className='w-5 h-5 text-green-500'>
                        <input type='checkbox' />
                      </div>
                      <div>
                        <Link
                          className='text-blue-600 font-semibold hover:underline'
                          href={`/todos/${todo.id}`}
                          target='_blank'
                        >
                          {todo.title}
                        </Link>
                        <div className='flex flex-wrap text-sm text-gray-600 mt-1'>
                          <span>#{todo.id} opened 3 days ago</span>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              ))
            )}
          </div>
        </div>
      </div>
    </>
  );
};
