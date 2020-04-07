import { createContext, Dispatch, SetStateAction } from "react";
import { DragItem } from "../components/MoveContextProvider";
import { Card } from "../model/model";

type MoveContextProps = {
  dropCard: () => void;
  overCard: (card: DragItem) => void;
  overedCard: DragItem;
  setOveredCard: Dispatch<SetStateAction<DragItem>>;
};

const defaultProps = {
  dropCard: () => {},
  overCard: () => {},
  overedCard: {} as Card,
  setOveredCard: () => {}
};

export const MoveContext = createContext<MoveContextProps>(defaultProps);
