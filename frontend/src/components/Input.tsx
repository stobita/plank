import styled from "styled-components";
import colors from "../colors";

export const Input = styled.input`
  border: 1px solid ${props => props.theme.border};
  background: ${colors.white};
  font-size: 16px;
  width: 100%;
  padding: 8px;
  box-sizing: border-box;
  border-radius: 4px;
  height: 36px;
`;
