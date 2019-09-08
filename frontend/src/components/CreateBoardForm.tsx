import React, { useContext } from "react";
import styled from "styled-components";
import boardsRepository, { CreateBoardPayload } from "../api/boardsRepository";
import { ButtonPair } from "./ButtonPair";
import { Button } from "./Button";
import { CloseButton } from "./CloseButton";
import { DataContext } from "../context/dataContext";
import { useForm } from "../hooks/useForm";
import { Input } from "./Input";

interface Props {
  onClickClose: () => void;
}

export const CreateBoardForm = (props: Props) => {
  const { setBoards } = useContext(DataContext);
  const createBoard = async () => {
    await boardsRepository.createBoard(formValue);
    const current = await boardsRepository.getBoards();
    setBoards(current);
  };
  const { formValue, handleOnChangeInput, handleOnSubmit } = useForm<
    CreateBoardPayload
  >(
    {
      name: ""
    },
    createBoard
  );
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
  background: ${props => props.theme.main};
  display: flex;
  flex-direction: column;
  position: absolute;
  width: 100%;
  height: 100%;
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
