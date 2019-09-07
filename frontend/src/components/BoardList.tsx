import React, { useEffect, useContext } from "react";
import styled from "styled-components";
import boardsRepository from "../api/boardsRepository";
import { AddButton } from "./AddButton";
import { DataContext } from "../context/dataContext";
import { ViewContext } from "../context/viewContext";

export const BoardList = () => {
  const { boards, setBoards } = useContext(DataContext);
  const { setCreateBoardActive } = useContext(ViewContext);

  useEffect(() => {
    boardsRepository.getBoards().then(items => {
      setBoards(items);
    });
  }, [setBoards]);

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
      <AddButton onClick={handleOnClickAddButton} />
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
