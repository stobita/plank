import React, { useContext, ReactNode, useEffect } from "react";
import { useDrop } from "react-dnd";
import { ItemTypes } from "../constants/dnd";
import { MoveContext } from "../context/moveContext";
import styled from "styled-components";
import { Card } from "../model/model";

interface Props {
  item: Card;
  children: ReactNode;
}

export const CardListItemBox = (props: Props) => {
  const { overCard } = useContext(MoveContext);
  const [{ isOver }, drop] = useDrop({
    accept: ItemTypes.CARD,
    canDrop: () => false,
    collect: monitor => ({
      isOver: !!monitor.isOver()
    })
  });

  useEffect(() => {
    if (isOver) {
      overCard(props.item);
    }
  }, [isOver]);

  return <Wrapper ref={drop}>{props.children}</Wrapper>;
};

const Wrapper = styled.div`
  margin-top: 8px;
`;
