import React, { useState, useEffect } from "react";
import styled, { css } from "styled-components";

interface Props {
  items: string[];
  inputValue: string;
  onClickItem: (id: number) => void;
  onMouseEnterItem: (id: number) => void;
  hoverIndex: number;
}

export const AutoSuggest = (props: Props) => {
  if (props.items.length < 1) return null;

  const handleOnClick = (e: React.MouseEvent<HTMLElement>) => {
    const targetIdx = e.currentTarget.dataset.idx;
    if (targetIdx) {
      props.onClickItem(Number(targetIdx));
    }
  };

  const handleOnMouseEnter = (e: React.MouseEvent<HTMLLIElement>) => {
    const targetIdx = e.currentTarget.dataset.idx;
    if (targetIdx) {
      props.onMouseEnterItem(Number(targetIdx));
    }
  };

  return (
    <Wrapper>
      <ul>
        {props.items.map((v, idx) => (
          <Item
            isSelect={idx === props.hoverIndex}
            key={idx}
            onClick={handleOnClick}
            data-idx={idx}
            onMouseEnter={handleOnMouseEnter}
          >
            {v}
          </Item>
        ))}
      </ul>
    </Wrapper>
  );
};

export const useSuggest = (items: string[], inputValue: string) => {
  const [selectedSuggestion, setSelectedSuggestion] = useState("");
  const [selectedSuggestionIndex, setSelectedSuggestionIndex] = useState(0);

  const suggestions = items.filter(
    (v) =>
      inputValue.length > 0 &&
      v.toLowerCase().startsWith(inputValue.toLowerCase())
  );

  useEffect(() => {
    setSelectedSuggestion(suggestions[selectedSuggestionIndex]);
  }, [selectedSuggestionIndex, suggestions]);

  const suggestUp = () => {
    if (selectedSuggestionIndex < 1) {
      setSelectedSuggestionIndex(suggestions.length - 1);
    } else {
      setSelectedSuggestionIndex((prev) => --prev);
    }
  };
  const suggestDown = () => {
    if (selectedSuggestionIndex > suggestions.length - 2) {
      setSelectedSuggestionIndex(0);
    } else {
      setSelectedSuggestionIndex((prev) => ++prev);
    }
  };

  return {
    suggestions,
    selectedSuggestion,
    setSelectedSuggestion,
    selectedSuggestionIndex,
    setSelectedSuggestionIndex,
    suggestUp,
    suggestDown,
  };
};

const Wrapper = styled.div`
  border: 1px solid ${(props) => props.theme.border};
  background: ${(props) => props.theme.main};
  margin-top: -4px;
`;

const Hovered = css`
  background: ${(props) => props.theme.weak};
`;

const Item = styled.li<{ isSelect: boolean }>`
  ${(props) => props.isSelect && Hovered}
  padding: 8px;
  color: ${(props) => props.theme.solid};
`;
