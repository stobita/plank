import React, { useEffect, useContext, useState } from "react";
import styled from "styled-components";
import { CardListItemBox } from "./CardListItemBox";
import { CardPanel } from "./CardPanel";
import { DragItem } from "./MoveContextProvider";
import { CardDropTarget } from "./CardDropTarget";
import { MoveContext } from "../context/moveContext";
import { Card } from "../model/model";
import { BoardContext } from "../context/boardContext";
import sectionsRepository from "../api/sectionsRepository";

interface Props {
  items: Card[];
  sectionId: number;
}

export const CardList = (props: Props) => {
  const { items } = props;
  const { overedCard } = useContext(MoveContext);
  const { removeCard } = useContext(BoardContext);
  const [list, setList] = useState<DragItem[]>([]);

  useEffect(() => {
    if (items !== null) {
      setList(items.slice());
    }
  }, [items]);

  useEffect(() => {
    // TODO: fix
    if (!("id" in overedCard)) {
    } else if (overedCard.section.id === props.sectionId) {
      const toIndex = list.findIndex(
        (v) => "id" in v && v.id === overedCard.id
      );
      const next = list.filter((v) => "id" in v);
      next.splice(toIndex, 0, { kind: "target" });
      setList(next);
    } else {
      setList((prev) => prev.filter((v) => "id" in v));
    }
  }, [overedCard]);

  const handleOnDrag = (card: Card, v: boolean) => {
    if (v) {
      const toIndex = list.findIndex((v) => "id" in v && v.id === card.id);
      const next = list.filter((v) => "id" in v);
      next.splice(toIndex, 0, { kind: "target" });
      setList(next);
    } else {
      setList((prev) => prev.filter((v) => "id" in v));
    }
  };

  const handleOnDrop = (card: Card) => {
    if (props.sectionId === card.section.id) {
      const fromIndex = list.findIndex((v) => "id" in v && v.id === card.id);
      const toIndex = list.findIndex((v) => !("id" in v));
      const prev = list[toIndex - 1];
      const next = list.filter((_, i) => i !== fromIndex);

      next.splice(toIndex, 0, card);
      setList(next);

      if (prev && "id" in prev) {
        sectionsRepository.updateCardPosition(
          props.sectionId,
          card.id,
          prev.id
        );
      } else {
        sectionsRepository.updateCardPosition(props.sectionId, card.id, null);
      }
    } else {
      const toIndex = list.findIndex((v) => !("id" in v));
      const next = list;
      next.splice(toIndex, 1, card);
      setList(next);
      removeCard(card);
    }
  };

  return (
    <Wrapper>
      {list.map((item) => (
        <Item key={"id" in item ? item.id : 0}>
          {!("id" in item) ? (
            <CardDropTarget
              active={!("id" in item)}
              key={0}
              onDrop={handleOnDrop}
            ></CardDropTarget>
          ) : (
            <CardListItemBox item={item}>
              <CardPanel
                key={item.id}
                card={item}
                onDrag={handleOnDrag}
              ></CardPanel>
            </CardListItemBox>
          )}
        </Item>
      ))}
    </Wrapper>
  );
};

const Wrapper = styled.div`
  display: flex;
  flex-direction: column;
`;

const Item = styled.div``;
