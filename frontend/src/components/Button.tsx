import styled, { css } from "styled-components";

interface Props {
  primary?: boolean;
  danger?: boolean;
}

export const Button = styled.button<Props>`
  font-size: 16px;
  width: 100%;
  padding: 8px;
  box-sizing: border-box;
  border-radius: 4px;
  font-weight: bold;
  color: ${props => props.theme.weak};
  border: 1px solid ${props => props.theme.weak};
  ${props => props.primary && PrimaryStyle}
  ${props => props.danger && DangerStyle};
`;

const PrimaryStyle = css`
  border: none;
  background: ${props => props.theme.primary};
  color: ${props => props.theme.primaryInner};
`;

const DangerStyle = css`
  border: none;
  background: ${props => props.theme.danger};
  color: ${props => props.theme.dangerInner};
`;
