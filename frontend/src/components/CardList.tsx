import React from "react";
import { Card } from "../model/model";
import styled from "styled-components";
import { CardPanel } from "./CardPanel";

interface Props {
  items: Card[];
}

export const CardList = (props: Props) => {
  return (
    <Wrapper>
      {props.items !== null &&
        props.items.map(card => <CardPanel card={card} key={card.id} />)}
    </Wrapper>
  );
};

const Wrapper = styled.div`
  display: flex;
  flex-direction: column;
`;
