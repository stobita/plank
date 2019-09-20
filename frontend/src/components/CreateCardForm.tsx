import React, { useContext } from "react";
import styled from "styled-components";
import { Input } from "./Input";
import { Button } from "./Button";
import boardsRepository from "../api/boardsRepository";
import { useForm } from "../hooks/useForm";
import sectionsRepository, {
  CreateCardPayload
} from "../api/sectionsRepository";
import { Section } from "../model/model";
import { DataContext } from "../context/dataContext";
import { ViewContext } from "../context/viewContext";
import { Textarea } from "./Textarea";

interface Props {
  section: Section;
  afterSubmit: () => void;
}

export const CreateCardForm = (props: Props) => {
  const { currentBoard } = useContext(ViewContext);
  const { setSections } = useContext(DataContext);
  const createCard = async () => {
    await sectionsRepository.createCard(props.section.id, formValue);
    const current = await boardsRepository.getBoardSections(currentBoard.id);
    setSections(current);
    initializeFormValue();
    props.afterSubmit();
  };
  const {
    formValue,
    handleOnChangeInput,
    handleOnSubmit,
    initializeFormValue
  } = useForm<CreateCardPayload>({ name: "", description: "" }, createCard);
  return (
    <Wrapper>
      <Form onSubmit={handleOnSubmit}>
        <Field>
          <Input
            onChange={handleOnChangeInput}
            name="name"
            value={formValue.name}
            type="text"
            placeholder="name"
          ></Input>
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
          <Button primary>Add</Button>
        </Field>
      </Form>
    </Wrapper>
  );
};

const Wrapper = styled.div``;

const Form = styled.form``;

const Field = styled.div`
  margin-top: 8px;
`;
