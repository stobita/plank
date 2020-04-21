import React, { useContext, useEffect } from "react";
import styled from "styled-components";
import { SectionPanel } from "./SectionPanel";
import { ViewContext } from "../context/viewContext";
import { EventContext } from "../context/eventContext";
import boardsRepository from "../api/boardsRepository";
import { Section, Card } from "../model/model";
import {
  DragDropContext,
  Droppable,
  Draggable,
  DropResult,
  DraggableLocation,
} from "react-beautiful-dnd";
import { DataContext } from "../context/dataContext";
import sectionsRepository from "../api/sectionsRepository";

export const BoardView = () => {
  const { sections, setSections } = useContext(DataContext);
  const { currentBoard } = useContext(ViewContext);
  const { updatedBoardIds } = useContext(EventContext);

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

  const reorderCard = async (
    section: Section,
    startIndex: number,
    endIndex: number
  ) => {
    const cards = sections.find((v: Section) => v.id === section.id)?.cards;
    if (!cards) return;
    const result = Array.from(cards);
    const [removed] = result.splice(startIndex, 1);
    result.splice(endIndex, 0, removed);

    const sectionIndex = sections.findIndex(
      (v: Section) => v.id === section.id
    );
    const sectionResult = Array.from(sections);
    sectionResult[sectionIndex].cards = result;
    setSections(sectionResult);

    await sectionsRepository.reorderCardPosition(
      section.id,
      removed.id,
      endIndex
    );
  };

  const moveCard = async (
    source: Section,
    destination: Section,
    droppableSource: DraggableLocation,
    droppableDestination: DraggableLocation
  ) => {
    const sourceClone = Array.from(source.cards);
    console.log(destination);
    let destClone: Card[];
    if (destination.cards === null) {
      destClone = [];
    } else {
      destClone = Array.from(destination.cards) || [];
    }
    const [removed] = sourceClone.splice(droppableSource.index, 1);

    destClone.splice(droppableDestination.index, 0, removed);

    const sourceSectionIndex = sections.findIndex((v) => v.id === source.id);

    const destSectionIndex = sections.findIndex(
      (v) => v.id === destination!.id
    );

    const sectionResult = Array.from(sections);

    sectionResult[sourceSectionIndex].cards = sourceClone;
    sectionResult[destSectionIndex].cards = destClone;
    setSections(sectionResult);

    await sectionsRepository.moveCard(
      source.id,
      removed.id,
      droppableDestination.index,
      destination.id
    );
  };

  const reorderSection = async (startIndex: number, endIndex: number) => {
    const result = Array.from(sections);
    const [removed] = result.splice(startIndex, 1);
    result.splice(endIndex, 0, removed);
    setSections(result);

    await boardsRepository.reorderSection(
      currentBoard.id,
      removed.id,
      endIndex
    );
  };

  const onDragEnd = (result: DropResult) => {
    const { source, destination, type } = result;
    if (type !== "BOARD") {
      if (source.droppableId === destination?.droppableId) {
        const section = sections.find(
          (v: Section) => String(v.id) === source.droppableId
        );
        if (!section) return;
        reorderCard(section, source.index, destination.index);
      } else {
        const sourceSection = sections.find(
          (v: Section) => String(v.id) === source.droppableId
        );
        const destSection = sections.find(
          (v: Section) => String(v.id) === destination?.droppableId
        );
        if (!sourceSection || !destSection || !destination) return;
        moveCard(sourceSection, destSection, source, destination);
      }
    } else {
      reorderSection(source.index, destination!.index);
    }
  };

  return (
    <DragDropContext onDragEnd={onDragEnd}>
      <Droppable
        droppableId={`board-${currentBoard.id}`}
        type="BOARD"
        direction="horizontal"
      >
        {(provided, snapshot) => (
          <Wrapper ref={provided.innerRef}>
            <Main>
              {sections.map((section: Section, index) => (
                <Draggable
                  draggableId={String(section.id)}
                  index={index}
                  key={section.id}
                >
                  {(provided, snapshot) => (
                    <div ref={provided.innerRef} {...provided.draggableProps}>
                      <SectionPanel
                        section={section}
                        key={section.id}
                        provided={provided}
                      />
                    </div>
                  )}
                </Draggable>
              ))}
              {provided.placeholder}
            </Main>
          </Wrapper>
        )}
      </Droppable>
    </DragDropContext>
  );
};

const Wrapper = styled.div`
  display: flex;
  flex-direction: column;
  flex: 1;
  padding: 16px;
`;

const Main = styled.div`
  display: flex;
`;
