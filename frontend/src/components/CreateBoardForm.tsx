import React, { useCallback, useState, useContext } from "react";
import styled from "styled-components";
import colors from "../colors";
import boardsRepository, { CreateBoardPayload } from "../api/boardsRepository";
import { ButtonPair } from "./ButtonPair";
import { Button } from "./Button";
import { DataContext } from "../Provider";

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
        <CloseButton onClick={props.onClickClose}>
          <ButtonInner />
          <ButtonInner />
        </CloseButton>
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
  position: absolute;
  background: ${props => props.theme.bg};
  height: 100%;
  width: 100%;
  top: 0;
  left: 0;
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

const CloseButton = styled.button`
  border: 1px solid ${props => props.theme.border};
  height: 36px;
  width: 36px;
  border-radius: 18px;
  position: relative;
  display: flex;
  justify-content: center;
  align-items: center;
  cursor: pointer;
`;

const ButtonInner = styled.span`
  position: absolute;
  width: 50%;
  height: 2px;
  background: ${props => props.theme.text};
  border-radius: 4px;
  &:nth-of-type(1) {
    transform: rotate(45deg);
  }
  &:nth-of-type(2) {
    transform: rotate(135deg);
  }
`;
