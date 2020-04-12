import React, { useEffect, useContext } from "react";
import styled from "styled-components";
import { useDrop } from "react-dnd";
import { ItemTypes } from "../constants/dnd";
import { MoveContext } from "../context/moveContext";
import { Card } from "../model/model";
import { DraggableCard } from "./MoveContextProvider";

interface Props {
  // https://github.com/react-dnd/react-dnd/issues/1435
  active: boolean;
  onDrop: (card: Card) => void;
}

export const CardDropTarget = (props: Props) => {
  const { overCard, overedCard } = useContext(MoveContext);
  const [{ isOver }, drop] = useDrop({
    accept: ItemTypes.CARD,
    drop: (item: DraggableCard) => {
      props.onDrop(item);
    },
    collect: (monitor) => ({
      isOver: !!monitor.isOver(),
    }),
  });
  useEffect(() => {
    if (isOver) {
      // TODO: fix without setTimeout
      // const timer = setTimeout(() => {
      overCard({ kind: "target" });
      // }, 50);
      // return () => clearTimeout(timer);
    }
  }, [isOver]);
  if (!props.active) {
    return null;
  }
  return <Wrapper ref={drop}>&nbsp;</Wrapper>;
};

const Wrapper = styled.div`
  background: ${(props) => props.theme.bg};
  border-radius: 4px;
  width: 100%;
  padding: 8px;
  box-sizing: border-box;
  height: 48px;
  margin-top: 8px;
`;
