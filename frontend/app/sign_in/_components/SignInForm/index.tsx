"use client";

import { FC, useState } from "react";
import { useForm } from "react-hook-form";
import { postSignIn } from "../../_actions/signIn";
import { useRouter } from "next/navigation";
import { AuthSignInInput } from "@/apis/types";

export const SignInForm: FC = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<AuthSignInInput>();

  const [validationError, setValidationError] = useState("");

  const router = useRouter();

  const onSubmit = handleSubmit(async (data) => {
    setValidationError("");

    const response = await postSignIn(data);

    // バリデーションエラーがなければ、確認画面へ遷移
    if (!response) {
      window.alert("ログインに成功しました!");
      router.push("/");
      return;
    }

    // NOTE: バリデーションエラーの格納と入力パスワードのリセット
    setValidationError(response);
  });

  return (
    <>
      <form onSubmit={onSubmit}>
        <div className='p-4 md:p-16'>
          <div className='w-4/5 md:w-3/5 mx-auto'>
            <h3 className='text-center text-2xl font-bold'>ログイン</h3>

            {validationError && (
              <div className='w-full pt-5 text-left'>
                <p className='text-red-400'>{validationError}</p>
              </div>
            )}

            <div className='mt-8'>
              <label
                htmlFor='email'
                className='block mb-2 text-sm font-medium text-gray-900 dark:text-white text-left'
              >
                <span className='font-bold'>メールアドレス</span>
              </label>
              <input
                type='email'
                className='bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500'
                {...register("email", { required: true })}
              />
              {errors.email && <span>メールアドレスは必須項目です。</span>}
            </div>

            <div className='mt-8'>
              <label
                htmlFor='email'
                className='block mb-2 text-sm font-medium text-gray-900 dark:text-white text-left'
              >
                <span className='font-bold'>パスワード</span>
              </label>
              <input
                type='password'
                className='bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500'
                {...register("password", { required: true })}
              />
              {errors.password && <span>パスワードは必須項目です。</span>}
            </div>

            <div className='w-full flex justify-center'>
              <div className='mt-16'>
                <button
                  type='submit'
                  className='py-2 px-8 border-green-500 bg-green-500 rounded-xl text-white'
                >
                  ログイン
                </button>
              </div>
            </div>
          </div>
        </div>
      </form>
    </>
  );
};
