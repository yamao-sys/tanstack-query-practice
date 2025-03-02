"use client";

import { FC } from "react";
import { useForm } from "react-hook-form";

export const SignUpForm: FC = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm();
  const onSubmit = handleSubmit((data) => console.log(data));

  return (
    <>
      <form onSubmit={onSubmit}>
        <input type='email' placeholder='Email' {...register("email", { required: true })} />
        {errors.email && <span>Emailは必須項目です。</span>}

        <button type='submit'>確認画面へ</button>
      </form>
    </>
  );
};
