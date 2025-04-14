"use client";

import { FC, useState } from "react";
import { useForm } from "react-hook-form";
import { postSignUp } from "../../_actions/signUp";
import { AuthSignUpInput, AuthSignUpValidationError } from "@/apis/types";

const INITIAL_VALIDATION_ERRORS = {
  firstName: [],
  lastName: [],
  email: [],
  password: [],
  birthday: [],
  frontIdentification: [],
  backIdentification: [],
};

export const SignUpForm: FC = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<AuthSignUpInput>();

  const [validationErrors, setValidationErrors] =
    useState<AuthSignUpValidationError>(INITIAL_VALIDATION_ERRORS);

  const onSubmit = handleSubmit(async (data) => {
    setValidationErrors(INITIAL_VALIDATION_ERRORS);

    const response = await postSignUp(data);

    // バリデーションエラーがなければ、確認画面へ遷移
    if (Object.keys(response.errors).length === 0) {
      window.alert("会員登録に成功しました!");
      return;
    }

    // NOTE: バリデーションエラーの格納と入力パスワードのリセット
    setValidationErrors(response.errors);
  });

  return (
    <>
      <form onSubmit={onSubmit}>
        <div className='p-4 md:p-16'>
          <div className='w-4/5 md:w-3/5 mx-auto'>
            <h3 className='text-center text-2xl font-bold'>会員登録</h3>

            <div className='flex justify-between'>
              <div className='mt-8' style={{ width: "45%" }}>
                <label
                  htmlFor='last-name'
                  className='block mb-2 text-sm font-medium text-gray-900 dark:text-white text-left'
                >
                  <span className='font-bold'>姓</span>
                </label>
                <input
                  type='text'
                  className='bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500'
                  {...register("lastName", { required: true })}
                />
                {errors.lastName && <span>姓は必須項目です。</span>}
                {validationErrors.lastName && (
                  <div className='w-full pt-5 text-left'>
                    {validationErrors.lastName.map((message, i) => (
                      <p key={i} className='text-red-400'>
                        {message}
                      </p>
                    ))}
                  </div>
                )}
              </div>
              <div className='mt-8' style={{ width: "45%" }}>
                <label
                  htmlFor='first-name'
                  className='block mb-2 text-sm font-medium text-gray-900 dark:text-white text-left'
                >
                  <span className='font-bold'>名</span>
                </label>
                <input
                  type='text'
                  className='bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500'
                  {...register("firstName", { required: true })}
                />
                {errors.firstName && <span>名は必須項目です。</span>}
                {validationErrors.lastName && (
                  <div className='w-full pt-5 text-left'>
                    {validationErrors.lastName.map((message, i) => (
                      <p key={i} className='text-red-400'>
                        {message}
                      </p>
                    ))}
                  </div>
                )}
              </div>
            </div>

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
              {validationErrors.email && (
                <div className='w-full pt-5 text-left'>
                  {validationErrors.email.map((message, i) => (
                    <p key={i} className='text-red-400'>
                      {message}
                    </p>
                  ))}
                </div>
              )}
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
              {validationErrors.password && (
                <div className='w-full pt-5 text-left'>
                  {validationErrors.password.map((message, i) => (
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
