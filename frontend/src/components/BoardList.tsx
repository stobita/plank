import React, { useState, useEffect, useContext } from "react";
import styled from "styled-components";
import boardsRepository from "../api/boardsRepository";
import { Board } from "../model/model";
import { CreateBoardForm } from "./CreateBoardForm";
import { ModalContext, DataContext } from "../Provider";

export const BoardList = () => {
  const { boards, setBoards } = useContext(DataContext);
  const { setCreateBoardActive } = useContext(ModalContext);

  useEffect(() => {
    boardsRepository.getBoards().then(items => {
      setBoards(items);
    });
  }, []);

  const handleOnClickAddButton = () => {
    setCreateBoardActive(true);
  };
  return (
    <Wrapper>
      <List>
        {boards.map(v => (
          <Item key={v.id}>{v.title}</Item>
        ))}
      </List>
      <Button onClick={handleOnClickAddButton}>
        <ButtonInner />
        <ButtonInner />
      </Button>
    </Wrapper>
  );
};

const Wrapper = styled.div``;

const List = styled.ul`
  list-style: none;
  padding: 0;
  margin: 0;
  margin-bottom: 16px;
`;

const Item = styled.li`
  display: flex;
  margin-bottom: 8px;
  font-size: 1.1rem;
`;

const Button = styled.button`
  border: 1px solid ${props => props.theme.border};
  background: ${props => props.theme.bg};
  height: 32px;
  width: 32px;
  border-radius: 16px;
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
    transform: rotate(90deg);
  }
`;
