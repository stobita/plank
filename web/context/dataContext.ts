import { createContext, Dispatch, SetStateAction } from 'react';
import { Board, Section, Label } from '../model/model';

type DataContextProps = {
  boards: Board[];
  setBoards: Dispatch<SetStateAction<Board[]>>;
  sections: Section[];
  setSections: Dispatch<SetStateAction<Section[]>>;

  labels: Label[];
  setLabels: Dispatch<SetStateAction<Label[]>>;
};

const defaultProps = {
  boards: [],
  setBoards: () => {},
  sections: [],
  setSections: () => {},
  labels: [],
  setLabels: () => {},
};

export const DataContext = createContext<DataContextProps>(defaultProps);
