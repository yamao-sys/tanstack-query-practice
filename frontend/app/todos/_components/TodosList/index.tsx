"use client";

import { Todo } from "@/apis/types";
import Link from "next/link";
import { FC, useState } from "react";

type Props = {
  todos: Todo[];
};

export const TodosList: FC<Props> = ({ todos }: Props) => {
  const [displayTodos, setDisplayTodos] = useState<Todo[]>(todos); // eslint-disable-line @typescript-eslint/no-unused-vars

  return (
    <>
      <div className='p-4 md:p-16'>
        <div className='w-4/5 md:w-3/5 mx-auto'>
          <h3 className='text-center text-2xl font-bold'>TODO一覧</h3>

          <div className='w-full mt-4 md:mt-16'>
            {displayTodos.length === 0 ? (
              <p>まだTODOが0件です</p>
            ) : (
              displayTodos.map((todo) => (
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
