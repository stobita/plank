import React from "react";
import styled, { css } from "styled-components";
import { ReactComponent as PlusIconImage } from "../assets/plus.svg";

interface Props {
  onClick: () => void;
  isClose?: boolean;
}

export const AddButton = (props: Props) => {
  return (
    <Button onClick={props.onClick}>
      <PlusIcon isClose={props.isClose} />
    </Button>
  );
};

const Button = styled.div`
  border: 1px solid ${props => props.theme.solid};
  background: ${props => props.theme.main};
  height: 32px;
  width: 32px;
  border-radius: 16px;
  position: relative;
  display: flex;
  justify-content: center;
  align-items: center;
  box-sizing: border-box;
  cursor: pointer;
`;

const PlusIcon = styled(PlusIconImage)<{ isClose?: boolean }>`
  fill: ${props => props.theme.solid};
  transition: all 0.3s;
  height: 24px;
  width: 24px;
  ${props =>
    props.isClose &&
    css`
      transform: rotate(45deg);
    `}
`;
