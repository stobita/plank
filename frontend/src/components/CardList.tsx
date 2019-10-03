import React from "react";
import styled from "styled-components";
import { CardListItemBox } from "./CardListItemBox";
import { CardPanel } from "./CardPanel";
import { DraggableCard } from "./MoveContextProvider";

interface Props {
  items: DraggableCard[];
}

export const CardList = (props: Props) => {
  return (
    <Wrapper>
      {props.items.map(card => (
        <CardListItemBox item={card} key={card.id}>
          <CardPanel card={card}></CardPanel>
        </CardListItemBox>
      ))}
    </Wrapper>
  );
};

const Wrapper = styled.div`
  display: flex;
  flex-direction: column;
`;
