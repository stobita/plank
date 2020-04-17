import { createContext, Dispatch, SetStateAction } from "react";
import { Board, Card, Section } from "../model/model";

type DataContextProps = {
  boards: Board[];
  setBoards: Dispatch<SetStateAction<Board[]>>;
  sections: Section[];
  setSections: Dispatch<SetStateAction<Section[]>>;
};

const defaultProps = {
  boards: [],
  setBoards: () => {},
  sections: [],
  setSections: () => {},
};

export const DataContext = createContext<DataContextProps>(defaultProps);
