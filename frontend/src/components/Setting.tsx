import React, { useContext } from "react";
import styled from "styled-components";
import { Switch } from "./Switch";
import { MainContext } from "../Provider";

export const Setting = () => {
  const { themeName, setThemeName } = useContext(MainContext);
  const isDarkTheme = themeName === "dark";
  const handleOnClickThemeSwitch = () => {
    setThemeName(prev => (prev === "light" ? "dark" : "light"));
  };
  return (
    <Wrapper>
      <Field>
        <label>Dark Theme</label>
        <Switch
          value={isDarkTheme}
          onChange={handleOnClickThemeSwitch}
        ></Switch>
      </Field>
    </Wrapper>
  );
};
const Wrapper = styled.div`
  padding: 8px 0;
`;

const Field = styled.div`
  display: flex;
  justify-content: space-between;
  align-items: center;
`;
