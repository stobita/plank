export type Board = {
  id: number;
  name: string;
};

export type Section = {
  id: number;
  name: string;
  cards: Card[];
};

export type Card = {
  id: number;
  name: string;
  description: string;
  section: Section;
};
