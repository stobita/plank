import React from "react";
import styled from "styled-components";
import { CardList } from "./CardList";
import { SectionPanelHead } from "./SectionPanelHead";
import { Section } from "../model/model";
import {
  DroppableProvided,
  Droppable,
  DraggableProvided,
} from "react-beautiful-dnd";

interface Props {
  section: Section;
  provided: DraggableProvided;
}

export const SectionPanel = (props: Props) => {
  return (
    <Droppable key={props.section.id} droppableId={String(props.section.id)}>
      {(provided, snapshot) => (
        <div ref={provided.innerRef}>
          <Wrapper>
            <Inner>
              <div {...props.provided.dragHandleProps}>
                <SectionPanelHead section={props.section}></SectionPanelHead>
              </div>
              <CardList
                items={props.section.cards}
                provided={provided}
              ></CardList>
            </Inner>
          </Wrapper>
        </div>
      )}
    </Droppable>
  );
};

const Wrapper = styled.div`
  padding: 0 8px;
  display: flex;
  flex-direction: column;
  min-width: 380px;
  max-width: 380px;
  box-sizing: border-box;
`;

const Inner = styled.div`
  background: ${(props) => props.theme.main};
  padding: 8px;
  box-sizing: border-box;
  border-radius: 4px;
`;
