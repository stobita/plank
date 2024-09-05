import React, { useContext } from 'react';
import styled from 'styled-components';
import { ViewContext } from '../context/viewContext';
import { SectionController } from './SectionController';
import FilterIconImage from '../assets/filter.svg';
import { SearchBoardCards } from './SearchBoardCards';

export const Header = () => {
  const { currentBoard } = useContext(ViewContext);
  const { setSectionFilterActive } = useContext(ViewContext);
  const handleOnClickFilter = () => {
    setSectionFilterActive((prev) => !prev);
  };
  return (
    <Wrapper>
      <Left>
        <FilterIcon onClick={handleOnClickFilter}></FilterIcon>
        <CreateSection>
          {currentBoard.id !== 0 && <SectionController />}
        </CreateSection>
      </Left>
      <Right>
        <SearchBoardCards></SearchBoardCards>
      </Right>
    </Wrapper>
  );
};

const Wrapper = styled.header`
  background: ${(props) => props.theme.main};
  height: 56px;
  border-bottom: 1px solid ${(props) => props.theme.border};
  display: flex;
`;

const Left = styled.div`
  flex: 1;
  display: flex;
  align-items: center;
`;

const Right = styled.div`
  flex: 1;
  display: flex;
  align-items: center;
  padding-right: 16px;
  justify-content: flex-end;
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
