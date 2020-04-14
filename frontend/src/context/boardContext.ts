import { createContext, Dispatch, SetStateAction } from "react";
import { Section, Card } from "../model/model";

type BoardContextProps = {
  sections: Section[];
  setSections: Dispatch<SetStateAction<Section[]>>;
  removeCard: (card: Card) => void;
};

const defaultProps = {
  sections: [],
  setSections: () => {},
  removeCard: (card: Card) => {},
};

export const BoardContext = createContext<BoardContextProps>(defaultProps);
