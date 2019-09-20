import React, { useContext } from "react";
import styled from "styled-components";
import boardsRepository, { CreateBoardPayload } from "../api/boardsRepository";
import { ButtonPair } from "./ButtonPair";
import { Button } from "./Button";
import { CloseButton } from "./CloseButton";
import { DataContext } from "../context/dataContext";
import { useForm } from "../hooks/useForm";
import { Input } from "./Input";
import { ViewContext } from "../context/viewContext";

interface Props {
  onClickCancel: () => void;
}

export const CreateBoardForm = (props: Props) => {
  const { setBoards } = useContext(DataContext);
  const { setCreateBoardActive, setCurrentBoard } = useContext(ViewContext);
  const createBoard = async () => {
    const res = await boardsRepository.createBoard(formValue);
    const current = await boardsRepository.getBoards();
    initializeFormValue();
    setCreateBoardActive(false);
    setBoards(current);
    setCurrentBoard(res);
  };
  const {
    formValue,
    handleOnChangeInput,
    handleOnSubmit,
    initializeFormValue
  } = useForm<CreateBoardPayload>(
    {
      name: ""
    },
    createBoard
  );
  return (
    <Wrapper>
      <Title>Create Board</Title>
      <form onSubmit={handleOnSubmit}>
        <Field>
          <Input
            name="name"
            type="text"
            placeholder="name"
            onChange={handleOnChangeInput}
            value={formValue.name}
          />
        </Field>
        <Field>
          <ButtonPair
            left={
              <Button type="submit" primary>
                Create
              </Button>
            }
            right={<Button onClick={props.onClickCancel}>Cancel</Button>}
          />
        </Field>
      </form>
    </Wrapper>
  );
};

const Wrapper = styled.div`
  flex: 1;
  display: flex;
  justify-content: center;
  flex-direction: column;
  margin-bottom: 128px;
`;

const Title = styled.h2`
  text-align: center;
`;

const Field = styled.div`
  display: flex;
  margin-bottom: 16px;
`;
