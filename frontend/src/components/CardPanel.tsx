import React, { useState, useEffect, useContext } from "react";
import styled from "styled-components";
import { CardPanelDetail } from "./CardPanelDetail";
import { useDrag } from "react-dnd";
import { DraggableCard } from "./MoveContextProvider";
import { MoveContext } from "../context/moveContext";
import { getEmptyImage } from "react-dnd-html5-backend";

interface Props {
  card: DraggableCard;
}
export const CardPanel = (props: Props) => {
  const { setDraggingCard, setIsDraggingCard } = useContext(MoveContext);
  const [expand, setExpand] = useState(false);
  const handleOnClick = () => {
    setExpand(prev => !prev);
  };
  const [{ isDragging }, drag, preview] = useDrag({
    item: props.card,
    previewOptions: {
      anchorX: 0
    },

    collect: monitor => ({
      isDragging: !!monitor.isDragging()
    })
  });
  const { card } = props;
  useEffect(() => {
    setIsDraggingCard(isDragging);
    setDraggingCard(props.card);
  }, [isDragging, props.card, setIsDraggingCard, setDraggingCard]);
  useEffect(() => {
    preview(getEmptyImage(), { captureDraggingState: false });
  }, []);
  return (
    <>
      {!isDragging && (
        <Wrapper ref={drag}>
          <Inner onClick={handleOnClick}>
            <Name>{card.name}</Name>
            <Detail expand={expand}>
              <CardPanelDetail item={card}></CardPanelDetail>
            </Detail>
          </Inner>
        </Wrapper>
      )}
    </>
  );
};

const Wrapper = styled.div`
  flex: 1;
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
