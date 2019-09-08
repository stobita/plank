import React from "react";
import styled, { css } from "styled-components";

interface Props {
  onClick: () => void;
  isClose?: boolean;
}

export const AddButton = (props: Props) => {
  return (
    <Button onClick={props.onClick}>
      <ButtonInner isClose={props.isClose} />
      <ButtonInner isClose={props.isClose} />
    </Button>
  );
};

const Button = styled.button`
  border: 1px solid ${props => props.theme.border};
  background: ${props => props.theme.main};
  height: 32px;
  width: 32px;
  border-radius: 16px;
  position: relative;
  display: flex;
  justify-content: center;
  align-items: center;
  cursor: pointer;
`;

const ButtonInner = styled.span<{ isClose?: boolean }>`
  transition: all 0.3s;
  position: absolute;
  width: 50%;
  height: 2px;
  background: ${props => props.theme.text};
  border-radius: 4px;
  &:nth-of-type(1) {
    transform: rotate(90deg);
  }
  ${props =>
    props.isClose &&
    css`
      &:nth-of-type(1) {
        transform: rotate(135deg);
      }
      &:nth-of-type(2) {
        transform: rotate(45deg);
      }
    `}
`;
