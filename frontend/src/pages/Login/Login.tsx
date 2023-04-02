import React, {useEffect, useState, FormEvent} from 'react';
import Avatar from '@mui/material/Avatar';
import Button from '@mui/material/Button';
import TextField from '@mui/material/TextField';
import FormControlLabel from '@mui/material/FormControlLabel';
import Checkbox from '@mui/material/Checkbox';
import Box from '@mui/material/Box';
import LockOutlinedIcon from '@mui/icons-material/LockOutlined';
import Typography from '@mui/material/Typography';
import { InputAdornment } from '@mui/material';
import { RemoveRedEye } from '@mui/icons-material';
import VisibilityOffIcon from '@mui/icons-material/VisibilityOff';
import useApi from '../../hooks/useApi';
import { useNavigate } from 'react-router-dom';
import { toast } from 'react-toastify';
import { useCookie } from '../../hooks/useCookie';

export const Login : React.FC = () => {
    const cookie = useCookie();
    const api = useApi();
    const navigate = useNavigate();
    const [showPassword, setShowPassword] = useState(false);
    const [saveToken, setSaveToken] = useState(false);

    const handleSubmit = (event: FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        const data = new FormData(event.currentTarget);
        const token = data.get('token') as string;

        if (!token) return;

        api.login(token, saveToken)
            .then(() => navigate("/stream"))
            .catch(() => toast("Something went wrong.", {type: "warning"}))
    };
    
    useEffect(()=> {
        const token = cookie.getToken();
        if (token) {
            api.login(token, false)
                .then(() => navigate("/stream"))
                .catch()
        }
    },[])


    return (
        <Box
            sx={{
                marginTop: -10,
                display: 'flex',
                flexDirection: 'column',
                alignItems: 'center',
            }}
        >
            <Avatar sx={{ m: 1, bgcolor: 'secondary.main' }}>
                <LockOutlinedIcon />
            </Avatar>
            <Typography component="h1" variant="h5">Sign in</Typography>
            <Box component="form" onSubmit={handleSubmit} sx={{ mt: 1 }}>
                <TextField
                    margin="normal"
                    required
                    fullWidth
                    name="token"
                    label="Token"
                    type={showPassword ? "text" : "password"}
                    id="token"
                    InputProps={{
                        endAdornment: (
                            <InputAdornment position="start" >
                                {showPassword
                                    ? <VisibilityOffIcon onClick={() => setShowPassword(false)} sx={{ cursor: "pointer" }} />
                                    : <RemoveRedEye onClick={() => setShowPassword(true)} sx={{ cursor: "pointer" }} />
                                }
                            </InputAdornment>
                        ),
                    }}
                    autoComplete="current-token"
                />
                <FormControlLabel
                    control={<Checkbox id="remember" color="primary" />}
                    onChange={() => setSaveToken(!saveToken)}
                    label="Remember token"
                />
                <Button
                    type="submit"
                    fullWidth
                    variant="contained"
                    sx={{ mt: 3, mb: 2 }}
                >
                    LogIn
                </Button>
            </Box>
        </Box>
    );
}