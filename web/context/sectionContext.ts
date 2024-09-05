import { Card } from "../model/model";
import { Dispatch, SetStateAction, createContext } from "react";

type SectionContextProps = {
  cards: Card[];
  setCards: Dispatch<SetStateAction<Card[]>>;
};
const defaultProps = {
  cards: [],
  setCards: () => {},
};

export const SectionContext = createContext<SectionContextProps>(defaultProps);
