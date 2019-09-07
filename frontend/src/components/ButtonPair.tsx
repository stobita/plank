import React from "react";
import styled from "styled-components";

interface Props {
  left: React.ReactNode;
  right: React.ReactNode;
}

export const ButtonPair = (props: Props) => {
  return (
    <Wrapper>
      <Column>{props.left}</Column>
      <Column>{props.right}</Column>
    </Wrapper>
  );
};

const Wrapper = styled.div`
  display: flex;
  width: 100%;
`;

const Column = styled.div`
  width: 100%;
  &:nth-child(1) {
    margin-right: 8px;
  }
  &:nth-last-child(1) {
    margin-left: 8px;
  }
`;
