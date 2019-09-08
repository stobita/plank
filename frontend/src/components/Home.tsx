import React, { useContext, useEffect } from "react";
import styled from "styled-components";
import { Header } from "./Header";
import { Sidebar } from "./Sidebar";
import { CreateBoardForm } from "./CreateBoardForm";
import { ViewContext } from "../context/viewContext";
import boardsRepository from "../api/boardsRepository";
import { DataContext } from "../context/dataContext";
import { BoardView } from "./BoardView";

export const Home = () => {
  const {
    createBoardActive,
    setCreateBoardActive,
    currentBoardId,
    setCurrentBoardId
  } = useContext(ViewContext);
  const { setSections, boards } = useContext(DataContext);

  useEffect(() => {
    if (currentBoardId !== 0) {
      boardsRepository.getBoardSections(currentBoardId).then(items => {
        setSections(items);
      });
    }
  }, [currentBoardId, setSections]);

  useEffect(() => {
    if (currentBoardId === 0 && boards.length > 0) {
      setCurrentBoardId(boards[0].id);
    }
  });

  const handleOnClickModalClose = () => {
    setCreateBoardActive(false);
  };

  return (
    <Wrapper>
      <Side>
        <Sidebar />
      </Side>
      <Body>
        <Header />
        <Main>
          <BoardView />
        </Main>
        {createBoardActive && (
          <CreateBoardForm onClickClose={handleOnClickModalClose} />
        )}
      </Body>
    </Wrapper>
  );
};

const Wrapper = styled.div`
  display: flex;
  min-height: 100vh;
  color: ${props => props.theme.solid};
  background: ${props => props.theme.bg};
`;

const Side = styled.div`
  flex: 1;
  display: flex;
`;

const Body = styled.div`
  flex: 5;
  display: flex;
  flex-direction: column;
  position: relative;
`;

const Main = styled.div`
  display: flex;
  justify-content: center;
  padding: 16px;
  flex: 1;
`;
