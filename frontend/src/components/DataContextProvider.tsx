import React, { ReactNode, useState } from "react";
import { DataContext } from "../context/dataContext";
import { Board, Section, Card } from "../model/model";

interface Props {
  children: ReactNode;
}

export const DataContextProvider = (props: Props) => {
  const [boards, setBoards] = useState<Board[]>([]);
  const [sections, setSections] = useState<Section[]>([]);
  const [cards, setCards] = useState<Card[]>([]);
  return (
    <DataContext.Provider
      value={{
        boards,
        setBoards,
        sections,
        setSections,
        cards,
        setCards
      }}
    >
      {props.children}
    </DataContext.Provider>
  );
};
