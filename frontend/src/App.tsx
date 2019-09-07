import React from "react";
import "./App.css";
import { Home } from "./components/Home";
import Provider from "./Provider";

const App: React.FC = () => {
  return (
    <div className="App">
      <Provider>
        <Home />
      </Provider>
    </div>
  );
};

export default App;
