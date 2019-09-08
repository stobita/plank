import React from "react";
import { Card } from "../model/model";
import styled from "styled-components";

interface Props {
  card: Card;
}
export const CardPanel = (props: Props) => {
  return (
    <Wrapper>
      <Inner>
        <Name>{props.card.name}</Name>
      </Inner>
    </Wrapper>
  );
};

const Wrapper = styled.div`
  flex: 1;
  padding-top: 8px;
`;

const Inner = styled.div`
  background: ${props => props.theme.main};
  border: 1px solid ${props => props.theme.border};
  border-radius: 4px;
  padding: 8px;
`;

const Name = styled.h4`
  margin: 0;
`;
