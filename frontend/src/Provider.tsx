import React, { ReactNode, Children } from "react";
import { ThemeType } from "./theme";

interface Props {
  children: ReactNode;
}

type MainContextType = {
  themeName: ThemeType;
  setThemeName: (t: ThemeType) => void;
};

export const MainContext = React.createContext<MainContextType>({
  themeName: "light",
  setThemeName: () => {}
});

export default (props: Props) => {
  const [themeName, setThemeName] = React.useState();
  return (
    <MainContext.Provider value={{ themeName, setThemeName }}>
      <div>{props.children}</div>
    </MainContext.Provider>
  );
};
