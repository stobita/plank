import React, { ReactNode, useState } from 'react';
import { DataContext } from '../context/dataContext';
import { Board, Section, Label } from '../model/model';

interface Props {
  children: ReactNode;
}

export const DataContextProvider = (props: Props) => {
  const [boards, setBoards] = useState<Board[]>([]);

  const [sections, setSections] = useState<Section[]>([]);
  const [labels, setLabels] = useState<Label[]>([]);

  return (
    <DataContext.Provider
      value={{
        boards,
        setBoards,
        sections,
        setSections,
        labels,
        setLabels,
      }}
    >
      {props.children}
    </DataContext.Provider>
  );
};
