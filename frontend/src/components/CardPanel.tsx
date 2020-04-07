import React, { useState, useEffect } from "react";
import styled from "styled-components";
import { CardPanelDetail } from "./CardPanelDetail";
import { useDrag } from "react-dnd";
import { getEmptyImage } from "react-dnd-html5-backend";
import { Card } from "../model/model";
import { ItemTypes } from "../constants/dnd";

interface Props {
  card: Card;
  onDrag?: (card: Card, v: boolean) => void;
}
export const CardPanel = (props: Props) => {
  const { card } = props;
  const [expand, setExpand] = useState(false);

  const handleOnClick = () => {
    setExpand(prev => !prev);
  };

  const [{ isDragging }, drag, preview] = useDrag({
    item: { ...props.card, type: ItemTypes.CARD },
    previewOptions: {
      anchorX: 0
    },

    collect: monitor => ({
      isDragging: !!monitor.isDragging()
    })
  });

  useEffect(() => {
    props.onDrag && props.onDrag(card, isDragging);
  }, [isDragging]);

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
