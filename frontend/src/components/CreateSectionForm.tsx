import React, { useContext } from "react";
import { Input } from "./Input";
import { Button } from "./Button";
import styled from "styled-components";
import { useForm } from "../hooks/useForm";
import boardsRepository, {
  CreateSectionPayload
} from "../api/boardsRepository";
import { ViewContext } from "../context/viewContext";
import { DataContext } from "../context/dataContext";

export const CreateSectionForm = () => {
  const { currentBoardId } = useContext(ViewContext);
  const { setSections } = useContext(DataContext);
  const createSection = async () => {
    await boardsRepository.createSection(currentBoardId, formValue);
    const current = await boardsRepository.getBoardSections(currentBoardId);
    setSections(current);
  };
  const { formValue, handleOnChangeInput, handleOnSubmit } = useForm<
    CreateSectionPayload
  >({ name: "" }, createSection);
  return (
    <Wrapper>
      <Form onSubmit={handleOnSubmit}>
        <Field>
          <NameInput name="name" type="text" onChange={handleOnChangeInput} />
        </Field>
        <Field>
          <Button>Create</Button>
        </Field>
      </Form>
    </Wrapper>
  );
};

const Wrapper = styled.div`
  padding-left: 8px;
`;

const Form = styled.form`
  display: flex;
`;

const NameInput = styled(Input)`
  width: 280px;
`;

const Field = styled.div`
  padding-left: 8px;
`;
