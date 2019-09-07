import React from "react";
import styled from "styled-components";

interface Props {
  onClick: () => void;
}

export const CloseButton = (props: Props) => {
  return (
    <Button onClick={props.onClick}>
      <ButtonInner />
      <ButtonInner />
    </Button>
  );
};

const Button = styled.button`
  border: 1px solid ${props => props.theme.border};
  height: 36px;
  width: 36px;
  border-radius: 18px;
  position: relative;
  display: flex;
  justify-content: center;
  align-items: center;
  cursor: pointer;
`;

const ButtonInner = styled.span`
  position: absolute;
  width: 50%;
  height: 2px;
  background: ${props => props.theme.text};
  border-radius: 4px;
  &:nth-of-type(1) {
    transform: rotate(45deg);
  }
  &:nth-of-type(2) {
    transform: rotate(135deg);
  }
`;
