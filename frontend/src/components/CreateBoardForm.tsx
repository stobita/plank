import React, { useCallback, useState, useContext } from "react";
import styled from "styled-components";
import colors from "../colors";
import boardsRepository, { CreateBoardPayload } from "../api/boardsRepository";
import { ButtonPair } from "./ButtonPair";
import { Button } from "./Button";
import { CloseButton } from "./CloseButton";
import { DataContext } from "../context/dataContext";

interface Props {
  onClickClose: () => void;
}

export const CreateBoardForm = (props: Props) => {
  const { setBoards } = useContext(DataContext);
  const [formValue, setFormValue] = useState<CreateBoardPayload>({ title: "" });
  const handleOnChangeInput = useCallback(
    (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
      e.persist();
      setFormValue(prev => ({ ...prev, [e.target.name]: e.target.value }));
    },
    []
  );
  const handleOnSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    createBoard();
  };
  const createBoard = async () => {
    await boardsRepository.createBoard(formValue);
    const current = await boardsRepository.getBoards();
    setBoards(current);
  };
  return (
    <Wrapper>
      <Head>
        <CloseButton onClick={props.onClickClose} />
      </Head>
      <Body>
        <Inner>
          <Title>Create Board</Title>
          <form onSubmit={handleOnSubmit}>
            <Field>
              <Input
                name="title"
                type="text"
                placeholder="title"
                onChange={handleOnChangeInput}
                value={formValue.title}
              ></Input>
            </Field>
            <Field>
              <ButtonPair
                left={
                  <Button type="submit" primary>
                    Create
                  </Button>
                }
                right={<Button onClick={props.onClickClose}>Cancel</Button>}
              />
            </Field>
          </form>
        </Inner>
      </Body>
    </Wrapper>
  );
};

const Wrapper = styled.div`
  flex: 1;
  background: ${props => props.theme.bg};
  display: flex;
  flex-direction: column;
`;

const Head = styled.div`
  display: flex;
  justify-content: flex-end;
  padding: 16px;
`;

const Body = styled.div`
  display: flex;
  justify-content: center;
`;

const Inner = styled.div`
  width: 40%;
  padding-top: 128px;
`;

const Title = styled.h2`
  text-align: center;
`;

const Field = styled.div`
  display: flex;
  margin-bottom: 16px;
`;

const Input = styled.input`
  border: 1px solid ${props => props.theme.border};
  background: ${colors.white};
  font-size: 16px;
  width: 100%;
  padding: 8px;
  box-sizing: border-box;
  border-radius: 4px;
  height: 36px;
`;
