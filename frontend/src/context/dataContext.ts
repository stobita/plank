import { createContext, Dispatch, SetStateAction } from "react";
import { Board, Card } from "../model/model";

type DataContextProps = {
  boards: Board[];
  setBoards: Dispatch<SetStateAction<Board[]>>;
};

const defaultProps = {
  boards: [],
  setBoards: () => {},
};

export const DataContext = createContext<DataContextProps>(defaultProps);
