import React, { useEffect, useContext } from "react";
import styled from "styled-components";
import boardsRepository from "../api/boardsRepository";
import { AddButton } from "./AddButton";
import { DataContext } from "../context/dataContext";
import { ViewContext } from "../context/viewContext";

export const BoardList = () => {
  const { boards, setBoards } = useContext(DataContext);
  const {
    createBoardActive,
    setCreateBoardActive,
    setCurrentBoardId,
    currentBoardId
  } = useContext(ViewContext);

  useEffect(() => {
    boardsRepository.getBoards().then(items => {
      setBoards(items);
    });
  }, [setBoards]);

  const handleOnClickAddButton = () => {
    setCreateBoardActive(prev => !prev);
  };

  const handleOnClickItem = (id: number) => {
    setCurrentBoardId(id);
  };

  return (
    <Wrapper>
      <List>
        {boards.map(v => (
          <Item
            key={v.id}
            onClick={() => handleOnClickItem(v.id)}
            selected={currentBoardId === v.id}
          >
            {v.name}
          </Item>
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

const Item = styled.li<{ selected: boolean }>`
  font-weight: bold;
  color: ${props => (props.selected ? props.theme.solid : props.theme.weak)};
  display: flex;
  margin-bottom: 12px;
  font-size: 1.1rem;
  cursor: pointer;
`;
