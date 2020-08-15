import React, { useContext, useCallback } from 'react';
import styled from 'styled-components';
import { Button } from './Button';
import { Input } from './Input';
import { Textarea } from './Textarea';
import sectionsRepository, {
  UpdateCardPayload,
} from '../api/sectionsRepository';
import { useForm } from '../hooks/useForm';
import { Card } from '../model/model';
import boardsRepository from '../api/boardsRepository';
import { ViewContext } from '../context/viewContext';
import { ButtonPair } from './ButtonPair';
import { DataContext } from '../context/dataContext';
import { TagInput } from './TagInput';
import { DatetimePicker } from './DatetimePicker';
import { useDropzone } from 'react-dropzone';
import { ImageUploader } from './ImageUploader';

interface Props {
  item: Card;
  onClickCancel: () => void;
  afterSubmit: () => void;
}
export const EditCardForm = (props: Props) => {
  const { currentBoard } = useContext(ViewContext);
  const { setSections, labels } = useContext(DataContext);
  const { item } = props;
  const updateCard = async () => {
    await sectionsRepository.updateCard(item.section.id, item.id, formValue);
    const current = await boardsRepository.getBoardSections(currentBoard.id);
    setSections(current);
    props.afterSubmit();
  };
  const {
    formValue,
    handleOnChangeInput,
    handleOnSubmit,
    onChangeLabel,
    onChangeLimitDate,
    onChangeImage,
  } = useForm<UpdateCardPayload>(
    {
      name: item.name,
      description: item.description,
      labels: item.labels!.map((v) => v.name),
      limitTime: item.limitTime,
      image: item.image,
    },
    updateCard,
  );
  const onClickWrapper = (e: React.MouseEvent<HTMLElement>) => {
    e.stopPropagation();
  };
  const getFileAsDataURL = (file: Blob): Promise<string> => {
    return new Promise((resolve) => {
      const reader = new FileReader();
      reader.onloadend = () => {
        console.log(reader.result);
        if (typeof reader.result === 'string') {
          resolve(reader.result);
        }
      };
      reader.readAsDataURL(file);
    });
  };

  const onDrop = useCallback((acceptedFiles) => {
    acceptedFiles.forEach(async (file: File) => {
      const image = await getFileAsDataURL(file);
      onChangeImage(image);
    });
  }, []);
  const { getRootProps, getInputProps, isDragActive } = useDropzone({ onDrop });

  return (
    <Wrapper onClick={onClickWrapper}>
      <Form onSubmit={handleOnSubmit}>
        <Field>
          <Input
            name="name"
            value={formValue.name}
            type="text"
            placeholder="name"
            onChange={handleOnChangeInput}
          ></Input>
        </Field>
        <Field>
          <DatetimePicker
            value={formValue.limitTime}
            onChange={onChangeLimitDate}
            placeholder="limit date time"
          ></DatetimePicker>
        </Field>
        <Field>
          <TagInput
            name="label"
            placeholder="label"
            value={formValue.labels ? formValue.labels : []}
            onChange={onChangeLabel}
            labels={labels}
          />
        </Field>
        <Field>
          <Textarea
            name="description"
            value={formValue.description}
            placeholder="description"
            onChange={handleOnChangeInput}
          />
        </Field>
        <Field>
          <ImageUploader
            image={formValue.image ? formValue.image : ''}
            onChange={onChangeImage}
          ></ImageUploader>
        </Field>
        <Field>
          <ButtonPair
            left={<Button primary>Save</Button>}
            right={
              <Button type="button" onClick={props.onClickCancel}>
                Cancel
              </Button>
            }
          />
        </Field>
      </Form>
    </Wrapper>
  );
};

const Wrapper = styled.div`
  flex: 1;
`;
const Form = styled.form``;
const Field = styled.div`
  margin-top: 8px;
`;
