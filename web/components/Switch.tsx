import React from "react";
import styled from "styled-components";
import { Transition } from "react-transition-group";
import { TransitionStatus } from "react-transition-group/Transition";

interface Props {
  value: boolean;
  onChange: () => void;
}

export const Switch = (props: Props) => {
  return (
    <Wrapper onClick={props.onChange}>
      <Left>
        <Transition in={props.value} timeout={500}>
          {state => <Toggle state={state}></Toggle>}
        </Transition>
      </Left>
      <Right></Right>
    </Wrapper>
  );
};
const Wrapper = styled.div`
  border: 1px solid ${props => props.theme.border};
  display: flex;
  cursor: pointer;
`;
const Left = styled.p`
  background: ${props => props.theme.primary};
  height: 24px;
  width: 24px;
  margin: 0;
`;
const Right = styled.p`
  background: ${props => props.theme.weak};
  height: 24px;
  width: 24px;
  margin: 0;
`;
const Toggle = styled.span<{ state: TransitionStatus }>`
  transition: 0.5s;
  height: 24px;
  width: 24px;
  display: block;
  background: ${props => props.theme.bg};

  transform: translateX(
    ${props =>
      props.state === "entering" || props.state === "entered" ? 24 : 0}px
  );
`;
