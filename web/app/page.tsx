'use client'

import { Home } from '../components/Home';
import { Provider } from "../components/Provider";

export default function Root() {
  return (
    <Provider>
      <Home />
    </Provider>
  );
}

