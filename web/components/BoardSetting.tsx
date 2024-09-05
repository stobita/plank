import React, { useContext, useState } from "react";
import styled from "styled-components";
import { ViewContext } from "../context/viewContext";
import { Input } from "./Input";
import boardsRepository, { UpdateBoardPayload } from "../api/boardsRepository";
import { useForm } from "../hooks/useForm";
import { DataContext } from "../context/dataContext";
import { ButtonPair } from "./ButtonPair";
import { Button } from "./Button";

export const BoardSetting = () => {
  const { setBoards } = useContext(DataContext);
  const { currentBoard, setBoardSettingActive, setCurrentBoard } = useContext(
    ViewContext
  );
  const [deleteActive, setDeleteActive] = useState(false);
  const updateBoard = async () => {
    await boardsRepository.updateBoard(currentBoard.id, formValue);
    const current = await boardsRepository.getBoards();
    setBoards(current);
  };
  const { formValue, handleOnChangeInput, handleOnSubmit } = useForm<
    UpdateBoardPayload
  >(
    {
      name: currentBoard.name
    },
    updateBoard
  );
  const handleOnClickCancel = () => {
    setBoardSettingActive(false);
  };
  const handleOnClickDelete = () => {
    setDeleteActive(true);
  };
  const handleOnClickDeleteCancel = () => {
    setDeleteActive(false);
  };
  const handleOnSubmitDelete = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    await boardsRepository.deleteBoard(currentBoard.id);
    setCurrentBoard({ id: 0, name: "" });
    const current = await boardsRepository.getBoards();
    setBoards(current);
  };
  return (
    <Wrapper>
      <Confirm active={deleteActive}>
        {deleteActive && (
          <ConfirmInner>
            <DangerMessage>Are you sure you want to delete it?</DangerMessage>
            <form onSubmit={handleOnSubmitDelete}>
              <ButtonPair
                left={
                  <Button danger type="submit" primary>
                    Delete
                  </Button>
                }
                right={
                  <Button onClick={handleOnClickDeleteCancel}>Cancel</Button>
                }
              />
            </form>
          </ConfirmInner>
        )}
      </Confirm>
      <Title>{currentBoard.name} Setting</Title>
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
                Save
              </Button>
            }
            right={<Button onClick={handleOnClickCancel}>Cancel</Button>}
          />
        </Field>
      </form>
      <Danger>
        {!deleteActive && (
          <Button danger type="button" onClick={handleOnClickDelete}>
            Delete this board
          </Button>
        )}
      </Danger>
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

const Danger = styled.div`
  margin-top: 32px;
`;

const Confirm = styled.div<{ active: boolean }>`
  transition: 0.5s;
  max-height: ${props => (props.active ? 240 : 0)}px;
  box-sizing: border-box;
`;

const ConfirmInner = styled.div`
  border: 1px solid ${props => props.theme.danger};
  padding: 32px;
  border-radius: 4px;
`;

const DangerMessage = styled.p`
  color: ${props => props.theme.danger};
  margin: 0;
  margin-bottom: 32px;
  font-weight: bold;
  font-size: 1.2rem;
`;
