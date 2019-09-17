import React, { useContext } from "react";
import styled from "styled-components";
import { EventContext } from "../context/eventContext";
import { ViewContext } from "../context/viewContext";
import { Board } from "../model/model";

interface Props {
  item: Board;
}
export const BoardListItem = (props: Props) => {
  const { item } = props;
  const { setCurrentBoardId, currentBoardId } = useContext(ViewContext);
  const { updatedBoardIds } = useContext(EventContext);
  const handleOnClickItem = (id: number) => {
    setCurrentBoardId(id);
  };
  return (
    <Wrapper
      key={item.id}
      onClick={() => handleOnClickItem(item.id)}
      selected={currentBoardId === item.id}
      updated={!!updatedBoardIds.find(id => id === item.id)}
    >
      {item.name}
    </Wrapper>
  );
};

const Wrapper = styled.li<{ selected: boolean; updated: boolean }>`
  font-weight: bold;
  color: ${props =>
    props.selected
      ? props.theme.solid
      : props.updated
      ? props.theme.primary
      : props.theme.weak};
  display: flex;
  margin-bottom: 12px;
  font-size: 1.1rem;
  cursor: pointer;
`;
