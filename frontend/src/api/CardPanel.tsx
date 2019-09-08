import React, { useState } from "react";
import { Card } from "../model/model";
import styled from "styled-components";

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
          <Description>
            {card.description !== "" ? card.description : "no description"}
          </Description>
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
  transition: 0.5s;
  max-height: ${props => (props.expand ? 64 : 0)}px;
  overflow: hidden;
`;

const Description = styled.p`
  padding-top: 8px;
  margin: 0;
`;
