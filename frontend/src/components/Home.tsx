import React, { useContext, useEffect } from "react";
import styled from "styled-components";
import { Header } from "./Header";
import { Sidebar } from "./Sidebar";
import { CreateBoardForm } from "./CreateBoardForm";
import { ViewContext } from "../context/viewContext";
import { DataContext } from "../context/dataContext";
import { BoardView } from "./BoardView";
import { EventContext } from "../context/eventContext";
import { BoardSetting } from "./BoardSetting";
import { HomeInterrupt } from "./HomeInterrupt";

export const Home = () => {
  const {
    createBoardActive,
    setCreateBoardActive,
    setBoardSettingActive,
    currentBoard,
    setCurrentBoard,
    boardSettingActive,
  } = useContext(ViewContext);
  const { updatedBoardIds, setUpdatedBoardIds } = useContext(EventContext);
  const { boards } = useContext(DataContext);

  useEffect(() => {
    if (updatedBoardIds.find((v) => v === currentBoard.id)) {
      setUpdatedBoardIds((prev) => prev.filter((v) => v !== currentBoard.id));
    }
  }, [currentBoard.id, setUpdatedBoardIds, updatedBoardIds]);

  useEffect(() => {
    if (currentBoard.id === 0 && boards.length > 0) {
      setCurrentBoard(boards[0]);
    }
  });

  const handleOnClickInterruptClose = () => {
    setCreateBoardActive(false);
    setBoardSettingActive(false);
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
          <HomeInterrupt onClickClose={handleOnClickInterruptClose}>
            <CreateBoardForm onClickCancel={handleOnClickInterruptClose} />
          </HomeInterrupt>
        )}
        {boardSettingActive && (
          <HomeInterrupt onClickClose={handleOnClickInterruptClose}>
            <BoardSetting></BoardSetting>
          </HomeInterrupt>
        )}
      </Body>
    </Wrapper>
  );
};

const Wrapper = styled.div`
  display: flex;
  color: ${(props) => props.theme.solid};
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
  background: ${(props) => props.theme.bg};
  margin-top: 56px;
  overflow: auto;
`;
