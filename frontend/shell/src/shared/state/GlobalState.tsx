import React, { createContext, useContext, useReducer, ReactNode } from 'react';

type State = {
  theme: 'light' | 'dark';
};

type Action = { type: 'SET_THEME'; theme: 'light' | 'dark' };

const initialState: State = {
  theme: 'light',
};

function reducer(state: State, action: Action): State {
  switch (action.type) {
    case 'SET_THEME':
      return { ...state, theme: action.theme };
    default:
      return state;
  }
}

const GlobalStateContext = createContext<{ state: State; dispatch: React.Dispatch<Action> } | undefined>(undefined);

export function GlobalStateProvider({ children }: { children: ReactNode }) {
  const [state, dispatch] = useReducer(reducer, initialState);
  return (
    <GlobalStateContext.Provider value={{ state, dispatch }}>
      {children}
    </GlobalStateContext.Provider>
  );
}

export function useGlobalState() {
  const context = useContext(GlobalStateContext);
  if (!context) throw new Error('useGlobalState must be used within GlobalStateProvider');
  return context;
}
