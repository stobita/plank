import React, { useEffect, useContext, useState } from "react";
import styled from "styled-components";
import { CardListItemBox } from "./CardListItemBox";
import { CardPanel } from "./CardPanel";
import { DragItem } from "./MoveContextProvider";
import { CardDropTarget } from "./CardDropTarget";
import { MoveContext } from "../context/moveContext";
import { Card } from "../model/model";

interface Props {
  items: Card[];
  sectionId: number;
}

export const CardList = (props: Props) => {
  const { items } = props;
  const { overedCard, dropCard } = useContext(MoveContext);
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
      const next = list.filter((_, i) => i !== fromIndex);
      next.splice(toIndex, 0, card);
      setList(next);
    } else {
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
