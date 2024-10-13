'use client'

import { Home } from '../components/Home';
import { Provider } from "../components/Provider";
import { ViewContext } from '../context/viewContext';
import { DataContext } from '../context/dataContext';
import { EventContext } from '../context/eventContext';
import { useContext, useEffect } from 'react';

export default function Root() {
  return (
    <Provider>
      <Home />
    </Provider>
  );
}

