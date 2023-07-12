import React, { createContext, useState } from 'react';

// Create the authentication context
export const AuthContext = createContext();

// Create a provider component to wrap your app and provide the authentication state
export const AuthProvider = ({ children }) => {
  const [isLoggedIn, setIsLoggedIn] = useState(false);

  return (
    <AuthContext.Provider value={{ isLoggedIn, setIsLoggedIn }}>
      {children}
    </AuthContext.Provider>
  );
};
