import React, { ReactNode } from "react";
import styled from "styled-components";
import { CloseButton } from "./CloseButton";

interface Props {
  children: ReactNode;
  onClickClose: () => void;
}
export const HomeInterrupt = (props: Props) => {
  return (
    <Wrapper>
      <Head>
        <CloseButton onClick={props.onClickClose} />
      </Head>
      <Body>
        <Content>{props.children}</Content>
      </Body>
    </Wrapper>
  );
};

const Wrapper = styled.div`
  z-index: 1;
  flex: 1;
  position: absolute;
  display: flex;
  flex-direction: column;
  background: ${props => props.theme.main};
  height: 100vh;
  width: calc(100vw - 240px);
`;
const Head = styled.div`
  position: absolute;
  padding: 16px;
  right: 0;
  top: 0;
`;
const Body = styled.div`
  flex: 1;
  display: flex;
  justify-content: center;
`;

const Content = styled.div`
  display: flex;
  align-items: center;
  width: 40%;
`;
