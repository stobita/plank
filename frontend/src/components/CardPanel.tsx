import React, { useState } from "react";
import { Card } from "../model/model";
import styled from "styled-components";
import { CardPanelDetail } from "./CardPanelDetail";

interface Props {
  card: Card;
}
export const CardPanel = (props: Props) => {
  const [expand, setExpand] = useState(false);
  const handleOnClick = () => {
    setExpand(prev => !prev);
  };
  const { card } = props;
  return (
    <Wrapper>
      <Inner onClick={handleOnClick}>
        <Name>{card.name}</Name>
        <Detail expand={expand}>
          <CardPanelDetail item={card}></CardPanelDetail>
        </Detail>
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
  cursor: pointer;
`;

const Name = styled.h4`
  margin: 0;
`;

const Detail = styled.div<{ expand: boolean }>`
  transition: 1s;
  max-height: ${props => (props.expand ? 240 : 0)}px;
  overflow: hidden;
`;
