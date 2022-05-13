import React, { useEffect } from 'react';
import axios from 'axios';

import { Center, MantineProvider } from "@mantine/core";
import Chat from './components/Chat';
import Template from './components/Template';

function App() {
  useEffect(() => {
  axios.get("http://localhost:5000/api/v1/messages/").then((data) => {
    console.log(data);
  });
}, []);
  return (
    <MantineProvider
      theme={{
        fontFamily: "Open Sans, sans serif",
        spacing: { xs: 15, sm: 20, md: 25, lg: 30, xl: 40 },
      }}
    >
      <Center >
        <Chat />
        <Template />
      </Center>
    </MantineProvider>
  );
}
export default App;
