import { Healthcheck } from "../types";
import { useCookie } from "./useCookie";
import { hookstate, useHookstate } from '@hookstate/core';


const REACT_APP_BASE_URL = process.env.REACT_APP_BASE_URL;
const token = hookstate<string | undefined>(undefined);

const useApi = () => {
    const cookie = useCookie();
    const state = useHookstate(token);
    
    if (state.get() === undefined) state.set(cookie.getToken());

    const login = async (token: string, saveSession?: boolean) => {
        if (saveSession) {
            cookie.setToken(token);
        }
        
        const res = await validateToken(token);
        if (res) {
            state.set(token);
            return true;
        }
    }

    const validateToken = async (token: string) => {
        const res = await fetch(`${REACT_APP_BASE_URL}/api/v1/auth/token`, {
            method: 'POST',
            headers: { 
                "Content-Type": "application/json",
                'Token' : token 
            }
        });

        if (res.status === 200) {
            return true;
        }
        else throw new Error('Something went wrong querying the server');
    }
    
    const healthchecks = async () => {
        try {
            const res = await fetch(`${REACT_APP_BASE_URL}/api/v1/healthcheck`, {
                method: 'GET',
                headers : {
                    "Content-Type": "application/json",
                },
            });
            return res.json() as Promise<Healthcheck>;
        }
        catch (e){
            return {
                status: false,
            }
        }
        
    }

    const getToken = () => state.get();
    
    return { login, healthchecks, getToken }

}

export default useApi;