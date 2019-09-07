import styled, { css } from "styled-components";
import colors from "../colors";

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
  color: ${colors.mainGray};
  border: 1px solid ${colors.mainGray};
  ${props => props.primary && PrimaryStyle}
  ${props => props.danger && DangerStyle};
`;

const PrimaryStyle = css`
  border: none;
  background: ${colors.primary};
  color: ${colors.mainWhite};
`;

const DangerStyle = css`
  border: none;
  background: ${colors.danger};
  color: ${colors.mainWhite};
`;
