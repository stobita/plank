import React, { ReactNode, useState } from "react";
import { DataContext } from "../context/dataContext";
import { Board, Section } from "../model/model";

interface Props {
  children: ReactNode;
}

export const DataContextProvider = (props: Props) => {
  const [boards, setBoards] = useState<Board[]>([]);
  const [sections, setSections] = useState<Section[]>([]);
  return (
    <DataContext.Provider
      value={{
        boards,
        setBoards,
        sections,
        setSections,
      }}
    >
      {props.children}
    </DataContext.Provider>
  );
};
