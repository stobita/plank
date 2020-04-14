import React, { useContext, useEffect } from "react";
import styled from "styled-components";
import { SectionPanel } from "./SectionPanel";
import { CardPanelDragPreview } from "./CardPanelDragPreview";
import { BoardContext } from "../context/boardContext";
import { ViewContext } from "../context/viewContext";
import { EventContext } from "../context/eventContext";
import boardsRepository from "../api/boardsRepository";
import { Section } from "../model/model";

export const BoardView = () => {
  const { sections } = useContext(BoardContext);
  const { currentBoard } = useContext(ViewContext);
  const { updatedBoardIds } = useContext(EventContext);
  const { setSections } = useContext(BoardContext);

  useEffect(() => {
    if (updatedBoardIds.find((v) => v === currentBoard.id)) {
      boardsRepository.getBoardSections(currentBoard.id).then((items) => {
        setSections(items);
      });
    }
  }, [currentBoard.id, updatedBoardIds, setSections]);

  useEffect(() => {
    if (currentBoard.id !== 0) {
      boardsRepository.getBoardSections(currentBoard.id).then((items) => {
        setSections(items);
      });
    }
  }, [currentBoard.id, setSections]);

  return (
    <Wrapper>
      <Main>
        {sections
          .slice()
          .sort((v) => v.position)
          .map((v: Section) => (
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
