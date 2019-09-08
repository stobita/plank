import styled from "styled-components";
import colors from "../colors";

export const Textarea = styled.textarea`
  border: 1px solid ${props => props.theme.border};
  background: ${colors.white};
  font-size: 16px;
  width: 100%;
  padding: 8px;
  box-sizing: border-box;
  border-radius: 4px;
  resize: none;
  min-height: 64px;
`;
