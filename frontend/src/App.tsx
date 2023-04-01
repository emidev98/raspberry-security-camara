import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import { createTheme, ThemeProvider } from '@mui/material/styles';
import { routes } from './routes';
import './App.scss';

const theme = createTheme({
  palette: {
    mode: 'dark',
    background: {
      default: '#1f1f1f',
    },
    primary: {
      main: '#fff',
    },
    secondary: {
      main: '#0ae98a',
    }
  },
});

const App = () => {
  const router = createBrowserRouter(routes());

  return (
    <div className="App">
      <ThemeProvider theme={theme}>
        <RouterProvider router={router} />
      </ThemeProvider>
    </div>
  );
}

export default App;