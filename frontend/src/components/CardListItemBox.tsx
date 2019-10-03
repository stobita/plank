import React, { useContext, ReactNode, useEffect } from "react";
import { useDrop } from "react-dnd";
import { ItemTypes } from "../constants/dnd";
import { MoveContext } from "../context/moveContext";
import { DraggableCard } from "./MoveContextProvider";
import styled from "styled-components";

interface Props {
  item: DraggableCard;
  children: ReactNode;
}

export const CardListItemBox = (props: Props) => {
  const { overCard, dropCard, isDraggingCard } = useContext(MoveContext);
  const [{ isOver }, drop] = useDrop({
    accept: ItemTypes.CARD,
    drop: () => {
      dropCard();
    },
    canDrop: () => props.item.isDropTarget,
    collect: monitor => ({
      isOver: !!monitor.isOver()
    })
  });

  useEffect(() => {
    if (isOver) {
      // TODO: fix without setTimeout
      const timer = setTimeout(() => {
        overCard(props.item);
      }, 50);
      return () => clearTimeout(timer);
    }
  }, [isOver, overCard, props.item]);

  return (
    <Wrapper ref={drop}>
      {!props.item.isDropTarget ? (
        <div>{props.children}</div>
      ) : isDraggingCard ? (
        <DropTarget>&nbsp;</DropTarget>
      ) : null}
    </Wrapper>
  );
};

const Wrapper = styled.div`
  margin-top: 8px;
`;

const DropTarget = styled.div`
  background: ${props => props.theme.bg};
  border-radius: 4px;
  width: 100%;
  padding: 8px;
  box-sizing: border-box;
  height: 48px;
`;
