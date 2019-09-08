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
        <Head>
          <Header />
        </Head>
        <Main>
          <BoardView />
        </Main>
        {createBoardActive && (
          <Interrupt>
            <CreateBoardForm onClickClose={handleOnClickModalClose} />
          </Interrupt>
        )}
      </Body>
    </Wrapper>
  );
};

const Wrapper = styled.div`
  display: flex;
  color: ${props => props.theme.solid};
  min-height: 100vh;
`;

const Side = styled.div`
  flex: 1;
  display: flex;
  position: fixed;
  z-index: 1;
  height: 100%;
  min-width: 240px;
`;

const Body = styled.div`
  flex: 5;
  display: flex;
  flex-direction: column;
  position: relative;
  margin-left: 240px;
`;

const Head = styled.div`
  position: fixed;
  width: 100%;
  z-index: 1;
`;

const Main = styled.div`
  display: flex;
  justify-content: center;
  padding: 16px;
  flex: 1;
  background: ${props => props.theme.bg};
  margin-top: 56px;
  overflow: auto;
`;

const Interrupt = styled.div`
  flex: 1;
  position: absolute;
  z-index: 1;
  height: 100vh;
  width: calc(100vw - 240px);
`;
