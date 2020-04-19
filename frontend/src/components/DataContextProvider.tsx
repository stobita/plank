import React, { ReactNode, useState, useEffect } from "react";
import { DataContext } from "../context/dataContext";
import { Board, Section, Card, Label } from "../model/model";
import labelsRepository from "../api/labelsRepository";

interface Props {
  children: ReactNode;
}

export const DataContextProvider = (props: Props) => {
  const [boards, setBoards] = useState<Board[]>([]);

  const [sections, setSections] = useState<Section[]>([]);
  const [labels, setLabels] = useState<Label[]>([]);

  useEffect(() => {
    labelsRepository.getLabels().then((items) => {
      setLabels(items);
    });
  }, []);

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
