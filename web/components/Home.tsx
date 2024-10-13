'use client'

import React, { useContext, useEffect } from 'react';
import styled from 'styled-components';
import { Header } from '../components/Header';
import { Sidebar } from '../components/Sidebar';
import { CreateBoardForm } from '../components/CreateBoardForm';
import { BoardView } from '../components/BoardView';
import { BoardSetting } from '../components/BoardSetting';
import { HomeInterrupt } from '../components/HomeInterrupt';
import { SectionFilter } from '../components/SectionFilter';
import { EventContext } from '@/context/eventContext';
import { DataContext } from '@/context/dataContext';
import { Board } from '@/model/model';
import { ViewContext } from '@/context/viewContext';

export const Home = () => {
  const {
    createBoardActive,
    setCreateBoardActive,
    setBoardSettingActive,
    currentBoard,
    setCurrentBoard,
    boardSettingActive,
    sectionFilterActive,
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
          {sectionFilterActive && <SectionFilter></SectionFilter>}
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
}


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
  width: calc(100vw - 240px);
  z-index: 1;
`;

const Main = styled.div`
  display: flex;
  justify-content: center;
  flex: 1;
  background: ${(props) => props.theme.bg};
  margin-top: 56px;
  overflow: auto;
`;
function setCurrentBoard(arg0: Board) {
  throw new Error('Function not implemented.');
}
