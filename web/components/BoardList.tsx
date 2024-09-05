import React, { useEffect, useContext } from "react";
import styled from "styled-components";
import boardsRepository from "../api/boardsRepository";
import { AddButton } from "./AddButton";
import { DataContext } from "../context/dataContext";
import { ViewContext } from "../context/viewContext";
import { BoardListItem } from "./BoardListItem";

export const BoardList = () => {
  const { boards, setBoards } = useContext(DataContext);
  const { createBoardActive, setCreateBoardActive } = useContext(ViewContext);

  useEffect(() => {
    boardsRepository.getBoards().then(items => {
      setBoards(items);
    });
  }, [setBoards]);

  const handleOnClickAddButton = () => {
    setCreateBoardActive(prev => !prev);
  };

  return (
    <Wrapper>
      <List>
        {boards.map(v => (
          <BoardListItem key={v.id} item={v}></BoardListItem>
        ))}
      </List>
      <AddButton onClick={handleOnClickAddButton} isClose={createBoardActive} />
    </Wrapper>
  );
};

const Wrapper = styled.div``;

const List = styled.ul`
  list-style: none;
  padding: 0;
  margin: 0;
  margin-bottom: 32px;
`;
