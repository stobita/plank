import { useState } from "react";

export const useForm = <P>(initFormValue: P, submitAction: () => void) => {
  const [formValue, setFormValue] = useState<P>(initFormValue);
  const handleOnChangeInput = (
    e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
  ) => {
    e.persist();
    setFormValue(prev => ({ ...prev, [e.target.name]: e.target.value }));
  };
  const handleOnSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    submitAction();
  };
  return { formValue, handleOnChangeInput, handleOnSubmit };
};
