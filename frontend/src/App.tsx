import React, { useContext } from "react";
import "./App.css";
import { Home } from "./components/Home";
import styled, { ThemeProvider } from "styled-components";
import Provider, { MainContext } from "./Provider";
import theme from "./theme";

const App: React.FC = () => {
  const { themeName } = useContext(MainContext);
  return (
    <div className="App">
      <Provider>
        <ThemeProvider theme={theme[themeName]}>
          <Home />
        </ThemeProvider>
      </Provider>
    </div>
  );
};

export default App;
