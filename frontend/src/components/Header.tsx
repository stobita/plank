import React, { useContext } from "react";
import styled from "styled-components";
import { ViewContext } from "../context/viewContext";
import { SectionController } from "./SectionController";
import { ReactComponent as FilterIconImage } from "../assets/filter.svg";

export const Header = () => {
  const { currentBoard } = useContext(ViewContext);
  const { setSectionFilterActive } = useContext(ViewContext);
  const handleOnClickFilter = () => {
    setSectionFilterActive((prev) => !prev);
  };
  return (
    <Wrapper>
      <FilterIcon onClick={handleOnClickFilter}></FilterIcon>
      <CreateSection>
        {currentBoard.id !== 0 && <SectionController />}
      </CreateSection>
    </Wrapper>
  );
};

const Wrapper = styled.header`
  background: ${(props) => props.theme.main};
  height: 56px;
  border-bottom: 1px solid ${(props) => props.theme.border};
  display: flex;
  align-items: center;
`;

const CreateSection = styled.div`
  padding-left: 16px;
  display: flex;
  align-items: center;
`;

const FilterIcon = styled(FilterIconImage)`
  cursor: pointer;
  fill: ${(props) => props.theme.solid};
  height: 36px;
  width: 36px;
  margin: 0 16px;
`;
