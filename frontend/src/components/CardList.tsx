import React, { useEffect, useState } from "react";
import styled from "styled-components";
import { CardPanel } from "./CardPanel";
import { Card } from "../model/model";
import { Draggable, DroppableProvided } from "react-beautiful-dnd";

interface Props {
  items: Card[];
  provided: DroppableProvided;
}

export const CardList = (props: Props) => {
  const { items } = props;
  const [list, setList] = useState<Card[]>([]);

  useEffect(() => {
    if (items !== null) {
      setList(items.slice());
    }
  }, [items]);

  return (
    <Wrapper>
      {list.map((item, index) => (
        <Item key={item.id}>
          <Draggable draggableId={`card-${item.id}`} index={index}>
            {(provided, snapshot) => (
              <CardItemBox
                ref={provided.innerRef}
                {...provided.draggableProps}
                {...provided.dragHandleProps}
              >
                <CardPanel card={item}></CardPanel>
              </CardItemBox>
            )}
          </Draggable>
        </Item>
      ))}
      {props.provided.placeholder}
    </Wrapper>
  );
};

const Wrapper = styled.div`
  display: flex;
  flex-direction: column;
`;

const Item = styled.div``;

const CardItemBox = styled.div`
  margin-top: 8px;
`;
