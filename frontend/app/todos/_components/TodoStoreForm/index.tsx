import { StoreTodoInput, StoreTodoValidationError } from "@/apis/types/todos";
import { FC } from "react";
import { UseFormRegister } from "react-hook-form";

type Props = {
  header: string;
  register: UseFormRegister<StoreTodoInput>;
  onSubmit: (e?: React.BaseSyntheticEvent) => Promise<void>;
  validationErrors: StoreTodoValidationError;
};

export const TodoStoreForm: FC<Props> = ({
  header,
  register,
  onSubmit,
  validationErrors,
}: Props) => {
  return (
    <>
      <form onSubmit={onSubmit}>
        <div className='p-4 md:p-16'>
          <div className='w-4/5 md:w-3/5 mx-auto'>
            <h3 className='text-center text-2xl font-bold'>{header}</h3>

            <div className='mt-8'>
              <label
                htmlFor='title'
                className='block mb-2 text-sm font-medium text-gray-900 dark:text-white text-left'
              >
                <span className='font-bold'>タイトル</span>
              </label>
              <input
                type='text'
                className='bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500'
                {...register("title")}
              />
              {validationErrors.title && (
                <div className='w-full pt-5 text-left'>
                  {validationErrors.title.map((message, i) => (
                    <p key={i} className='text-red-400'>
                      {message}
                    </p>
                  ))}
                </div>
              )}
            </div>

            <div className='mt-8'>
              <label
                htmlFor='content'
                className='block mb-2 text-sm font-medium text-gray-900 dark:text-white text-left'
              >
                <span className='font-bold'>内容</span>
              </label>
              <textarea
                id='content'
                rows={16}
                className='block p-2.5 w-full text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500 resize-none'
                {...register("content")}
              />
              {validationErrors.content && (
                <div className='w-full pt-5 text-left'>
                  {validationErrors.content.map((message, i) => (
                    <p key={i} className='text-red-400'>
                      {message}
                    </p>
                  ))}
                </div>
              )}
            </div>

            <div className='w-full flex justify-center'>
              <div className='mt-16'>
                <button
                  type='submit'
                  className='py-2 px-8 border-green-500 bg-green-500 rounded-xl text-white'
                >
                  登録
                </button>
              </div>
            </div>
          </div>
        </div>
      </form>
    </>
  );
};
