import { createContext, Dispatch, SetStateAction } from "react";
import { Board, Section, Card } from "../model/model";

type DataContextProps = {
  boards: Board[];
  setBoards: Dispatch<SetStateAction<Board[]>>;

  sections: Section[];
  setSections: Dispatch<SetStateAction<Section[]>>;

  cards: Card[];
  setCards: Dispatch<SetStateAction<Card[]>>;
};

const defaultProps = {
  boards: [],
  setBoards: () => {},

  sections: [],
  setSections: () => {},

  cards: [],
  setCards: () => {}
};

export const DataContext = createContext<DataContextProps>(defaultProps);
