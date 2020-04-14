import React, { ReactNode, useState } from "react";
import { Section, Card } from "../model/model";
import { BoardContext } from "../context/boardContext";

interface Props {
  children: ReactNode;
}

export const BoardContextProvider = (props: Props) => {
  const [sections, setSections] = useState<Section[]>([]);
  const removeCard = (card: Card) => {
    const section = sections.find((v) => v.id === card.section.id);
    if (!section) {
      return;
    }
    const cards = section.cards.filter((v) => v.id !== card.id);
    section.cards = cards;

    const next = [...sections.filter((v) => v.id !== section.id), section];

    setSections(next);
  };
  return (
    <BoardContext.Provider
      value={{
        sections,
        setSections,
        removeCard,
      }}
    >
      {props.children}
    </BoardContext.Provider>
  );
};
