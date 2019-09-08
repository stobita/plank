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
  const { createBoardActive, setCreateBoardActive } = useContext(ViewContext);
  const { currentBoardId } = useContext(ViewContext);
  const { setSections } = useContext(DataContext);

  useEffect(() => {
    boardsRepository.getBoardSections(currentBoardId).then(items => {
      setSections(items);
    });
  }, [currentBoardId, setSections]);

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
          {createBoardActive && (
            <CreateBoardForm onClickClose={handleOnClickModalClose} />
          )}
        </Main>
      </Body>
    </Wrapper>
  );
};

const Wrapper = styled.div`
  display: flex;
  min-height: 100vh;
  color: ${props => props.theme.text};
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
`;

const Main = styled.div`
  display: flex;
  justify-content: center;
  padding: 16px;
  position: relative;
  flex: 1;
`;
