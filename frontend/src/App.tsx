import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import { createTheme, ThemeProvider } from '@mui/material/styles';
import { routes } from './routes';
import './App.scss';
import useApi from './hooks/useApi';
import { useEffect } from 'react';
import Container from '@mui/material/Container';
import { CssBaseline } from '@mui/material';
import { ToastContainer } from 'react-toastify';

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
    const api = useApi();

    useEffect(() => {
        api.healthchecks().then((res) => {
            if (!res.status) router.navigate('/error');
        })
    }, [api, router])


    return (
        <div className="App">
            <ThemeProvider theme={theme}>
                <Container component="main" maxWidth="xs">
                    <CssBaseline />
                    <RouterProvider router={router} />
                </Container>
                <ToastContainer theme="dark"
                    position="top-left" />
            </ThemeProvider>
        </div>
    );
}

export default App;