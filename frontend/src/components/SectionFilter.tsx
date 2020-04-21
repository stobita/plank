import React, { useContext, useState } from "react";
import styled from "styled-components";
import { ReactComponent as CloseIconImage } from "../assets/close.svg";
import { DataContext } from "../context/dataContext";
import { ViewContext } from "../context/viewContext";
import boardsRepository from "../api/boardsRepository";
import { Label } from "../model/model";
import sectionsRepository from "../api/sectionsRepository";

export const SectionFilter = () => {
  const { setSections, labels } = useContext(DataContext);
  const [selectedLabelId, setSelectedLabelId] = useState(0);
  const { currentBoard, setSectionFilterActive } = useContext(ViewContext);
  const handleOnClickClose = () => {
    setSectionFilterActive(false);
  };

  const handleOnClickLabel = async (labelId: number) => {
    const res = await boardsRepository.getLabelSections(
      currentBoard.id,
      labelId
    );
    setSections([]);
    setSections(res);
    setSelectedLabelId(labelId);
  };
  return (
    <Wrapper>
      <Header>
        <CloseIcon onClick={handleOnClickClose}></CloseIcon>
      </Header>
      <Heading>labels</Heading>
      <List>
        {labels.map((v) => (
          <Item
            selected={selectedLabelId === v.id}
            onClick={() => handleOnClickLabel(v.id)}
            key={v.id}
          >
            {v.name}
          </Item>
        ))}
      </List>
    </Wrapper>
  );
};

const Wrapper = styled.div`
  border-right: 1px solid ${(props) => props.theme.border};
  background: ${(props) => props.theme.main};
  width: 180px;
  padding: 16px;
`;

const Header = styled.div`
  display: flex;
  justify-content: flex-end;
`;

const Heading = styled.p`
  font-size: 20px;
  font-weight: bold;
`;

const CloseIcon = styled(CloseIconImage)`
  fill: ${(props) => props.theme.solid};
  height: 24px;
  width: 24px;
  cursor: pointer;
`;

const List = styled.ul``;
const Item = styled.li<{ selected: boolean }>`
  font-weight: ${(props) => props.selected && "bold"};
  padding: 4px;
  cursor: pointer;
`;
