import React, { useContext } from "react";
import styled from "styled-components";
import { SectionPanel } from "./SectionPanel";
import { CardPanelDragPreview } from "./CardPanelDragPreview";
import { DataContext } from "../context/dataContext";

export const BoardView = () => {
  const { sections } = useContext(DataContext);
  return (
    <Wrapper>
      <Main>
        {sections.map(v => (
          <SectionPanel section={v} key={v.id} />
        ))}
      </Main>
      <CardPanelDragPreview></CardPanelDragPreview>
    </Wrapper>
  );
};

const Wrapper = styled.div`
  display: flex;
  flex-direction: column;
  flex: 1;
`;

const Main = styled.div`
  display: flex;
`;
