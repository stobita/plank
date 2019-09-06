import React from "react";
import styled from "styled-components";
import { Header } from "./Header";

export const Home = () => {
  return (
    <Wrapper>
      <Header />
      <Main>main</Main>
    </Wrapper>
  );
};

const Wrapper = styled.div`
  display: flex;
  flex-direction: column;
  min-height: 100vh;
  color: ${props => props.theme.text};
  background: ${props => props.theme.bg};
`;

const Main = styled.div`
  display: flex;
  justify-content: center;
  padding: 16px;
`;
