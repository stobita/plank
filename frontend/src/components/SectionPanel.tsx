import React from "react";
import styled from "styled-components";
import { CardList } from "./CardList";
import { SectionPanelHead } from "./SectionPanelHead";
import { DraggableSection } from "./MoveContextProvider";

interface Props {
  section: DraggableSection;
}

export const SectionPanel = (props: Props) => {
  return (
    <Wrapper>
      <Inner>
        <SectionPanelHead section={props.section}></SectionPanelHead>
        <CardList items={props.section.cards}></CardList>
      </Inner>
    </Wrapper>
  );
};

const Wrapper = styled.div`
  padding: 0 8px;
  display: flex;
  flex-direction: column;
  min-width: 380px;
  box-sizing: border-box;
`;

const Inner = styled.div`
  background: ${props => props.theme.main};
  padding: 8px;
  box-sizing: border-box;
  border-radius: 4px;
`;
