import React, { useContext } from "react";
import styled from "styled-components";
import { Header } from "./Header";
import { Sidebar } from "./Sidebar";
import { ModalContext } from "../Provider";
import { CreateBoardForm } from "./CreateBoardForm";

export const Home = () => {
  const { createBoardActive, setCreateBoardActive } = useContext(ModalContext);
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
