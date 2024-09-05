import React from "react";
import styled from "styled-components";
import CloseIconImage from "../assets/close.svg";

interface Props {
  onClick: () => void;
}

export const CloseButton = (props: Props) => {
  return (
    <Button onClick={props.onClick}>
      <CloseIcon />
    </Button>
  );
};

const Button = styled.button`
  border: 1px solid ${props => props.theme.solid};
  height: 32px;
  width: 32px;
  border-radius: 16px;
  position: relative;
  display: flex;
  justify-content: center;
  align-items: center;
  cursor: pointer;
`;

const CloseIcon = styled(CloseIconImage)`
  fill: ${props => props.theme.solid};
  height: 24px;
  width: 24px;
`;
