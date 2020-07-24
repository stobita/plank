import React, { useState, useContext } from 'react';
import { Input } from './Input';
import styled from 'styled-components';
import { ReactComponent as SearchIconImage } from '../assets/search.svg';
import boardsRepository from '../api/boardsRepository';
import { ViewContext } from '../context/viewContext';
import { DataContext } from '../context/dataContext';

interface Props {}

export const SearchBoardCards = (props: Props) => {
  const { currentBoard } = useContext(ViewContext);
  const { setSections } = useContext(DataContext);
  const [word, setWord] = useState('');
  const handleOnChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setWord(e.target.value);
  };
  const handleOnSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    boardsRepository
      .searchBoardSections(currentBoard.id, word)
      .then((items) => {
        setSections([]);
        setSections(items);
      });
  };
  return (
    <Wrapper>
      <SearchIcon></SearchIcon>
      <form onSubmit={handleOnSubmit}>
        <Input
          type="text"
          style={{
            paddingLeft: '36px',
          }}
          placeholder="Search Cards"
          onChange={handleOnChange}
        ></Input>
      </form>
    </Wrapper>
  );
};

const Wrapper = styled.div`
  width: 256px;
  display: flex;
  align-items: center;
  position: relative;
`;

const SearchIcon = styled(SearchIconImage)`
  fill: ${(props) => props.theme.weak};
  height: 24px;
  width: 24px;
  position: absolute;
  left: 8px;
`;
