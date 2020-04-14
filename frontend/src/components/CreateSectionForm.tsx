import React, { useContext } from "react";
import { Input } from "./Input";
import { Button } from "./Button";
import styled from "styled-components";
import { useForm } from "../hooks/useForm";
import boardsRepository, {
  CreateSectionPayload,
} from "../api/boardsRepository";
import { ViewContext } from "../context/viewContext";
import { BoardContext } from "../context/boardContext";

interface Props {
  afterSubmit: () => void;
}

export const CreateSectionForm = (props: Props) => {
  const { currentBoard } = useContext(ViewContext);
  const { setSections } = useContext(BoardContext);
  const createSection = async () => {
    await boardsRepository.createSection(currentBoard.id, formValue);
    const current = await boardsRepository.getBoardSections(currentBoard.id);
    setSections(current);
    props.afterSubmit();
    initializeFormValue();
  };
  const {
    formValue,
    handleOnChangeInput,
    handleOnSubmit,
    initializeFormValue,
  } = useForm<CreateSectionPayload>({ name: "" }, createSection);
  return (
    <Wrapper>
      <Form onSubmit={handleOnSubmit}>
        <Field>
          <NameInput
            value={formValue.name}
            name="name"
            type="text"
            onChange={handleOnChangeInput}
          />
        </Field>
        <Field>
          <Button primary>Create</Button>
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
