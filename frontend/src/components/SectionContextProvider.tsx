import React, { ReactNode, useState } from "react";
import { Card } from "../model/model";
import { SectionContext } from "../context/sectionContext";

interface Props {
  children: ReactNode;
}

export const SectionContextProvider = (props: Props) => {
  const [cards, setCards] = useState<Card[]>([]);
  return (
    <SectionContext.Provider
      value={{
        cards,
        setCards,
      }}
    >
      {props.children}
    </SectionContext.Provider>
  );
};
