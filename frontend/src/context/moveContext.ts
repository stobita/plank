import { createContext, Dispatch, SetStateAction } from "react";
import {
  DraggableCard,
  DraggableSection
} from "../components/MoveContextProvider";

type MoveContextProps = {
  draggingCard: DraggableCard | undefined;
  setDraggingCard: Dispatch<SetStateAction<DraggableCard | undefined>>;
  dropCard: () => void;
  overCard: (card: DraggableCard) => void;
  sections: DraggableSection[];
  isDraggingCard: boolean;
  setIsDraggingCard: Dispatch<SetStateAction<boolean>>;
};

const defaultProps = {
  draggingCard: {} as DraggableCard,
  setDraggingCard: () => {},
  dropCard: () => {},
  overCard: () => {},
  sections: [],
  isDraggingCard: false,
  setIsDraggingCard: () => {}
};

export const MoveContext = createContext<MoveContextProps>(defaultProps);
