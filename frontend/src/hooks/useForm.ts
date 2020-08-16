import { useState, useCallback } from 'react';

export const useForm = <P>(initFormValue: P, submitAction: () => void) => {
  const [formValue, setFormValue] = useState<P>(initFormValue);
  const handleOnChangeInput = (
    e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>,
  ) => {
    e.persist();
    setFormValue((prev) => ({ ...prev, [e.target.name]: e.target.value }));
  };
  const handleOnSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    submitAction();
  };

  const onChangeLabel = useCallback((items: string[]) => {
    setFormValue((prev) => ({ ...prev, labels: items }));
  }, []);

  const onChangeLimitDate = (limitTime: number) => {
    setFormValue((prev) => ({ ...prev, limitTime: limitTime }));
  };

  const onChangeImage = (image: string) => {
    setFormValue((prev) => ({ ...prev, image }));
  };

  const initializeFormValue = () => {
    setFormValue(initFormValue);
  };
  return {
    formValue,
    handleOnChangeInput,
    handleOnSubmit,
    initializeFormValue,
    onChangeLabel,
    onChangeLimitDate,
    onChangeImage,
  };
};
