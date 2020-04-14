import React, { ReactNode, useState } from "react";
import { DataContext } from "../context/dataContext";
import { Board } from "../model/model";

interface Props {
  children: ReactNode;
}

export const DataContextProvider = (props: Props) => {
  const [boards, setBoards] = useState<Board[]>([]);
  return (
    <DataContext.Provider
      value={{
        boards,
        setBoards,
      }}
    >
      {props.children}
    </DataContext.Provider>
  );
};
